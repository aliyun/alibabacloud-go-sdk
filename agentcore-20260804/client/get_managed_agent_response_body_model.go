// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetManagedAgentResponseBody
	GetCode() *string
	SetData(v *GetManagedAgentResponseBodyData) *GetManagedAgentResponseBody
	GetData() *GetManagedAgentResponseBodyData
	SetHttpStatusCode(v int32) *GetManagedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetManagedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetManagedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetManagedAgentResponseBody
	GetSuccess() *bool
}

type GetManagedAgentResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the managed agent.
	Data *GetManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetManagedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetManagedAgentResponseBody) GetData() *GetManagedAgentResponseBodyData {
	return s.Data
}

func (s *GetManagedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetManagedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetManagedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetManagedAgentResponseBody) SetCode(v string) *GetManagedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetManagedAgentResponseBody) SetData(v *GetManagedAgentResponseBodyData) *GetManagedAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetManagedAgentResponseBody) SetHttpStatusCode(v int32) *GetManagedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetManagedAgentResponseBody) SetMessage(v string) *GetManagedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetManagedAgentResponseBody) SetRequestId(v string) *GetManagedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedAgentResponseBody) SetSuccess(v bool) *GetManagedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetManagedAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyData struct {
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
	// The environment configuration.
	Environment *GetManagedAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
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
	// The model configuration.
	Model *GetManagedAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration.
	Network *GetManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration.
	Runtime *GetManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*GetManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of sub-agent configurations.
	SubAgents []*GetManagedAgentResponseBodyDataSubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The template configuration.
	Template *GetManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*GetManagedAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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

func (s GetManagedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetManagedAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetManagedAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetManagedAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *GetManagedAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetManagedAgentResponseBodyData) GetEnvironment() *GetManagedAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *GetManagedAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *GetManagedAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *GetManagedAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *GetManagedAgentResponseBodyData) GetModel() *GetManagedAgentResponseBodyDataModel {
	return s.Model
}

func (s *GetManagedAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyData) GetNetwork() *GetManagedAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *GetManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetManagedAgentResponseBodyData) GetRuntime() *GetManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *GetManagedAgentResponseBodyData) GetSkills() []*GetManagedAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *GetManagedAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetManagedAgentResponseBodyData) GetSubAgents() []*GetManagedAgentResponseBodyDataSubAgents {
	return s.SubAgents
}

func (s *GetManagedAgentResponseBodyData) GetTemplate() *GetManagedAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *GetManagedAgentResponseBodyData) GetTools() []*GetManagedAgentResponseBodyDataTools {
	return s.Tools
}

