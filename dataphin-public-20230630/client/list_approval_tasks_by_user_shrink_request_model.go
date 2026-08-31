// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalTasksByUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListApprovalTasksByUserShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListApprovalTasksByUserShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListApprovalTasksByUserShrinkRequest
	GetOpUserId() *string
}

type ListApprovalTasksByUserShrinkRequest struct {
	// The query conditions.
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
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListApprovalTasksByUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListApprovalTasksByUserShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListApprovalTasksByUserShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListApprovalTasksByUserShrinkRequest) SetListQueryShrink(v string) *ListApprovalTasksByUserShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListApprovalTasksByUserShrinkRequest) SetOpTenantId(v int64) *ListApprovalTasksByUserShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListApprovalTasksByUserShrinkRequest) SetOpUserId(v string) *ListApprovalTasksByUserShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListApprovalTasksByUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
