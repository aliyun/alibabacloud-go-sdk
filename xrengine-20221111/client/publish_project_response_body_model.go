// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishProjectResponseBody
	GetCode() *string
	SetErrorName(v string) *PublishProjectResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *PublishProjectResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *PublishProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishProjectResponseBody
	GetSuccess() *bool
}

type PublishProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PublishProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishProjectResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *PublishProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *PublishProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishProjectResponseBody) SetCode(v string) *PublishProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PublishProjectResponseBody) SetErrorName(v string) *PublishProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PublishProjectResponseBody) SetHttpCode(v int32) *PublishProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PublishProjectResponseBody) SetMessage(v string) *PublishProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PublishProjectResponseBody) SetRequestId(v string) *PublishProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishProjectResponseBody) SetSuccess(v bool) *PublishProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PublishProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
