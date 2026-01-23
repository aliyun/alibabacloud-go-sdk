// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardTemplateShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateStandardTemplateShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardTemplateShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardTemplateShrinkRequest) SetOpTenantId(v int64) *UpdateStandardTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardTemplateShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardTemplateShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
