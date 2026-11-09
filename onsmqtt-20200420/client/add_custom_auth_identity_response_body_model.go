// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddCustomAuthIdentityResponseBody
	GetCode() *int32
	SetMessage(v string) *AddCustomAuthIdentityResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomAuthIdentityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCustomAuthIdentityResponseBody
	GetSuccess() *bool
}

type AddCustomAuthIdentityResponseBody struct {
	// The status code. A value of 200 indicates success.
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
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCustomAuthIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomAuthIdentityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddCustomAuthIdentityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomAuthIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomAuthIdentityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCustomAuthIdentityResponseBody) SetCode(v int32) *AddCustomAuthIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *AddCustomAuthIdentityResponseBody) SetMessage(v string) *AddCustomAuthIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomAuthIdentityResponseBody) SetRequestId(v string) *AddCustomAuthIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomAuthIdentityResponseBody) SetSuccess(v bool) *AddCustomAuthIdentityResponseBody {
	s.Success = &v
	return s
}

func (s *AddCustomAuthIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
