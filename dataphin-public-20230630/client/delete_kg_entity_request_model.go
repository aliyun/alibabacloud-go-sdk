// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *DeleteKgEntityRequest
	GetEntityId() *string
	SetEntityType(v string) *DeleteKgEntityRequest
	GetEntityType() *string
	SetOpTenantId(v int64) *DeleteKgEntityRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *DeleteKgEntityRequest
	GetWorkspaceId() *string
}

type DeleteKgEntityRequest struct {
	// The entity record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteKgEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteKgEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DeleteKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteKgEntityRequest) SetEntityId(v string) *DeleteKgEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteKgEntityRequest) SetEntityType(v string) *DeleteKgEntityRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteKgEntityRequest) SetOpTenantId(v int64) *DeleteKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteKgEntityRequest) SetWorkspaceId(v string) *DeleteKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteKgEntityRequest) Validate() error {
	return dara.Validate(s)
}
