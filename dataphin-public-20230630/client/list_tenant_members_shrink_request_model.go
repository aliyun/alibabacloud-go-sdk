// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListTenantMembersShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListTenantMembersShrinkRequest
	GetOpTenantId() *int64
}

type ListTenantMembersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListTenantMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTenantMembersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListTenantMembersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListTenantMembersShrinkRequest) SetListQueryShrink(v string) *ListTenantMembersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListTenantMembersShrinkRequest) SetOpTenantId(v int64) *ListTenantMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTenantMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
