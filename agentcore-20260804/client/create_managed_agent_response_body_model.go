// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManagedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateManagedAgentResponseBody
	GetCode() *string
	SetData(v *CreateManagedAgentResponseBodyData) *CreateManagedAgentResponseBody
	GetData() *CreateManagedAgentResponseBodyData
	SetHttpStatusCode(v int32) *CreateManagedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateManagedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateManagedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateManagedAgentResponseBody
	GetSuccess() *bool
}

type CreateManagedAgentResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information about the managed agent after creation.
	Data *CreateManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateManagedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateManagedAgentResponseBody) GetData() *CreateManagedAgentResponseBodyData {
	return s.Data
}

func (s *CreateManagedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateManagedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateManagedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateManagedAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateManagedAgentResponseBody) SetCode(v string) *CreateManagedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateManagedAgentResponseBody) SetData(v *CreateManagedAgentResponseBodyData) *CreateManagedAgentResponseBody {
	s.Data = v
	return s
}

func (s *CreateManagedAgentResponseBody) SetHttpStatusCode(v int32) *CreateManagedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateManagedAgentResponseBody) SetMessage(v string) *CreateManagedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateManagedAgentResponseBody) SetRequestId(v string) *CreateManagedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateManagedAgentResponseBody) SetSuccess(v bool) *CreateManagedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateManagedAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyData struct {
	// The managed agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The creation mode.
	//
	// example:
	//
	// Managed
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The creation time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment type.
	//
	// example:
	//
	// Managed
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the managed agent.
	//
	// example:
	//
	// An agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration information.
	Environment *CreateManagedAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent instruction that guides the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The latest specification version number.
	//
	// example:
	//
	// 1
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The status of the latest version.
	//
	// example:
	//
	// succeeded
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The model configuration information.
	Model *CreateManagedAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration information.
	Network *CreateManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration information.
	Runtime *CreateManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*CreateManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of sub-agent configurations.
	SubAgents []*CreateManagedAgentResponseBodyDataSubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The template configuration information.
	Template *CreateManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*CreateManagedAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateManagedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateManagedAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateManagedAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateManagedAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateManagedAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateManagedAgentResponseBodyData) GetEnvironment() *CreateManagedAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *CreateManagedAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *CreateManagedAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *CreateManagedAgentResponseBodyData) GetModel() *CreateManagedAgentResponseBodyDataModel {
	return s.Model
}

func (s *CreateManagedAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyData) GetNetwork() *CreateManagedAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *CreateManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateManagedAgentResponseBodyData) GetRuntime() *CreateManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *CreateManagedAgentResponseBodyData) GetSkills() []*CreateManagedAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *CreateManagedAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateManagedAgentResponseBodyData) GetSubAgents() []*CreateManagedAgentResponseBodyDataSubAgents {
	return s.SubAgents
}

func (s *CreateManagedAgentResponseBodyData) GetTemplate() *CreateManagedAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *CreateManagedAgentResponseBodyData) GetTools() []*CreateManagedAgentResponseBodyDataTools {
	return s.Tools
}

