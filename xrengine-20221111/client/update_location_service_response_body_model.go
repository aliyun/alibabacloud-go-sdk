// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLocationServiceResponseBody
	GetCode() *string
	SetData(v bool) *UpdateLocationServiceResponseBody
	GetData() *bool
	SetErrorName(v string) *UpdateLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *UpdateLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *UpdateLocationServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLocationServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLocationServiceResponseBody
	GetSuccess() *bool
}

type UpdateLocationServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLocationServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *UpdateLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLocationServiceResponseBody) SetCode(v string) *UpdateLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetData(v bool) *UpdateLocationServiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetErrorName(v string) *UpdateLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetHttpCode(v int32) *UpdateLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetMessage(v string) *UpdateLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetRequestId(v string) *UpdateLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) SetSuccess(v bool) *UpdateLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLocationServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
