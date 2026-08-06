// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListKgRelationRequestListQuery) *ListKgRelationRequest
	GetListQuery() *ListKgRelationRequestListQuery
	SetOpTenantId(v int64) *ListKgRelationRequest
	GetOpTenantId() *int64
	SetRelationType(v string) *ListKgRelationRequest
	GetRelationType() *string
	SetWorkspaceId(v string) *ListKgRelationRequest
	GetWorkspaceId() *string
}

type ListKgRelationRequest struct {
	// The paged search filter conditions.
	ListQuery *ListKgRelationRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s ListKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationRequest) GoString() string {
	return s.String()
}

func (s *ListKgRelationRequest) GetListQuery() *ListKgRelationRequestListQuery {
	return s.ListQuery
}

func (s *ListKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListKgRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ListKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKgRelationRequest) SetListQuery(v *ListKgRelationRequestListQuery) *ListKgRelationRequest {
	s.ListQuery = v
	return s
}

func (s *ListKgRelationRequest) SetOpTenantId(v int64) *ListKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListKgRelationRequest) SetRelationType(v string) *ListKgRelationRequest {
	s.RelationType = &v
	return s
}

func (s *ListKgRelationRequest) SetWorkspaceId(v string) *ListKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKgRelationRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKgRelationRequestListQuery struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source entity record ID.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf81f
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
	// The target entity record ID.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf82f
	TargetEntityId *string `json:"TargetEntityId,omitempty" xml:"TargetEntityId,omitempty"`
}

func (s ListKgRelationRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListKgRelationRequestListQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListKgRelationRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKgRelationRequestListQuery) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *ListKgRelationRequestListQuery) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *ListKgRelationRequestListQuery) SetPageNum(v int32) *ListKgRelationRequestListQuery {
	s.PageNum = &v
	return s
}

func (s *ListKgRelationRequestListQuery) SetPageSize(v int32) *ListKgRelationRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListKgRelationRequestListQuery) SetSourceEntityId(v string) *ListKgRelationRequestListQuery {
	s.SourceEntityId = &v
	return s
}

func (s *ListKgRelationRequestListQuery) SetTargetEntityId(v string) *ListKgRelationRequestListQuery {
	s.TargetEntityId = &v
	return s
}

func (s *ListKgRelationRequestListQuery) Validate() error {
	return dara.Validate(s)
}
