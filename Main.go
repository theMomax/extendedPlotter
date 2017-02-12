package main

import(
    "image"
    "fmt"
    "image/png"
    "image/color"
    "os"
    "imgProcessor/data"
    "log"
    "project-x/scanner"
)

var filenameIn, filenameOut, identifier, separator string
var accuracy int

//set default settings and start main menu
func main(){
    fmt.Println("-------------------------------------------------------------------------")
    fmt.Println("----------------------- Welcome to ImageProcessor -----------------------")
    fmt.Println("-----------------------   (C)2017 Max Obermeier   -----------------------")
    fmt.Println("-------------------------------------------------------------------------")
    filenameIn = "output.txt"
    identifier = "$Data"
    separator = "/"
    accuracy = 0
    menu()
}

//main menu
func menu(){
    for {
        fmt.Println()
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","license","colors","settings","process","exit")
        if input == "help" {
            help()
        }else if input == "license" {
            license()
        }else if input == "exit" {
            os.Exit(0)
        }else if input == "colors"{
            listColors()
        }else if input == "settings" {
            filenameIn, filenameOut, identifier, separator, accuracy = getParameters()
        }else if input == "process" {
            createImg(filenameIn, filenameOut, identifier, separator, accuracy)
        }
    }
}

//print license
func license(){
    fmt.Println("MIT License")
    fmt.Println("Copyright (c) 2017 Max Obermeier")
    fmt.Println("")
    fmt.Println(`Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:"`)
    fmt.Println("")
    fmt.Println("The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.")
    fmt.Println("")
    fmt.Println(`THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`)
    fmt.Println("")
}

//list colors and according values
func listColors(){
    fmt.Println("The third parameter of each highlighted line is portraied as a color.")
    fmt.Println("Here is a list of the colors, with its according value.")
    fmt.Println("  - 0 \t \t=> white")
    fmt.Println("  - 1 \t \t=> red")
    fmt.Println("  - 2 \t \t=> blue")
    fmt.Println("  - 3 \t \t=> green")
    fmt.Println("  - 4 \t \t=> turquoise")
    fmt.Println("  - 5 \t \t=> purple")
    fmt.Println("  - 6 \t \t=> yellow")
    fmt.Println("  - 7 \t \t=> black")
    fmt.Println("  - > 7 \t=> white")

}

//list all commands
func help(){
    fmt.Println("List of options:")
    fmt.Println("  - help \t \t=> Show list of options")
    fmt.Println("  - license \t \t=> Show license")
    fmt.Println("  - settings \t \t=> Set processing parameters")
    fmt.Println("  - process \t \t=> Start image creating process")
    fmt.Println("  - colors \t \t=> Show list of colors")
    fmt.Println("  - exit \t \t=> Exit program")
}

//change settings
func getParameters() (filenameIn, filenameOut, identifier, separator string, accuracy int){
    fmt.Println("Enter the filename of the input file (ending with .txt):")
    filenameIn = scanner.GetString()
    fmt.Println("Enter the identifier, the lines containing data start with:")
    identifier = scanner.GetString()
    fmt.Println("Enter the separator, the values are separated with:")
    separator = scanner.GetString()
    fmt.Println("Enter the number of decimal places the coordinates are cut off after:")
    accuracy = scanner.GetI("><",0,10)
    fmt.Println("Enter the filename of the output file (ending with .png):")
    filenameOut = scanner.GetString()
    return
}

//create image
func createImg(filenameIn, filenameOut, identifier, separator string, accuracy int){
    //build color slice
    var colors []color.RGBA
    colors = append(colors, color.RGBA{255,255,255,255})
    colors = append(colors, color.RGBA{255,0,0,255})
    colors = append(colors, color.RGBA{0,0,255,255})
    colors = append(colors, color.RGBA{0,255,0,255})
    colors = append(colors, color.RGBA{0,255,255,255})
    colors = append(colors, color.RGBA{255,0,255,255})
    colors = append(colors, color.RGBA{255,255,0,255})
    colors = append(colors, color.RGBA{0,0,0,255})
    //create data from file
    d := data.NewData()
    err := d.CreateFromFile(filenameIn, identifier, separator, accuracy)
    //check for any errors
    if err != nil {
        log.Println(err)
        return
    }
    //create rectangle with fitting dimensions (rectangle is used to build a rgbaImage)
    rect := image.Rectangle{image.Point{0, 0}, image.Point{d.X, d.Y}}
    //build image form rectangle
    img := image.NewRGBA(rect)

    //range over x and y coordinates of data.Img
    for x := range d.Img {
        for y := range d.Img[x]{
            //if there is a specific color for this value in the colors slice use it
            if d.Img[x][y] < len(colors) {
                //invert y - coordinates, because (0|0) of a image/png is at the top left corner and not at the bottom left as in a coordinate system
                img.SetRGBA(x, d.Y - y, colors[d.Img[x][y]])
            //if there is no color specified just use white
            }else {
                //invert y - coordinates, because (0|0) of a image/png is at the top left corner and not at the bottom left as in a coordinate system
                img.SetRGBA(x, d.Y - y, colors[0])
            }
        }
    }

    //finally save the image with the given name
    f, _ := os.Create(filenameOut)
    defer f.Close()
    err = png.Encode(f, img)
    if err != nil {
        log.Println(err)
        return
    }
}
