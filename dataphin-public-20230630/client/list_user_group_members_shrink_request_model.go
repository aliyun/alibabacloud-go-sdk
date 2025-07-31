// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListUserGroupMembersShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListUserGroupMembersShrinkRequest
	GetOpTenantId() *int64
}

type ListUserGroupMembersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListUserGroupMembersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserGroupMembersShrinkRequest) SetListQueryShrink(v string) *ListUserGroupMembersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserGroupMembersShrinkRequest) SetOpTenantId(v int64) *ListUserGroupMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserGroupMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
