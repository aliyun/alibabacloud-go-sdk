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
	// The business status code. The value is SUCCESS when the operation succeeds.
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
	// The message returned for the request.
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
	// The AgenticFS additional mount list. The total number of items combined with ossMounts cannot exceed 10.
	AgenticFsMounts []*UpdateManagedAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// Contains only skills that are added or overridden by the user. Skills inherited from templates are not included. The resource model reads this field to preserve update semantics. The skills field in the request is still used for creation and update operations.
	ConfiguredSkills []*UpdateManagedAgentResponseBodyDataConfiguredSkills `json:"configuredSkills,omitempty" xml:"configuredSkills,omitempty" type:"Repeated"`
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
	Environment *UpdateManagedAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent runtime harness.
	Harness *UpdateManagedAgentResponseBodyDataHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
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
	Model *UpdateManagedAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration.
	Network *UpdateManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The OSS mount list. A maximum of 10 items are supported.
	OssMounts []*UpdateManagedAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration information.
	Runtime *UpdateManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The instance counts of the managed agent grouped by sandbox phase. Current keys: PENDING (being created or initialized), RUNNING (running), HIBERNATING (entering hibernation), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (being terminated), FAILED (runtime failure). Only phases that actually occur are returned. Missing keys are treated as 0. This field is a dynamic mapping and new keys may be added in the future. The frontend can use FAILED > 0 to determine whether abnormal instances exist.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The skill configuration list.
	Skills []*UpdateManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-agent configuration list.
	SubAgents []*UpdateManagedAgentResponseBodyDataSubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The template configuration information.
	Template *UpdateManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The tool configuration list.
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

