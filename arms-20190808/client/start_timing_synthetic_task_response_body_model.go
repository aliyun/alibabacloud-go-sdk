// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StartTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v bool) *StartTimingSyntheticTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StartTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartTimingSyntheticTaskResponseBody
	GetRequestId() *string
}

type StartTimingSyntheticTaskResponseBody struct {
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
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StartTimingSyntheticTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTimingSyntheticTaskResponseBody) SetCode(v int64) *StartTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartTimingSyntheticTaskResponseBody) SetData(v bool) *StartTimingSyntheticTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartTimingSyntheticTaskResponseBody) SetMessage(v string) *StartTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartTimingSyntheticTaskResponseBody) SetRequestId(v string) *StartTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTimingSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
