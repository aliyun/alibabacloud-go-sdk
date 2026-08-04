// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountProfileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAccountProfileInfoResponseBody
	GetCode() *string
	SetMessage(v string) *CreateAccountProfileInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAccountProfileInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAccountProfileInfoResponseBody
	GetSuccess() *bool
}

type CreateAccountProfileInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountProfileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountProfileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountProfileInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAccountProfileInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAccountProfileInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountProfileInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAccountProfileInfoResponseBody) SetCode(v string) *CreateAccountProfileInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccountProfileInfoResponseBody) SetMessage(v string) *CreateAccountProfileInfoResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountProfileInfoResponseBody) SetRequestId(v string) *CreateAccountProfileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountProfileInfoResponseBody) SetSuccess(v bool) *CreateAccountProfileInfoResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAccountProfileInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
