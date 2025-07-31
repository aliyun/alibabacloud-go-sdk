// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatePipelineNodeCommandShrink(v string) *CreatePipelineNodeShrinkRequest
	GetCreatePipelineNodeCommandShrink() *string
	SetOpTenantId(v int64) *CreatePipelineNodeShrinkRequest
	GetOpTenantId() *int64
}

type CreatePipelineNodeShrinkRequest struct {
	// This parameter is required.
	CreatePipelineNodeCommandShrink *string `json:"CreatePipelineNodeCommand,omitempty" xml:"CreatePipelineNodeCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreatePipelineNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeShrinkRequest) GetCreatePipelineNodeCommandShrink() *string {
	return s.CreatePipelineNodeCommandShrink
}

func (s *CreatePipelineNodeShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineNodeShrinkRequest) SetCreatePipelineNodeCommandShrink(v string) *CreatePipelineNodeShrinkRequest {
	s.CreatePipelineNodeCommandShrink = &v
	return s
}

func (s *CreatePipelineNodeShrinkRequest) SetOpTenantId(v int64) *CreatePipelineNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
