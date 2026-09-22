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
	// The business status code. The value is SUCCESS when the operation succeeds.
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
	// The list of additional AgenticFS mounts. The total number of AgenticFS mounts and OSS mounts cannot exceed 10.
	AgenticFsMounts []*GetManagedAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The skills that are explicitly added or overridden by the user. This field does not include skills inherited from the template. The resource model reads this field to preserve update semantics. The skills field in the request is still used for create and update operations.
	ConfiguredSkills []*GetManagedAgentResponseBodyDataConfiguredSkills `json:"configuredSkills,omitempty" xml:"configuredSkills,omitempty" type:"Repeated"`
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
	// The agent runtime harness.
	Harness *GetManagedAgentResponseBodyDataHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
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
	// The list of OSS mounts. A maximum of 10 items are allowed.
	OssMounts []*GetManagedAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration.
	Runtime *GetManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The instance counts of the managed agent grouped by sandbox phase. Current keys: PENDING (creating or initializing), RUNNING (running), HIBERNATING (hibernating), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (terminating), FAILED (runtime failure). Only phases that actually occur are returned. Missing keys should be treated as 0. This field is a dynamic map and new keys may be added in the future. Use FAILED > 0 to determine whether abnormal instances exist.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The list of skill configurations.
	Skills []*GetManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of child agent configurations.
	SubAgents []*GetManagedAgentResponseBodyDataSubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The template configuration.
	Template *GetManagedAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*GetManagedAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// The time when the resource was last updated, in RFC 3339 format.
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

