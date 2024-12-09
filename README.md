This repository demonstrates how to profile Go applications using the [Day 5 - Part 2](https://adventofcode.com/2024/day/5) challenge from Advent of Code 2024 as an example.

The repository includes both unoptimized and optimized versions of the code to demonstrate performance improvements through profiling.

Usage:

1.	Build the application:
```
go build -o profiling-app
```

2.	Run with profiling enabled:
```
./profiling-app
```

3.	Visualize the profile:
```
go tool pprof -http=":8000" ./profiling-app ./cpu.pprof
```

Performance Results:
- Original implementation (Part2): ~4.6s execution time
- Optimized version (Part2Fixes): ~0.42s execution time

Tools Used
- pprof
- Flamegraph.com
- Speedscope.app

For a detailed explanation of the profiling process and visualization tools, check out the blog post: ["Enhancing Go performance: Profiling applications with flamegraphs"](https://blog.alexoglou.com/posts/profiling-golang/)
