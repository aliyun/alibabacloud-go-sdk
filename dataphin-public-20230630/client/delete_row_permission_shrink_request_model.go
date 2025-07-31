// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteRowPermissionCommandShrink(v string) *DeleteRowPermissionShrinkRequest
	GetDeleteRowPermissionCommandShrink() *string
	SetOpTenantId(v int64) *DeleteRowPermissionShrinkRequest
	GetOpTenantId() *int64
}

type DeleteRowPermissionShrinkRequest struct {
	// This parameter is required.
	DeleteRowPermissionCommandShrink *string `json:"DeleteRowPermissionCommand,omitempty" xml:"DeleteRowPermissionCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteRowPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRowPermissionShrinkRequest) GetDeleteRowPermissionCommandShrink() *string {
	return s.DeleteRowPermissionCommandShrink
}

func (s *DeleteRowPermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteRowPermissionShrinkRequest) SetDeleteRowPermissionCommandShrink(v string) *DeleteRowPermissionShrinkRequest {
	s.DeleteRowPermissionCommandShrink = &v
	return s
}

func (s *DeleteRowPermissionShrinkRequest) SetOpTenantId(v int64) *DeleteRowPermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteRowPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
