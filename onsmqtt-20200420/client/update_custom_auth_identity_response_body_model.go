// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateCustomAuthIdentityResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateCustomAuthIdentityResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomAuthIdentityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomAuthIdentityResponseBody
	GetSuccess() *bool
}

type UpdateCustomAuthIdentityResponseBody struct {
	// The result code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F00084A-7F07-4B15-BADA-8903A4FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomAuthIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthIdentityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateCustomAuthIdentityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomAuthIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomAuthIdentityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomAuthIdentityResponseBody) SetCode(v int32) *UpdateCustomAuthIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomAuthIdentityResponseBody) SetMessage(v string) *UpdateCustomAuthIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomAuthIdentityResponseBody) SetRequestId(v string) *UpdateCustomAuthIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomAuthIdentityResponseBody) SetSuccess(v bool) *UpdateCustomAuthIdentityResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomAuthIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
