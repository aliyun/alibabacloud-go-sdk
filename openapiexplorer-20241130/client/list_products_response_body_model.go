// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody
	GetProducts() []*ListProductsResponseBodyProducts
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
}

type ListProductsResponseBody struct {
	Products []*ListProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetProducts() []*ListProductsResponseBodyProducts {
	return s.Products
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyProducts struct {
	Category2Name *string `json:"category2Name,omitempty" xml:"category2Name,omitempty"`
	CategoryName  *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// example:
	//
	// Ecs
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 2014-05-26
	DefaultVersion *string `json:"defaultVersion,omitempty" xml:"defaultVersion,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Group          *string `json:"group,omitempty" xml:"group,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ECS
	ShortName *string `json:"shortName,omitempty" xml:"shortName,omitempty"`
	// example:
	//
	// RPC
	Style    *string   `json:"style,omitempty" xml:"style,omitempty"`
	Versions []*string `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) GetCategory2Name() *string {
	return s.Category2Name
}

func (s *ListProductsResponseBodyProducts) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListProductsResponseBodyProducts) GetCode() *string {
	return s.Code
}

func (s *ListProductsResponseBodyProducts) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListProductsResponseBodyProducts) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyProducts) GetGroup() *string {
	return s.Group
}

func (s *ListProductsResponseBodyProducts) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyProducts) GetShortName() *string {
	return s.ShortName
}

func (s *ListProductsResponseBodyProducts) GetStyle() *string {
	return s.Style
}

func (s *ListProductsResponseBodyProducts) GetVersions() []*string {
	return s.Versions
}

func (s *ListProductsResponseBodyProducts) SetCategory2Name(v string) *ListProductsResponseBodyProducts {
	s.Category2Name = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetCategoryName(v string) *ListProductsResponseBodyProducts {
	s.CategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetCode(v string) *ListProductsResponseBodyProducts {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetDefaultVersion(v string) *ListProductsResponseBodyProducts {
	s.DefaultVersion = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetDescription(v string) *ListProductsResponseBodyProducts {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetGroup(v string) *ListProductsResponseBodyProducts {
	s.Group = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetName(v string) *ListProductsResponseBodyProducts {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetShortName(v string) *ListProductsResponseBodyProducts {
	s.ShortName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetStyle(v string) *ListProductsResponseBodyProducts {
	s.Style = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetVersions(v []*string) *ListProductsResponseBodyProducts {
	s.Versions = v
	return s
}

func (s *ListProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}
