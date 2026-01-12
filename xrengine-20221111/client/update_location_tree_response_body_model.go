// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLocationTreeResponseBody
	GetCode() *string
	SetData(v bool) *UpdateLocationTreeResponseBody
	GetData() *bool
	SetErrorName(v string) *UpdateLocationTreeResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *UpdateLocationTreeResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *UpdateLocationTreeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLocationTreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLocationTreeResponseBody
	GetSuccess() *bool
}

type UpdateLocationTreeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLocationTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationTreeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocationTreeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLocationTreeResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateLocationTreeResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *UpdateLocationTreeResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateLocationTreeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLocationTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLocationTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLocationTreeResponseBody) SetCode(v string) *UpdateLocationTreeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetData(v bool) *UpdateLocationTreeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetErrorName(v string) *UpdateLocationTreeResponseBody {
	s.ErrorName = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetHttpCode(v int32) *UpdateLocationTreeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetMessage(v string) *UpdateLocationTreeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetRequestId(v string) *UpdateLocationTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) SetSuccess(v bool) *UpdateLocationTreeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLocationTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