func (s *GetManagedAgentResponseBodyData) GetAgenticFsMounts() []*GetManagedAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *GetManagedAgentResponseBodyData) GetConfiguredSkills() []*GetManagedAgentResponseBodyDataConfiguredSkills {
	return s.ConfiguredSkills
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

func (s *GetManagedAgentResponseBodyData) GetHarness() *GetManagedAgentResponseBodyDataHarness {
	return s.Harness
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

func (s *GetManagedAgentResponseBodyData) GetOssMounts() []*GetManagedAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *GetManagedAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetManagedAgentResponseBodyData) GetRuntime() *GetManagedAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *GetManagedAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
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

func (s *GetManagedAgentResponseBodyData) SetAgenticFsMounts(v []*GetManagedAgentResponseBodyDataAgenticFsMounts) *GetManagedAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *GetManagedAgentResponseBodyData) SetConfiguredSkills(v []*GetManagedAgentResponseBodyDataConfiguredSkills) *GetManagedAgentResponseBodyData {
	s.ConfiguredSkills = v
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

func (s *GetManagedAgentResponseBodyData) SetHarness(v *GetManagedAgentResponseBodyDataHarness) *GetManagedAgentResponseBodyData {
	s.Harness = v
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

func (s *GetManagedAgentResponseBodyData) SetOssMounts(v []*GetManagedAgentResponseBodyDataOssMounts) *GetManagedAgentResponseBodyData {
	s.OssMounts = v
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

func (s *GetManagedAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *GetManagedAgentResponseBodyData {
	s.SandboxPhaseCounts = v
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

type GetManagedAgentResponseBodyDataAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. This field is required for each mount item as validated by the backend. Mount targets must not be duplicated or have parent-child overlaps.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty relative directory that exists under the AccessPoint. This field is required for each mount item as validated by the backend. Root directories, absolute paths, and parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false. This is not a RAM role read-only policy.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The AccessPoint domain name. This field is required for each mount item as validated by the backend. The value does not include protocol, port, or path. Use the DomainName from the NAS ListAccessPoints response.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s GetManagedAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *GetManagedAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *GetManagedAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *GetManagedAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *GetManagedAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataConfiguredSkills struct {
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
	// - REFERENCE: references the AI Registry.
	//
	// - STATIC: statically provided with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// **[Deprecated]*	- Legacy compatibility field. Use sourceType and versionSelector for new requests.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version selector for the referenced skill. If omitted, the default value is LABEL/latest. Currently, only LABEL/latest is supported.
	VersionSelector *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s GetManagedAgentResponseBodyDataConfiguredSkills) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataConfiguredSkills) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) GetVersion() *string {
	return s.Version
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) GetVersionSelector() *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	return s.VersionSelector
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) SetName(v string) *GetManagedAgentResponseBodyDataConfiguredSkills {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) SetSourceType(v string) *GetManagedAgentResponseBodyDataConfiguredSkills {
	s.SourceType = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) SetVersion(v string) *GetManagedAgentResponseBodyDataConfiguredSkills {
	s.Version = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) SetVersionSelector(v *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) *GetManagedAgentResponseBodyDataConfiguredSkills {
	s.VersionSelector = v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector struct {
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
	// The selector value. When the type is LABEL, this is the label name (such as latest). When the type is VERSION, this is the specific version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetType(v string) *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetValue(v string) *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
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

type GetManagedAgentResponseBodyDataHarness struct {
	// The runtime harness configuration.
	Configuration *GetManagedAgentResponseBodyDataHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime harness type.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetManagedAgentResponseBodyDataHarness) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataHarness) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataHarness) GetConfiguration() *GetManagedAgentResponseBodyDataHarnessConfiguration {
	return s.Configuration
}

func (s *GetManagedAgentResponseBodyDataHarness) GetType() *string {
	return s.Type
}

func (s *GetManagedAgentResponseBodyDataHarness) SetConfiguration(v *GetManagedAgentResponseBodyDataHarnessConfiguration) *GetManagedAgentResponseBodyDataHarness {
	s.Configuration = v
	return s
}

func (s *GetManagedAgentResponseBodyDataHarness) SetType(v string) *GetManagedAgentResponseBodyDataHarness {
	s.Type = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataHarness) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataHarnessConfiguration struct {
	// The Connector Service Account Key.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The Connector Service Account Name.
	//
	// example:
	//
	// my-connector-key
	ConnectorServiceAccountName *string `json:"connectorServiceAccountName,omitempty" xml:"connectorServiceAccountName,omitempty"`
}

func (s GetManagedAgentResponseBodyDataHarnessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataHarnessConfiguration) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountKey() *string {
	return s.ConnectorServiceAccountKey
}

func (s *GetManagedAgentResponseBodyDataHarnessConfiguration) GetConnectorServiceAccountName() *string {
	return s.ConnectorServiceAccountName
}

func (s *GetManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountKey(v string) *GetManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountKey = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataHarnessConfiguration) SetConnectorServiceAccountName(v string) *GetManagedAgentResponseBodyDataHarnessConfiguration {
	s.ConnectorServiceAccountName = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataHarnessConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataModel struct {
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
	Quota *GetManagedAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
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

func (s *GetManagedAgentResponseBodyDataModel) GetQuota() *GetManagedAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *GetManagedAgentResponseBodyDataModel) SetModelConnectionId(v string) *GetManagedAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModel) SetModelName(v string) *GetManagedAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModel) SetQuota(v *GetManagedAgentResponseBodyDataModelQuota) *GetManagedAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *GetManagedAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataModelQuota struct {
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
	// The quota statistical period. Valid values:
	//
	// - day: daily.
	//
	// - month: monthly.
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

func (s GetManagedAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *GetManagedAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetEnabled(v bool) *GetManagedAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetLimitType(v string) *GetManagedAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *GetManagedAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetPeriodType(v string) *GetManagedAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *GetManagedAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *GetManagedAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *GetManagedAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataModelQuota) Validate() error {
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

type GetManagedAgentResponseBodyDataOssMounts struct {
	// The OSS bucket name. This field is required for each mount item as validated by the backend.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path inside the container. This field is validated as required by the backend for each mount entry.
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
	// Specifies whether to mount as read-only. Default value: false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetManagedAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *GetManagedAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *GetManagedAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *GetManagedAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetManagedAgentResponseBodyDataOssMounts) SetBucketName(v string) *GetManagedAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataOssMounts) SetMountPath(v string) *GetManagedAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataOssMounts) SetPath(v string) *GetManagedAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *GetManagedAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *GetManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The sandbox auto scaling and session configuration.
	Hpa *GetManagedAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
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

func (s *GetManagedAgentResponseBodyDataRuntime) GetHpa() *GetManagedAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *GetManagedAgentResponseBodyDataRuntime) GetSessionPolicy() *GetManagedAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *GetManagedAgentResponseBodyDataRuntime) SetCompute(v *GetManagedAgentResponseBodyDataRuntimeCompute) *GetManagedAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntime) SetHpa(v *GetManagedAgentResponseBodyDataRuntimeHpa) *GetManagedAgentResponseBodyDataRuntime {
	s.Hpa = v
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

type GetManagedAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether to enable auto scaling. This field is validated as required by the backend when hpa is present.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of active sessions per sandbox. This field is validated as required by the backend when hpa is present.
	//
	// example:
	//
	// 5
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of sandboxes. This field is required when HPA is enabled and must be no less than the minimum value.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of sandboxes. This field is required when HPA is enabled.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The time-to-live in seconds for a session after inactivity. This field is validated as required by the backend when hpa is present.
	//
	// example:
	//
	// 3600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s GetManagedAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *GetManagedAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *GetManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *GetManagedAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *GetManagedAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *GetManagedAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The HTTP header name used for session affinity. This field takes effect when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
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
	// The version currently in effect at runtime. This field is read-only.
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
	// - REFERENCE: references the AI Registry.
	//
	// - STATIC: statically provided with the package.
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
	VersionSelector *GetManagedAgentResponseBodyDataSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s GetManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataSkills) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *GetManagedAgentResponseBodyDataSkills) GetFromTemplate() *bool {
	return s.FromTemplate
}

