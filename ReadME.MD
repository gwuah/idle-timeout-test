# fly idle timeout test


# Howoto
- Run fly deploy (don't forget to change the app-name)
- Run `time go run client/main.go https://<your-app-name>.fly.dev/longconnection`
- I set a timeout of 10 minutes, my server only responds after 11 minutes, but i'm somehow getting a response from my server. (which means the timeout isn't being applied)
- See sample run below 
```
g@MacBook-Air longconnection $ %{time go run client/main.go https://longconnection.fly.dev/longconnection
longconnection
go run client/main.go https://longconnection.fly.dev/longconnection  0.25s user 0.28s system 0% cpu 11:01.73 total
```