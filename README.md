# go-chuck-facts

go-chuck-facts is a cli tool build in golang to fetch random or categorised facts about the legend Chuck Norris from [chucknorris.io](https://api.chucknorris.io)

## Installation

Clone the repository and make sure that golang is installed.

```bash
$ git clone https://github.com/nga1hte/go-chuck-facts.git
$ go build gocf.go
```
export the binary to PATH for easy calling.
```bash
export PATH=$PATH:$(pwd)
```


## Usage
To run the program execute
```
$ ./gocf
> Chuck Norris frequently signs up for beginner karate classes so he can accidentally roundhouse kick kids in the neck.
```
Flags and help
```
 ./gocf -h
Usage of ./gcf:
  -c string
    	categories: animal, career, celebrity, dev, explicit, fashion, food, history, money, movie, music, political, religion, science, sport, travel
  -n int
    	Number of facts to fetch (default 1)

```


## Contributing

Any contribution is welcomed. Project created for practice.

## License

[MIT](https://choosealicense.com/licenses/mit/)
