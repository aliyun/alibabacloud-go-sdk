// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListModelApisRequest
	GetGatewayId() *string
	SetIncludeLegacy(v bool) *ListModelApisRequest
	GetIncludeLegacy() *bool
	SetKeyword(v string) *ListModelApisRequest
	GetKeyword() *string
	SetName(v string) *ListModelApisRequest
	GetName() *string
	SetPageNumber(v int32) *ListModelApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelApisRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListModelApisRequest
	GetResourceGroupId() *string
}

type ListModelApisRequest struct {
	// The gateway instance ID used for filtering.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Specifies whether to return legacy model APIs created by scenario. If omitted or set to false, only APIs created by selecting a model are returned. If set to true, APIs created by both methods are returned on the same page, excluding common HTTP, MCP, or Agent APIs.
	IncludeLegacy *bool `json:"includeLegacy,omitempty" xml:"includeLegacy,omitempty"`
	// The search keyword. Supports fuzzy search by API name or exact search by API ID.
	//
	// example:
	//
	// beta
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The API name for exact search.
	//
	// example:
	//
	// beta-model-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group used for filtering.
	//
	// example:
	//
	// rg-acf0000000000000000
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListModelApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelApisRequest) GoString() string {
	return s.String()
}

func (s *ListModelApisRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListModelApisRequest) GetIncludeLegacy() *bool {
	return s.IncludeLegacy
}

func (s *ListModelApisRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModelApisRequest) GetName() *string {
	return s.Name
}

func (s *ListModelApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelApisRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListModelApisRequest) SetGatewayId(v string) *ListModelApisRequest {
	s.GatewayId = &v
	return s
}

func (s *ListModelApisRequest) SetIncludeLegacy(v bool) *ListModelApisRequest {
	s.IncludeLegacy = &v
	return s
}

func (s *ListModelApisRequest) SetKeyword(v string) *ListModelApisRequest {
	s.Keyword = &v
	return s
}

func (s *ListModelApisRequest) SetName(v string) *ListModelApisRequest {
	s.Name = &v
	return s
}

func (s *ListModelApisRequest) SetPageNumber(v int32) *ListModelApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelApisRequest) SetPageSize(v int32) *ListModelApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelApisRequest) SetResourceGroupId(v string) *ListModelApisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListModelApisRequest) Validate() error {
	return dara.Validate(s)
}
