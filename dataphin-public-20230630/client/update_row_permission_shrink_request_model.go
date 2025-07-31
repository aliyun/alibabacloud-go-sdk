// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRowPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateRowPermissionShrinkRequest
	GetOpTenantId() *int64
	SetUpdateRowPermissionCommandShrink(v string) *UpdateRowPermissionShrinkRequest
	GetUpdateRowPermissionCommandShrink() *string
}

type UpdateRowPermissionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateRowPermissionCommandShrink *string `json:"UpdateRowPermissionCommand,omitempty" xml:"UpdateRowPermissionCommand,omitempty"`
}

func (s UpdateRowPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateRowPermissionShrinkRequest) GetUpdateRowPermissionCommandShrink() *string {
	return s.UpdateRowPermissionCommandShrink
}

func (s *UpdateRowPermissionShrinkRequest) SetOpTenantId(v int64) *UpdateRowPermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateRowPermissionShrinkRequest) SetUpdateRowPermissionCommandShrink(v string) *UpdateRowPermissionShrinkRequest {
	s.UpdateRowPermissionCommandShrink = &v
	return s
}

func (s *UpdateRowPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
