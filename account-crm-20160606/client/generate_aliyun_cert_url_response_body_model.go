// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAliyunCertUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateAliyunCertUrlResponseBody
	GetCode() *string
	SetData(v string) *GenerateAliyunCertUrlResponseBody
	GetData() *string
	SetHttpCode(v string) *GenerateAliyunCertUrlResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GenerateAliyunCertUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateAliyunCertUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateAliyunCertUrlResponseBody
	GetSuccess() *bool
}

type GenerateAliyunCertUrlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateAliyunCertUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAliyunCertUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAliyunCertUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateAliyunCertUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateAliyunCertUrlResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GenerateAliyunCertUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateAliyunCertUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAliyunCertUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateAliyunCertUrlResponseBody) SetCode(v string) *GenerateAliyunCertUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) SetData(v string) *GenerateAliyunCertUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) SetHttpCode(v string) *GenerateAliyunCertUrlResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) SetMessage(v string) *GenerateAliyunCertUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) SetRequestId(v string) *GenerateAliyunCertUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) SetSuccess(v bool) *GenerateAliyunCertUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateAliyunCertUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
