// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTenantMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveTenantMemberShrinkRequest
	GetOpTenantId() *int64
	SetRemoveCommandShrink(v string) *RemoveTenantMemberShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveTenantMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveTenantMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTenantMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveTenantMemberShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveTenantMemberShrinkRequest) SetOpTenantId(v int64) *RemoveTenantMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveTenantMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveTenantMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveTenantMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
