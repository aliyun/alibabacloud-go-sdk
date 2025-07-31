// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersBySourceUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddTenantMembersBySourceUserShrinkRequest
	GetAddCommandShrink() *string
	SetOpTenantId(v int64) *AddTenantMembersBySourceUserShrinkRequest
	GetOpTenantId() *int64
}

type AddTenantMembersBySourceUserShrinkRequest struct {
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersBySourceUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddTenantMembersBySourceUserShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddTenantMembersBySourceUserShrinkRequest) SetAddCommandShrink(v string) *AddTenantMembersBySourceUserShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddTenantMembersBySourceUserShrinkRequest) SetOpTenantId(v int64) *AddTenantMembersBySourceUserShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddTenantMembersBySourceUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
