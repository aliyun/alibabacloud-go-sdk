// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineIntegratedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *StopPipelineIntegratedTaskRequestContext) *StopPipelineIntegratedTaskRequest
	GetContext() *StopPipelineIntegratedTaskRequestContext
	SetOpTenantId(v int64) *StopPipelineIntegratedTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *StopPipelineIntegratedTaskRequest
	GetOpUserId() *string
	SetStopCommand(v *StopPipelineIntegratedTaskRequestStopCommand) *StopPipelineIntegratedTaskRequest
	GetStopCommand() *StopPipelineIntegratedTaskRequestStopCommand
}

type StopPipelineIntegratedTaskRequest struct {
	// This parameter is required.
	Context *StopPipelineIntegratedTaskRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
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
	StopCommand *StopPipelineIntegratedTaskRequestStopCommand `json:"StopCommand,omitempty" xml:"StopCommand,omitempty" type:"Struct"`
}

func (s StopPipelineIntegratedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskRequest) GetContext() *StopPipelineIntegratedTaskRequestContext {
	return s.Context
}

func (s *StopPipelineIntegratedTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *StopPipelineIntegratedTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *StopPipelineIntegratedTaskRequest) GetStopCommand() *StopPipelineIntegratedTaskRequestStopCommand {
	return s.StopCommand
}

func (s *StopPipelineIntegratedTaskRequest) SetContext(v *StopPipelineIntegratedTaskRequestContext) *StopPipelineIntegratedTaskRequest {
	s.Context = v
	return s
}

func (s *StopPipelineIntegratedTaskRequest) SetOpTenantId(v int64) *StopPipelineIntegratedTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *StopPipelineIntegratedTaskRequest) SetOpUserId(v string) *StopPipelineIntegratedTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *StopPipelineIntegratedTaskRequest) SetStopCommand(v *StopPipelineIntegratedTaskRequestStopCommand) *StopPipelineIntegratedTaskRequest {
	s.StopCommand = v
	return s
}

func (s *StopPipelineIntegratedTaskRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.StopCommand != nil {
		if err := s.StopCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopPipelineIntegratedTaskRequestContext struct {
	// This parameter is required.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s StopPipelineIntegratedTaskRequestContext) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskRequestContext) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskRequestContext) GetEnv() *string {
	return s.Env
}

func (s *StopPipelineIntegratedTaskRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StopPipelineIntegratedTaskRequestContext) SetEnv(v string) *StopPipelineIntegratedTaskRequestContext {
	s.Env = &v
	return s
}

func (s *StopPipelineIntegratedTaskRequestContext) SetProjectId(v int64) *StopPipelineIntegratedTaskRequestContext {
	s.ProjectId = &v
	return s
}

func (s *StopPipelineIntegratedTaskRequestContext) Validate() error {
	return dara.Validate(s)
}

type StopPipelineIntegratedTaskRequestStopCommand struct {
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StopPipelineIntegratedTaskRequestStopCommand) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskRequestStopCommand) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskRequestStopCommand) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *StopPipelineIntegratedTaskRequestStopCommand) SetTaskIds(v []*string) *StopPipelineIntegratedTaskRequestStopCommand {
	s.TaskIds = v
	return s
}

func (s *StopPipelineIntegratedTaskRequestStopCommand) Validate() error {
	return dara.Validate(s)
}
