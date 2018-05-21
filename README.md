# timeit

go-timeit is an utility that calculate the absolute real time from a given milliseconds quantity.

# Usage

Get the package:

```bash
$ go get github.com/sanchezpaco/go-timeit/timeit
```

Import the package into your project: 

```go
import "github.com/sanchezpaco/go-timeit/timeit"
```

Use it!

```go
myTimeInMilliseconds := 28809412823.0

timePassed := timeit.TimeIt(myTimeInMilliseconds)

// Result:
// {
//  "Years": 1,
//  "Months": 9,
//  "Weeks": 0,
//  "Days": 1,
//  "Hours": 6,
//  "Minutes": 36,
//  "Seconds": 41
// }
```

## Versioning ##

In general, go-github follows [semver](https://semver.org/) 

# License

This library is distributed under the GNU General Public License found in the [LICENSE](./LICENSE) file.