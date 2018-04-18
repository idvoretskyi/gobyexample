# Коли ми запускаємо прогрумму повідомлення `"ping"`
# успішно передається з однтєї горутини до нашої
# программи через наш канал.
$ go run channels.go
ping

# Відповідно налаштувань по замовчуванню _надсилання_
# та _отримання_ блокують виконання програми допоки
# _відправкник_ або _отримувач_ не будуть готові.
# Ця особливість дозволяє нашій программі зачекати
# на повідомлення  `"ping"` без використання інших
# типів синхронізації.