package main

import (
    "fmt"
    "os"
    qr "github.com/skip2/go-qrcode"
    term "github.com/nsf/termbox-go"
)

func reset() {
     term.Sync() // cosmestic purpose
}

func main() {
    textBlocks := [4]string{
        "Hello Pilot.\nThis is your Captain. If you are reading this message, I am already out of duties." +
            " That means that from now you are responsible for Galaxy №20-18.\nBad news: one of your planets, Earth 1.0 (Solar system, 3rd from Sun)," +
            " has been damaged.\nDescription: Earth 1.0. has 3 spaces: of place, people minds and time.\n" +
            "Problem: Each of the spaces has a toxic content in it.\n1 level task: Run the program: \"ReFill_Space\" to save the planet." +
            " Each step is about replacing bad components with the healing ones.\nFollow my instructions and you will complete it.\n" +
            "Keep your pants dry, Pilot.\nPress \"E\" to start.\nP.S. write down the keys I'm telling you, you will need them later.",
        "Step 1. Space of Place.\nSubjects: buildings, cities, countries, land, atmosphere\n" +
            "Toxic components: commercialization + insecurity, traffic jams, corrupt government, tons of trash, CO2;\n" +
            "Heeling components: social living + safety, comfortable social and traffic structures, second use of materials, electro mobiles;\n" +
            "To refill press \"A\"",
        "Step 2. Space of Minds.\n" +
            "Subjects: thoughts, memory, attitude;\n" +
            "Toxic components: problems, rush, harm, intolerance, disregard\n" +
            "Heeling components: plans, dreams, mindfulness, kindness;\nTo refill press \"S\"",
        "Step 3. Space of Time.\n" +
            "Subjects: energy\nToxic components: fear, darkness\n" +
            "Heeling components: pure light\nTo refill press \"A\"",
    }

    rightKeys := [4]string{
        "e",
        "a",
        "s",
        "a",
    }

    var q *qr.QRCode
    q, _ = qr.New("https://ru.pinterest.com/pin/400679698092731991/", qr.Medium)

    err := term.Init()
    if err != nil {
        fmt.Println(err)
    }

    defer term.Close()

    for i, quote := range textBlocks {
        reset()
        fmt.Println(quote)

        keyPressListenerLoop:
        for {
            switch ev := term.PollEvent(); ev.Type {
                case term.EventKey:
                switch ev.Ch {
                    case rune(rightKeys[i][0]):
                    break keyPressListenerLoop
                    default:
                    fmt.Println("Wrong key..")
                }
                case term.EventError:
                    panic(ev.Err)
            }
        }
    }
    reset()
    fmt.Print("You did it. Thank you, pilot.\nNow there is one thing I haven't told you - one more step to finish this program: \n" +
        "Save your own butt.\n" +
        "if you = architect then\n" +
        "    do it good, make it comfortable;\n" +
        "if you = young then\n" +
        "    stay curious, brave, naive, wide open and true;\n" +
        "if you = human then\n" +
        "    hug the loved ones (:= whole world) and drink tea;\n" +
        "else\n" +
        "    do what you ought to, come what may;\n" +
        "Put all the collected letters then press Enter to complete.\n")

    finalWordLoop:
    for {
        word := ""
        fmt.Scanf("%s", &word)
        if word == "easa" {
            break finalWordLoop
        }
        fmt.Println("Try again...")
    }

    reset()
    fmt.Println("\nGood job. Scan this QR code with your smartphone to see your score.")
    fmt.Println(q.ToString(true))
    fmt.Scanf("%s")
    os.Exit(0)
}