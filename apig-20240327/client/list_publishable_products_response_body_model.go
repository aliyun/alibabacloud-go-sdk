// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishableProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPublishableProductsResponseBody
	GetCode() *string
	SetData(v *ListPublishableProductsResponseBodyData) *ListPublishableProductsResponseBody
	GetData() *ListPublishableProductsResponseBodyData
	SetMessage(v string) *ListPublishableProductsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPublishableProductsResponseBody
	GetRequestId() *string
}

type ListPublishableProductsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListPublishableProductsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPublishableProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishableProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishableProductsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPublishableProductsResponseBody) GetData() *ListPublishableProductsResponseBodyData {
	return s.Data
}

func (s *ListPublishableProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublishableProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishableProductsResponseBody) SetCode(v string) *ListPublishableProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishableProductsResponseBody) SetData(v *ListPublishableProductsResponseBodyData) *ListPublishableProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishableProductsResponseBody) SetMessage(v string) *ListPublishableProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishableProductsResponseBody) SetRequestId(v string) *ListPublishableProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishableProductsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishableProductsResponseBodyData struct {
	// The list of API products that can be published to the portal. Each element in the list is named productInfo and contains the information about an API product.
	Content []*ListPublishableProductsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPublishableProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPublishableProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishableProductsResponseBodyData) GetContent() []*ListPublishableProductsResponseBodyDataContent {
	return s.Content
}

func (s *ListPublishableProductsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishableProductsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishableProductsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishableProductsResponseBodyData) SetContent(v []*ListPublishableProductsResponseBodyDataContent) *ListPublishableProductsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListPublishableProductsResponseBodyData) SetPageNumber(v int32) *ListPublishableProductsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPublishableProductsResponseBodyData) SetPageSize(v int32) *ListPublishableProductsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishableProductsResponseBodyData) SetTotalCount(v int32) *ListPublishableProductsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPublishableProductsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishableProductsResponseBodyDataContent struct {
	// The description of the API product. This corresponds to the Description property of the Product resource.
	//
	// example:
	//
	// An API product that provides weather query capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the API product. This corresponds to the Name property of the Product resource.
	//
	// example:
	//
	// weather-query-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the API product. This is the primary ID of the Product resource. You can use this ID in the PublishProduct operation at /himarket/v1/products/{productId}/publications to publish the product to the portal. You can also use it to locate the product in operations such as GetProduct and ListProductPublications.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s ListPublishableProductsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListPublishableProductsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListPublishableProductsResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListPublishableProductsResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListPublishableProductsResponseBodyDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *ListPublishableProductsResponseBodyDataContent) SetDescription(v string) *ListPublishableProductsResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListPublishableProductsResponseBodyDataContent) SetName(v string) *ListPublishableProductsResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListPublishableProductsResponseBodyDataContent) SetProductId(v string) *ListPublishableProductsResponseBodyDataContent {
	s.ProductId = &v
	return s
}

func (s *ListPublishableProductsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
