# Run a checker in env test

This is an executable specification file. This file follows markdown syntax.
Every heading in this file denotes a scenario. Every bulleted point denotes a step.

To execute this specification, run

    gauge specs


* If I make a "GET" query to "http://google.com".

## Checking http response

tags: single word

* The status code should return "200".
