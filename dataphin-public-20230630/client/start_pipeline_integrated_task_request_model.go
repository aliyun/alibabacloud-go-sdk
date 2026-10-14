// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineIntegratedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *StartPipelineIntegratedTaskRequestContext) *StartPipelineIntegratedTaskRequest
	GetContext() *StartPipelineIntegratedTaskRequestContext
	SetOpTenantId(v int64) *StartPipelineIntegratedTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *StartPipelineIntegratedTaskRequest
	GetOpUserId() *string
	SetStartCommand(v *StartPipelineIntegratedTaskRequestStartCommand) *StartPipelineIntegratedTaskRequest
	GetStartCommand() *StartPipelineIntegratedTaskRequestStartCommand
}

type StartPipelineIntegratedTaskRequest struct {
	// The request context.
	//
	// This parameter is required.
	Context *StartPipelineIntegratedTaskRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30110121
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The startup parameters.
	//
	// This parameter is required.
	StartCommand *StartPipelineIntegratedTaskRequestStartCommand `json:"StartCommand,omitempty" xml:"StartCommand,omitempty" type:"Struct"`
}

func (s StartPipelineIntegratedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskRequest) GetContext() *StartPipelineIntegratedTaskRequestContext {
	return s.Context
}

func (s *StartPipelineIntegratedTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *StartPipelineIntegratedTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *StartPipelineIntegratedTaskRequest) GetStartCommand() *StartPipelineIntegratedTaskRequestStartCommand {
	return s.StartCommand
}

func (s *StartPipelineIntegratedTaskRequest) SetContext(v *StartPipelineIntegratedTaskRequestContext) *StartPipelineIntegratedTaskRequest {
	s.Context = v
	return s
}

func (s *StartPipelineIntegratedTaskRequest) SetOpTenantId(v int64) *StartPipelineIntegratedTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequest) SetOpUserId(v string) *StartPipelineIntegratedTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequest) SetStartCommand(v *StartPipelineIntegratedTaskRequestStartCommand) *StartPipelineIntegratedTaskRequest {
	s.StartCommand = v
	return s
}

func (s *StartPipelineIntegratedTaskRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.StartCommand != nil {
		if err := s.StartCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartPipelineIntegratedTaskRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD: Production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the integration pipeline node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s StartPipelineIntegratedTaskRequestContext) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskRequestContext) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskRequestContext) GetEnv() *string {
	return s.Env
}

func (s *StartPipelineIntegratedTaskRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StartPipelineIntegratedTaskRequestContext) SetEnv(v string) *StartPipelineIntegratedTaskRequestContext {
	s.Env = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestContext) SetProjectId(v int64) *StartPipelineIntegratedTaskRequestContext {
	s.ProjectId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestContext) Validate() error {
	return dara.Validate(s)
}

type StartPipelineIntegratedTaskRequestStartCommand struct {
	// The bandwidth limit for the offline full task. Unit: MB/s.
	//
	// example:
	//
	// 10
	ByteSpeed *int32 `json:"ByteSpeed,omitempty" xml:"ByteSpeed,omitempty"`
	// The specified time checkpoint for starting in pure incremental mode.
	//
	// example:
	//
	// 2026-09-22 15:28:31
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The concurrency for the offline full task.
	//
	// example:
	//
	// 10
	Concurrent *int32 `json:"Concurrent,omitempty" xml:"Concurrent,omitempty"`
	// The startup mode for the full table task. Valid values:
	//
	// - INIT
	//
	// - RELAY
	//
	// example:
	//
	// RELAY
	FullTaskMode *string `json:"FullTaskMode,omitempty" xml:"FullTaskMode,omitempty"`
	// The incremental instance ID required when restarting the instance.
	//
	// example:
	//
	// t_1234567890_123123
	IncrementalTaskId *string `json:"IncrementalTaskId,omitempty" xml:"IncrementalTaskId,omitempty"`
	// The memory for the offline full task. Unit: MB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the real-time node required when starting from a node.
	//
	// example:
	//
	// n_1234567890
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the resource group used by the full table instance.
	//
	// example:
	//
	// default
	QuotaGroupId *string `json:"QuotaGroupId,omitempty" xml:"QuotaGroupId,omitempty"`
	// The synchronization mode. Valid values:
	//
	// - DI_DF
	//
	// - DI
	//
	// example:
	//
	// DI_DF
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s StartPipelineIntegratedTaskRequestStartCommand) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskRequestStartCommand) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetByteSpeed() *int32 {
	return s.ByteSpeed
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetConcurrent() *int32 {
	return s.Concurrent
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetFullTaskMode() *string {
	return s.FullTaskMode
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetIncrementalTaskId() *string {
	return s.IncrementalTaskId
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetMemory() *int32 {
	return s.Memory
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetNodeId() *string {
	return s.NodeId
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetQuotaGroupId() *string {
	return s.QuotaGroupId
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) GetSyncMode() *string {
	return s.SyncMode
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetByteSpeed(v int32) *StartPipelineIntegratedTaskRequestStartCommand {
	s.ByteSpeed = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetCheckpoint(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.Checkpoint = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetConcurrent(v int32) *StartPipelineIntegratedTaskRequestStartCommand {
	s.Concurrent = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetFullTaskMode(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.FullTaskMode = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetIncrementalTaskId(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.IncrementalTaskId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetMemory(v int32) *StartPipelineIntegratedTaskRequestStartCommand {
	s.Memory = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetNodeId(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.NodeId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetQuotaGroupId(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.QuotaGroupId = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) SetSyncMode(v string) *StartPipelineIntegratedTaskRequestStartCommand {
	s.SyncMode = &v
	return s
}

func (s *StartPipelineIntegratedTaskRequestStartCommand) Validate() error {
	return dara.Validate(s)
}
