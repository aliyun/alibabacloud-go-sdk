// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAligenieUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UnbindAligenieUserResponseBody
	GetCode() *int32
	SetMessage(v string) *UnbindAligenieUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindAligenieUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindAligenieUserResponseBody
	GetSuccess() *bool
}

type UnbindAligenieUserResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAligenieUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAligenieUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindAligenieUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindAligenieUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAligenieUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindAligenieUserResponseBody) SetCode(v int32) *UnbindAligenieUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetMessage(v string) *UnbindAligenieUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetRequestId(v string) *UnbindAligenieUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) SetSuccess(v bool) *UnbindAligenieUserResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindAligenieUserResponseBody) Validate() error {
	return dara.Validate(s)
}
