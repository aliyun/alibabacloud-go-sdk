// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateTenantMemberShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateTenantMemberShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateTenantMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateTenantMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateTenantMemberShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateTenantMemberShrinkRequest) SetOpTenantId(v int64) *UpdateTenantMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantMemberShrinkRequest) SetUpdateCommandShrink(v string) *UpdateTenantMemberShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateTenantMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
