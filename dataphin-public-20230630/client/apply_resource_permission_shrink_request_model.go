// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourcePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommandShrink(v string) *ApplyResourcePermissionShrinkRequest
	GetApplyCommandShrink() *string
	SetOpTenantId(v int64) *ApplyResourcePermissionShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ApplyResourcePermissionShrinkRequest
	GetOpUserId() *string
}

type ApplyResourcePermissionShrinkRequest struct {
	// The permission application request.
	//
	// This parameter is required.
	ApplyCommandShrink *string `json:"ApplyCommand,omitempty" xml:"ApplyCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ApplyResourcePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionShrinkRequest) GetApplyCommandShrink() *string {
	return s.ApplyCommandShrink
}

func (s *ApplyResourcePermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyResourcePermissionShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ApplyResourcePermissionShrinkRequest) SetApplyCommandShrink(v string) *ApplyResourcePermissionShrinkRequest {
	s.ApplyCommandShrink = &v
	return s
}

func (s *ApplyResourcePermissionShrinkRequest) SetOpTenantId(v int64) *ApplyResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyResourcePermissionShrinkRequest) SetOpUserId(v string) *ApplyResourcePermissionShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ApplyResourcePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
