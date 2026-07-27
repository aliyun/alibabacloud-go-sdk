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
	SetRelationType(v string) *ListKgRelationShrinkRequest
	GetRelationType() *string
	SetWorkspaceId(v string) *ListKgRelationShrinkRequest
	GetWorkspaceId() *string
}

type ListKgRelationShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
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
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
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
