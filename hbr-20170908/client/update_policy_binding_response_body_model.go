// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePolicyBindingResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePolicyBindingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePolicyBindingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePolicyBindingResponseBody
	GetSuccess() *bool
}

type UpdatePolicyBindingResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6F24C46-54B9-519B-9AB8-A8988D705E67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePolicyBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePolicyBindingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolicyBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyBindingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePolicyBindingResponseBody) SetCode(v string) *UpdatePolicyBindingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyBindingResponseBody) SetMessage(v string) *UpdatePolicyBindingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolicyBindingResponseBody) SetRequestId(v string) *UpdatePolicyBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyBindingResponseBody) SetSuccess(v bool) *UpdatePolicyBindingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePolicyBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
