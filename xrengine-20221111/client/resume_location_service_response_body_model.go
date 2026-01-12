// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeLocationServiceResponseBody
	GetCode() *string
	SetData(v bool) *ResumeLocationServiceResponseBody
	GetData() *bool
	SetErrorName(v string) *ResumeLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ResumeLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ResumeLocationServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeLocationServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeLocationServiceResponseBody
	GetSuccess() *bool
}

type ResumeLocationServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeLocationServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *ResumeLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ResumeLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ResumeLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeLocationServiceResponseBody) SetCode(v string) *ResumeLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetData(v bool) *ResumeLocationServiceResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetErrorName(v string) *ResumeLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetHttpCode(v int32) *ResumeLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetMessage(v string) *ResumeLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetRequestId(v string) *ResumeLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) SetSuccess(v bool) *ResumeLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeLocationServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
