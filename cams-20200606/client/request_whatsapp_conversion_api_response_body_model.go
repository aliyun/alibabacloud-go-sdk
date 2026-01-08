// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestWhatsappConversionApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RequestWhatsappConversionApiResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RequestWhatsappConversionApiResponseBody
	GetCode() *string
	SetMessage(v string) *RequestWhatsappConversionApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *RequestWhatsappConversionApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RequestWhatsappConversionApiResponseBody
	GetSuccess() *bool
}

type RequestWhatsappConversionApiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestWhatsappConversionApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RequestWhatsappConversionApiResponseBody) GoString() string {
	return s.String()
}

func (s *RequestWhatsappConversionApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RequestWhatsappConversionApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *RequestWhatsappConversionApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RequestWhatsappConversionApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RequestWhatsappConversionApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RequestWhatsappConversionApiResponseBody) SetAccessDeniedDetail(v string) *RequestWhatsappConversionApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RequestWhatsappConversionApiResponseBody) SetCode(v string) *RequestWhatsappConversionApiResponseBody {
	s.Code = &v
	return s
}

func (s *RequestWhatsappConversionApiResponseBody) SetMessage(v string) *RequestWhatsappConversionApiResponseBody {
	s.Message = &v
	return s
}

func (s *RequestWhatsappConversionApiResponseBody) SetRequestId(v string) *RequestWhatsappConversionApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestWhatsappConversionApiResponseBody) SetSuccess(v bool) *RequestWhatsappConversionApiResponseBody {
	s.Success = &v
	return s
}

func (s *RequestWhatsappConversionApiResponseBody) Validate() error {
	return dara.Validate(s)
}
