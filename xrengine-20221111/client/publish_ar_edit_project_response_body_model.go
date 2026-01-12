// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishArEditProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishArEditProjectResponseBody
	GetCode() *string
	SetData(v bool) *PublishArEditProjectResponseBody
	GetData() *bool
	SetErrorName(v string) *PublishArEditProjectResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *PublishArEditProjectResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *PublishArEditProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishArEditProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishArEditProjectResponseBody
	GetSuccess() *bool
}

type PublishArEditProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishArEditProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishArEditProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PublishArEditProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishArEditProjectResponseBody) GetData() *bool {
	return s.Data
}

func (s *PublishArEditProjectResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *PublishArEditProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *PublishArEditProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishArEditProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishArEditProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishArEditProjectResponseBody) SetCode(v string) *PublishArEditProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetData(v bool) *PublishArEditProjectResponseBody {
	s.Data = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetErrorName(v string) *PublishArEditProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetHttpCode(v int32) *PublishArEditProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetMessage(v string) *PublishArEditProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetRequestId(v string) *PublishArEditProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishArEditProjectResponseBody) SetSuccess(v bool) *PublishArEditProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PublishArEditProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
