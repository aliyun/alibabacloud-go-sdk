// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasLike(v string) *ListEnvironmentsRequest
	GetAliasLike() *string
	SetGatewayId(v string) *ListEnvironmentsRequest
	GetGatewayId() *string
	SetGatewayNameLike(v string) *ListEnvironmentsRequest
	GetGatewayNameLike() *string
	SetGatewayType(v string) *ListEnvironmentsRequest
	GetGatewayType() *string
	SetNameLike(v string) *ListEnvironmentsRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListEnvironmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnvironmentsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListEnvironmentsRequest
	GetResourceGroupId() *string
}

type ListEnvironmentsRequest struct {
	// Environment alias, fuzzy search.
	//
	// example:
	//
	// production
	AliasLike *string `json:"aliasLike,omitempty" xml:"aliasLike,omitempty"`
	// Gateway ID, exact search.
	//
	// example:
	//
	// gw-cptv6ktlhtgnqr73h8d1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Gateway name, fuzzy search.
	//
	// example:
	//
	// test-gw
	GatewayNameLike *string `json:"gatewayNameLike,omitempty" xml:"gatewayNameLike,omitempty"`
	// The gateway type
	//
	// example:
	//
	// APIGateway
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// Environment name, fuzzy search.
	//
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// Page number, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek2sy66mftleiq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) GetAliasLike() *string {
	return s.AliasLike
}

func (s *ListEnvironmentsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListEnvironmentsRequest) GetGatewayNameLike() *string {
	return s.GatewayNameLike
}

func (s *ListEnvironmentsRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListEnvironmentsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListEnvironmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnvironmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnvironmentsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEnvironmentsRequest) SetAliasLike(v string) *ListEnvironmentsRequest {
	s.AliasLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetGatewayId(v string) *ListEnvironmentsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetGatewayNameLike(v string) *ListEnvironmentsRequest {
	s.GatewayNameLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetGatewayType(v string) *ListEnvironmentsRequest {
	s.GatewayType = &v
	return s
}

func (s *ListEnvironmentsRequest) SetNameLike(v string) *ListEnvironmentsRequest {
	s.NameLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageNumber(v int32) *ListEnvironmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int32) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsRequest) SetResourceGroupId(v string) *ListEnvironmentsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEnvironmentsRequest) Validate() error {
	return dara.Validate(s)
}
