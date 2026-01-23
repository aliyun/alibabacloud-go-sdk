// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardSetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardSetShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateStandardSetShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardSetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardSetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardSetShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardSetShrinkRequest) SetOpTenantId(v int64) *UpdateStandardSetShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardSetShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardSetShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardSetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
