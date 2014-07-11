# PushFinish

PushFinish is a small tool that notifies you via Pushover when a command line application has finished its operation.

## Set up

Note: You will need a Pushover account in order to use PushFinish.

1. Compile pushfinish.go: `go build pushfinish.go`
2. Put the binary in some directory that is listed in $PATH.
3. [Register the application with Pushover](https://pushover.net/apps/build). Set the environmental variable PUSHFINISH_TOKEN to the received application token.
4. Set the environmental variable PUSHFINISH_USER to your user key. You will find that after logging in at [Pushover](https://pushover.net/).

Top Tip: In order to set those environmental variables permanently add them to ~/.bashrc: `export PUSHFINISH_USER=your key here` and `export PUSHFINISH_TOKEN=your application token here`.

## Usage

`pushfinish ./your_long_task your_parameters`

It will automatically push a notification as soon as the application finishes.
