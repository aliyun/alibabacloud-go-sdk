// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManagedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateManagedAgentRequestBody) *CreateManagedAgentRequest
	GetBody() *CreateManagedAgentRequestBody
	SetClientToken(v string) *CreateManagedAgentRequest
	GetClientToken() *string
}

type CreateManagedAgentRequest struct {
	// The request body.
	Body *CreateManagedAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The reserved idempotency token. The backend does not provide idempotency guarantees in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateManagedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequest) GetBody() *CreateManagedAgentRequestBody {
	return s.Body
}

func (s *CreateManagedAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateManagedAgentRequest) SetBody(v *CreateManagedAgentRequestBody) *CreateManagedAgentRequest {
	s.Body = v
	return s
}

func (s *CreateManagedAgentRequest) SetClientToken(v string) *CreateManagedAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateManagedAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBody struct {
	// The description of the managed agent.
	//
	// example:
	//
	// An agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration.
	Environment *CreateManagedAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent instruction that guides the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration.
	//
	// This parameter is required.
	Model *CreateManagedAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration.
	Network *CreateManagedAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The runtime configuration.
	//
	// This parameter is required.
	Runtime *CreateManagedAgentRequestBodyRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*CreateManagedAgentRequestBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The list of sub-agent configurations.
	SubAgents []*CreateManagedAgentRequestBodySubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The agent template configuration.
	Template *CreateManagedAgentRequestBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*CreateManagedAgentRequestBodyTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateManagedAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBody) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateManagedAgentRequestBody) GetEnvironment() *CreateManagedAgentRequestBodyEnvironment {
	return s.Environment
}

func (s *CreateManagedAgentRequestBody) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentRequestBody) GetModel() *CreateManagedAgentRequestBodyModel {
	return s.Model
}

func (s *CreateManagedAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBody) GetNetwork() *CreateManagedAgentRequestBodyNetwork {
	return s.Network
}

func (s *CreateManagedAgentRequestBody) GetRuntime() *CreateManagedAgentRequestBodyRuntime {
	return s.Runtime
}

func (s *CreateManagedAgentRequestBody) GetSkills() []*CreateManagedAgentRequestBodySkills {
	return s.Skills
}

func (s *CreateManagedAgentRequestBody) GetSubAgents() []*CreateManagedAgentRequestBodySubAgents {
	return s.SubAgents
}

func (s *CreateManagedAgentRequestBody) GetTemplate() *CreateManagedAgentRequestBodyTemplate {
	return s.Template
}

func (s *CreateManagedAgentRequestBody) GetTools() []*CreateManagedAgentRequestBodyTools {
	return s.Tools
}

