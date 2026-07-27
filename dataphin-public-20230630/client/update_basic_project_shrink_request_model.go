// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBasicProjectShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateBasicProjectShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBasicProjectShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBasicProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBasicProjectShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBasicProjectShrinkRequest) SetOpTenantId(v int64) *UpdateBasicProjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBasicProjectShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBasicProjectShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBasicProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
