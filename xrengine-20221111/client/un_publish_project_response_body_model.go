// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnPublishProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnPublishProjectResponseBody
	GetCode() *string
	SetErrorName(v string) *UnPublishProjectResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *UnPublishProjectResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *UnPublishProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnPublishProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnPublishProjectResponseBody
	GetSuccess() *bool
}

type UnPublishProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnPublishProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnPublishProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UnPublishProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnPublishProjectResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *UnPublishProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UnPublishProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnPublishProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnPublishProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnPublishProjectResponseBody) SetCode(v string) *UnPublishProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UnPublishProjectResponseBody) SetErrorName(v string) *UnPublishProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *UnPublishProjectResponseBody) SetHttpCode(v int32) *UnPublishProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UnPublishProjectResponseBody) SetMessage(v string) *UnPublishProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UnPublishProjectResponseBody) SetRequestId(v string) *UnPublishProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnPublishProjectResponseBody) SetSuccess(v bool) *UnPublishProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UnPublishProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
