// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListUserGroupsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListUserGroupsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListUserGroupsShrinkRequest
	GetOpUserId() *string
}

type ListUserGroupsShrinkRequest struct {
	// The paged query parameters.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListUserGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListUserGroupsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserGroupsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListUserGroupsShrinkRequest) SetListQueryShrink(v string) *ListUserGroupsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) SetOpTenantId(v int64) *ListUserGroupsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) SetOpUserId(v string) *ListUserGroupsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
