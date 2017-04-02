package eventsource

import "fmt"

const (
	AggregateNil      = "AggregateNil"
	DuplicateID       = "DuplicateID"
	DuplicateVersion  = "DuplicateVersion"
	DuplicateAt       = "DuplicateAt"
	DuplicateType     = "DuplicateType"
	InvalidID         = "InvalidID"
	InvalidAt         = "InvalidAt"
	InvalidVersion    = "InvalidVersion"
	InvalidEncoding   = "InvalidEncoding"
	UnboundEventType  = "UnboundEventType"
	AggregateNotFound = "AggregateNotFound"
	UnhandledEvent    = "UnhandledEvent"
)

type Error interface {
	error

	// Returns the original error if one was set.  Nil is returned if not set.
	Cause() error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string
}

type baseErr struct {
	cause   error
	code    string
	message string
}

func (b *baseErr) Cause() error    { return b.cause }
func (b *baseErr) Code() string    { return b.code }
func (b *baseErr) Message() string { return b.message }
func (b *baseErr) Error() string   { return fmt.Sprintf("[%v] %v - %v", b.code, b.message, b.cause) }
func (b *baseErr) String() string  { return b.Error() }

func NewError(err error, code, format string, args ...interface{}) error {
	return &baseErr{
		code:    code,
		message: fmt.Sprintf(format, args...),
		cause:   err,
	}
}
