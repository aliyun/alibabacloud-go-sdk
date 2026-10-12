// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProductCategoryResponseBody
	GetCode() *string
	SetData(v *CreateProductCategoryResponseBodyData) *CreateProductCategoryResponseBody
	GetData() *CreateProductCategoryResponseBodyData
	SetMessage(v string) *CreateProductCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProductCategoryResponseBody
	GetRequestId() *string
}

type CreateProductCategoryResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateProductCategoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateProductCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProductCategoryResponseBody) GetData() *CreateProductCategoryResponseBodyData {
	return s.Data
}

func (s *CreateProductCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProductCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductCategoryResponseBody) SetCode(v string) *CreateProductCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductCategoryResponseBody) SetData(v *CreateProductCategoryResponseBodyData) *CreateProductCategoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductCategoryResponseBody) SetMessage(v string) *CreateProductCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductCategoryResponseBody) SetRequestId(v string) *CreateProductCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProductCategoryResponseBodyData struct {
	// The ID of the newly created product category. It can be used as the categoryId path parameter in GetProductCategory, UpdateProductCategory, and DeleteProductCategory, or as an element in the category ID array for SetProductCategories.
	//
	// example:
	//
	// category-xxxxxxxx
	CategoryId *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// The name of the product category.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProductCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProductCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductCategoryResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *CreateProductCategoryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateProductCategoryResponseBodyData) SetCategoryId(v string) *CreateProductCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *CreateProductCategoryResponseBodyData) SetName(v string) *CreateProductCategoryResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateProductCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
