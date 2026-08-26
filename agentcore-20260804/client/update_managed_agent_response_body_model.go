// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateManagedAgentResponseBody
	GetCode() *string
	SetData(v *UpdateManagedAgentResponseBodyData) *UpdateManagedAgentResponseBody
	GetData() *UpdateManagedAgentResponseBodyData
	SetHttpStatusCode(v int32) *UpdateManagedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateManagedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateManagedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateManagedAgentResponseBody
	GetSuccess() *bool
}

type UpdateManagedAgentResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the managed agent.
	Data *UpdateManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
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

func (s UpdateManagedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateManagedAgentResponseBody) GetData() *UpdateManagedAgentResponseBodyData {
	return s.Data
}

func (s *UpdateManagedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateManagedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateManagedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateManagedAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateManagedAgentResponseBody) SetCode(v string) *UpdateManagedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateManagedAgentResponseBody) SetData(v *UpdateManagedAgentResponseBodyData) *UpdateManagedAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateManagedAgentResponseBody) SetHttpStatusCode(v int32) *UpdateManagedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateManagedAgentResponseBody) SetMessage(v string) *UpdateManagedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateManagedAgentResponseBody) SetRequestId(v string) *UpdateManagedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateManagedAgentResponseBody) SetSuccess(v bool) *UpdateManagedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateManagedAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyData struct {
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
	Environment *UpdateManagedAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
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
	// The latest version status.
	//
	// example:
	//
	// succeeded
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The model configuration information.
	Model *UpdateManagedAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration information.
	Network *UpdateManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration information.
	Runtime *UpdateManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*UpdateManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of sub-agent configurations.
	SubAgents []*UpdateManagedAgentResponseBodyDataSubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The template configuration information.
	Template *UpdateManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*UpdateManagedAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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

func (s UpdateManagedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateManagedAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *UpdateManagedAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateManagedAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateManagedAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateManagedAgentResponseBodyData) GetEnvironment() *UpdateManagedAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *UpdateManagedAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateManagedAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *UpdateManagedAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *UpdateManagedAgentResponseBodyData) GetModel() *UpdateManagedAgentResponseBodyDataModel {
	return s.Model
}

func (s *UpdateManagedAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyData) GetNetwork() *UpdateManagedAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *UpdateManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateManagedAgentResponseBodyData) GetRuntime() *UpdateManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *UpdateManagedAgentResponseBodyData) GetSkills() []*UpdateManagedAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *UpdateManagedAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateManagedAgentResponseBodyData) GetSubAgents() []*UpdateManagedAgentResponseBodyDataSubAgents {
	return s.SubAgents
}

func (s *UpdateManagedAgentResponseBodyData) GetTemplate() *UpdateManagedAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *UpdateManagedAgentResponseBodyData) GetTools() []*UpdateManagedAgentResponseBodyDataTools {
	return s.Tools
}

