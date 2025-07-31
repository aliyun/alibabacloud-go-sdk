// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListProjectMembersShrinkRequest
	GetId() *int64
	SetListQueryShrink(v string) *ListProjectMembersShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListProjectMembersShrinkRequest
	GetOpTenantId() *int64
}

type ListProjectMembersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListProjectMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ListProjectMembersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListProjectMembersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListProjectMembersShrinkRequest) SetId(v int64) *ListProjectMembersShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetListQueryShrink(v string) *ListProjectMembersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetOpTenantId(v int64) *ListProjectMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
