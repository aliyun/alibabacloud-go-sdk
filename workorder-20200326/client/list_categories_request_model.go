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
	SetProductCode(v string) *ListCategoriesRequest
	GetProductCode() *string
}

type ListCategoriesRequest struct {
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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

func (s *ListCategoriesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListCategoriesRequest) SetLanguage(v string) *ListCategoriesRequest {
	s.Language = &v
	return s
}

func (s *ListCategoriesRequest) SetProductCode(v string) *ListCategoriesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
