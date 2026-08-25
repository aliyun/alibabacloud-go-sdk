// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListResourceTypesShrinkRequest
	GetAcceptLanguage() *string
	SetKeyword(v string) *ListResourceTypesShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListResourceTypesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesShrinkRequest
	GetNextToken() *string
	SetProduct(v string) *ListResourceTypesShrinkRequest
	GetProduct() *string
	SetSort(v string) *ListResourceTypesShrinkRequest
	GetSort() *string
	SetStatus(v string) *ListResourceTypesShrinkRequest
	GetStatus() *string
	SetSubcategory(v string) *ListResourceTypesShrinkRequest
	GetSubcategory() *string
	SetSupportTerraformer(v bool) *ListResourceTypesShrinkRequest
	GetSupportTerraformer() *bool
	SetTerraformProviderVersion(v string) *ListResourceTypesShrinkRequest
	GetTerraformProviderVersion() *string
	SetTerraformResourceTypesShrink(v string) *ListResourceTypesShrinkRequest
	GetTerraformResourceTypesShrink() *string
}

type ListResourceTypesShrinkRequest struct {
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
	// The keyword for searching resource code or name. Fuzzy match is supported.
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
	// - Top: returned in order of popular access.
	//
	// example:
	//
	// Normal
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// The status list for filtering. Valid values:
	//
	// - Available
	//
	// - Deprecated
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
	// The Terraform Provider version. If this parameter is left empty, the latest version is used by default.
	//
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The list of Terraform resources.
	TerraformResourceTypesShrink *string `json:"terraformResourceTypes,omitempty" xml:"terraformResourceTypes,omitempty"`
}

func (s ListResourceTypesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListResourceTypesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceTypesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *ListResourceTypesShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourceTypesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceTypesShrinkRequest) GetSubcategory() *string {
	return s.Subcategory
}

func (s *ListResourceTypesShrinkRequest) GetSupportTerraformer() *bool {
	return s.SupportTerraformer
}

func (s *ListResourceTypesShrinkRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListResourceTypesShrinkRequest) GetTerraformResourceTypesShrink() *string {
	return s.TerraformResourceTypesShrink
}

func (s *ListResourceTypesShrinkRequest) SetAcceptLanguage(v string) *ListResourceTypesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetKeyword(v string) *ListResourceTypesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetMaxResults(v int32) *ListResourceTypesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetNextToken(v string) *ListResourceTypesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetProduct(v string) *ListResourceTypesShrinkRequest {
	s.Product = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetSort(v string) *ListResourceTypesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetStatus(v string) *ListResourceTypesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetSubcategory(v string) *ListResourceTypesShrinkRequest {
	s.Subcategory = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetSupportTerraformer(v bool) *ListResourceTypesShrinkRequest {
	s.SupportTerraformer = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetTerraformProviderVersion(v string) *ListResourceTypesShrinkRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetTerraformResourceTypesShrink(v string) *ListResourceTypesShrinkRequest {
	s.TerraformResourceTypesShrink = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
