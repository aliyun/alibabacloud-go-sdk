// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProductCategoriesResponseBody
	GetCode() *string
	SetData(v *ListProductCategoriesResponseBodyData) *ListProductCategoriesResponseBody
	GetData() *ListProductCategoriesResponseBodyData
	SetMessage(v string) *ListProductCategoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProductCategoriesResponseBody
	GetRequestId() *string
}

type ListProductCategoriesResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListProductCategoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ListProductCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductCategoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductCategoriesResponseBody) GetData() *ListProductCategoriesResponseBodyData {
	return s.Data
}

func (s *ListProductCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductCategoriesResponseBody) SetCode(v string) *ListProductCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductCategoriesResponseBody) SetData(v *ListProductCategoriesResponseBodyData) *ListProductCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListProductCategoriesResponseBody) SetMessage(v string) *ListProductCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductCategoriesResponseBody) SetRequestId(v string) *ListProductCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductCategoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductCategoriesResponseBodyData struct {
	// The list of product categories. Each element in the list is named category and contains the information about a product category.
	Content []*ListProductCategoriesResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
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
	// The total count.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductCategoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductCategoriesResponseBodyData) GetContent() []*ListProductCategoriesResponseBodyDataContent {
	return s.Content
}

func (s *ListProductCategoriesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductCategoriesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductCategoriesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProductCategoriesResponseBodyData) SetContent(v []*ListProductCategoriesResponseBodyDataContent) *ListProductCategoriesResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProductCategoriesResponseBodyData) SetPageNumber(v int32) *ListProductCategoriesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListProductCategoriesResponseBodyData) SetPageSize(v int32) *ListProductCategoriesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductCategoriesResponseBodyData) SetTotalCount(v int64) *ListProductCategoriesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProductCategoriesResponseBodyData) Validate() error {
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

type ListProductCategoriesResponseBodyDataContent struct {
	// The product category ID. You can use this ID as the categoryId path parameter for GetProductCategory, UpdateProductCategory, and DeleteProductCategory. You can also use it as an element in the categories parameter, which is a set of category IDs, for SetProductCategories and CreateProduct.
	//
	// example:
	//
	// category-xxxxxxxx
	CategoryId *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// The product category description.
	//
	// example:
	//
	// Product category for AI-related capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The product category icon, which consists of the icon type (type) and icon information (value) fields.
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The product category name.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProductCategoriesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListProductCategoriesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListProductCategoriesResponseBodyDataContent) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListProductCategoriesResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListProductCategoriesResponseBodyDataContent) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *ListProductCategoriesResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListProductCategoriesResponseBodyDataContent) SetCategoryId(v string) *ListProductCategoriesResponseBodyDataContent {
	s.CategoryId = &v
	return s
}

func (s *ListProductCategoriesResponseBodyDataContent) SetDescription(v string) *ListProductCategoriesResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListProductCategoriesResponseBodyDataContent) SetIcon(v *HiMarketIcon) *ListProductCategoriesResponseBodyDataContent {
	s.Icon = v
	return s
}

func (s *ListProductCategoriesResponseBodyDataContent) SetName(v string) *ListProductCategoriesResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListProductCategoriesResponseBodyDataContent) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}
