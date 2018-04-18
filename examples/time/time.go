// Go має об’ємну підтримку часових значеннь
// та значеннь тривалості. Розглянемо наступні приклади.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Розмочнемо з отримання поточного часу.
    now := time.Now()
    p(now)

    // Ми можемо збудувати структуру часу надаючи значення
    // рік, місяць, день та інше. Час завжди ассоційований
    // з Розташуванням, інакше кажучи з часовою зоною.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // Ми можемо скористатись з різних частинок структури
    // значення часу.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // Також доступне для використання день тижня за
    // допомогою функції `Weekday`.
    p(then.Weekday())

    // Наступні методи порівнють два часи, перевіряючи
    // чи трапилась подію у відповідності до іншого часу
    // до, після або в тоже момент (з точністю до секунди) відповідно
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // `Sub` або різниця метод поверне `Duration` (тривалість) що
    // предсталвятиме інтервал між двома подіями.
    diff := now.Sub(then)
    p(diff)

    // Ми навіть можемо порахувати довжину тривалості у різноманітних
    // вечелинах.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // Ми можемо навіть скористатись методом `Add` для збільчення часу на
    // певну тривалість або зменшеня якщо буде використано
    // тривалість зі знаком мінус.
    p(then.Add(diff))
    p(then.Add(-diff))
}