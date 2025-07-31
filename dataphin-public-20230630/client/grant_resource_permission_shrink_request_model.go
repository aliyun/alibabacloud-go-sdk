// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResourcePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantCommandShrink(v string) *GrantResourcePermissionShrinkRequest
	GetGrantCommandShrink() *string
	SetOpTenantId(v int64) *GrantResourcePermissionShrinkRequest
	GetOpTenantId() *int64
}

type GrantResourcePermissionShrinkRequest struct {
	// This parameter is required.
	GrantCommandShrink *string `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GrantResourcePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionShrinkRequest) GetGrantCommandShrink() *string {
	return s.GrantCommandShrink
}

func (s *GrantResourcePermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GrantResourcePermissionShrinkRequest) SetGrantCommandShrink(v string) *GrantResourcePermissionShrinkRequest {
	s.GrantCommandShrink = &v
	return s
}

func (s *GrantResourcePermissionShrinkRequest) SetOpTenantId(v int64) *GrantResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GrantResourcePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
