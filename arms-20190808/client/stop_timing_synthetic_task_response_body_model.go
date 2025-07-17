// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StopTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v bool) *StopTimingSyntheticTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StopTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopTimingSyntheticTaskResponseBody
	GetRequestId() *string
}

type StopTimingSyntheticTaskResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StopTimingSyntheticTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTimingSyntheticTaskResponseBody) SetCode(v int64) *StopTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopTimingSyntheticTaskResponseBody) SetData(v bool) *StopTimingSyntheticTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopTimingSyntheticTaskResponseBody) SetMessage(v string) *StopTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopTimingSyntheticTaskResponseBody) SetRequestId(v string) *StopTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTimingSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
