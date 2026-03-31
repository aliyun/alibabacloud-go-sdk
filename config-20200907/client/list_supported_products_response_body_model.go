// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListSupportedProductsResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListSupportedProductsResponseBody
	GetNextToken() *string
	SetProducts(v []*ListSupportedProductsResponseBodyProducts) *ListSupportedProductsResponseBody
	GetProducts() []*ListSupportedProductsResponseBodyProducts
	SetRequestId(v string) *ListSupportedProductsResponseBody
	GetRequestId() *string
}

type ListSupportedProductsResponseBody struct {
	// The maximum number of entries to return for a single request. Valid values: 1 to 500.
	//
	// example:
	//
	// 100
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// D3AjqMNSy0ls7zBNCf3a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The cloud services that are supported by Cloud Config.
	Products []*ListSupportedProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 610B0276-ABEE-57DF-9C13-C2324FADA9D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSupportedProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportedProductsResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListSupportedProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportedProductsResponseBody) GetProducts() []*ListSupportedProductsResponseBodyProducts {
	return s.Products
}

func (s *ListSupportedProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportedProductsResponseBody) SetMaxResults(v string) *ListSupportedProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupportedProductsResponseBody) SetNextToken(v string) *ListSupportedProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupportedProductsResponseBody) SetProducts(v []*ListSupportedProductsResponseBodyProducts) *ListSupportedProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListSupportedProductsResponseBody) SetRequestId(v string) *ListSupportedProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportedProductsResponseBody) Validate() error {
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

type ListSupportedProductsResponseBodyProducts struct {
	// The English name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service
	ProductNameEn *string `json:"ProductNameEn,omitempty" xml:"ProductNameEn,omitempty"`
	// The Chinese name of the Alibaba Cloud service.
	//
	// example:
	//
	// 云服务器ECS
	ProductNameZh *string `json:"ProductNameZh,omitempty" xml:"ProductNameZh,omitempty"`
	// The resource types that are supported by Cloud Config.
	ResourceTypeList []*ListSupportedProductsResponseBodyProductsResourceTypeList `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
}

func (s ListSupportedProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListSupportedProductsResponseBodyProducts) GetProductNameEn() *string {
	return s.ProductNameEn
}

func (s *ListSupportedProductsResponseBodyProducts) GetProductNameZh() *string {
	return s.ProductNameZh
}

func (s *ListSupportedProductsResponseBodyProducts) GetResourceTypeList() []*ListSupportedProductsResponseBodyProductsResourceTypeList {
	return s.ResourceTypeList
}

func (s *ListSupportedProductsResponseBodyProducts) SetProductNameEn(v string) *ListSupportedProductsResponseBodyProducts {
	s.ProductNameEn = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProducts) SetProductNameZh(v string) *ListSupportedProductsResponseBodyProducts {
	s.ProductNameZh = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProducts) SetResourceTypeList(v []*ListSupportedProductsResponseBodyProductsResourceTypeList) *ListSupportedProductsResponseBodyProducts {
	s.ResourceTypeList = v
	return s
}

func (s *ListSupportedProductsResponseBodyProducts) Validate() error {
	if s.ResourceTypeList != nil {
		for _, item := range s.ResourceTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupportedProductsResponseBodyProductsResourceTypeList struct {
	// The identifier of the resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The English name of the resource type.
	//
	// example:
	//
	// Ecs Instance
	TypeNameEn *string `json:"TypeNameEn,omitempty" xml:"TypeNameEn,omitempty"`
	// The Chinese name of the resource type.
	//
	// example:
	//
	// ECS实例
	TypeNameZh *string `json:"TypeNameZh,omitempty" xml:"TypeNameZh,omitempty"`
	// The URL of the resource type in the console.
	//
	// example:
	//
	// https://ecs.console.aliyun.com/#/server/@{ResourceId}/detail?regionId=@{RegionId}
	TypePageLink *string `json:"TypePageLink,omitempty" xml:"TypePageLink,omitempty"`
}

func (s ListSupportedProductsResponseBodyProductsResourceTypeList) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedProductsResponseBodyProductsResourceTypeList) GoString() string {
	return s.String()
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) GetTypeNameEn() *string {
	return s.TypeNameEn
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) GetTypeNameZh() *string {
	return s.TypeNameZh
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) GetTypePageLink() *string {
	return s.TypePageLink
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) SetResourceType(v string) *ListSupportedProductsResponseBodyProductsResourceTypeList {
	s.ResourceType = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) SetTypeNameEn(v string) *ListSupportedProductsResponseBodyProductsResourceTypeList {
	s.TypeNameEn = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) SetTypeNameZh(v string) *ListSupportedProductsResponseBodyProductsResourceTypeList {
	s.TypeNameZh = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) SetTypePageLink(v string) *ListSupportedProductsResponseBodyProductsResourceTypeList {
	s.TypePageLink = &v
	return s
}

func (s *ListSupportedProductsResponseBodyProductsResourceTypeList) Validate() error {
	return dara.Validate(s)
}
