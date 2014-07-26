# NASDAQ OMX simulator

NASDAQ OMX simulator that replays NASDAQ OMX Nordic TotalView ITCH SoapFILE
files over MoldUDP protocol.

## Installation

To install the simulator:

```
$ go get github.com/penberg/nasdaq-omx-simulator
```

## Usage

The simulator needs access to an ITCH file that follows the SoupFILE format.

You can start the simulator with:

```
$ nasdaq-omx-simulator <ITCH-file>
```

The simulator then starts UDP multicasting the ITCH messages using MoldUDP
protocol.

## License

Copyright © 2014 Pekka Enberg and contributors

NASDAQ OMX simulator is distributed under the 2-clause BSD license.
