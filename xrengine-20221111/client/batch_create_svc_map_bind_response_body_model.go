// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateSvcMapBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateSvcMapBindResponseBody
	GetCode() *string
	SetData(v bool) *BatchCreateSvcMapBindResponseBody
	GetData() *bool
	SetErrorName(v string) *BatchCreateSvcMapBindResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *BatchCreateSvcMapBindResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *BatchCreateSvcMapBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateSvcMapBindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateSvcMapBindResponseBody
	GetSuccess() *bool
}

type BatchCreateSvcMapBindResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateSvcMapBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateSvcMapBindResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateSvcMapBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateSvcMapBindResponseBody) GetData() *bool {
	return s.Data
}

func (s *BatchCreateSvcMapBindResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *BatchCreateSvcMapBindResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *BatchCreateSvcMapBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateSvcMapBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateSvcMapBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateSvcMapBindResponseBody) SetCode(v string) *BatchCreateSvcMapBindResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetData(v bool) *BatchCreateSvcMapBindResponseBody {
	s.Data = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetErrorName(v string) *BatchCreateSvcMapBindResponseBody {
	s.ErrorName = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetHttpCode(v int32) *BatchCreateSvcMapBindResponseBody {
	s.HttpCode = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetMessage(v string) *BatchCreateSvcMapBindResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetRequestId(v string) *BatchCreateSvcMapBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) SetSuccess(v bool) *BatchCreateSvcMapBindResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateSvcMapBindResponseBody) Validate() error {
	return dara.Validate(s)
}
