# ascii-art-web-stylize

## Цель задания 
Ascii-art-web заключается в создании и запуске сервера, на котором можно будет использовать веб- GUI (графический пользовательский интерфейс) версию последнего проекта, ascii-art.

На веб-странице должно быть разрешено использование различных баннеров:
   *standard
   *shadow
   *thinkertoy
# Авторы
@abaltaba
@nzharylk
#####
Должно быть реализованы следующие конечные точки HTTP:
   *GET /: Отправляет HTML-ответ, главную страницу.
   1.1. Совет GET: перейдите в шаблоны для получения и отображения данных с сервера.

   *POST /ascii-art: отправляет данные на сервер Go (текст и баннер)
   2.1. Совет POST: используйте форму и другие типы тегов , чтобы отправить запрос на публикацию.

На главной странице должны быть:
   *ввод текста
   *радиокнопки, выберите объект или что-нибудь еще, чтобы переключаться между баннерами
   *Кнопка, которая отправляет запрос POST в '/ascii-art' и выводит результат на страницу.


## Как запустить программу?

`go run main.go` или  `go run .`
 
`перейти на сайт по ссылке http://localhost:4040`