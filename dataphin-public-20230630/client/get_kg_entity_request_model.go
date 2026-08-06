// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *GetKgEntityRequest
	GetEntityId() *string
	SetEntityType(v string) *GetKgEntityRequest
	GetEntityType() *string
	SetOpTenantId(v int64) *GetKgEntityRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *GetKgEntityRequest
	GetWorkspaceId() *string
}

type GetKgEntityRequest struct {
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

func (s GetKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgEntityRequest) GoString() string {
	return s.String()
}

func (s *GetKgEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKgEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgEntityRequest) SetEntityId(v string) *GetKgEntityRequest {
	s.EntityId = &v
	return s
}

func (s *GetKgEntityRequest) SetEntityType(v string) *GetKgEntityRequest {
	s.EntityType = &v
	return s
}

func (s *GetKgEntityRequest) SetOpTenantId(v int64) *GetKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgEntityRequest) SetWorkspaceId(v string) *GetKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgEntityRequest) Validate() error {
	return dara.Validate(s)
}
