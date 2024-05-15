# Простой echo сервер

## Пример работы с middleware
В проекте используется Веб-фреймворк Echo https://github.com/labstack/echo


В ответе возвращается число дней до 2025 года. При этом, если в заголовке передается роль admin, то логгер выводит сообщение про красную кнопку.

```shell
curl --location --request GET '127.0.0.1:8080/status' \
--header 'User-Role: admin'
```  
В ответе:
```
Days left: 230
```

В логе(консоле):
```
2024/05/15 18:06:54 red button user detected
```


По материалам открытого урока:
* https://www.youtube.com/watch?v=Lsh3ylmXdJ8
