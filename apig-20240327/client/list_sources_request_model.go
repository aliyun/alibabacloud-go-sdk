// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListSourcesRequest
	GetGatewayId() *string
	SetPageNumber(v int32) *ListSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSourcesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListSourcesRequest
	GetResourceGroupId() *string
	SetType(v string) *ListSourcesRequest
	GetType() *string
}

type ListSourcesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The type by which you want to query sources.
	//
	// Valid values:
	//
	// 	- K8S
	//
	// 	- MSE_NACOS
	//
	// example:
	//
	// MSE_NACOS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListSourcesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListSourcesRequest) SetGatewayId(v string) *ListSourcesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListSourcesRequest) SetPageNumber(v int32) *ListSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSourcesRequest) SetPageSize(v int32) *ListSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourcesRequest) SetResourceGroupId(v string) *ListSourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSourcesRequest) SetType(v string) *ListSourcesRequest {
	s.Type = &v
	return s
}

func (s *ListSourcesRequest) Validate() error {
	return dara.Validate(s)
}
