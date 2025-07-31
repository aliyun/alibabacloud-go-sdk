// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRowPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateRowPermissionCommandShrink(v string) *CreateRowPermissionShrinkRequest
	GetCreateRowPermissionCommandShrink() *string
	SetOpTenantId(v int64) *CreateRowPermissionShrinkRequest
	GetOpTenantId() *int64
}

type CreateRowPermissionShrinkRequest struct {
	// This parameter is required.
	CreateRowPermissionCommandShrink *string `json:"CreateRowPermissionCommand,omitempty" xml:"CreateRowPermissionCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateRowPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionShrinkRequest) GetCreateRowPermissionCommandShrink() *string {
	return s.CreateRowPermissionCommandShrink
}

func (s *CreateRowPermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateRowPermissionShrinkRequest) SetCreateRowPermissionCommandShrink(v string) *CreateRowPermissionShrinkRequest {
	s.CreateRowPermissionCommandShrink = &v
	return s
}

func (s *CreateRowPermissionShrinkRequest) SetOpTenantId(v int64) *CreateRowPermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateRowPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
