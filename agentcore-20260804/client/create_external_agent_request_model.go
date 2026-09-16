// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateExternalAgentRequestBody) *CreateExternalAgentRequest
	GetBody() *CreateExternalAgentRequestBody
	SetClientToken(v string) *CreateExternalAgentRequest
	GetClientToken() *string
}

type CreateExternalAgentRequest struct {
	// The request body.
	Body *CreateExternalAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The reserved idempotency token. The backend does not provide idempotency guarantee in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateExternalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequest) GetBody() *CreateExternalAgentRequestBody {
	return s.Body
}

func (s *CreateExternalAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExternalAgentRequest) SetBody(v *CreateExternalAgentRequestBody) *CreateExternalAgentRequest {
	s.Body = v
	return s
}

func (s *CreateExternalAgentRequest) SetClientToken(v string) *CreateExternalAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExternalAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentRequestBody struct {
	// The description of the external agent.
	//
	// example:
	//
	// A code review agent running in the user environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The agent instruction that guides the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration. Available only when modelSource is set to PLATFORM.
	Model *CreateExternalAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The model configuration source. PLATFORM indicates that the platform parses and delivers the model configuration. RUNTIME indicates that the external runtime manages the model independently, and the model parameter cannot be specified at the same time. Valid values:
	//
	// - PLATFORM: platform model.
	//
	// - RUNTIME: runtime model.
	//
	// example:
	//
	// PLATFORM
	ModelSource *string `json:"modelSource,omitempty" xml:"modelSource,omitempty"`
	// The name of the external agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-external-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of skill configurations.
	Skills []*CreateExternalAgentRequestBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The agent template configuration.
	Template *CreateExternalAgentRequestBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*CreateExternalAgentRequestBodyTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateExternalAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBody) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateExternalAgentRequestBody) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateExternalAgentRequestBody) GetModel() *CreateExternalAgentRequestBodyModel {
	return s.Model
}

func (s *CreateExternalAgentRequestBody) GetModelSource() *string {
	return s.ModelSource
}

func (s *CreateExternalAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentRequestBody) GetSkills() []*CreateExternalAgentRequestBodySkills {
	return s.Skills
}

func (s *CreateExternalAgentRequestBody) GetTemplate() *CreateExternalAgentRequestBodyTemplate {
	return s.Template
}

func (s *CreateExternalAgentRequestBody) GetTools() []*CreateExternalAgentRequestBodyTools {
	return s.Tools
}

func (s *CreateExternalAgentRequestBody) SetDescription(v string) *CreateExternalAgentRequestBody {
	s.Description = &v
	return s
}

func (s *CreateExternalAgentRequestBody) SetInstruction(v string) *CreateExternalAgentRequestBody {
	s.Instruction = &v
	return s
}

func (s *CreateExternalAgentRequestBody) SetModel(v *CreateExternalAgentRequestBodyModel) *CreateExternalAgentRequestBody {
	s.Model = v
	return s
}

func (s *CreateExternalAgentRequestBody) SetModelSource(v string) *CreateExternalAgentRequestBody {
	s.ModelSource = &v
	return s
}

func (s *CreateExternalAgentRequestBody) SetName(v string) *CreateExternalAgentRequestBody {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentRequestBody) SetSkills(v []*CreateExternalAgentRequestBodySkills) *CreateExternalAgentRequestBody {
	s.Skills = v
	return s
}

func (s *CreateExternalAgentRequestBody) SetTemplate(v *CreateExternalAgentRequestBodyTemplate) *CreateExternalAgentRequestBody {
	s.Template = v
	return s
}

func (s *CreateExternalAgentRequestBody) SetTools(v []*CreateExternalAgentRequestBodyTools) *CreateExternalAgentRequestBody {
	s.Tools = v
	return s
}

func (s *CreateExternalAgentRequestBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExternalAgentRequestBodyModel struct {
	// The model connection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// The upstream model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration. If not specified, no quota is configured.
	Quota *CreateExternalAgentRequestBodyModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s CreateExternalAgentRequestBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodyModel) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodyModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateExternalAgentRequestBodyModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateExternalAgentRequestBodyModel) GetQuota() *CreateExternalAgentRequestBodyModelQuota {
	return s.Quota
}

