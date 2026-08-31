// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddTenantMembersShrinkRequest
	GetAddCommandShrink() *string
	SetOpTenantId(v int64) *AddTenantMembersShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *AddTenantMembersShrinkRequest
	GetOpUserId() *string
}

type AddTenantMembersShrinkRequest struct {
	// The request command.
	//
	// This parameter is required.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
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
}

func (s AddTenantMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddTenantMembersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddTenantMembersShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *AddTenantMembersShrinkRequest) SetAddCommandShrink(v string) *AddTenantMembersShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddTenantMembersShrinkRequest) SetOpTenantId(v int64) *AddTenantMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddTenantMembersShrinkRequest) SetOpUserId(v string) *AddTenantMembersShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *AddTenantMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
