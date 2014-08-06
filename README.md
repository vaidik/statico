# statico

Go's HTTP server for serving static files brought to the CLI for ease-of-use.
And it is the beginning of my journey in the world of Go.

## Usage

For serving static files from the current directory:

    $ statico
    Serving static files from directory /home/vaidik on 0.0.0.0:8080

### Serve different directory

    $ statico -path /home/vaidik/movies
    Serving static files from directory /home/vaidik/movies on 0.0.0.0:8080

### Listen to different host and/or port

    $ statico -host 192.168.1.2 -port 4000
    Serving static files from directory /home/vaidik on 192.168.1.2:4000

### Get help

    $ statico -help

## Installation

    git clone https://github.com/vaidik/statico.git
    cd statico
    go build statico.go
    mv statico /usr/bin

## License

`statico` is MIT licensed. See [LICENSE.md][1].

[1]: https://github.com/vaidik/statico/blob/master/LICENSE.md
