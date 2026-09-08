// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLLogContextLine interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *RLLogContextLine
	GetMessage() *string
	SetTimestampMs(v int64) *RLLogContextLine
	GetTimestampMs() *int64
}

type RLLogContextLine struct {
	// The log text (<= 2000 characters, with ANSI escape codes stripped).
	//
	// example:
	//
	// CUDA out of memory. Tried to allocate 2.00 GiB
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The millisecond timestamp of the log line.
	//
	// example:
	//
	// 1787293208012
	TimestampMs *int64 `json:"TimestampMs,omitempty" xml:"TimestampMs,omitempty"`
}

func (s RLLogContextLine) String() string {
	return dara.Prettify(s)
}

func (s RLLogContextLine) GoString() string {
	return s.String()
}

func (s *RLLogContextLine) GetMessage() *string {
	return s.Message
}

func (s *RLLogContextLine) GetTimestampMs() *int64 {
	return s.TimestampMs
}

func (s *RLLogContextLine) SetMessage(v string) *RLLogContextLine {
	s.Message = &v
	return s
}

func (s *RLLogContextLine) SetTimestampMs(v int64) *RLLogContextLine {
	s.TimestampMs = &v
	return s
}

func (s *RLLogContextLine) Validate() error {
	return dara.Validate(s)
}
