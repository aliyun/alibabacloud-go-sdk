// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateUserGroupShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateUserGroupShrinkRequest
	GetOpTenantId() *int64
}

type CreateUserGroupShrinkRequest struct {
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUserGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateUserGroupShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateUserGroupShrinkRequest) SetCreateCommandShrink(v string) *CreateUserGroupShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateUserGroupShrinkRequest) SetOpTenantId(v int64) *CreateUserGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateUserGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
