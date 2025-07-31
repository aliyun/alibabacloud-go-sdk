// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListResourcePermissionsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListResourcePermissionsShrinkRequest
	GetOpTenantId() *int64
}

type ListResourcePermissionsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListResourcePermissionsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListResourcePermissionsShrinkRequest) SetListQueryShrink(v string) *ListResourcePermissionsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListResourcePermissionsShrinkRequest) SetOpTenantId(v int64) *ListResourcePermissionsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListResourcePermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
