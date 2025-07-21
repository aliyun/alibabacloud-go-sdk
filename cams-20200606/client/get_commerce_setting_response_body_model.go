// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommerceSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCommerceSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCommerceSettingResponseBody
	GetCode() *string
	SetData(v *GetCommerceSettingResponseBodyData) *GetCommerceSettingResponseBody
	GetData() *GetCommerceSettingResponseBodyData
	SetMessage(v string) *GetCommerceSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCommerceSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCommerceSettingResponseBody
	GetSuccess() *bool
}

type GetCommerceSettingResponseBody struct {
	// Access denied for detailed information.
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
	// The returned data.
	Data *GetCommerceSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCommerceSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommerceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCommerceSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCommerceSettingResponseBody) GetData() *GetCommerceSettingResponseBodyData {
	return s.Data
}

func (s *GetCommerceSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCommerceSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommerceSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCommerceSettingResponseBody) SetAccessDeniedDetail(v string) *GetCommerceSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetCode(v string) *GetCommerceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetData(v *GetCommerceSettingResponseBodyData) *GetCommerceSettingResponseBody {
	s.Data = v
	return s
}

func (s *GetCommerceSettingResponseBody) SetMessage(v string) *GetCommerceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetRequestId(v string) *GetCommerceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetSuccess(v bool) *GetCommerceSettingResponseBody {
	s.Success = &v
	return s
}

func (s *GetCommerceSettingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCommerceSettingResponseBodyData struct {
	// Indicates whether the shopping cart button is displayed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	CartEnable *bool `json:"CartEnable,omitempty" xml:"CartEnable,omitempty"`
	// Indicates whether the catalog button is displayed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	CatalogVisible *bool `json:"CatalogVisible,omitempty" xml:"CatalogVisible,omitempty"`
}

func (s GetCommerceSettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCommerceSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponseBodyData) GetCartEnable() *bool {
	return s.CartEnable
}

func (s *GetCommerceSettingResponseBodyData) GetCatalogVisible() *bool {
	return s.CatalogVisible
}

func (s *GetCommerceSettingResponseBodyData) SetCartEnable(v bool) *GetCommerceSettingResponseBodyData {
	s.CartEnable = &v
	return s
}

func (s *GetCommerceSettingResponseBodyData) SetCatalogVisible(v bool) *GetCommerceSettingResponseBodyData {
	s.CatalogVisible = &v
	return s
}

func (s *GetCommerceSettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
