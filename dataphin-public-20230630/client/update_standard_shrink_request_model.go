// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateStandardShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardShrinkRequest) SetOpTenantId(v int64) *UpdateStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
