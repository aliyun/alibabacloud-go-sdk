// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListResourceTypesRequest
	GetAcceptLanguage() *string
	SetKeyword(v string) *ListResourceTypesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListResourceTypesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesRequest
	GetNextToken() *string
	SetProduct(v string) *ListResourceTypesRequest
	GetProduct() *string
	SetSort(v string) *ListResourceTypesRequest
	GetSort() *string
	SetStatus(v string) *ListResourceTypesRequest
	GetStatus() *string
	SetSubcategory(v string) *ListResourceTypesRequest
	GetSubcategory() *string
	SetSupportTerraformer(v bool) *ListResourceTypesRequest
	GetSupportTerraformer() *bool
	SetTerraformProviderVersion(v string) *ListResourceTypesRequest
	GetTerraformProviderVersion() *string
	SetTerraformResourceTypes(v []*string) *ListResourceTypesRequest
	GetTerraformResourceTypes() []*string
}

type ListResourceTypesRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// example:
	//
	// vpc
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// tokenForNextPage
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// example:
	//
	// Normal
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// Available,Deprecated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// compute
	Subcategory *string `json:"subcategory,omitempty" xml:"subcategory,omitempty"`
	// example:
	//
	// true
	SupportTerraformer *bool `json:"supportTerraformer,omitempty" xml:"supportTerraformer,omitempty"`
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string   `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TerraformResourceTypes   []*string `json:"terraformResourceTypes,omitempty" xml:"terraformResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListResourceTypesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceTypesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesRequest) GetProduct() *string {
	return s.Product
}

func (s *ListResourceTypesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourceTypesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceTypesRequest) GetSubcategory() *string {
	return s.Subcategory
}

func (s *ListResourceTypesRequest) GetSupportTerraformer() *bool {
	return s.SupportTerraformer
}

func (s *ListResourceTypesRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListResourceTypesRequest) GetTerraformResourceTypes() []*string {
	return s.TerraformResourceTypes
}

func (s *ListResourceTypesRequest) SetAcceptLanguage(v string) *ListResourceTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesRequest) SetKeyword(v string) *ListResourceTypesRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceTypesRequest) SetMaxResults(v int32) *ListResourceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesRequest) SetNextToken(v string) *ListResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesRequest) SetProduct(v string) *ListResourceTypesRequest {
	s.Product = &v
	return s
}

func (s *ListResourceTypesRequest) SetSort(v string) *ListResourceTypesRequest {
	s.Sort = &v
	return s
}

func (s *ListResourceTypesRequest) SetStatus(v string) *ListResourceTypesRequest {
	s.Status = &v
	return s
}

func (s *ListResourceTypesRequest) SetSubcategory(v string) *ListResourceTypesRequest {
	s.Subcategory = &v
	return s
}

func (s *ListResourceTypesRequest) SetSupportTerraformer(v bool) *ListResourceTypesRequest {
	s.SupportTerraformer = &v
	return s
}

func (s *ListResourceTypesRequest) SetTerraformProviderVersion(v string) *ListResourceTypesRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListResourceTypesRequest) SetTerraformResourceTypes(v []*string) *ListResourceTypesRequest {
	s.TerraformResourceTypes = v
	return s
}

func (s *ListResourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
