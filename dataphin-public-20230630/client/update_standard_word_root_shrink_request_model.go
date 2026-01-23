// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardWordRootShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardWordRootShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateStandardWordRootShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardWordRootShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardWordRootShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardWordRootShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardWordRootShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardWordRootShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardWordRootShrinkRequest) SetOpTenantId(v int64) *UpdateStandardWordRootShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardWordRootShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardWordRootShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardWordRootShrinkRequest) Validate() error {
	return dara.Validate(s)
}
