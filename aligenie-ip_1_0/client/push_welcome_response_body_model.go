// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *PushWelcomeResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushWelcomeResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushWelcomeResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *PushWelcomeResponseBody
	GetStatusCode() *int32
}

type PushWelcomeResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PushWelcomeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeResponseBody) GoString() string {
	return s.String()
}

func (s *PushWelcomeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushWelcomeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushWelcomeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushWelcomeResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushWelcomeResponseBody) SetMessage(v string) *PushWelcomeResponseBody {
	s.Message = &v
	return s
}

func (s *PushWelcomeResponseBody) SetRequestId(v string) *PushWelcomeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushWelcomeResponseBody) SetResult(v bool) *PushWelcomeResponseBody {
	s.Result = &v
	return s
}

func (s *PushWelcomeResponseBody) SetStatusCode(v int32) *PushWelcomeResponseBody {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeResponseBody) Validate() error {
	return dara.Validate(s)
}
