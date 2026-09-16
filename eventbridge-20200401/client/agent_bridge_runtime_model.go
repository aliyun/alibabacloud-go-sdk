// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentBridgeRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *AgentBridgeRuntime
	GetCu() *int32
	SetErrorCode(v string) *AgentBridgeRuntime
	GetErrorCode() *string
	SetErrorMessage(v string) *AgentBridgeRuntime
	GetErrorMessage() *string
	SetName(v string) *AgentBridgeRuntime
	GetName() *string
	SetProgress(v int32) *AgentBridgeRuntime
	GetProgress() *int32
	SetStage(v string) *AgentBridgeRuntime
	GetStage() *string
	SetStatus(v string) *AgentBridgeRuntime
	GetStatus() *string
	SetTargetCu(v int32) *AgentBridgeRuntime
	GetTargetCu() *int32
}

type AgentBridgeRuntime struct {
	// The number of CUs that last took effect successfully for the AgentBridge runtime.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The stable error code returned when a creation or specification change operation fails.
	//
	// example:
	//
	// RUNTIME_OPERATION_TIMEOUT
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The desensitized error message returned when a creation or specification change operation fails.
	//
	// example:
	//
	// Runtime operation timed out
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the AgentBridge runtime. The initial name is typically default.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The progress of the current creation or specification change operation. Valid values: 0 to 100.
	//
	// example:
	//
	// 80
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The current stage of the creation or specification change operation.
	//
	// example:
	//
	// RUNTIME_HEALTH_CHECK
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The current status of the AgentBridge runtime. RUNNING indicates that the runtime is ready and can serve queries. Valid values: CREATING, RUNNING, UPDATING, RECOVERING, CLOSED, CREATE_FAILED, and UPDATE_FAILED.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The target number of CUs during a creation, specification change, or corresponding failure state. This parameter is not returned when the runtime is running stably.
	//
	// example:
	//
	// 2
	TargetCu *int32 `json:"TargetCu,omitempty" xml:"TargetCu,omitempty"`
}

func (s AgentBridgeRuntime) String() string {
	return dara.Prettify(s)
}

func (s AgentBridgeRuntime) GoString() string {
	return s.String()
}

func (s *AgentBridgeRuntime) GetCu() *int32 {
	return s.Cu
}

func (s *AgentBridgeRuntime) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AgentBridgeRuntime) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AgentBridgeRuntime) GetName() *string {
	return s.Name
}

func (s *AgentBridgeRuntime) GetProgress() *int32 {
	return s.Progress
}

func (s *AgentBridgeRuntime) GetStage() *string {
	return s.Stage
}

func (s *AgentBridgeRuntime) GetStatus() *string {
	return s.Status
}

func (s *AgentBridgeRuntime) GetTargetCu() *int32 {
	return s.TargetCu
}

func (s *AgentBridgeRuntime) SetCu(v int32) *AgentBridgeRuntime {
	s.Cu = &v
	return s
}

func (s *AgentBridgeRuntime) SetErrorCode(v string) *AgentBridgeRuntime {
	s.ErrorCode = &v
	return s
}

func (s *AgentBridgeRuntime) SetErrorMessage(v string) *AgentBridgeRuntime {
	s.ErrorMessage = &v
	return s
}

func (s *AgentBridgeRuntime) SetName(v string) *AgentBridgeRuntime {
	s.Name = &v
	return s
}

func (s *AgentBridgeRuntime) SetProgress(v int32) *AgentBridgeRuntime {
	s.Progress = &v
	return s
}

func (s *AgentBridgeRuntime) SetStage(v string) *AgentBridgeRuntime {
	s.Stage = &v
	return s
}

func (s *AgentBridgeRuntime) SetStatus(v string) *AgentBridgeRuntime {
	s.Status = &v
	return s
}

func (s *AgentBridgeRuntime) SetTargetCu(v int32) *AgentBridgeRuntime {
	s.TargetCu = &v
	return s
}

func (s *AgentBridgeRuntime) Validate() error {
	return dara.Validate(s)
}