func (s *UpdateManagedAgentResponseBodyData) GetAgenticFsMounts() []*UpdateManagedAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *UpdateManagedAgentResponseBodyData) GetConfiguredSkills() []*UpdateManagedAgentResponseBodyDataConfiguredSkills {
	return s.ConfiguredSkills
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

func (s *UpdateManagedAgentResponseBodyData) GetHarness() *UpdateManagedAgentResponseBodyDataHarness {
	return s.Harness
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

func (s *UpdateManagedAgentResponseBodyData) GetOssMounts() []*UpdateManagedAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *UpdateManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateManagedAgentResponseBodyData) GetRuntime() *UpdateManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *UpdateManagedAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
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

func (s *UpdateManagedAgentResponseBodyData) SetAgenticFsMounts(v []*UpdateManagedAgentResponseBodyDataAgenticFsMounts) *UpdateManagedAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *UpdateManagedAgentResponseBodyData) SetConfiguredSkills(v []*UpdateManagedAgentResponseBodyDataConfiguredSkills) *UpdateManagedAgentResponseBodyData {
	s.ConfiguredSkills = v
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

func (s *UpdateManagedAgentResponseBodyData) SetHarness(v *UpdateManagedAgentResponseBodyDataHarness) *UpdateManagedAgentResponseBodyData {
	s.Harness = v
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

func (s *UpdateManagedAgentResponseBodyData) SetOssMounts(v []*UpdateManagedAgentResponseBodyDataOssMounts) *UpdateManagedAgentResponseBodyData {
	s.OssMounts = v
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

func (s *UpdateManagedAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *UpdateManagedAgentResponseBodyData {
	s.SandboxPhaseCounts = v
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
	if s.AgenticFsMounts != nil {
		for _, item := range s.AgenticFsMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfiguredSkills != nil {
		for _, item := range s.ConfiguredSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type UpdateManagedAgentResponseBodyDataAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. Required for each mount entry as validated by the backend. Mount targets must not be duplicated or have parent-child overlaps.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty relative directory that exists under the AccessPoint. Required for each mount entry as validated by the backend. Root directory, absolute paths, and parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false. This is not the RAM role read-only policy.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The AccessPoint domain name. Required for each mount entry as validated by the backend. Do not include the protocol, port, or path. Use the DomainName from the NAS ListAccessPoints response.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *UpdateManagedAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *UpdateManagedAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *UpdateManagedAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *UpdateManagedAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataConfiguredSkills struct {
	// The skill name in the Workspace AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill source type. Valid values:
	//
	// - REFERENCE: references AI Registry.
	//
	// - STATIC: statically bundled with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// A legacy compatibility field. Use sourceType and versionSelector for new requests.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The referenced version selector. Defaults to LABEL/latest if omitted. Currently supports LABEL/latest.
	VersionSelector *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentResponseBodyDataConfiguredSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataConfiguredSkills) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) GetVersionSelector() *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	return s.VersionSelector
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) SetName(v string) *UpdateManagedAgentResponseBodyDataConfiguredSkills {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) SetSourceType(v string) *UpdateManagedAgentResponseBodyDataConfiguredSkills {
	s.SourceType = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) SetVersion(v string) *UpdateManagedAgentResponseBodyDataConfiguredSkills {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) SetVersionSelector(v *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) *UpdateManagedAgentResponseBodyDataConfiguredSkills {
	s.VersionSelector = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector struct {
	// The version selector type. Valid values:
	//
	// - LABEL: selects by label.
	//
	// - VERSION: selects by specific version.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The selector value. If the type is LABEL, specify a label name such as latest. If the type is VERSION, specify a specific version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetType(v string) *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetValue(v string) *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
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

type UpdateManagedAgentResponseBodyDataHarness struct {
	// The runtime harness configuration.
	Configuration *UpdateManagedAgentResponseBodyDataHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime harness type.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataHarness) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataHarness) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataHarness) GetConfiguration() *UpdateManagedAgentResponseBodyDataHarnessConfiguration {
	return s.Configuration
}

func (s *UpdateManagedAgentResponseBodyDataHarness) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentResponseBodyDataHarness) SetConfiguration(v *UpdateManagedAgentResponseBodyDataHarnessConfiguration) *UpdateManagedAgentResponseBodyDataHarness {
	s.Configuration = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataHarness) SetType(v string) *UpdateManagedAgentResponseBodyDataHarness {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataHarness) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyDataHarnessConfiguration struct {
	// The connector service account key.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The connector service account name.
	//
	// example:
	//
	// my-connector-key
	ConnectorServiceAccountName *string `json:"connectorServiceAccountName,omitempty" xml:"connectorServiceAccountName,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataHarnessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataHarnessConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountKey() *string {
	return s.ConnectorServiceAccountKey
}

func (s *UpdateManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountName() *string {
	return s.ConnectorServiceAccountName
}

func (s *UpdateManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountKey(v string) *UpdateManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountKey = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountName(v string) *UpdateManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountName = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataHarnessConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataModel struct {
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
	// The model token quota configuration and the quota usage status for the current period. This field is empty if no quota is configured.
	Quota *UpdateManagedAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
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

func (s *UpdateManagedAgentResponseBodyDataModel) GetQuota() *UpdateManagedAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *UpdateManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *UpdateManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModel) SetModelName(v string) *UpdateManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModel) SetQuota(v *UpdateManagedAgentResponseBodyDataModelQuota) *UpdateManagedAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyDataModelQuota struct {
	// Indicates whether the quota is enabled. This field is not returned if no quota is configured.
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
	// Indicates whether the quota has been exceeded in the current period. This field is read-only and returned by the backend.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The quota statistical period. A value of day indicates daily. A value of month indicates monthly.
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The gateway quota rule status. This field is read-only and returned by the backend.
	//
	// example:
	//
	// ACTIVE
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// The maximum number of tokens that can be consumed in a single period.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The number of tokens consumed in the current period. This field is read-only and returned by the backend.
	//
	// example:
	//
	// 12345
	UsedAmount *int64 `json:"usedAmount,omitempty" xml:"usedAmount,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetEnabled(v bool) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetLimitType(v string) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetPeriodType(v string) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *UpdateManagedAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataModelQuota) Validate() error {
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
	// Specifies whether to allow public network access.
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
	// Specifies whether to allow VPC access.
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

type UpdateManagedAgentResponseBodyDataOssMounts struct {
	// The OSS bucket name. Required for each mount entry as validated by the backend.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path in the container. Required for each mount entry as validated by the backend.
	//
	// example:
	//
	// /mnt/oss/datasets
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The relative object prefix within the bucket. If not specified, the entire bucket is mounted.
	//
	// example:
	//
	// datasets
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) SetBucketName(v string) *UpdateManagedAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) SetMountPath(v string) *UpdateManagedAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) SetPath(v string) *UpdateManagedAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *UpdateManagedAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *UpdateManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The sandbox auto scaling and session configuration.
	Hpa *UpdateManagedAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
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

func (s *UpdateManagedAgentResponseBodyDataRuntime) GetHpa() *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) SetCompute(v *UpdateManagedAgentResponseBodyDataRuntimeCompute) *UpdateManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntime) SetHpa(v *UpdateManagedAgentResponseBodyDataRuntimeHpa) *UpdateManagedAgentResponseBodyDataRuntime {
	s.Hpa = v
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

type UpdateManagedAgentResponseBodyDataRuntimeCompute struct {
	// The compute class.
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

type UpdateManagedAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether to enable auto scaling. Required when hpa is present, as validated by the backend.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of active sessions per sandbox. Required when hpa is present, as validated by the backend.
	//
	// example:
	//
	// 5
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of sandboxes. Required when HPA is enabled and must be no less than the minimum value.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of sandboxes. Required when HPA is enabled.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The time in seconds before an inactive session is reclaimed. Required when hpa is present, as validated by the backend.
	//
	// example:
	//
	// 3600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *UpdateManagedAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The HTTP header name used for session affinity. Takes effect when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
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
	// The version that has taken effect at runtime. This field is read-only.
	//
	// example:
	//
	// 1.0.0
	AppliedVersion *string `json:"appliedVersion,omitempty" xml:"appliedVersion,omitempty"`
	// Indicates whether the skill originates from a fixed template. This field is read-only. Template items cannot be removed.
	//
	// example:
	//
	// false
	FromTemplate *bool `json:"fromTemplate,omitempty" xml:"fromTemplate,omitempty"`
	// The skill name.
	//
	// example:
	//
	// code-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The current target version. This field is read-only.
	//
	// example:
	//
	// 1.0.0
	ResolvedVersion *string `json:"resolvedVersion,omitempty" xml:"resolvedVersion,omitempty"`
	// The skill source type. Valid values:
	//
	// - REFERENCE: references AI Registry.
	//
	// - STATIC: statically bundled with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The referenced version selector. Defaults to LABEL/latest if omitted.
	VersionSelector *UpdateManagedAgentResponseBodyDataSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetFromTemplate() *bool {
	return s.FromTemplate
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetResolvedVersion() *string {
	return s.ResolvedVersion
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentResponseBodyDataSkills) GetVersionSelector() *UpdateManagedAgentResponseBodyDataSkillsVersionSelector {
	return s.VersionSelector
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetAppliedVersion(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.AppliedVersion = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetFromTemplate(v bool) *UpdateManagedAgentResponseBodyDataSkills {
	s.FromTemplate = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetName(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetResolvedVersion(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.ResolvedVersion = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetSourceType(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.SourceType = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetVersion(v string) *UpdateManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) SetVersionSelector(v *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) *UpdateManagedAgentResponseBodyDataSkills {
	s.VersionSelector = v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentResponseBodyDataSkillsVersionSelector struct {
	// The version selector type. Valid values:
	//
	// - LABEL: selects by label.
	//
	// - VERSION: selects by specific version.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The selector value. If the type is LABEL, specify a label name such as latest. If the type is VERSION, specify a specific version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateManagedAgentResponseBodyDataSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponseBodyDataSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) SetType(v string) *UpdateManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) SetValue(v string) *UpdateManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *UpdateManagedAgentResponseBodyDataSkillsVersionSelector) Validate() error {
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
	// The AI Registry template configuration.
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
