// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterUserResponseBody
	GetCode() *string
	SetErrorName(v string) *RegisterUserResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *RegisterUserResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *RegisterUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *RegisterUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterUserResponseBody
	GetSuccess() *bool
}

type RegisterUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterUserResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterUserResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *RegisterUserResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *RegisterUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterUserResponseBody) SetCode(v string) *RegisterUserResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterUserResponseBody) SetErrorName(v string) *RegisterUserResponseBody {
	s.ErrorName = &v
	return s
}

func (s *RegisterUserResponseBody) SetHttpCode(v int32) *RegisterUserResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RegisterUserResponseBody) SetMessage(v string) *RegisterUserResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterUserResponseBody) SetRequestId(v string) *RegisterUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUserResponseBody) SetSuccess(v bool) *RegisterUserResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterUserResponseBody) Validate() error {
	return dara.Validate(s)
}
