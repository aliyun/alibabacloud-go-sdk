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
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// Default value: zh-CN.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// The keyword for searching resource codes or names. Fuzzy match is supported.
	//
	// example:
	//
	// vpc
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries per page. Valid values: 0 to 200. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// tokenForNextPage
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The product code. Fuzzy match is supported.
	//
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The order in which resource types are returned. Valid values:
	//
	// - Normal (default): returned in normal order.
	//
	// - Top: returned in order of popularity.
	//
	// example:
	//
	// Normal
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// The status filter list. Valid values:
	//
	// - Available
	//
	// - Deprecated.
	//
	// example:
	//
	// Available,Deprecated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The product subcategory in Terraform.
	//
	// example:
	//
	// compute
	Subcategory *string `json:"subcategory,omitempty" xml:"subcategory,omitempty"`
	// Specifies whether Terraformer is supported.
	//
	// example:
	//
	// true
	SupportTerraformer *bool `json:"supportTerraformer,omitempty" xml:"supportTerraformer,omitempty"`
	// The Terraform provider version. If this parameter is left empty, the latest version is used by default.
	//
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform resources.
	TerraformResourceTypes []*string `json:"terraformResourceTypes,omitempty" xml:"terraformResourceTypes,omitempty" type:"Repeated"`
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
