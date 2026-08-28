// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExternalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateExternalAgentRequestBody) *UpdateExternalAgentRequest
	GetBody() *UpdateExternalAgentRequestBody
	SetClientToken(v string) *UpdateExternalAgentRequest
	GetClientToken() *string
}

type UpdateExternalAgentRequest struct {
	// The request body.
	Body *UpdateExternalAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// A reserved idempotency token. The backend does not guarantee idempotency in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateExternalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequest) GetBody() *UpdateExternalAgentRequestBody {
	return s.Body
}

func (s *UpdateExternalAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExternalAgentRequest) SetBody(v *UpdateExternalAgentRequestBody) *UpdateExternalAgentRequest {
	s.Body = v
	return s
}

func (s *UpdateExternalAgentRequest) SetClientToken(v string) *UpdateExternalAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExternalAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExternalAgentRequestBody struct {
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
	// The model configuration. This parameter is available only when modelSource is set to PLATFORM.
	Model *UpdateExternalAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The source of the model configuration. Valid values:
	//
	// - PLATFORM: The model configuration is parsed and distributed by the platform. You can specify the model parameter.
	//
	// - RUNTIME: The model is managed by the external runtime. You cannot specify the model parameter at the same time.
	//
	// example:
	//
	// PLATFORM
	ModelSource *string `json:"modelSource,omitempty" xml:"modelSource,omitempty"`
	// The name of the external agent.
	//
	// example:
	//
	// my-external-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of skill configurations.
	Skills []*UpdateExternalAgentRequestBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The agent template configuration.
	Template *UpdateExternalAgentRequestBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*UpdateExternalAgentRequestBodyTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s UpdateExternalAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateExternalAgentRequestBody) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateExternalAgentRequestBody) GetModel() *UpdateExternalAgentRequestBodyModel {
	return s.Model
}

func (s *UpdateExternalAgentRequestBody) GetModelSource() *string {
	return s.ModelSource
}

func (s *UpdateExternalAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentRequestBody) GetSkills() []*UpdateExternalAgentRequestBodySkills {
	return s.Skills
}

func (s *UpdateExternalAgentRequestBody) GetTemplate() *UpdateExternalAgentRequestBodyTemplate {
	return s.Template
}

func (s *UpdateExternalAgentRequestBody) GetTools() []*UpdateExternalAgentRequestBodyTools {
	return s.Tools
}

func (s *UpdateExternalAgentRequestBody) SetDescription(v string) *UpdateExternalAgentRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetInstruction(v string) *UpdateExternalAgentRequestBody {
	s.Instruction = &v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetModel(v *UpdateExternalAgentRequestBodyModel) *UpdateExternalAgentRequestBody {
	s.Model = v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetModelSource(v string) *UpdateExternalAgentRequestBody {
	s.ModelSource = &v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetName(v string) *UpdateExternalAgentRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetSkills(v []*UpdateExternalAgentRequestBodySkills) *UpdateExternalAgentRequestBody {
	s.Skills = v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetTemplate(v *UpdateExternalAgentRequestBodyTemplate) *UpdateExternalAgentRequestBody {
	s.Template = v
	return s
}

func (s *UpdateExternalAgentRequestBody) SetTools(v []*UpdateExternalAgentRequestBodyTools) *UpdateExternalAgentRequestBody {
	s.Tools = v
	return s
}

func (s *UpdateExternalAgentRequestBody) Validate() error {
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

type UpdateExternalAgentRequestBodyModel struct {
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
}

func (s UpdateExternalAgentRequestBodyModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBodyModel) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBodyModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *UpdateExternalAgentRequestBodyModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateExternalAgentRequestBodyModel) SetModelConnectionId(v string) *UpdateExternalAgentRequestBodyModel {
	s.ModelConnectionId = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyModel) SetModelName(v string) *UpdateExternalAgentRequestBodyModel {
	s.ModelName = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyModel) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentRequestBodySkills struct {
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

func (s UpdateExternalAgentRequestBodySkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBodySkills) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBodySkills) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentRequestBodySkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateExternalAgentRequestBodySkills) SetName(v string) *UpdateExternalAgentRequestBodySkills {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentRequestBodySkills) SetVersion(v string) *UpdateExternalAgentRequestBodySkills {
	s.Version = &v
	return s
}

func (s *UpdateExternalAgentRequestBodySkills) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentRequestBodyTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *UpdateExternalAgentRequestBodyTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s UpdateExternalAgentRequestBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBodyTemplate) GetAiRegistry() *UpdateExternalAgentRequestBodyTemplateAiRegistry {
	return s.AiRegistry
}

func (s *UpdateExternalAgentRequestBodyTemplate) SetAiRegistry(v *UpdateExternalAgentRequestBodyTemplateAiRegistry) *UpdateExternalAgentRequestBodyTemplate {
	s.AiRegistry = v
	return s
}

func (s *UpdateExternalAgentRequestBodyTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExternalAgentRequestBodyTemplateAiRegistry struct {
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
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateExternalAgentRequestBodyTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBodyTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBodyTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentRequestBodyTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *UpdateExternalAgentRequestBodyTemplateAiRegistry) SetName(v string) *UpdateExternalAgentRequestBodyTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyTemplateAiRegistry) SetVersion(v string) *UpdateExternalAgentRequestBodyTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentRequestBodyTools struct {
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

func (s UpdateExternalAgentRequestBodyTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentRequestBodyTools) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentRequestBodyTools) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentRequestBodyTools) GetType() *string {
	return s.Type
}

func (s *UpdateExternalAgentRequestBodyTools) SetName(v string) *UpdateExternalAgentRequestBodyTools {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyTools) SetType(v string) *UpdateExternalAgentRequestBodyTools {
	s.Type = &v
	return s
}

func (s *UpdateExternalAgentRequestBodyTools) Validate() error {
	return dara.Validate(s)
}
