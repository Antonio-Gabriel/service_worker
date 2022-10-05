# Service Worker

I simulate a miltithreads in golang and understand how he run in the backgroung.
In fact I really liked of way that the GO to work with threads and solve a multiple process at once.

I saw too that is possible to specify how many cors you have at your machine to allow the process runs more fast or faster.

So nice yh. I recomend.

```go
channel := make(chan int)

go func() {
    for i := 1; i <= 10; i++ {
        go worker(channel)
    }
}()
```

Resume, the secret for runs at once is here, you can create `n` machines, I think that depend on the process that you want to trait,
next paste the data on channel (The channel can being anything), and when the worker receive the data will process of anyway.

I really understood in this simple example:

Worker func:

```go
func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i, " done")
		time.Sleep(time.Second * 5)
	}
}
```
It's **GO Way**
