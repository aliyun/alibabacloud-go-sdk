// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListDomainsRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListDomainsRequest
	GetGatewayType() *string
	SetNameLike(v string) *ListDomainsRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListDomainsRequest
	GetResourceGroupId() *string
}

type ListDomainsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The domain name keyword for fuzzy search.
	//
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ahr5uil8raz0rq3b
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListDomainsRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListDomainsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDomainsRequest) SetGatewayId(v string) *ListDomainsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListDomainsRequest) SetGatewayType(v string) *ListDomainsRequest {
	s.GatewayType = &v
	return s
}

func (s *ListDomainsRequest) SetNameLike(v string) *ListDomainsRequest {
	s.NameLike = &v
	return s
}

func (s *ListDomainsRequest) SetPageNumber(v int32) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int32) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainsRequest) SetResourceGroupId(v string) *ListDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDomainsRequest) Validate() error {
	return dara.Validate(s)
}