func (s *GetManagedAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetManagedAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetManagedAgentResponseBodyData) SetAgentId(v string) *GetManagedAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetCreateMode(v string) *GetManagedAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetCreatedAt(v string) *GetManagedAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetDeployType(v string) *GetManagedAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetDescription(v string) *GetManagedAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetEnvironment(v *GetManagedAgentResponseBodyDataEnvironment) *GetManagedAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetInstruction(v string) *GetManagedAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetLatestSpecVersion(v int64) *GetManagedAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetLatestVersionStatus(v string) *GetManagedAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetModel(v *GetManagedAgentResponseBodyDataModel) *GetManagedAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetName(v string) *GetManagedAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetNetwork(v *GetManagedAgentResponseBodyDataNetwork) *GetManagedAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetRegionId(v string) *GetManagedAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetRuntime(v *GetManagedAgentResponseBodyDataRuntime) *GetManagedAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetSkills(v []*GetManagedAgentResponseBodyDataSkills) *GetManagedAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetStatus(v string) *GetManagedAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetSubAgents(v []*GetManagedAgentResponseBodyDataSubAgents) *GetManagedAgentResponseBodyData {
	s.SubAgents = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetTemplate(v *GetManagedAgentResponseBodyDataTemplate) *GetManagedAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetTools(v []*GetManagedAgentResponseBodyDataTools) *GetManagedAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetUpdatedAt(v string) *GetManagedAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetWorkspaceId(v string) *GetManagedAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetManagedAgentResponseBodyData) Validate() error {
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

type GetManagedAgentResponseBodyDataEnvironment struct {
	// The list of credential references.
	CredentialReferences []*GetManagedAgentResponseBodyDataEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*GetManagedAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetManagedAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataEnvironment) GetCredentialReferences() []*GetManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *GetManagedAgentResponseBodyDataEnvironment) GetVariables() []*GetManagedAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *GetManagedAgentResponseBodyDataEnvironment) SetCredentialReferences(v []*GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) *GetManagedAgentResponseBodyDataEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *GetManagedAgentResponseBodyDataEnvironment) SetVariables(v []*GetManagedAgentResponseBodyDataEnvironmentVariables) *GetManagedAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *GetManagedAgentResponseBodyDataEnvironment) Validate() error {
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

type GetManagedAgentResponseBodyDataEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) SetCredentialId(v string) *GetManagedAgentResponseBodyDataEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataEnvironmentVariables struct {
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

func (s GetManagedAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *GetManagedAgentResponseBodyDataEnvironmentVariables) SetName(v string) *GetManagedAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *GetManagedAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataModel struct {
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

func (s GetManagedAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *GetManagedAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *GetManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *GetManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModel) SetModelName(v string) *GetManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataNetwork struct {
	// The public network access configuration.
	AccessInternet *GetManagedAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *GetManagedAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s GetManagedAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataNetwork) GetAccessInternet() *GetManagedAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *GetManagedAgentResponseBodyDataNetwork) GetAccessVpc() *GetManagedAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *GetManagedAgentResponseBodyDataNetwork) SetAccessInternet(v *GetManagedAgentResponseBodyDataNetworkAccessInternet) *GetManagedAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *GetManagedAgentResponseBodyDataNetwork) SetAccessVpc(v *GetManagedAgentResponseBodyDataNetworkAccessVpc) *GetManagedAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *GetManagedAgentResponseBodyDataNetwork) Validate() error {
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

type GetManagedAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether public network access is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetManagedAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *GetManagedAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether VPC access is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetManagedAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *GetManagedAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *GetManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *GetManagedAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s GetManagedAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataRuntime) GetCompute() *GetManagedAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *GetManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *GetManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *GetManagedAgentResponseBodyDataRuntime) SetCompute(v *GetManagedAgentResponseBodyDataRuntimeCompute) *GetManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntime) SetSessionPolicy(v *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) *GetManagedAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntime) Validate() error {
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

type GetManagedAgentResponseBodyDataRuntimeCompute struct {
	// The compute specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s GetManagedAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *GetManagedAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *GetManagedAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
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

func (s GetManagedAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *GetManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *GetManagedAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataSkills struct {
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

func (s GetManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *GetManagedAgentResponseBodyDataSkills) SetName(v string) *GetManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetVersion(v string) *GetManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataSubAgents struct {
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

func (s GetManagedAgentResponseBodyDataSubAgents) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataSubAgents) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataSubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *GetManagedAgentResponseBodyDataSubAgents) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataSubAgents) SetInstruction(v string) *GetManagedAgentResponseBodyDataSubAgents {
	s.Instruction = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgents) SetName(v string) *GetManagedAgentResponseBodyDataSubAgents {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgents) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataTemplate struct {
	// The AI registry template configuration.
	AiRegistry *GetManagedAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s GetManagedAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataTemplate) GetAiRegistry() *GetManagedAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *GetManagedAgentResponseBodyDataTemplate) SetAiRegistry(v *GetManagedAgentResponseBodyDataTemplateAiRegistry) *GetManagedAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *GetManagedAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataTemplateAiRegistry struct {
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

func (s GetManagedAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *GetManagedAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *GetManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *GetManagedAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataTools struct {
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

func (s GetManagedAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *GetManagedAgentResponseBodyDataTools) SetName(v string) *GetManagedAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataTools) SetType(v string) *GetManagedAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
