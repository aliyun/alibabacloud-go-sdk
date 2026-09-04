// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetObjectIds(v []*string) *BatchRemoveOperatingObjectFavoritesRequest
	GetObjectIds() []*string
	SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type BatchRemoveOperatingObjectFavoritesRequest struct {
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The list of primary object business IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract-001
	ObjectIds []*string `json:"objectIds,omitempty" xml:"objectIds,omitempty" type:"Repeated"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The operating object name.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly in winnexo-cli by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s BatchRemoveOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetObjectIds(v []*string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.ObjectIds = v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
