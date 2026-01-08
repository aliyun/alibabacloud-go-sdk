// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhatsappConversionApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateWhatsappConversionApiResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateWhatsappConversionApiResponseBody
	GetCode() *string
	SetMessage(v string) *CreateWhatsappConversionApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWhatsappConversionApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWhatsappConversionApiResponseBody
	GetSuccess() *bool
}

type CreateWhatsappConversionApiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWhatsappConversionApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWhatsappConversionApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhatsappConversionApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateWhatsappConversionApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWhatsappConversionApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWhatsappConversionApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWhatsappConversionApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWhatsappConversionApiResponseBody) SetAccessDeniedDetail(v string) *CreateWhatsappConversionApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateWhatsappConversionApiResponseBody) SetCode(v string) *CreateWhatsappConversionApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWhatsappConversionApiResponseBody) SetMessage(v string) *CreateWhatsappConversionApiResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWhatsappConversionApiResponseBody) SetRequestId(v string) *CreateWhatsappConversionApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhatsappConversionApiResponseBody) SetSuccess(v bool) *CreateWhatsappConversionApiResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWhatsappConversionApiResponseBody) Validate() error {
	return dara.Validate(s)
}
