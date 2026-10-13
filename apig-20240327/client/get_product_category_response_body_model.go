// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProductCategoryResponseBody
	GetCode() *string
	SetData(v *GetProductCategoryResponseBodyData) *GetProductCategoryResponseBody
	GetData() *GetProductCategoryResponseBodyData
	SetMessage(v string) *GetProductCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProductCategoryResponseBody
	GetRequestId() *string
}

type GetProductCategoryResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetProductCategoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetProductCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProductCategoryResponseBody) GetData() *GetProductCategoryResponseBodyData {
	return s.Data
}

func (s *GetProductCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProductCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductCategoryResponseBody) SetCode(v string) *GetProductCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductCategoryResponseBody) SetData(v *GetProductCategoryResponseBodyData) *GetProductCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetProductCategoryResponseBody) SetMessage(v string) *GetProductCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductCategoryResponseBody) SetRequestId(v string) *GetProductCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductCategoryResponseBodyData struct {
	// The ID of the product category. This ID can be used as the categoryId path parameter for GetProductCategory, UpdateProductCategory, and DeleteProductCategory, or as an element in the categories array for SetProductCategories and CreateProduct.
	//
	// example:
	//
	// category-xxxxxxxx
	CategoryId *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// The description of the product category.
	//
	// example:
	//
	// Product category containing AI-related capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The icon of the product category. It consists of two fields: icon type (type) and icon value (value).
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The name of the product category.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetProductCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProductCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductCategoryResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetProductCategoryResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetProductCategoryResponseBodyData) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *GetProductCategoryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetProductCategoryResponseBodyData) SetCategoryId(v string) *GetProductCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetProductCategoryResponseBodyData) SetDescription(v string) *GetProductCategoryResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductCategoryResponseBodyData) SetIcon(v *HiMarketIcon) *GetProductCategoryResponseBodyData {
	s.Icon = v
	return s
}

func (s *GetProductCategoryResponseBodyData) SetName(v string) *GetProductCategoryResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductCategoryResponseBodyData) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}
