// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineIntegratedTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *StopPipelineIntegratedTaskShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *StopPipelineIntegratedTaskShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *StopPipelineIntegratedTaskShrinkRequest
	GetOpUserId() *string
	SetStopCommandShrink(v string) *StopPipelineIntegratedTaskShrinkRequest
	GetStopCommandShrink() *string
}

type StopPipelineIntegratedTaskShrinkRequest struct {
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30121101
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// This parameter is required.
	StopCommandShrink *string `json:"StopCommand,omitempty" xml:"StopCommand,omitempty"`
}

func (s StopPipelineIntegratedTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *StopPipelineIntegratedTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *StopPipelineIntegratedTaskShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *StopPipelineIntegratedTaskShrinkRequest) GetStopCommandShrink() *string {
	return s.StopCommandShrink
}

func (s *StopPipelineIntegratedTaskShrinkRequest) SetContextShrink(v string) *StopPipelineIntegratedTaskShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *StopPipelineIntegratedTaskShrinkRequest) SetOpTenantId(v int64) *StopPipelineIntegratedTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *StopPipelineIntegratedTaskShrinkRequest) SetOpUserId(v string) *StopPipelineIntegratedTaskShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *StopPipelineIntegratedTaskShrinkRequest) SetStopCommandShrink(v string) *StopPipelineIntegratedTaskShrinkRequest {
	s.StopCommandShrink = &v
	return s
}

func (s *StopPipelineIntegratedTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
