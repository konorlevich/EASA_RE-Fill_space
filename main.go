package main

import (
    "fmt"
    qr "github.com/skip2/go-qrcode"
    term "github.com/nsf/termbox-go"
    "github.com/konorlevich/EASA_RE-Fill_space/questions"
)

var finalQuestion = questions.FinalQuestion{"You did it. Thank you, pilot.\n" +
    "Now there is one thing I haven't told you - one more step to finish this program: \n" +
    "Save your own butt.\n" +
    "if you = architect then\n" +
    "    do it good, make it comfortable\n" +
    "if you = young then\n" +
    "    stay curious, brave, naive, wide open and true\n" +
    "if you = human then\n" +
    "    hug the loved ones (:= whole world) and drink tea\n" +
    "else\n" +
    "    do what you ought to, come what may;\n" +
    "Put all the collected letters then press Enter to complete.",
    "easa",
}

var simpleQuestions = [4] questions.SimpleQuestion{
    {
        "Hello Pilot.\n" +
            "This is your Captain. If you are reading this message, I am already out of duties." +
            " That means that from now you are responsible for Galaxy №20-18.\n" +
            "Bad news: one of your planets, Earth 1.0 (Solar system, 3rd from Sun), has been damaged.\n" +
            "Description: Earth 1.0. has 3 spaces: of place, people minds and time.\n" +
            "Problem: Each of the spaces has a toxic content in it.\n" +
            "1 level task: Run the program: \"ReFill_Space\" to save the planet." +
            " Each step is about replacing bad components with the healing ones.\n" +
            "Follow my instructions and you will complete it.\n" +
            "Keep your pants dry, Pilot.\nPress \"E\" to start.\n" +
            "P.S. write down the keys I'm telling you, you will need them later.",
        'e'},
    {
        "Step 1. Space of Place.\n" +
            "Subjects: buildings, cities, countries, land, atmosphere\n" +
            "Toxic components: commercialization + insecurity, traffic jams, corrupt government, tons of trash, CO2\n" +
            "Heeling components: social living + safety, comfortable social and traffic structures, second use of materials, electromobiles\n" +
            "To refill press \"A\"",
        'a'},
    {
        "Step 2. Space of Minds.\n" +
            "Subjects: thoughts, memory, attitude\n" +
            "Toxic components: problems, rush, harm, intolerance, disregard\n" +
            "Heeling components: plans, dreams, mindfulness, kindness\n" +
            "To refill press \"S\"",
        's'},
    {
        "Step 3. Space of Time.\n" +
            "Subjects: energy\n" +
            "Toxic components: fear, darkness\n" +
            "Heeling components: pure light\n" +
            "To refill press \"A\"",
        'a'},
}

type QR struct {
    url string
}

func (q QR) print() {
    painter, _ := qr.New(q.url, qr.Medium)
    fmt.Println(painter.ToString(true))
}

var code = QR{"https://ru.pinterest.com/pin/400679698092731991/",}

func init() {
   err := term.Init()
   if err != nil {
       fmt.Println(err)
   }
}

func main() {
    for _, question := range simpleQuestions {
        question.Ask()
    }
    finalQuestion.Ask()
    term.Sync()
    fmt.Println("\nGood job. Scan this QR code with your smartphone to see your score.")
    code.print()
    fmt.Println("Press any key to exit")
    fmt.Scanf("%f")
    term.Close()
}
