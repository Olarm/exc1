### What is concurrency? What is parallelism? What's the difference?
> Concurrency is performing several tasks in the same time period, e.g starting several tasks and then finishing them after they are started. Parallelism is when tasks run in paralell, eg. threads running at the same time

# Why have machines become increasingly multicore in the past decade?
> There are a number of physical issues with increasing CPU speed, perhaps most important is that the energy requirement to run a faster CPU currently scales more or less exponentially with the gain in frequency, which also produces exponentially increasing heat compared to CPU gain. With multicore processors the issues of heat and energy is diminished while a program utilizing several cores can still run efficiently even though each single core runs at a lower speed.

 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > If 1 running task requires waiting then other tasks can be run in the meantime, making better use of hardware

 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > probably both, it adds better usage of available hardware but also added complexity to the code

 ### What are the differences between processes, threads, green threads, and coroutines?
 > Processes are independent sequences of execution in separate memory space. Threads are also independent sequences of execution but run in a process and several threads in a process also share memory space. A green thread is a thread executed by a virtual machine. Coroutines are similar to concurrency in that several execution sequences are run in the same time frame but not in parallel. The difference is that coroutines within a program are dependent upon eachother.

 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > 'pthread_create()' creates a new thread, 'threading.thread()' creates a new thread and 'go' creates coroutines

 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > The GIL is a mutex preventing multiple threads accessing python bytecodes at once. In practice this is a bottleneck when executing more then 1 thread

 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > The multiprocessing library provides a workaround through a multiprocessing approach. This gives each python process its own interpreter and memory space.

 ### What does `func GOMAXPROCS(n int) int` change?
 > The number of threads allocated to the program
