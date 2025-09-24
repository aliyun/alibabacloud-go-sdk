// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryProductListResponseBody
	GetCode() *string
	SetData(v *QueryProductListResponseBodyData) *QueryProductListResponseBody
	GetData() *QueryProductListResponseBodyData
	SetMessage(v string) *QueryProductListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryProductListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProductListResponseBody
	GetSuccess() *bool
}

type QueryProductListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about all Alibaba Cloud services.
	Data *QueryProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// This API is not applicable for caller.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 94858229-2758-4663-A7D0-99490D541F15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryProductListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryProductListResponseBody) GetData() *QueryProductListResponseBodyData {
	return s.Data
}

func (s *QueryProductListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryProductListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProductListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProductListResponseBody) SetCode(v string) *QueryProductListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProductListResponseBody) SetData(v *QueryProductListResponseBodyData) *QueryProductListResponseBody {
	s.Data = v
	return s
}

func (s *QueryProductListResponseBody) SetMessage(v string) *QueryProductListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProductListResponseBody) SetRequestId(v string) *QueryProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductListResponseBody) SetSuccess(v bool) *QueryProductListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProductListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryProductListResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service definitions.
	ProductList *QueryProductListResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Struct"`
	// The total number of services.
	//
	// example:
	//
	// 449
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryProductListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryProductListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryProductListResponseBodyData) GetProductList() *QueryProductListResponseBodyDataProductList {
	return s.ProductList
}

func (s *QueryProductListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryProductListResponseBodyData) SetPageNum(v int32) *QueryProductListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryProductListResponseBodyData) SetPageSize(v int32) *QueryProductListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryProductListResponseBodyData) SetProductList(v *QueryProductListResponseBodyDataProductList) *QueryProductListResponseBodyData {
	s.ProductList = v
	return s
}

func (s *QueryProductListResponseBodyData) SetTotalCount(v int32) *QueryProductListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryProductListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryProductListResponseBodyDataProductList struct {
	Product []*QueryProductListResponseBodyDataProductListProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
}

func (s QueryProductListResponseBodyDataProductList) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyDataProductList) GetProduct() []*QueryProductListResponseBodyDataProductListProduct {
	return s.Product
}

func (s *QueryProductListResponseBodyDataProductList) SetProduct(v []*QueryProductListResponseBodyDataProductListProduct) *QueryProductListResponseBodyDataProductList {
	s.Product = v
	return s
}

func (s *QueryProductListResponseBodyDataProductList) Validate() error {
	return dara.Validate(s)
}

type QueryProductListResponseBodyDataProductListProduct struct {
	// The code of the service.
	//
	// example:
	//
	// cdn
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// CDN (Pay-as-you-go)
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// CDN
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryProductListResponseBodyDataProductListProduct) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListResponseBodyDataProductListProduct) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyDataProductListProduct) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryProductListResponseBodyDataProductListProduct) GetProductName() *string {
	return s.ProductName
}

func (s *QueryProductListResponseBodyDataProductListProduct) GetProductType() *string {
	return s.ProductType
}

func (s *QueryProductListResponseBodyDataProductListProduct) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductCode(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductCode = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductName(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductName = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductType(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductType = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetSubscriptionType(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.SubscriptionType = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) Validate() error {
	return dara.Validate(s)
}
