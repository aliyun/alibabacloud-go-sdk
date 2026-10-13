// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v string) *ListProductsRequest
	GetCategoryIds() *string
	SetMaxResults(v int32) *ListProductsRequest
	GetMaxResults() *int32
	SetName(v string) *ListProductsRequest
	GetName() *string
	SetNextToken(v string) *ListProductsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsRequest
	GetPageSize() *int32
	SetPortalId(v string) *ListProductsRequest
	GetPortalId() *string
	SetStatus(v string) *ListProductsRequest
	GetStatus() *string
	SetType(v string) *ListProductsRequest
	GetType() *string
}

type ListProductsRequest struct {
	// The product category IDs, which correspond to the Categories property of the API product resource. The members are category IDs.
	//
	// example:
	//
	// category-xxxxxxxx
	CategoryIds *string `json:"categoryIds,omitempty" xml:"categoryIds,omitempty"`
	// The maximum number of entries to return per request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the API product, which is the Name property of the API product resource. This is the same field as name in the response list items.
	//
	// example:
	//
	// weather-query-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The cursor token for the next page.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the HiMarket developer portal where the API product is published, which corresponds to the PortalId property of the API product resource. This property is used only for list queries in the product resource definition.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The status of the API product, which corresponds to the Status property of the API product resource. This is the same field as status in the response list items. Valid values: Pending, Ready, and Published.
	//
	// example:
	//
	// Published
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the API product, which corresponds to the Type property of the API product resource. This is the same field as type in the response list items. The valid values are not declared in the API definition. Refer to the types displayed in the console.
	//
	// example:
	//
	// API
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetCategoryIds() *string {
	return s.CategoryIds
}

func (s *ListProductsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsRequest) GetName() *string {
	return s.Name
}

func (s *ListProductsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsRequest) GetPortalId() *string {
	return s.PortalId
}

func (s *ListProductsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProductsRequest) GetType() *string {
	return s.Type
}

func (s *ListProductsRequest) SetCategoryIds(v string) *ListProductsRequest {
	s.CategoryIds = &v
	return s
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetName(v string) *ListProductsRequest {
	s.Name = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsRequest) SetPageNumber(v int32) *ListProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsRequest) SetPageSize(v int32) *ListProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsRequest) SetPortalId(v string) *ListProductsRequest {
	s.PortalId = &v
	return s
}

func (s *ListProductsRequest) SetStatus(v string) *ListProductsRequest {
	s.Status = &v
	return s
}

func (s *ListProductsRequest) SetType(v string) *ListProductsRequest {
	s.Type = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
