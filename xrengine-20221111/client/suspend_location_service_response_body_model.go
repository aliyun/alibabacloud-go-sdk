// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendLocationServiceResponseBody
	GetCode() *string
	SetData(v bool) *SuspendLocationServiceResponseBody
	GetData() *bool
	SetErrorName(v string) *SuspendLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *SuspendLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *SuspendLocationServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendLocationServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendLocationServiceResponseBody
	GetSuccess() *bool
}

type SuspendLocationServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendLocationServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *SuspendLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *SuspendLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SuspendLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendLocationServiceResponseBody) SetCode(v string) *SuspendLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetData(v bool) *SuspendLocationServiceResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetErrorName(v string) *SuspendLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetHttpCode(v int32) *SuspendLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetMessage(v string) *SuspendLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetRequestId(v string) *SuspendLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) SetSuccess(v bool) *SuspendLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendLocationServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
