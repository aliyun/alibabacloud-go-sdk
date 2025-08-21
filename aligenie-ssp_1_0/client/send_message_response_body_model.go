// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendMessageResponseBody
	GetCode() *string
	SetMessage(v string) *SendMessageResponseBody
	GetMessage() *string
	SetResult(v bool) *SendMessageResponseBody
	GetResult() *bool
}

type SendMessageResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendMessageResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SendMessageResponseBody) SetCode(v string) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetResult(v bool) *SendMessageResponseBody {
	s.Result = &v
	return s
}

func (s *SendMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
