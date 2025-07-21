// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConnectionCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetWhatsappConnectionCatalogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetWhatsappConnectionCatalogResponseBody
	GetCode() *string
	SetMessage(v string) *GetWhatsappConnectionCatalogResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *GetWhatsappConnectionCatalogResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *GetWhatsappConnectionCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWhatsappConnectionCatalogResponseBody
	GetSuccess() *bool
}

type GetWhatsappConnectionCatalogResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {"id":"200292992"}
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWhatsappConnectionCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConnectionCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWhatsappConnectionCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetAccessDeniedDetail(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetCode(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetMessage(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetModel(v map[string]interface{}) *GetWhatsappConnectionCatalogResponseBody {
	s.Model = v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetRequestId(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetSuccess(v bool) *GetWhatsappConnectionCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}
