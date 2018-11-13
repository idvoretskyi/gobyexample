// На відміну від [таймерів](timers), призначення яких
// відстежувати час до одноразової події в майбутньому,
// _маятники_ мають своєю метою виконання події через
// чітко задані інтервали.

package main

import "time"
import "fmt"

func main() {

    // Маятники використовують канал до якого надсилають значення,
    // це дуже схоже на те, як працюють таймери. Тут ми використаємо
    // `range` для каналу маятника і відразу виводимо значення
    // яке прибуває з каналу (раз на пів секунди).
    ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        var f = "15:04:05.999999999"
        for t := range ticker.C {
            fmt.Println("Коливання @", t.Format(f))
        }
    }()

    // Після того як таймер зупинено (зауважте аналогію з зупинкою таймера),
    // з його каналу перестають надходити повідомлення.
    // Зупинимо наш таймер через 1600 мілісекунд
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    fmt.Println("Маятник зупинено")
}
