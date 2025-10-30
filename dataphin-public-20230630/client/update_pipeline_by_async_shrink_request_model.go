// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineByAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *UpdatePipelineByAsyncShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *UpdatePipelineByAsyncShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdatePipelineByAsyncShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdatePipelineByAsyncShrinkRequest struct {
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

func (s UpdatePipelineByAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *UpdatePipelineByAsyncShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdatePipelineByAsyncShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdatePipelineByAsyncShrinkRequest) SetContextShrink(v string) *UpdatePipelineByAsyncShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *UpdatePipelineByAsyncShrinkRequest) SetOpTenantId(v int64) *UpdatePipelineByAsyncShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdatePipelineByAsyncShrinkRequest) SetUpdateCommandShrink(v string) *UpdatePipelineByAsyncShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdatePipelineByAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
