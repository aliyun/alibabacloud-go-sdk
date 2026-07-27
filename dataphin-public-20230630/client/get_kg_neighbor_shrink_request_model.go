// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgNeighborShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityDataId(v string) *GetKgNeighborShrinkRequest
	GetEntityDataId() *string
	SetEntityType(v string) *GetKgNeighborShrinkRequest
	GetEntityType() *string
	SetNeighborsQueryShrink(v string) *GetKgNeighborShrinkRequest
	GetNeighborsQueryShrink() *string
	SetOpTenantId(v int64) *GetKgNeighborShrinkRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *GetKgNeighborShrinkRequest
	GetWorkspaceId() *string
}

type GetKgNeighborShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityDataId *string `json:"EntityDataId,omitempty" xml:"EntityDataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Student
	EntityType           *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	NeighborsQueryShrink *string `json:"NeighborsQuery,omitempty" xml:"NeighborsQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKgNeighborShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetKgNeighborShrinkRequest) GetEntityDataId() *string {
	return s.EntityDataId
}

func (s *GetKgNeighborShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKgNeighborShrinkRequest) GetNeighborsQueryShrink() *string {
	return s.NeighborsQueryShrink
}

func (s *GetKgNeighborShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgNeighborShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgNeighborShrinkRequest) SetEntityDataId(v string) *GetKgNeighborShrinkRequest {
	s.EntityDataId = &v
	return s
}

func (s *GetKgNeighborShrinkRequest) SetEntityType(v string) *GetKgNeighborShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *GetKgNeighborShrinkRequest) SetNeighborsQueryShrink(v string) *GetKgNeighborShrinkRequest {
	s.NeighborsQueryShrink = &v
	return s
}

func (s *GetKgNeighborShrinkRequest) SetOpTenantId(v int64) *GetKgNeighborShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgNeighborShrinkRequest) SetWorkspaceId(v string) *GetKgNeighborShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgNeighborShrinkRequest) Validate() error {
	return dara.Validate(s)
}