func (s *CreateManagedAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateManagedAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateManagedAgentResponseBodyData) SetAgentId(v string) *CreateManagedAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetCreateMode(v string) *CreateManagedAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetCreatedAt(v string) *CreateManagedAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetDeployType(v string) *CreateManagedAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetDescription(v string) *CreateManagedAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetEnvironment(v *CreateManagedAgentResponseBodyDataEnvironment) *CreateManagedAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetInstruction(v string) *CreateManagedAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetLatestSpecVersion(v int64) *CreateManagedAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetLatestVersionStatus(v string) *CreateManagedAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetModel(v *CreateManagedAgentResponseBodyDataModel) *CreateManagedAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetName(v string) *CreateManagedAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetNetwork(v *CreateManagedAgentResponseBodyDataNetwork) *CreateManagedAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetRegionId(v string) *CreateManagedAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetRuntime(v *CreateManagedAgentResponseBodyDataRuntime) *CreateManagedAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetSkills(v []*CreateManagedAgentResponseBodyDataSkills) *CreateManagedAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetStatus(v string) *CreateManagedAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetSubAgents(v []*CreateManagedAgentResponseBodyDataSubAgents) *CreateManagedAgentResponseBodyData {
	s.SubAgents = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetTemplate(v *CreateManagedAgentResponseBodyDataTemplate) *CreateManagedAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetTools(v []*CreateManagedAgentResponseBodyDataTools) *CreateManagedAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetUpdatedAt(v string) *CreateManagedAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetWorkspaceId(v string) *CreateManagedAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) Validate() error {
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

type CreateManagedAgentResponseBodyDataEnvironment struct {
	// The list of credential references.
	CredentialReferences []*CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*CreateManagedAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateManagedAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataEnvironment) GetCredentialReferences() []*CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *CreateManagedAgentResponseBodyDataEnvironment) GetVariables() []*CreateManagedAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *CreateManagedAgentResponseBodyDataEnvironment) SetCredentialReferences(v []*CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) *CreateManagedAgentResponseBodyDataEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEnvironment) SetVariables(v []*CreateManagedAgentResponseBodyDataEnvironmentVariables) *CreateManagedAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEnvironment) Validate() error {
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

type CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) SetCredentialId(v string) *CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataEnvironmentVariables struct {
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

func (s CreateManagedAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentVariables) SetName(v string) *CreateManagedAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *CreateManagedAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataModel struct {
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

func (s CreateManagedAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateManagedAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *CreateManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModel) SetModelName(v string) *CreateManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataNetwork struct {
	// The public network access configuration.
	AccessInternet *CreateManagedAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *CreateManagedAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s CreateManagedAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataNetwork) GetAccessInternet() *CreateManagedAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *CreateManagedAgentResponseBodyDataNetwork) GetAccessVpc() *CreateManagedAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *CreateManagedAgentResponseBodyDataNetwork) SetAccessInternet(v *CreateManagedAgentResponseBodyDataNetworkAccessInternet) *CreateManagedAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataNetwork) SetAccessVpc(v *CreateManagedAgentResponseBodyDataNetworkAccessVpc) *CreateManagedAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataNetwork) Validate() error {
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

type CreateManagedAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether to allow access to the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether to allow access to the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s CreateManagedAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataRuntime) GetCompute() *CreateManagedAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *CreateManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateManagedAgentResponseBodyDataRuntime) SetCompute(v *CreateManagedAgentResponseBodyDataRuntimeCompute) *CreateManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntime) SetSessionPolicy(v *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) *CreateManagedAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntime) Validate() error {
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

type CreateManagedAgentResponseBodyDataRuntimeCompute struct {
	// The compute specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *CreateManagedAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *CreateManagedAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
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

func (s CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataSkills struct {
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

func (s CreateManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetName(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetVersion(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataSubAgents struct {
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

func (s CreateManagedAgentResponseBodyDataSubAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataSubAgents) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) SetInstruction(v string) *CreateManagedAgentResponseBodyDataSubAgents {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) SetName(v string) *CreateManagedAgentResponseBodyDataSubAgents {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataTemplate struct {
	// The AI registry template configuration.
	AiRegistry *CreateManagedAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s CreateManagedAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataTemplate) GetAiRegistry() *CreateManagedAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *CreateManagedAgentResponseBodyDataTemplate) SetAiRegistry(v *CreateManagedAgentResponseBodyDataTemplateAiRegistry) *CreateManagedAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataTemplateAiRegistry struct {
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

func (s CreateManagedAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *CreateManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *CreateManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataTools struct {
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

func (s CreateManagedAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataTools) SetName(v string) *CreateManagedAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataTools) SetType(v string) *CreateManagedAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
