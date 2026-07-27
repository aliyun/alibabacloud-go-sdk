// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListKgEntityShrinkRequest
	GetEntityType() *string
	SetListQueryShrink(v string) *ListKgEntityShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListKgEntityShrinkRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *ListKgEntityShrinkRequest
	GetWorkspaceId() *string
}

type ListKgEntityShrinkRequest struct {
	// The entity type code.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The paged query filter conditions.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
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

func (s ListKgEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKgEntityShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKgEntityShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListKgEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListKgEntityShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKgEntityShrinkRequest) SetEntityType(v string) *ListKgEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *ListKgEntityShrinkRequest) SetListQueryShrink(v string) *ListKgEntityShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListKgEntityShrinkRequest) SetOpTenantId(v int64) *ListKgEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListKgEntityShrinkRequest) SetWorkspaceId(v string) *ListKgEntityShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKgEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
