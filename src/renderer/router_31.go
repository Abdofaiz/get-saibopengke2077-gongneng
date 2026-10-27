package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func raDqeMRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQCJs5zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZykV9s1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LvFpW5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSJTjvorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsNl99xjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joPr59AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMTbLO1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwIgJEYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6VMuLoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuKLKR8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAMZwvIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3piIn2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMFeU43KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiW7xcQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j36U1Vv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3QH14HYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1Kso7XYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLxXpdTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNxhvo0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Os49oaHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcBHs8dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzOL3TF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYeg3myYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFUSsooUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cK43oS25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 196rpQbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsSlcPiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjjRoF2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6D3uWiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hh1IHAaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDAGIyirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u31687jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTjBmzLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dN2HZ1FyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsmMtZAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2SmgTJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqQqqe5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFRUWFA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hw1FibR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7sXzMZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCw41QN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSgCCr7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KoHF79Z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ze3WB9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeARWeJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flVrbYbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y14wdGbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WPzq49YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cUdg36fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ff3o5XUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAYPPlOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEzI02z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWjkLsYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr05q4NqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcRlpc7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbeXaUahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXmPJ6JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7HbOH2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e4lGInNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSXh7MIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZuLeCOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esAF5lk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBghZdnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0x1zmZKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdJadYOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pt9SmwGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppL9pwBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PQOXCx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AW8n2WirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAbj6wkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGfHhQyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cFS6HiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhgTPxsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asUiDZ9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpfK0C3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yZ2R14ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5ysb9ArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZB4qN9n2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQXbRT8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmg9y307Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHcxNC7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IerrATsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Lgu6Ck5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vEJ97OYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvRHQvSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3V4rx21QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pceeKQ3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZV56KZldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKOcatKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UR12hmmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZmhZEWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xabvYU1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMCwRizmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RE274uxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPR4PvLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSSM16GVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDTHNsoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2k3En8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBode74eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBl8xtlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QB4viBQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfbErn6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClvStAosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvkgnTDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqV0TQkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZokVYW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SK4Ljb4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6D9cii9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mu6WhaElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaGHZK4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZkRrehkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQ4LlbaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wuh4ITXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VcB7CBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYoT14CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8NxQiK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCcRVSaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQUFBWyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QZES5TPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJ1eZmvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTWRtwCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPDJYIsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7M9LXzUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHZXfIyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g8JBHQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XApTqgv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0CItDlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbNyuXLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIaT9z0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhbOEhkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99KAAHQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcXI2W33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgXjtFvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnvchR78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuU4gXSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZyDkD4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func somTXtP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kv3MMl8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnLOUgXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIJugXmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAxzBuRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYEpa69wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cI4ERVdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iK02NXU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfkOEV06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTwrb7qHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkCLubQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19RnnlhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAJOQjdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bj5n3B5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbp9uFDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 378OZV0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTV3ln0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vlDb96pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rUj6FsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dbblXRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtdCVVBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1hnZhWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrI62MndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhAVKvxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKg9Ow08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apGla9JYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXuqSyShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsdT9HNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGXG5icBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKE0YodsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WmFUVKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YouxtYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8x1brUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqKvRADeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERDZkWNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPZZieSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYvA9tcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QyWnxgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlqRoHboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6Jmo2xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAFiFhntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SKdPlLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func To1ecrXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q6OA1QTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QmvEYBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0a28fhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXQV3WCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QTOlIIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHVYoeHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilZl7ifqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFxqV0zdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51IampATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aGq1nbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPUGAZ37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdqvvVkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sbzxYBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9T60RdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L426c4a7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFvShP0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2d8eP9cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7WkxuNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54vL5ShwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rf1eqtZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQcuDqXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUNhFNWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYClWaBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3tTYeJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmnv4JwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCgbDhjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLVY9bAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Diz3SYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zM8hL7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fv62wZjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm30oJ6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8aLassXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SurZm5irWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6YBQ88CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqdHsxwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0cGbvqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fp5LKiySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9K964sNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DUwHm6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgDATXIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkoVRuCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49PBJJZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tf4hjvnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNwkA7r1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDX2IsXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTXXdvneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNfA9LAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1hC2prIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcERWFSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9j44oyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KN2TQ21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iP3mRMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSbbJFQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6LaICysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpb6nCPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWjrOON5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQybAr7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G97mvJThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TR5AZqXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyMrmDF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipZVOUuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypGGX7NjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwPzoowLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ea2Q7q2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5eeMDgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func we6SmGpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AB5fwrbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mbKG6g8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMDFLoPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGdRBsSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4y3dGmojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3hZo75AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7Ag4yGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0THVTCl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MlCVeOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6DjbKB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3663cB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAsjLMoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMiPajtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ggv4c6WvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROPrZPYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRrjXiCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daK1zLjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNDPb9OMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D25QCLbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGhWaqqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHPxNAoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tW71j7ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZTEkJzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R1nZv7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWPY9A97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMGXuwGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydc1S2J1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psPMkgJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nak5cxOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BwrWPX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdKtOgVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrnqc7dCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykoEwX3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DH5VBwfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhrxjU3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46DkShtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLWbT8y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKIqAzwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3iM0efkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnYrHWS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diN67fMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeUTzunAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oTB3Em6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vbojrWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suA8HEYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N6UrZUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okR1uK5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6F0rUgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVqAnypkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6b86VIF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZY4bjCNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmLL7VlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vISgnybRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wI2Oaz79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2N31d3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDQQBvP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0XnSGUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxznEonwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytJl52eDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DlrlMN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkkJQIjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OalhnPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDkX8IaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98lPDIXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goj2SOfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5M8J6cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DogHKOTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPsxv79IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qO6zVuihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Jakf0aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nc7d2BjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yUZvKuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmvZhQrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gm5H3YqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWeXrO1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbZDV8ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBe49s6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nibEkQE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lncXEpqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNKU3WEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTRFmPD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mg6Gt0lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7bMW7jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHq4cFlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMPDUo2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uuKHMPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EF3aArshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KN4ynGaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgoRio25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKQvKl3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiHmkbsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBcag0gYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIn2Sc1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgUTYdUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CHRKbHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEsyWZNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3qzaXfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBsbLDAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rU6hGhh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDNINZK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWhcF7dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIIh2wqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63ZX44yvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXwtgczVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPWEL1JSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func In8QjMJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TBjhDK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOkrqGwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dODNwIyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ig2zSOKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sS8tontjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQsjQNPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzxFtoyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdsDG4lVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6d3d4IcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0aWmzVozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6LuRVbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7d2Q0dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDFuVoK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrT4hvCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYAQrjCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdUNqScvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkcrIIKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWmvc5V0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5oWzgTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcZjW8OXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvF665Y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t78ajg6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7w6Xm8ZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiLG3WkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgd2jcdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QT8gwd5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WMn086CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86EKyYFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnhZXzKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OY4wsouJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dfv6583pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKSSbfcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUfWsw2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQWhU385Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8NESqZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYmsokEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jmr2IoUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meiOlirpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfNOlPjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy5FtxkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2yEWD9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8g9Mk5nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROiNFc7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1kMATXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIguiYwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsWse4cTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezvdgVeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3a0i6Az1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ8aOQXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k70F48GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCq8HIiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2YhxnyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5sYiQRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JOOh0J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwsfinjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LI21tkT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYb4uXLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eo5OL5zcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4aSEaI9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vi9sYw2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2MdElpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCM3SbuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEFaMk46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2JXWnzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXYxIacQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2C0ev79iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0EnYt29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPGZMl2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szzs3f6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bY0TCjWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJRlUr8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FUDOPjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODmIispJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eG9xZ3A7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTMjwwB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpzaTKh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhAYCwNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38la4UbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42HKGo4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhkHWmH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7V3wygLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LHBg7WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DF1G0T8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVmlssq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2P5baX2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoZh3RF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28hu0eUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qy9xZiZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiggDyueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkOyZ5NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Isz6fZjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2WRZAMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LENHD9LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pQ4FmGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoFSVb1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyEhFaw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JE65wo3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTurP0VQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Jmwg2JdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cB7hl4euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXqb9W59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFXXYqNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTb5q0eyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5PtU334Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2lGMZfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BH6K3BVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bw5YiKlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfmkHteOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTASy1zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6w768yxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayV0UgfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7vfyw5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmkEgKIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QY8MgtAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaFnlweqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99ITFhLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfzWjKVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCK74KybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpsyFJMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5ydGe1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSyg0p9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLZYUgI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oD5ajx9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WVZFKRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEWiTt8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btfqbDKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLYAgNWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hd8jL7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ak3EQxX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntesYDY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBjZFtfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDjMkM9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1resbECvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJTXHGVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8amG4nvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cb3SBbLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQU7GmNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rfL2vt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHaol4xnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjn6il5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eX6nE7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEle8BjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQl87FLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0G2BzsCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8Rjcm6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vcje9XnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE3y93pTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZQa3IeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAxuB2bcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iopblPTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IfVysgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLzhuXVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odBVhu33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMCNydWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbdh1BtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kk1Jqj5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqRIIRrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1U0ebZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nci02nNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkBng4QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgGln5MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Nrg8WNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hcape87NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cypWSEKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZf198LFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdDLtPAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5ACPqkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5mAW7KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rKqllYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjaiqWiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OL8agx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4A8sYXssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtbF4cDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKfb9W8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoMNKn5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSdoQaaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2xEfQKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJ68gDgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyWFXJLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oS4dPLiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X084irgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9dOLGVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njSlXxSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gc5SgoHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6z5y0BDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NblO7NcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVjlJp0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfoOf0gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zP1GyzHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmz5lrqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFSgo2ROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFhvhTYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O82jRidwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YRgxzijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vu0NQYG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZIDZIJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CRV7kVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ1QexVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3OZnIwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMo2Q9boWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmgHuBgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxhxrwvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmUCGJr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgDGY2p4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExBzVbYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EqZ1YiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyZmoQRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k26kywdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNQm34OiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nklAVsdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrvlMcEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4XD7BtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func heSjTVRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1tvnOcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG4S90XFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cx0oIIP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ah3dqMuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPFStzWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gu8TT6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3kkAfNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHYhT3y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ytfIH4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIhlYhwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SPeThUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QL4qSjHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OoYc51QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEuuBX1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLTXIbiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4GiV8TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLJj0jm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78uGSsxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLiunn8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX9ufuZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBXLop09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zz3XzUgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdosPXrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umKVYffwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlWdUHWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ATHXRMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 196HgLcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAogFK1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5MdTMI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UDuaSDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iWDvNboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HU2j0zVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNLkqmXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hiagqJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4QrjdbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTJeF0cGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0il4JEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flqgpvNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMD1ezY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOJIFqmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJtAyvfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCVd2JVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQmowTewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJkjhflcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdqq8fEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsxY3kwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTU9NOVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAwIxyhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kyo1L5coWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whjie7TcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGYL4UsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func purmxALsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cATPnojQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOSLrz2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHYUpFzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwPp6k3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbLJZzHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FhEEJHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aksgy3D8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQykLuCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZKrbCm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvYZjpKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0ttPYHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5NrPdGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tI1fZujDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdKn3NCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Y7hIxWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeBcV20TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMoBqAjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWH6wsE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IIQReuw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijjqcGPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNoLsj0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTaNinWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlSAvycYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSXXbOzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZtWdwyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2esgFokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEAFivReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tDg8G6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBMgQ5OuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AK0CBfdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yd9Q6k77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGUdIkHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CE3RmCrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knbgFzigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opomnzEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPQlwwrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zl1iFUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRGgqIgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjNFcoawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSrXXlTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DqUOTklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7qf1aZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVk6zs7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3MHfWxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJlWfqZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

