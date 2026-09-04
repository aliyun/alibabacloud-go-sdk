// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *ClearOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetObjectType(v string) *ClearOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *ClearOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type ClearOperatingObjectFavoritesRequest struct {
	// The graph name. You can call listGraphs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The operating object name, such as customer_1.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID to take effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ClearOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *ClearOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ClearOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ClearOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ClearOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearOperatingObjectFavoritesRequest) SetGraphName(v string) *ClearOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetObjectType(v string) *ClearOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetTenantId(v string) *ClearOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
