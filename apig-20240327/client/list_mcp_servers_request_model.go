// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateFromTypes(v string) *ListMcpServersRequest
	GetCreateFromTypes() *string
	SetDeployStatuses(v string) *ListMcpServersRequest
	GetDeployStatuses() *string
	SetGatewayId(v string) *ListMcpServersRequest
	GetGatewayId() *string
	SetNameLike(v string) *ListMcpServersRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListMcpServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMcpServersRequest
	GetPageSize() *int32
	SetType(v string) *ListMcpServersRequest
	GetType() *string
}

type ListMcpServersRequest struct {
	// example:
	//
	// ApiGatewayHttpToMCP
	CreateFromTypes *string `json:"createFromTypes,omitempty" xml:"createFromTypes,omitempty"`
	// example:
	//
	// Deployed
	DeployStatuses *string `json:"deployStatuses,omitempty" xml:"deployStatuses,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-co370icmjeu****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// RealMCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMcpServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersRequest) GoString() string {
	return s.String()
}

func (s *ListMcpServersRequest) GetCreateFromTypes() *string {
	return s.CreateFromTypes
}

func (s *ListMcpServersRequest) GetDeployStatuses() *string {
	return s.DeployStatuses
}

func (s *ListMcpServersRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListMcpServersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListMcpServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMcpServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcpServersRequest) GetType() *string {
	return s.Type
}

func (s *ListMcpServersRequest) SetCreateFromTypes(v string) *ListMcpServersRequest {
	s.CreateFromTypes = &v
	return s
}

func (s *ListMcpServersRequest) SetDeployStatuses(v string) *ListMcpServersRequest {
	s.DeployStatuses = &v
	return s
}

func (s *ListMcpServersRequest) SetGatewayId(v string) *ListMcpServersRequest {
	s.GatewayId = &v
	return s
}

func (s *ListMcpServersRequest) SetNameLike(v string) *ListMcpServersRequest {
	s.NameLike = &v
	return s
}

func (s *ListMcpServersRequest) SetPageNumber(v int32) *ListMcpServersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMcpServersRequest) SetPageSize(v int32) *ListMcpServersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcpServersRequest) SetType(v string) *ListMcpServersRequest {
	s.Type = &v
	return s
}

func (s *ListMcpServersRequest) Validate() error {
	return dara.Validate(s)
}
