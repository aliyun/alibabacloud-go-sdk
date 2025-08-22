// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProductsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProductsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsRequest
	GetNextToken() *string
	SetSort(v string) *ListProductsRequest
	GetSort() *string
	SetStatus(v string) *ListProductsRequest
	GetStatus() *string
	SetSupportTerraformer(v bool) *ListProductsRequest
	GetSupportTerraformer() *bool
	SetTerraformProviderVersion(v string) *ListProductsRequest
	GetTerraformProviderVersion() *string
}

type ListProductsRequest struct {
	// example:
	//
	// ECS
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJ0b2tlbiI6IjEwMjM0NTY3ODkwIn0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// Normal
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	SupportTerraformer *bool `json:"supportTerraformer,omitempty" xml:"supportTerraformer,omitempty"`
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProductsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListProductsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProductsRequest) GetSupportTerraformer() *bool {
	return s.SupportTerraformer
}

func (s *ListProductsRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListProductsRequest) SetKeyword(v string) *ListProductsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsRequest) SetSort(v string) *ListProductsRequest {
	s.Sort = &v
	return s
}

func (s *ListProductsRequest) SetStatus(v string) *ListProductsRequest {
	s.Status = &v
	return s
}

func (s *ListProductsRequest) SetSupportTerraformer(v bool) *ListProductsRequest {
	s.SupportTerraformer = &v
	return s
}

func (s *ListProductsRequest) SetTerraformProviderVersion(v string) *ListProductsRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
