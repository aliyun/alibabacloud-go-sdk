// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardTemplateShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardTemplateShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardTemplateShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardTemplateShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardTemplateShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardTemplateShrinkRequest) SetOpTenantId(v int64) *CreateStandardTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