func (s *UpdateManagedAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateManagedAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateManagedAgentResponseBodyData) SetAgentId(v string) *UpdateManagedAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetCreateMode(v string) *UpdateManagedAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetCreatedAt(v string) *UpdateManagedAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetDeployType(v string) *UpdateManagedAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetDescription(v string) *UpdateManagedAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetEnvironment(v *UpdateManagedAgentResponseBodyDataEnvironment) *UpdateManagedAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetInstruction(v string) *UpdateManagedAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetLatestSpecVersion(v int64) *UpdateManagedAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetLatestVersionStatus(v string) *UpdateManagedAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetModel(v *UpdateManagedAgentResponseBodyDataModel) *UpdateManagedAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetName(v string) *UpdateManagedAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetNetwork(v *UpdateManagedAgentResponseBodyDataNetwork) *UpdateManagedAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetRegionId(v string) *UpdateManagedAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetRuntime(v *UpdateManagedAgentResponseBodyDataRuntime) *UpdateManagedAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetSkills(v []*UpdateManagedAgentResponseBodyDataSkills) *UpdateManagedAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetStatus(v string) *UpdateManagedAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetSubAgents(v []*UpdateManagedAgentResponseBodyDataSubAgents) *UpdateManagedAgentResponseBodyData {
	s.SubAgents = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetTemplate(v *UpdateManagedAgentResponseBodyDataTemplate) *UpdateManagedAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetTools(v []*UpdateManagedAgentResponseBodyDataTools) *UpdateManagedAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetUpdatedAt(v string) *UpdateManagedAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetWorkspaceId(v string) *UpdateManagedAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) Validate() error {
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

type UpdateManagedAgentResponseBodyDataEnvironment struct {
	// The list of credential references.
	CredentialReferences []*UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*UpdateManagedAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateManagedAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataEnvironment) GetCredentialReferences() []*UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *UpdateManagedAgentResponseBodyDataEnvironment) GetVariables() []*UpdateManagedAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *UpdateManagedAgentResponseBodyDataEnvironment) SetCredentialReferences(v []*UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) *UpdateManagedAgentResponseBodyDataEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataEnvironment) SetVariables(v []*UpdateManagedAgentResponseBodyDataEnvironmentVariables) *UpdateManagedAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataEnvironment) Validate() error {
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

type UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) SetCredentialId(v string) *UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataEnvironmentVariables struct {
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

func (s UpdateManagedAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentVariables) SetName(v string) *UpdateManagedAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *UpdateManagedAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataModel struct {
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

func (s UpdateManagedAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *UpdateManagedAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *UpdateManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModel) SetModelName(v string) *UpdateManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataNetwork struct {
	// The public network access configuration.
	AccessInternet *UpdateManagedAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *UpdateManagedAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataNetwork) GetAccessInternet() *UpdateManagedAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *UpdateManagedAgentResponseBodyDataNetwork) GetAccessVpc() *UpdateManagedAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *UpdateManagedAgentResponseBodyDataNetwork) SetAccessInternet(v *UpdateManagedAgentResponseBodyDataNetworkAccessInternet) *UpdateManagedAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataNetwork) SetAccessVpc(v *UpdateManagedAgentResponseBodyDataNetworkAccessVpc) *UpdateManagedAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataNetwork) Validate() error {
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

type UpdateManagedAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether to allow access to the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *UpdateManagedAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether to allow access to the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *UpdateManagedAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *UpdateManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) GetCompute() *UpdateManagedAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) SetCompute(v *UpdateManagedAgentResponseBodyDataRuntimeCompute) *UpdateManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) SetSessionPolicy(v *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) *UpdateManagedAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) Validate() error {
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

type UpdateManagedAgentResponseBodyDataRuntimeCompute struct {
	// The compute specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *UpdateManagedAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
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

func (s UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataSkills struct {
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

func (s UpdateManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetName(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetVersion(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataSubAgents struct {
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

func (s UpdateManagedAgentResponseBodyDataSubAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataSubAgents) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataSubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateManagedAgentResponseBodyDataSubAgents) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataSubAgents) SetInstruction(v string) *UpdateManagedAgentResponseBodyDataSubAgents {
	s.Instruction = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSubAgents) SetName(v string) *UpdateManagedAgentResponseBodyDataSubAgents {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSubAgents) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataTemplate struct {
	// The AI registry template configuration.
	AiRegistry *UpdateManagedAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataTemplate) GetAiRegistry() *UpdateManagedAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *UpdateManagedAgentResponseBodyDataTemplate) SetAiRegistry(v *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) *UpdateManagedAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyDataTemplateAiRegistry struct {
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

func (s UpdateManagedAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *UpdateManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *UpdateManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataTools struct {
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

func (s UpdateManagedAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentResponseBodyDataTools) SetName(v string) *UpdateManagedAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataTools) SetType(v string) *UpdateManagedAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
