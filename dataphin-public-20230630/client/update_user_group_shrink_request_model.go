// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateUserGroupShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateUserGroupShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateUserGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId          *int64  `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateUserGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateUserGroupShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateUserGroupShrinkRequest) SetOpTenantId(v int64) *UpdateUserGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupShrinkRequest) SetUpdateCommandShrink(v string) *UpdateUserGroupShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateUserGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
