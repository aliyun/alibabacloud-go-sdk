// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetKgRelationRequest
	GetOpTenantId() *int64
	SetRelationId(v string) *GetKgRelationRequest
	GetRelationId() *string
	SetRelationType(v string) *GetKgRelationRequest
	GetRelationType() *string
	SetWorkspaceId(v string) *GetKgRelationRequest
	GetWorkspaceId() *string
}

type GetKgRelationRequest struct {
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
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
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

func (s GetKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgRelationRequest) GoString() string {
	return s.String()
}

func (s *GetKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgRelationRequest) GetRelationId() *string {
	return s.RelationId
}

func (s *GetKgRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *GetKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgRelationRequest) SetOpTenantId(v int64) *GetKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgRelationRequest) SetRelationId(v string) *GetKgRelationRequest {
	s.RelationId = &v
	return s
}

func (s *GetKgRelationRequest) SetRelationType(v string) *GetKgRelationRequest {
	s.RelationType = &v
	return s
}

func (s *GetKgRelationRequest) SetWorkspaceId(v string) *GetKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgRelationRequest) Validate() error {
	return dara.Validate(s)
}
