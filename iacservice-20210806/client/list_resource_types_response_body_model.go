// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceTypesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceTypesResponseBody
	GetRequestId() *string
	SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody
	GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes
	SetTotalCount(v int32) *ListResourceTypesResponseBody
	GetTotalCount() *int32
}

type ListResourceTypesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// LC4NJL3Ru2bIiRdnbADPQp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 9bcaac3c-420d-4303-87ab-7638c07b0a0b
	RequestId     *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 93
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypesResponseBody) GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes {
	return s.ResourceTypes
}

func (s *ListResourceTypesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceTypesResponseBody) SetMaxResults(v int32) *ListResourceTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetNextToken(v string) *ListResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) SetTotalCount(v int32) *ListResourceTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceTypesResponseBodyResourceTypes struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// VPC
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// https://vpc.console.aliyun.com/vpc/${RegionId}/route-tables/${RouteTableId}
	ResourceDetailPageUrl *string `json:"resourceDetailPageUrl,omitempty" xml:"resourceDetailPageUrl,omitempty"`
	// example:
	//
	// https://vpc.console.aliyun.com/vpc/${RegionId}/route-tables
	ResourceListPageUrl *string `json:"resourceListPageUrl,omitempty" xml:"resourceListPageUrl,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1.0.0
	StatusStartVersion *string `json:"statusStartVersion,omitempty" xml:"statusStartVersion,omitempty"`
	// example:
	//
	// VPC
	Subcategory *string `json:"subcategory,omitempty" xml:"subcategory,omitempty"`
	// example:
	//
	// true
	SupportTerraformer *string `json:"supportTerraformer,omitempty" xml:"supportTerraformer,omitempty"`
	// example:
	//
	// 1.248.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// alicloud_route_table
	TerraformResourceType *string `json:"terraformResourceType,omitempty" xml:"terraformResourceType,omitempty"`
	Title                 *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetDescription() *string {
	return s.Description
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetProduct() *string {
	return s.Product
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetProductName() *string {
	return s.ProductName
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetResourceDetailPageUrl() *string {
	return s.ResourceDetailPageUrl
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetResourceListPageUrl() *string {
	return s.ResourceListPageUrl
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetStatus() *string {
	return s.Status
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetStatusStartVersion() *string {
	return s.StatusStartVersion
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetSubcategory() *string {
	return s.Subcategory
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetSupportTerraformer() *string {
	return s.SupportTerraformer
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetTerraformResourceType() *string {
	return s.TerraformResourceType
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetTitle() *string {
	return s.Title
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProduct(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Product = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceDetailPageUrl(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceDetailPageUrl = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceListPageUrl(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceListPageUrl = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetStatus(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Status = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetStatusStartVersion(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.StatusStartVersion = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetSubcategory(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Subcategory = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetSupportTerraformer(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.SupportTerraformer = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetTerraformProviderVersion(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetTerraformResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.TerraformResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetTitle(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Title = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) Validate() error {
	return dara.Validate(s)
}
