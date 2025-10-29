// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsResponseBody
	GetNextToken() *string
	SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody
	GetProducts() []*ListProductsResponseBodyProducts
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsResponseBody
	GetTotalCount() *int32
}

type ListProductsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 30BaZ9ekYWXJdqshYecA++coNg7qT1Zbm3RfLyFIZeY=
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Products  []*ListProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 9bcaac3c-420d-4303-87ab-7638c07b0a0b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 134
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsResponseBody) GetProducts() []*ListProductsResponseBodyProducts {
	return s.Products
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
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
	FirstCategoryName *string `json:"firstCategoryName,omitempty" xml:"firstCategoryName,omitempty"`
	// example:
	//
	// Enterprise application
	FirstCategoryNameEn *string `json:"firstCategoryNameEn,omitempty" xml:"firstCategoryNameEn,omitempty"`
	// example:
	//
	// MSE
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// Microservices Engine
	ProductNameEn      *string `json:"productNameEn,omitempty" xml:"productNameEn,omitempty"`
	SecondCategoryName *string `json:"secondCategoryName,omitempty" xml:"secondCategoryName,omitempty"`
	// example:
	//
	// Application service
	SecondCategoryNameEn *string `json:"secondCategoryNameEn,omitempty" xml:"secondCategoryNameEn,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// Microservice_Engine(MSE)
	Subcategory *string `json:"subcategory,omitempty" xml:"subcategory,omitempty"`
	// example:
	//
	// true
	SupportTerraformer *bool `json:"supportTerraformer,omitempty" xml:"supportTerraformer,omitempty"`
	// example:
	//
	// 1.229.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *ListProductsResponseBodyProducts) GetFirstCategoryNameEn() *string {
	return s.FirstCategoryNameEn
}

func (s *ListProductsResponseBodyProducts) GetProduct() *string {
	return s.Product
}

func (s *ListProductsResponseBodyProducts) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsResponseBodyProducts) GetProductNameEn() *string {
	return s.ProductNameEn
}

func (s *ListProductsResponseBodyProducts) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *ListProductsResponseBodyProducts) GetSecondCategoryNameEn() *string {
	return s.SecondCategoryNameEn
}

func (s *ListProductsResponseBodyProducts) GetStatus() *string {
	return s.Status
}

func (s *ListProductsResponseBodyProducts) GetSubcategory() *string {
	return s.Subcategory
}

func (s *ListProductsResponseBodyProducts) GetSupportTerraformer() *bool {
	return s.SupportTerraformer
}

func (s *ListProductsResponseBodyProducts) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListProductsResponseBodyProducts) SetFirstCategoryName(v string) *ListProductsResponseBodyProducts {
	s.FirstCategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetFirstCategoryNameEn(v string) *ListProductsResponseBodyProducts {
	s.FirstCategoryNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProduct(v string) *ListProductsResponseBodyProducts {
	s.Product = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductName(v string) *ListProductsResponseBodyProducts {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductNameEn(v string) *ListProductsResponseBodyProducts {
	s.ProductNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetSecondCategoryName(v string) *ListProductsResponseBodyProducts {
	s.SecondCategoryName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetSecondCategoryNameEn(v string) *ListProductsResponseBodyProducts {
	s.SecondCategoryNameEn = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetStatus(v string) *ListProductsResponseBodyProducts {
	s.Status = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetSubcategory(v string) *ListProductsResponseBodyProducts {
	s.Subcategory = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetSupportTerraformer(v bool) *ListProductsResponseBodyProducts {
	s.SupportTerraformer = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetTerraformProviderVersion(v string) *ListProductsResponseBodyProducts {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}
