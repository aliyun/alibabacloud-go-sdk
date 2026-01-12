// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSvcMapBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeleteSvcMapBindResponseBody
	GetCode() *string
	SetData(v bool) *BatchDeleteSvcMapBindResponseBody
	GetData() *bool
	SetErrorName(v string) *BatchDeleteSvcMapBindResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *BatchDeleteSvcMapBindResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *BatchDeleteSvcMapBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteSvcMapBindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteSvcMapBindResponseBody
	GetSuccess() *bool
}

type BatchDeleteSvcMapBindResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteSvcMapBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSvcMapBindResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteSvcMapBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeleteSvcMapBindResponseBody) GetData() *bool {
	return s.Data
}

func (s *BatchDeleteSvcMapBindResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *BatchDeleteSvcMapBindResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *BatchDeleteSvcMapBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteSvcMapBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteSvcMapBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteSvcMapBindResponseBody) SetCode(v string) *BatchDeleteSvcMapBindResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetData(v bool) *BatchDeleteSvcMapBindResponseBody {
	s.Data = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetErrorName(v string) *BatchDeleteSvcMapBindResponseBody {
	s.ErrorName = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetHttpCode(v int32) *BatchDeleteSvcMapBindResponseBody {
	s.HttpCode = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetMessage(v string) *BatchDeleteSvcMapBindResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetRequestId(v string) *BatchDeleteSvcMapBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) SetSuccess(v bool) *BatchDeleteSvcMapBindResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponseBody) Validate() error {
	return dara.Validate(s)
}
