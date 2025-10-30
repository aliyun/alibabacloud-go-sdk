// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineByAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *CreatePipelineByAsyncShrinkRequest
	GetContextShrink() *string
	SetCreateCommandShrink(v string) *CreatePipelineByAsyncShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreatePipelineByAsyncShrinkRequest
	GetOpTenantId() *int64
}

type CreatePipelineByAsyncShrinkRequest struct {
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreatePipelineByAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *CreatePipelineByAsyncShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreatePipelineByAsyncShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineByAsyncShrinkRequest) SetContextShrink(v string) *CreatePipelineByAsyncShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *CreatePipelineByAsyncShrinkRequest) SetCreateCommandShrink(v string) *CreatePipelineByAsyncShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreatePipelineByAsyncShrinkRequest) SetOpTenantId(v int64) *CreatePipelineByAsyncShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineByAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
