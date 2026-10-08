// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineIntegratedTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *StartPipelineIntegratedTaskShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *StartPipelineIntegratedTaskShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *StartPipelineIntegratedTaskShrinkRequest
	GetOpUserId() *string
	SetStartCommandShrink(v string) *StartPipelineIntegratedTaskShrinkRequest
	GetStartCommandShrink() *string
}

type StartPipelineIntegratedTaskShrinkRequest struct {
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
	// 30110121
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// This parameter is required.
	StartCommandShrink *string `json:"StartCommand,omitempty" xml:"StartCommand,omitempty"`
}

func (s StartPipelineIntegratedTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *StartPipelineIntegratedTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *StartPipelineIntegratedTaskShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *StartPipelineIntegratedTaskShrinkRequest) GetStartCommandShrink() *string {
	return s.StartCommandShrink
}

func (s *StartPipelineIntegratedTaskShrinkRequest) SetContextShrink(v string) *StartPipelineIntegratedTaskShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *StartPipelineIntegratedTaskShrinkRequest) SetOpTenantId(v int64) *StartPipelineIntegratedTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *StartPipelineIntegratedTaskShrinkRequest) SetOpUserId(v string) *StartPipelineIntegratedTaskShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *StartPipelineIntegratedTaskShrinkRequest) SetStartCommandShrink(v string) *StartPipelineIntegratedTaskShrinkRequest {
	s.StartCommandShrink = &v
	return s
}

func (s *StartPipelineIntegratedTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
