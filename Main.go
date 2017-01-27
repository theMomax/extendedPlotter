package main

import(
    "image"
    "fmt"
    "image/png"
    "image/color"
    "os"
    "imgProcessor/simplexOut"
    "project-x/scanner"
)

func main(){
    test()
    fmt.Println("---------------------------------------------------------------")
    fmt.Println("------------------ Welcome to ImageProcessor ------------------")
    fmt.Println("------------------   (C)2017 Max Obermeier   ------------------")
    fmt.Println("---------------------------------------------------------------")
    var filename, identifier, separator string
    var accuracy int
    filename = "output.txt"
    identifier = "$Data"
    separator = "/"
    accuracy = 0
    for {
        fmt.Println()
        fmt.Println("Enter help to get a list of options or type in any other command.")
        input := scanner.GetS("==","help","colors","settings","process","exit")
        if input == "help" {
            help()
        }else if input == "exit" {
            os.Exit(0)
        }else if input == "colors"{
            listColors()
        }else if input == "settings" {
            filename, identifier, separator, accuracy = getParameters()
        }else if input == "process" {
            createImg(filename,identifier,separator,accuracy)
        }
    }

}

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
    fmt.Println("  - > 7 \t \t=> white")

}

func help(){
    fmt.Println("List of options:")
    fmt.Println("  - help \t \t=> Show list of options")
    fmt.Println("  - settings \t \t=> Set processing parameters")
    fmt.Println("  - process \t \t=> Start image creating process")
    fmt.Println("  - colors \t \t=> Show list of colors")
    fmt.Println("  - exit \t \t=> Exit program")
}

func getParameters() (filename, identifier, separator string, accuracy int){
    fmt.Println("Enter the filename of the input file:")
    filename = scanner.GetString()
    fmt.Println("Enter the identifier, the lines containing data start with:")
    identifier = scanner.GetString()
    fmt.Println("Enter the separator, the values are separated with:")
    separator = scanner.GetString()
    fmt.Println("Enter the number of decimal places the coordinates are cut off after:")
    accuracy = scanner.GetI("><",0,10)
    return
}

func test(){
    var colors []color.RGBA
    colors = append(colors, color.RGBA{255,255,255,255})
    colors = append(colors, color.RGBA{255,0,0,255})
    colors = append(colors, color.RGBA{0,0,255,255})
    colors = append(colors, color.RGBA{0,255,0,255})
    colors = append(colors, color.RGBA{0,255,255,255})
    colors = append(colors, color.RGBA{255,0,255,255})
    colors = append(colors, color.RGBA{255,255,0,255})
    colors = append(colors, color.RGBA{0,0,0,255})
    rect := image.Rectangle{image.Point{0, 0}, image.Point{10, 9}}
    img := image.NewRGBA(rect)

    for i := range colors {
        for j := 0; j < 10; j++ {
            img.SetRGBA(j, i, colors[i])

        }

    }
    f, _ := os.Create("out.png")
    defer f.Close()
    err := png.Encode(f, img)
    if err != nil {
        fmt.Println(err)
    }
}

func createImg(filename, identifier, separator string, accuracy int){
    var colors []color.RGBA
    colors = append(colors, color.RGBA{255,255,255,255})
    colors = append(colors, color.RGBA{255,0,0,255})
    colors = append(colors, color.RGBA{0,0,255,255})
    colors = append(colors, color.RGBA{0,255,0,255})
    colors = append(colors, color.RGBA{0,255,255,255})
    colors = append(colors, color.RGBA{255,0,255,255})
    colors = append(colors, color.RGBA{255,255,0,255})
    colors = append(colors, color.RGBA{0,0,0,255})
    d := simplexOut.NewData()
    d.CreateFromFile(filename, identifier, separator, accuracy)
    rect := image.Rectangle{image.Point{0, 0}, image.Point{len(d.Img), d.GetWidth()}}
    img := image.NewRGBA(rect)

    for i := range d.Img {
        for j := range d.Img[i]{
            if d.Img[i][j] < len(colors) {
                img.SetRGBA(i, j, colors[d.Img[i][j]])
            }else {
                img.SetRGBA(i, j, colors[0])
            }
        }
    }
    f, _ := os.Create("out.png")
    defer f.Close()
    err := png.Encode(f, img)
    if err != nil {
        fmt.Println(err)
    }
}
