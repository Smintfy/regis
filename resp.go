package main

/*
 *	https://redis.io/docs/reference/protocol-spec/
 */
const (
	String     = "+"
	Error      = "-"
	Integer    = ":"
	BulkString = "$"
	Array      = "*"
)