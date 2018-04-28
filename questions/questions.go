package questions

import (
    "fmt"
    term "github.com/nsf/termbox-go"
    )

type SimpleQuestion struct {
    Text   string
    Answer rune
}

func (q SimpleQuestion) Ask() {
    term.Sync()
    fmt.Println(q.Text)

keyPressListenerLoop:
    for {
        switch ev := term.PollEvent(); ev.Type {
        case term.EventKey:
            switch ev.Ch {
            case q.Answer:
                break keyPressListenerLoop
            default:
                fmt.Println("Wrong key..")
            }
        case term.EventError:
            panic(ev.Err)
        }
    }
}

type FinalQuestion struct {
    Text   string
    Answer string
}

func (q FinalQuestion) Ask() {
    term.Sync()
    fmt.Println(q.Text)

finalWordLoop:
    for {
        var word string
        fmt.Scanf("%s", &word)
        if word == q.Answer {
            break finalWordLoop
        }
        fmt.Println("Try again...")
    }

}