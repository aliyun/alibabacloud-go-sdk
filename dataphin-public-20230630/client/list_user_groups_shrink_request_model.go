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
}

type ListUserGroupsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListUserGroupsShrinkRequest) SetListQueryShrink(v string) *ListUserGroupsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) SetOpTenantId(v int64) *ListUserGroupsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
