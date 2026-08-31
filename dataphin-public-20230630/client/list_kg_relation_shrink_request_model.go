// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListKgRelationShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListKgRelationShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListKgRelationShrinkRequest
	GetOpUserId() *string
	SetRelationType(v string) *ListKgRelationShrinkRequest
	GetRelationType() *string
	SetWorkspaceId(v string) *ListKgRelationShrinkRequest
	GetWorkspaceId() *string
}

type ListKgRelationShrinkRequest struct {
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
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s ListKgRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKgRelationShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListKgRelationShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListKgRelationShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListKgRelationShrinkRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ListKgRelationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKgRelationShrinkRequest) SetListQueryShrink(v string) *ListKgRelationShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListKgRelationShrinkRequest) SetOpTenantId(v int64) *ListKgRelationShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListKgRelationShrinkRequest) SetOpUserId(v string) *ListKgRelationShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListKgRelationShrinkRequest) SetRelationType(v string) *ListKgRelationShrinkRequest {
	s.RelationType = &v
	return s
}

func (s *ListKgRelationShrinkRequest) SetWorkspaceId(v string) *ListKgRelationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKgRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
