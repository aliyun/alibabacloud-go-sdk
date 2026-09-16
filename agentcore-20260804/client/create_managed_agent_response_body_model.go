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
	// The business status code. The value SUCCESS is returned if the operation is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information about the created managed agent.
	Data *CreateManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The harness for the managed agent. Valid values: qwenpaw and qodercli.
	Harness *CreateManagedAgentResponseBodyDataHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
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
	Model *CreateManagedAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration information.
	Network *CreateManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS mounts. A maximum of 10 entries are supported.
	OssMounts []*CreateManagedAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration information.
	Runtime *CreateManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The number of managed agent instances grouped by sandbox phase. Current keys: PENDING (being created or initialized), RUNNING (running), HIBERNATING (entering hibernation), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (being terminated), and FAILED (runtime failure). Only phases that actually occur are returned. Missing keys are treated as 0. This field is a dynamic mapping, and new keys may be added in the future. The frontend can use FAILED > 0 to determine whether abnormal instances exist.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
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
	// The template configuration.
	Template *CreateManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The tool configuration list.
	Tools []*CreateManagedAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// The time when the managed agent was last updated, in RFC 3339 format.
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

func (s *CreateManagedAgentResponseBodyData) GetHarness() *CreateManagedAgentResponseBodyDataHarness {
	return s.Harness
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

func (s *CreateManagedAgentResponseBodyData) GetOssMounts() []*CreateManagedAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *CreateManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateManagedAgentResponseBodyData) GetRuntime() *CreateManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *CreateManagedAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
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

func (s *CreateManagedAgentResponseBodyData) SetHarness(v *CreateManagedAgentResponseBodyDataHarness) *CreateManagedAgentResponseBodyData {
	s.Harness = v
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

func (s *CreateManagedAgentResponseBodyData) SetOssMounts(v []*CreateManagedAgentResponseBodyDataOssMounts) *CreateManagedAgentResponseBodyData {
	s.OssMounts = v
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

func (s *CreateManagedAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *CreateManagedAgentResponseBodyData {
	s.SandboxPhaseCounts = v
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
	if s.Harness != nil {
		if err := s.Harness.Validate(); err != nil {
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
	if s.OssMounts != nil {
		for _, item := range s.OssMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	// The name of the environment variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable.
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

type CreateManagedAgentResponseBodyDataHarness struct {
	// The Connector binding configuration for the qodercli harness.
	Configuration *CreateManagedAgentResponseBodyDataHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The harness type. Valid values: qwenpaw and qodercli. When the type is qodercli, binding is performed based on configuration.connectorServiceAccountKey, and the name is also populated during queries.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataHarness) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataHarness) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataHarness) GetConfiguration() *CreateManagedAgentResponseBodyDataHarnessConfiguration {
	return s.Configuration
}

func (s *CreateManagedAgentResponseBodyDataHarness) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetConfiguration(v *CreateManagedAgentResponseBodyDataHarnessConfiguration) *CreateManagedAgentResponseBodyDataHarness {
	s.Configuration = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetType(v string) *CreateManagedAgentResponseBodyDataHarness {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarness) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataHarnessConfiguration struct {
	// The Key ID used to bind a Service Account Key of the QoderCLI Connector. This parameter can be omitted when only one key exists, but is required when multiple keys exist.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The Connector Key name that is populated during queries. This parameter is not used as a binding criterion during writes.
	//
	// example:
	//
	// my-connector-key
	ConnectorServiceAccountName *string `json:"connectorServiceAccountName,omitempty" xml:"connectorServiceAccountName,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataHarnessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataHarnessConfiguration) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountKey() *string {
	return s.ConnectorServiceAccountKey
}

func (s *CreateManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountName() *string {
	return s.ConnectorServiceAccountName
}

func (s *CreateManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountKey(v string) *CreateManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountKey = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountName(v string) *CreateManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountName = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarnessConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataModel struct {
	// The model connection ID.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// The upstream model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration and the quota usage status in the current cycle. This parameter is empty if no quota is configured.
	Quota *CreateManagedAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
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

func (s *CreateManagedAgentResponseBodyDataModel) GetQuota() *CreateManagedAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *CreateManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *CreateManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModel) SetModelName(v string) *CreateManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModel) SetQuota(v *CreateManagedAgentResponseBodyDataModelQuota) *CreateManagedAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataModelQuota struct {
	// Indicates whether the quota is enabled. This parameter is not returned if no quota is configured.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. Currently, only token is supported.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// Indicates whether the quota has been exceeded in the current cycle. This is a read-only field returned by the backend.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The quota statistical period. A value of day indicates a daily period. A value of month indicates a monthly period.
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The gateway quota rule status. This is a read-only field returned by the backend.
	//
	// example:
	//
	// ACTIVE
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// The maximum number of tokens that can be consumed within a single cycle.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The number of tokens consumed in the current cycle. This is a read-only field returned by the backend.
	//
	// example:
	//
	// 12345
	UsedAmount *int64 `json:"usedAmount,omitempty" xml:"usedAmount,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetLimitType(v string) *CreateManagedAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *CreateManagedAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetPeriodType(v string) *CreateManagedAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *CreateManagedAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *CreateManagedAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *CreateManagedAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataModelQuota) Validate() error {
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
	// Specifies whether to allow public network access.
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
	// Specifies whether to allow VPC access.
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

type CreateManagedAgentResponseBodyDataOssMounts struct {
	// The OSS bucket name. This parameter is required by backend validation for each mount entry.
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path inside the container. This parameter is required by backend validation for each mount entry.
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The relative object prefix within the bucket. If this parameter is not specified, the entire bucket is mounted.
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) SetBucketName(v string) *CreateManagedAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) SetMountPath(v string) *CreateManagedAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) SetPath(v string) *CreateManagedAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *CreateManagedAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The sandbox auto scaling and session configuration.
	Hpa *CreateManagedAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
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

func (s *CreateManagedAgentResponseBodyDataRuntime) GetHpa() *CreateManagedAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *CreateManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *CreateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateManagedAgentResponseBodyDataRuntime) SetCompute(v *CreateManagedAgentResponseBodyDataRuntimeCompute) *CreateManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntime) SetHpa(v *CreateManagedAgentResponseBodyDataRuntimeHpa) *CreateManagedAgentResponseBodyDataRuntime {
	s.Hpa = v
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
	if s.Hpa != nil {
		if err := s.Hpa.Validate(); err != nil {
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
	// The compute class.
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

type CreateManagedAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether to enable auto scaling. This parameter is required by backend validation when hpa is present.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of active sessions per sandbox. This parameter is required by backend validation when hpa is present.
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of sandboxes. This parameter is required when HPA is enabled and the value must be no less than the minimum value.
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of sandboxes. This parameter is required when HPA is enabled.
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The session reclamation time after inactivity, in seconds. This parameter is required by backend validation when hpa is present.
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *CreateManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *CreateManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *CreateManagedAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *CreateManagedAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The name of the HTTP header used for session affinity. This parameter takes effect when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
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
	// Review the code
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
	// This parameter is required.
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
