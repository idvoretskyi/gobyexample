# Коли ми запустимо цю программу, ми побачимо
# спершу результат блокуючого виклику, а вже
# потім результат виконання  двох горутин.
# Як ми можете помітити, обидві горутини
# виконуються одночасно і контролюється
# середовищем виконання Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# _Канали_, є доповненням до горутин у
# механізмах одночасного виконання реалізованих у Go.