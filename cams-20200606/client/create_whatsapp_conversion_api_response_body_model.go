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
	// The details about the access denial.
	//
	// example:
	//
	// None.
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Sample value.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// Sample value.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
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
