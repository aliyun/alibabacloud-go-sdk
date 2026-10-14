// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOptimizeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]interface{}) *CreateOptimizeTaskRequest
	GetConfig() map[string]interface{}
	SetOptimizeTaskName(v string) *CreateOptimizeTaskRequest
	GetOptimizeTaskName() *string
	SetStatus(v string) *CreateOptimizeTaskRequest
	GetStatus() *string
	SetType(v string) *CreateOptimizeTaskRequest
	GetType() *string
	SetClientToken(v string) *CreateOptimizeTaskRequest
	GetClientToken() *string
}

type CreateOptimizeTaskRequest struct {
	// The type-specific configurations. Optional prompts for ExperienceMining and SkillOptimize only support inline Markdown. The main text is saved to inputs.prompt.content (non-empty UTF-8, up to 16 KiB).
	//
	// example:
	//
	// {"ioVersion":2,"inputs":{"experienceLibrary":{"type":"dataset","dataset":"experiences"},"trajectory":{"type":"sls","project":"my-project","logstore":"agent-trajectory","format":"atif"},"prompt":{"type":"inline","format":"markdown","content":"# Guidance\\\\nExtract evidence-backed tool contracts."}},"outputs":{"changePlan":{"type":"dataset","dataset":"experience-plans"}},"operators":["tool_contract"],"application":{"mode":"manual"}}
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the optimization task. The name of a new service task is specified by the caller and must be unique within the same AgentSpace.
	//
	// example:
	//
	// efficiency-analysis
	OptimizeTaskName *string `json:"optimizeTaskName,omitempty" xml:"optimizeTaskName,omitempty"`
	// The status of the task. Valid values: Enabled, Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the optimization task. Valid values: ExperienceMining, SkillOptimize, EfficiencyAnalysis, SOPGeneration, EvalMining, PolicyMining, and AssetMining. The memory base uses the current ContextStore and the original Worker, and is not integrated with Mining.
	//
	// example:
	//
	// EfficiencyAnalysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateOptimizeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOptimizeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOptimizeTaskRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateOptimizeTaskRequest) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *CreateOptimizeTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateOptimizeTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateOptimizeTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOptimizeTaskRequest) SetConfig(v map[string]interface{}) *CreateOptimizeTaskRequest {
	s.Config = v
	return s
}

func (s *CreateOptimizeTaskRequest) SetOptimizeTaskName(v string) *CreateOptimizeTaskRequest {
	s.OptimizeTaskName = &v
	return s
}

func (s *CreateOptimizeTaskRequest) SetStatus(v string) *CreateOptimizeTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateOptimizeTaskRequest) SetType(v string) *CreateOptimizeTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateOptimizeTaskRequest) SetClientToken(v string) *CreateOptimizeTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOptimizeTaskRequest) Validate() error {
	return dara.Validate(s)
}
