// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePolicyBindingResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePolicyBindingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePolicyBindingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePolicyBindingResponseBody
	GetSuccess() *bool
}

type DeletePolicyBindingResponseBody struct {
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
	// 3E961A5E-C5C6-566D-BFC3-0362A6A52EBA
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

func (s DeletePolicyBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyBindingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePolicyBindingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolicyBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyBindingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePolicyBindingResponseBody) SetCode(v string) *DeletePolicyBindingResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyBindingResponseBody) SetMessage(v string) *DeletePolicyBindingResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyBindingResponseBody) SetRequestId(v string) *DeletePolicyBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyBindingResponseBody) SetSuccess(v bool) *DeletePolicyBindingResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePolicyBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