func (s *CreateExternalAgentRequestBodyModel) SetModelConnectionId(v string) *CreateExternalAgentRequestBodyModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModel) SetModelName(v string) *CreateExternalAgentRequestBodyModel {
	s.ModelName = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModel) SetQuota(v *CreateExternalAgentRequestBodyModelQuota) *CreateExternalAgentRequestBodyModel {
	s.Quota = v
	return s
}

func (s *CreateExternalAgentRequestBodyModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentRequestBodyModelQuota struct {
	// Specifies whether to enable token quota. Defaults to true if not specified. Set to false to disable and delete existing quota rules.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. Required by backend validation when quota is enabled. Fixed value: token.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The quota statistical period. Required by backend validation when quota is enabled. Valid values: day (daily) and month (monthly).
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The maximum number of tokens that can be consumed within a single period. Required by backend validation when quota is enabled. The value must be greater than 0.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
}

func (s CreateExternalAgentRequestBodyModelQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodyModelQuota) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodyModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateExternalAgentRequestBodyModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateExternalAgentRequestBodyModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateExternalAgentRequestBodyModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateExternalAgentRequestBodyModelQuota) SetEnabled(v bool) *CreateExternalAgentRequestBodyModelQuota {
	s.Enabled = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModelQuota) SetLimitType(v string) *CreateExternalAgentRequestBodyModelQuota {
	s.LimitType = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModelQuota) SetPeriodType(v string) *CreateExternalAgentRequestBodyModelQuota {
	s.PeriodType = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModelQuota) SetUsageLimit(v int64) *CreateExternalAgentRequestBodyModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *CreateExternalAgentRequestBodyModelQuota) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentRequestBodySkills struct {
	// The skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateExternalAgentRequestBodySkills) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodySkills) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodySkills) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentRequestBodySkills) GetVersion() *string {
	return s.Version
}

func (s *CreateExternalAgentRequestBodySkills) SetName(v string) *CreateExternalAgentRequestBodySkills {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentRequestBodySkills) SetVersion(v string) *CreateExternalAgentRequestBodySkills {
	s.Version = &v
	return s
}

func (s *CreateExternalAgentRequestBodySkills) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentRequestBodyTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *CreateExternalAgentRequestBodyTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s CreateExternalAgentRequestBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodyTemplate) GetAiRegistry() *CreateExternalAgentRequestBodyTemplateAiRegistry {
	return s.AiRegistry
}

func (s *CreateExternalAgentRequestBodyTemplate) SetAiRegistry(v *CreateExternalAgentRequestBodyTemplateAiRegistry) *CreateExternalAgentRequestBodyTemplate {
	s.AiRegistry = v
	return s
}

func (s *CreateExternalAgentRequestBodyTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentRequestBodyTemplateAiRegistry struct {
	// The name of the template in AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateExternalAgentRequestBodyTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodyTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodyTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentRequestBodyTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *CreateExternalAgentRequestBodyTemplateAiRegistry) SetName(v string) *CreateExternalAgentRequestBodyTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentRequestBodyTemplateAiRegistry) SetVersion(v string) *CreateExternalAgentRequestBodyTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *CreateExternalAgentRequestBodyTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentRequestBodyTools struct {
	// The tool name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tool type. Valid values:
	//
	// - MCP: MCP tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateExternalAgentRequestBodyTools) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentRequestBodyTools) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentRequestBodyTools) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentRequestBodyTools) GetType() *string {
	return s.Type
}

func (s *CreateExternalAgentRequestBodyTools) SetName(v string) *CreateExternalAgentRequestBodyTools {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentRequestBodyTools) SetType(v string) *CreateExternalAgentRequestBodyTools {
	s.Type = &v
	return s
}

func (s *CreateExternalAgentRequestBodyTools) Validate() error {
	return dara.Validate(s)
}
