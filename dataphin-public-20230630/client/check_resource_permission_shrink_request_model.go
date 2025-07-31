// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourcePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommandShrink(v string) *CheckResourcePermissionShrinkRequest
	GetCheckCommandShrink() *string
	SetOpTenantId(v int64) *CheckResourcePermissionShrinkRequest
	GetOpTenantId() *int64
}

type CheckResourcePermissionShrinkRequest struct {
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckResourcePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionShrinkRequest) GetCheckCommandShrink() *string {
	return s.CheckCommandShrink
}

func (s *CheckResourcePermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckResourcePermissionShrinkRequest) SetCheckCommandShrink(v string) *CheckResourcePermissionShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckResourcePermissionShrinkRequest) SetOpTenantId(v int64) *CheckResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckResourcePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
