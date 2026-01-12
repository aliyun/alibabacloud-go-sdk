// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLocationServiceResponseBody
	GetCode() *string
	SetData(v bool) *CreateLocationServiceResponseBody
	GetData() *bool
	SetErrorName(v string) *CreateLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *CreateLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *CreateLocationServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLocationServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLocationServiceResponseBody
	GetSuccess() *bool
}

type CreateLocationServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLocationServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *CreateLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLocationServiceResponseBody) SetCode(v string) *CreateLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetData(v bool) *CreateLocationServiceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetErrorName(v string) *CreateLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetHttpCode(v int32) *CreateLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetMessage(v string) *CreateLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetRequestId(v string) *CreateLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocationServiceResponseBody) SetSuccess(v bool) *CreateLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLocationServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
