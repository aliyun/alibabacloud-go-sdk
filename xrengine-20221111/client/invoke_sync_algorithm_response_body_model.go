// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSyncAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeSyncAlgorithmResponseBody
	GetCode() *string
	SetData(v interface{}) *InvokeSyncAlgorithmResponseBody
	GetData() interface{}
	SetErrorName(v string) *InvokeSyncAlgorithmResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *InvokeSyncAlgorithmResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *InvokeSyncAlgorithmResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeSyncAlgorithmResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokeSyncAlgorithmResponseBody
	GetSuccess() *bool
}

type InvokeSyncAlgorithmResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string     `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32      `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvokeSyncAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeSyncAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeSyncAlgorithmResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeSyncAlgorithmResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InvokeSyncAlgorithmResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *InvokeSyncAlgorithmResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *InvokeSyncAlgorithmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeSyncAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeSyncAlgorithmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokeSyncAlgorithmResponseBody) SetCode(v string) *InvokeSyncAlgorithmResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetData(v interface{}) *InvokeSyncAlgorithmResponseBody {
	s.Data = v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetErrorName(v string) *InvokeSyncAlgorithmResponseBody {
	s.ErrorName = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetHttpCode(v int32) *InvokeSyncAlgorithmResponseBody {
	s.HttpCode = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetMessage(v string) *InvokeSyncAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetRequestId(v string) *InvokeSyncAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) SetSuccess(v bool) *InvokeSyncAlgorithmResponseBody {
	s.Success = &v
	return s
}

func (s *InvokeSyncAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
