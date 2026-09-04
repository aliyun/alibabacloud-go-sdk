// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *ListOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetNextToken(v string) *ListOperatingObjectFavoritesRequest
	GetNextToken() *string
	SetObjectType(v string) *ListOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *ListOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetPageSize(v int64) *ListOperatingObjectFavoritesRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type ListOperatingObjectFavoritesRequest struct {
	// The graph name. Call listGraphs to retrieve available graphs.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJ2IjoxLCJpZCI6MTAwMX0.c2lnbmF0dXJlX2V4YW1wbGU
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The digital employee name (operating object name). Optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page size.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID to take effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *ListOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListOperatingObjectFavoritesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOperatingObjectFavoritesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOperatingObjectFavoritesRequest) SetGraphName(v string) *ListOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetNextToken(v string) *ListOperatingObjectFavoritesRequest {
	s.NextToken = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetObjectType(v string) *ListOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *ListOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetPageSize(v int64) *ListOperatingObjectFavoritesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetTenantId(v string) *ListOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
