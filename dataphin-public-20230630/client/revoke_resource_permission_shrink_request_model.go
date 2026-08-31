// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourcePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RevokeResourcePermissionShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *RevokeResourcePermissionShrinkRequest
	GetOpUserId() *string
	SetRevokeCommandShrink(v string) *RevokeResourcePermissionShrinkRequest
	GetRevokeCommandShrink() *string
}

type RevokeResourcePermissionShrinkRequest struct {
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
	// The authorization revocation request.
	//
	// This parameter is required.
	RevokeCommandShrink *string `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty"`
}

func (s RevokeResourcePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RevokeResourcePermissionShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *RevokeResourcePermissionShrinkRequest) GetRevokeCommandShrink() *string {
	return s.RevokeCommandShrink
}

func (s *RevokeResourcePermissionShrinkRequest) SetOpTenantId(v int64) *RevokeResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeResourcePermissionShrinkRequest) SetOpUserId(v string) *RevokeResourcePermissionShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *RevokeResourcePermissionShrinkRequest) SetRevokeCommandShrink(v string) *RevokeResourcePermissionShrinkRequest {
	s.RevokeCommandShrink = &v
	return s
}

func (s *RevokeResourcePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