func (s *GetManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataSkills) GetResolvedVersion() *string {
	return s.ResolvedVersion
}

func (s *GetManagedAgentResponseBodyDataSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *GetManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *GetManagedAgentResponseBodyDataSkills) GetVersionSelector() *GetManagedAgentResponseBodyDataSkillsVersionSelector {
	return s.VersionSelector
}

func (s *GetManagedAgentResponseBodyDataSkills) SetAppliedVersion(v string) *GetManagedAgentResponseBodyDataSkills {
	s.AppliedVersion = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetFromTemplate(v bool) *GetManagedAgentResponseBodyDataSkills {
	s.FromTemplate = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetName(v string) *GetManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetResolvedVersion(v string) *GetManagedAgentResponseBodyDataSkills {
	s.ResolvedVersion = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetSourceType(v string) *GetManagedAgentResponseBodyDataSkills {
	s.SourceType = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetVersion(v string) *GetManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) SetVersionSelector(v *GetManagedAgentResponseBodyDataSkillsVersionSelector) *GetManagedAgentResponseBodyDataSkills {
	s.VersionSelector = v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataSkillsVersionSelector struct {
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
	// The selector value. When the type is LABEL, this is the label name (such as latest). When the type is VERSION, this is the specific version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetManagedAgentResponseBodyDataSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *GetManagedAgentResponseBodyDataSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *GetManagedAgentResponseBodyDataSkillsVersionSelector) SetType(v string) *GetManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkillsVersionSelector) SetValue(v string) *GetManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type GetManagedAgentResponseBodyDataSubAgents struct {
	// The child agent instruction.
	//
	// This parameter is required.
	//
	// example:
	//
	// Review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The child agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skills and actual versions used by the child agent. The version field is not returned when the template package does not contain version information.
	Skills []*GetManagedAgentResponseBodyDataSubAgentsSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
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

func (s *GetManagedAgentResponseBodyDataSubAgents) GetSkills() []*GetManagedAgentResponseBodyDataSubAgentsSkills {
	return s.Skills
}

func (s *GetManagedAgentResponseBodyDataSubAgents) SetInstruction(v string) *GetManagedAgentResponseBodyDataSubAgents {
	s.Instruction = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgents) SetName(v string) *GetManagedAgentResponseBodyDataSubAgents {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgents) SetSkills(v []*GetManagedAgentResponseBodyDataSubAgentsSkills) *GetManagedAgentResponseBodyDataSubAgents {
	s.Skills = v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgents) Validate() error {
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetManagedAgentResponseBodyDataSubAgentsSkills struct {
	// The skill name used by the child agent. Declared as optional for compatibility, but validated as required by the backend for each entry.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The optional version number. If omitted, null, or blank, the latest version is resolved.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetManagedAgentResponseBodyDataSubAgentsSkills) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponseBodyDataSubAgentsSkills) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponseBodyDataSubAgentsSkills) GetName() *string {
	return s.Name
}

func (s *GetManagedAgentResponseBodyDataSubAgentsSkills) GetVersion() *string {
	return s.Version
}

func (s *GetManagedAgentResponseBodyDataSubAgentsSkills) SetName(v string) *GetManagedAgentResponseBodyDataSubAgentsSkills {
	s.Name = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgentsSkills) SetVersion(v string) *GetManagedAgentResponseBodyDataSubAgentsSkills {
	s.Version = &v
	return s
}

func (s *GetManagedAgentResponseBodyDataSubAgentsSkills) Validate() error {
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
