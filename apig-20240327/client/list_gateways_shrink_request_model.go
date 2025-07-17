// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaysShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListGatewaysShrinkRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListGatewaysShrinkRequest
	GetGatewayType() *string
	SetKeyword(v string) *ListGatewaysShrinkRequest
	GetKeyword() *string
	SetName(v string) *ListGatewaysShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListGatewaysShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewaysShrinkRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListGatewaysShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListGatewaysShrinkRequest
	GetTagShrink() *string
}

type ListGatewaysShrinkRequest struct {
	// The instance ID. If you specify an ID, an exact search is performed.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The search keyword. A full match is performed. The search is case-insensitive.
	//
	// example:
	//
	// dev
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The instance name. If you specify a name, an exact search is performed.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz3wes3hnre5a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags that you want to use for the search.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListGatewaysShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysShrinkRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewaysShrinkRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewaysShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGatewaysShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListGatewaysShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewaysShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewaysShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGatewaysShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListGatewaysShrinkRequest) SetGatewayId(v string) *ListGatewaysShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetGatewayType(v string) *ListGatewaysShrinkRequest {
	s.GatewayType = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetKeyword(v string) *ListGatewaysShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetName(v string) *ListGatewaysShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetPageNumber(v int32) *ListGatewaysShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetPageSize(v int32) *ListGatewaysShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetResourceGroupId(v string) *ListGatewaysShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetTagShrink(v string) *ListGatewaysShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListGatewaysShrinkRequest) Validate() error {
	return dara.Validate(s)
}
