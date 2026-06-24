// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateOpenSourceAccountResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateOpenSourceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOpenSourceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOpenSourceAccountResponseBody
	GetSuccess() *bool
}

type UpdateOpenSourceAccountResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8BFB1C9D-08A2-4859-A47C-403C9EFA2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOpenSourceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateOpenSourceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOpenSourceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOpenSourceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOpenSourceAccountResponseBody) SetCode(v int32) *UpdateOpenSourceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOpenSourceAccountResponseBody) SetMessage(v string) *UpdateOpenSourceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOpenSourceAccountResponseBody) SetRequestId(v string) *UpdateOpenSourceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOpenSourceAccountResponseBody) SetSuccess(v bool) *UpdateOpenSourceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOpenSourceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
