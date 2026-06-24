// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSourceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateOpenSourceAccountResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateOpenSourceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOpenSourceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOpenSourceAccountResponseBody
	GetSuccess() *bool
}

type CreateOpenSourceAccountResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEBA5E0C-50D0-4FA6-A794-4901E5465***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOpenSourceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSourceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateOpenSourceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOpenSourceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSourceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOpenSourceAccountResponseBody) SetCode(v int32) *CreateOpenSourceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOpenSourceAccountResponseBody) SetMessage(v string) *CreateOpenSourceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOpenSourceAccountResponseBody) SetRequestId(v string) *CreateOpenSourceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSourceAccountResponseBody) SetSuccess(v bool) *CreateOpenSourceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOpenSourceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
