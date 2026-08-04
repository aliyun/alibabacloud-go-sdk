// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageResponseBody
	GetCode() *string
	SetMessage(v string) *ReadMessageResponseBody
	GetMessage() *string
	SetResult(v bool) *ReadMessageResponseBody
	GetResult() *bool
}

type ReadMessageResponseBody struct {
	// Status code returned by the service. Only "SUCCESS" indicates success; all other values indicate failure.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 外部userId映射关系不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Succeeded in marking as read
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReadMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ReadMessageResponseBody) SetCode(v string) *ReadMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageResponseBody) SetMessage(v string) *ReadMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageResponseBody) SetResult(v bool) *ReadMessageResponseBody {
	s.Result = &v
	return s
}

func (s *ReadMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
