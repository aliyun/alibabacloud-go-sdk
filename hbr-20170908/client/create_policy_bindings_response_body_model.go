// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePolicyBindingsResponseBody
	GetCode() *string
	SetMessage(v string) *CreatePolicyBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePolicyBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePolicyBindingsResponseBody
	GetSuccess() *bool
}

type CreatePolicyBindingsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 82CC5B6C-72F7-5D39-92F6-67887DF9AD46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePolicyBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePolicyBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolicyBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePolicyBindingsResponseBody) SetCode(v string) *CreatePolicyBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyBindingsResponseBody) SetMessage(v string) *CreatePolicyBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyBindingsResponseBody) SetRequestId(v string) *CreatePolicyBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyBindingsResponseBody) SetSuccess(v bool) *CreatePolicyBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePolicyBindingsResponseBody) Validate() error {
	return dara.Validate(s)
}
