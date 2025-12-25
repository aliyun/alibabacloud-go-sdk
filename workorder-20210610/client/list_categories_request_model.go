// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListCategoriesRequest
	GetLanguage() *string
	SetName(v string) *ListCategoriesRequest
	GetName() *string
	SetProductId(v int64) *ListCategoriesRequest
	GetProductId() *int64
}

type ListCategoriesRequest struct {
	// Multi-language, support, Chinese, English. Value definition: zh: Chinese, en: English.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the classification question. Fuzzy search is supported.
	//
	// example:
	//
	// ecs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the product. You can call the ListProducts operation to obtain the product ID. The ProductId parameter is the ID of an Alibaba Cloud product. Multiple Categories are displayed for each product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18550
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListCategoriesRequest) GetName() *string {
	return s.Name
}

func (s *ListCategoriesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListCategoriesRequest) SetLanguage(v string) *ListCategoriesRequest {
	s.Language = &v
	return s
}

func (s *ListCategoriesRequest) SetName(v string) *ListCategoriesRequest {
	s.Name = &v
	return s
}

func (s *ListCategoriesRequest) SetProductId(v int64) *ListCategoriesRequest {
	s.ProductId = &v
	return s
}

func (s *ListCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