func (s *CreateManagedAgentRequestBody) SetDescription(v string) *CreateManagedAgentRequestBody {
	s.Description = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetEnvironment(v *CreateManagedAgentRequestBodyEnvironment) *CreateManagedAgentRequestBody {
	s.Environment = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetInstruction(v string) *CreateManagedAgentRequestBody {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetModel(v *CreateManagedAgentRequestBodyModel) *CreateManagedAgentRequestBody {
	s.Model = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetName(v string) *CreateManagedAgentRequestBody {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetNetwork(v *CreateManagedAgentRequestBodyNetwork) *CreateManagedAgentRequestBody {
	s.Network = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetRuntime(v *CreateManagedAgentRequestBodyRuntime) *CreateManagedAgentRequestBody {
	s.Runtime = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetSkills(v []*CreateManagedAgentRequestBodySkills) *CreateManagedAgentRequestBody {
	s.Skills = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetSubAgents(v []*CreateManagedAgentRequestBodySubAgents) *CreateManagedAgentRequestBody {
	s.SubAgents = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetTemplate(v *CreateManagedAgentRequestBodyTemplate) *CreateManagedAgentRequestBody {
	s.Template = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetTools(v []*CreateManagedAgentRequestBodyTools) *CreateManagedAgentRequestBody {
	s.Tools = v
	return s
}

func (s *CreateManagedAgentRequestBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
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
	if s.SubAgents != nil {
		for _, item := range s.SubAgents {
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

type CreateManagedAgentRequestBodyEnvironment struct {
	// The list of credential references.
	CredentialReferences []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*CreateManagedAgentRequestBodyEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateManagedAgentRequestBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironment) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironment) GetCredentialReferences() []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *CreateManagedAgentRequestBodyEnvironment) GetVariables() []*CreateManagedAgentRequestBodyEnvironmentVariables {
	return s.Variables
}

func (s *CreateManagedAgentRequestBodyEnvironment) SetCredentialReferences(v []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences) *CreateManagedAgentRequestBodyEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironment) SetVariables(v []*CreateManagedAgentRequestBodyEnvironmentVariables) *CreateManagedAgentRequestBodyEnvironment {
	s.Variables = v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironment) Validate() error {
	if s.CredentialReferences != nil {
		for _, item := range s.CredentialReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s CreateManagedAgentRequestBodyEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) SetCredentialId(v string) *CreateManagedAgentRequestBodyEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyEnvironmentVariables struct {
	// The environment variable name.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The environment variable value.
	//
	// This parameter is required.
	//
	// example:
	//
	// sk-xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateManagedAgentRequestBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) SetName(v string) *CreateManagedAgentRequestBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) SetValue(v string) *CreateManagedAgentRequestBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyModel struct {
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

func (s CreateManagedAgentRequestBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyModel) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateManagedAgentRequestBodyModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateManagedAgentRequestBodyModel) SetModelConnectionId(v string) *CreateManagedAgentRequestBodyModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModel) SetModelName(v string) *CreateManagedAgentRequestBodyModel {
	s.ModelName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModel) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyNetwork struct {
	// The public network access configuration.
	AccessInternet *CreateManagedAgentRequestBodyNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *CreateManagedAgentRequestBodyNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetwork) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetwork) GetAccessInternet() *CreateManagedAgentRequestBodyNetworkAccessInternet {
	return s.AccessInternet
}

func (s *CreateManagedAgentRequestBodyNetwork) GetAccessVpc() *CreateManagedAgentRequestBodyNetworkAccessVpc {
	return s.AccessVpc
}

func (s *CreateManagedAgentRequestBodyNetwork) SetAccessInternet(v *CreateManagedAgentRequestBodyNetworkAccessInternet) *CreateManagedAgentRequestBodyNetwork {
	s.AccessInternet = v
	return s
}

func (s *CreateManagedAgentRequestBodyNetwork) SetAccessVpc(v *CreateManagedAgentRequestBodyNetworkAccessVpc) *CreateManagedAgentRequestBodyNetwork {
	s.AccessVpc = v
	return s
}

func (s *CreateManagedAgentRequestBodyNetwork) Validate() error {
	if s.AccessInternet != nil {
		if err := s.AccessInternet.Validate(); err != nil {
			return err
		}
	}
	if s.AccessVpc != nil {
		if err := s.AccessVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyNetworkAccessInternet struct {
	// Specifies whether to allow access to the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentRequestBodyNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) SetEnabled(v bool) *CreateManagedAgentRequestBodyNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyNetworkAccessVpc struct {
	// Specifies whether to allow access to the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentRequestBodyNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) SetEnabled(v bool) *CreateManagedAgentRequestBodyNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *CreateManagedAgentRequestBodyRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntime) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntime) GetCompute() *CreateManagedAgentRequestBodyRuntimeCompute {
	return s.Compute
}

func (s *CreateManagedAgentRequestBodyRuntime) GetSessionPolicy() *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateManagedAgentRequestBodyRuntime) SetCompute(v *CreateManagedAgentRequestBodyRuntimeCompute) *CreateManagedAgentRequestBodyRuntime {
	s.Compute = v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntime) SetSessionPolicy(v *CreateManagedAgentRequestBodyRuntimeSessionPolicy) *CreateManagedAgentRequestBodyRuntime {
	s.SessionPolicy = v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntime) Validate() error {
	if s.Compute != nil {
		if err := s.Compute.Validate(); err != nil {
			return err
		}
	}
	if s.SessionPolicy != nil {
		if err := s.SessionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyRuntimeCompute struct {
	// The compute specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s CreateManagedAgentRequestBodyRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntimeCompute) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) SetComputeClass(v string) *CreateManagedAgentRequestBodyRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntimeSessionPolicy struct {
	// The HTTP header name used for session affinity. This parameter takes effect only when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISABLED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) SetHeaderName(v string) *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) SetType(v string) *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodySkills struct {
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

func (s CreateManagedAgentRequestBodySkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodySkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentRequestBodySkills) SetName(v string) *CreateManagedAgentRequestBodySkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) SetVersion(v string) *CreateManagedAgentRequestBodySkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodySubAgents struct {
	// The sub-agent instruction.
	//
	// This parameter is required.
	//
	// example:
	//
	// Please review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The sub-agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateManagedAgentRequestBodySubAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySubAgents) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentRequestBodySubAgents) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodySubAgents) SetInstruction(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) SetName(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyTemplate struct {
	// The AI registry template configuration.
	AiRegistry *CreateManagedAgentRequestBodyTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTemplate) GetAiRegistry() *CreateManagedAgentRequestBodyTemplateAiRegistry {
	return s.AiRegistry
}

func (s *CreateManagedAgentRequestBodyTemplate) SetAiRegistry(v *CreateManagedAgentRequestBodyTemplateAiRegistry) *CreateManagedAgentRequestBodyTemplate {
	s.AiRegistry = v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyTemplateAiRegistry struct {
	// The name of the template in the AI registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in the AI registry.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateManagedAgentRequestBodyTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) SetName(v string) *CreateManagedAgentRequestBodyTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) SetVersion(v string) *CreateManagedAgentRequestBodyTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyTools struct {
	// The tool name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tool type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyTools) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTools) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTools) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyTools) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyTools) SetName(v string) *CreateManagedAgentRequestBodyTools {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTools) SetType(v string) *CreateManagedAgentRequestBodyTools {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTools) Validate() error {
	return dara.Validate(s)
}
