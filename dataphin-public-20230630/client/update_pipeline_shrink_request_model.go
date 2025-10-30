// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *UpdatePipelineShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *UpdatePipelineShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdatePipelineShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdatePipelineShrinkRequest struct {
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdatePipelineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *UpdatePipelineShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdatePipelineShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdatePipelineShrinkRequest) SetContextShrink(v string) *UpdatePipelineShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *UpdatePipelineShrinkRequest) SetOpTenantId(v int64) *UpdatePipelineShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdatePipelineShrinkRequest) SetUpdateCommandShrink(v string) *UpdatePipelineShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdatePipelineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
