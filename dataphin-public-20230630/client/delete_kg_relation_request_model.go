// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *DeleteKgRelationRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteKgRelationRequest
	GetOpUserId() *string
	SetRelationId(v string) *DeleteKgRelationRequest
	GetRelationId() *string
	SetRelationType(v string) *DeleteKgRelationRequest
	GetRelationType() *string
	SetWorkspaceId(v string) *DeleteKgRelationRequest
	GetWorkspaceId() *string
}

type DeleteKgRelationRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The relationship record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteKgRelationRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteKgRelationRequest) GetRelationId() *string {
	return s.RelationId
}

func (s *DeleteKgRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *DeleteKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteKgRelationRequest) SetOpTenantId(v int64) *DeleteKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteKgRelationRequest) SetOpUserId(v string) *DeleteKgRelationRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteKgRelationRequest) SetRelationId(v string) *DeleteKgRelationRequest {
	s.RelationId = &v
	return s
}

func (s *DeleteKgRelationRequest) SetRelationType(v string) *DeleteKgRelationRequest {
	s.RelationType = &v
	return s
}

func (s *DeleteKgRelationRequest) SetWorkspaceId(v string) *DeleteKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteKgRelationRequest) Validate() error {
	return dara.Validate(s)
}
