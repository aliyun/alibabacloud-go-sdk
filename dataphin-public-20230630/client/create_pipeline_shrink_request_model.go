// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *CreatePipelineShrinkRequest
	GetContextShrink() *string
	SetCreateCommandShrink(v string) *CreatePipelineShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreatePipelineShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreatePipelineShrinkRequest
	GetOpUserId() *string
}

type CreatePipelineShrinkRequest struct {
	// The request context information.
	//
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The configuration for creating a pipeline or workflow node.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreatePipelineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *CreatePipelineShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreatePipelineShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreatePipelineShrinkRequest) SetContextShrink(v string) *CreatePipelineShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *CreatePipelineShrinkRequest) SetCreateCommandShrink(v string) *CreatePipelineShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreatePipelineShrinkRequest) SetOpTenantId(v int64) *CreatePipelineShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineShrinkRequest) SetOpUserId(v string) *CreatePipelineShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreatePipelineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
