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
	// The information about the managed Harness agent after creation.
	Data *CreateManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing result message.
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
	// Indicates whether the request is successful.
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
	// The inbound access mode.
	//
	// example:
	//
	// API_KEY
	AccessMode *string `json:"accessMode,omitempty" xml:"accessMode,omitempty"`
	// The user access scope.
	//
	// example:
	//
	// ALL
	AccessScope *string `json:"accessScope,omitempty" xml:"accessScope,omitempty"`
	// The managed Harness agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The NAS storage mount configuration. Currently, the AgenticFS type is supported, which is used to mount file directories in the access point to the agent runtime instance. For NAS mounts, a maximum of 10 items are allowed in total including ossMounts.
	AgenticFsMounts []*CreateManagedAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The explicitly configured skills.
	ConfiguredSkills []*CreateManagedAgentResponseBodyDataConfiguredSkills `json:"configuredSkills,omitempty" xml:"configuredSkills,omitempty" type:"Repeated"`
	// The target configuration accepted by the managed agent, which does not mean it has taken effect at runtime. This field is used in the detail responses of CreateManagedAgent, UpdateManagedAgent, and GetManagedAgent. In the configuration of a fixed version, it represents the configuration snapshot saved for that version. A returned object with enabled:false indicates that the configuration is disabled. The absence of the entire object indicates that the response does not provide the configuration.
	ContextConfiguration *CreateManagedAgentResponseBodyDataContextConfiguration `json:"contextConfiguration,omitempty" xml:"contextConfiguration,omitempty" type:"Struct"`
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
	// The description of the managed Harness agent.
	//
	// example:
	//
	// Agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Read-only. This field is provided only by GetManagedAgent when an effective version exists, and corresponds to the configuration snapshot of effectiveSpecVersion. When a new version is being prepared or fails to be published, it remains the original effective configuration. It is not returned before the first time it takes effect, or in CreateManagedAgent, UpdateManagedAgent, and the configuration of a fixed version. Omission does not mean it is disabled, and the client cannot automatically fill it with enabled:false.
	EffectiveContextConfiguration *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration `json:"effectiveContextConfiguration,omitempty" xml:"effectiveContextConfiguration,omitempty" type:"Struct"`
	// The environment configuration information.
	Environment *CreateManagedAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent runtime framework.
	Harness *CreateManagedAgentResponseBodyDataHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
	// The agent instruction, which is used to guide the behavior of the agent.
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
	// The name of the managed Harness agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration information.
	Network *CreateManagedAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS mounts. A maximum of 10 items are supported.
	OssMounts []*CreateManagedAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The prompt variable declaration.
	PromptVariables []*CreateManagedAgentResponseBodyDataPromptVariables `json:"promptVariables,omitempty" xml:"promptVariables,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime configuration information.
	Runtime *CreateManagedAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The instance counts of managed Harness agents by sandbox phase. Current keys: PENDING (creating or initializing), RUNNING (running), HIBERNATING (hibernating), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (terminating), and FAILED (failed at runtime). Only phases that actually occur are returned, and missing keys are treated as 0. This field is a dynamic map and new keys may be added in the future. The frontend can use FAILED > 0 to determine whether abnormal instances exist.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The list of skill configurations.
	Skills []*CreateManagedAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed Harness agent.
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

func (s *CreateManagedAgentResponseBodyData) GetAccessMode() *string {
	return s.AccessMode
}

func (s *CreateManagedAgentResponseBodyData) GetAccessScope() *string {
	return s.AccessScope
}

func (s *CreateManagedAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateManagedAgentResponseBodyData) GetAgenticFsMounts() []*CreateManagedAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *CreateManagedAgentResponseBodyData) GetConfiguredSkills() []*CreateManagedAgentResponseBodyDataConfiguredSkills {
	return s.ConfiguredSkills
}

func (s *CreateManagedAgentResponseBodyData) GetContextConfiguration() *CreateManagedAgentResponseBodyDataContextConfiguration {
	return s.ContextConfiguration
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

func (s *CreateManagedAgentResponseBodyData) GetEffectiveContextConfiguration() *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration {
	return s.EffectiveContextConfiguration
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

func (s *CreateManagedAgentResponseBodyData) GetPromptVariables() []*CreateManagedAgentResponseBodyDataPromptVariables {
	return s.PromptVariables
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

func (s *CreateManagedAgentResponseBodyData) SetAccessMode(v string) *CreateManagedAgentResponseBodyData {
	s.AccessMode = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetAccessScope(v string) *CreateManagedAgentResponseBodyData {
	s.AccessScope = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetAgentId(v string) *CreateManagedAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetAgenticFsMounts(v []*CreateManagedAgentResponseBodyDataAgenticFsMounts) *CreateManagedAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetConfiguredSkills(v []*CreateManagedAgentResponseBodyDataConfiguredSkills) *CreateManagedAgentResponseBodyData {
	s.ConfiguredSkills = v
	return s
}

func (s *CreateManagedAgentResponseBodyData) SetContextConfiguration(v *CreateManagedAgentResponseBodyDataContextConfiguration) *CreateManagedAgentResponseBodyData {
	s.ContextConfiguration = v
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

func (s *CreateManagedAgentResponseBodyData) SetEffectiveContextConfiguration(v *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) *CreateManagedAgentResponseBodyData {
	s.EffectiveContextConfiguration = v
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

func (s *CreateManagedAgentResponseBodyData) SetPromptVariables(v []*CreateManagedAgentResponseBodyDataPromptVariables) *CreateManagedAgentResponseBodyData {
	s.PromptVariables = v
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
	if s.ContextConfiguration != nil {
		if err := s.ContextConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.EffectiveContextConfiguration != nil {
		if err := s.EffectiveContextConfiguration.Validate(); err != nil {
			return err
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
	if s.PromptVariables != nil {
		for _, item := range s.PromptVariables {
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

type CreateManagedAgentResponseBodyDataAgenticFsMounts struct {
	// The absolute mount path in the container, such as `/data/train`. Mount targets cannot be duplicated or form parent-child overwrites, and must avoid platform-reserved directories.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The existing non-empty relative directory under the access point. This is a required field validated by the backend for each mount item. Root directories, absolute paths, or parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. If omitted, the default value is false. This is not the read-only policy of a RAM role.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The access point domain name. This is a required field validated by the backend for each mount item. It does not contain a protocol, port, or path. Use the `DomainName` from the NAS `ListAccessPoints` response.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *CreateManagedAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *CreateManagedAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *CreateManagedAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *CreateManagedAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataConfiguredSkills struct {
	// The name of the skill in the Workspace AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type of the skill.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The legacy compatibility field. New requests use sourceType and versionSelector.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version selector for the skill reference. Supports LABEL for tags and VERSION for specific versions. Defaults to LABEL/latest when omitted. Not applicable for STATIC sources.
	VersionSelector *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s CreateManagedAgentResponseBodyDataConfiguredSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataConfiguredSkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) GetVersionSelector() *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	return s.VersionSelector
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) SetName(v string) *CreateManagedAgentResponseBodyDataConfiguredSkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) SetSourceType(v string) *CreateManagedAgentResponseBodyDataConfiguredSkills {
	s.SourceType = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) SetVersion(v string) *CreateManagedAgentResponseBodyDataConfiguredSkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) SetVersionSelector(v *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) *CreateManagedAgentResponseBodyDataConfiguredSkills {
	s.VersionSelector = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector struct {
	// The type of the version selector.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the selector. For LABEL, specify the tag name, such as latest. For VERSION, specify the exact version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetType(v string) *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) SetValue(v string) *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataConfiguredSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataContextConfiguration struct {
	// Read-only resource references. This field may be omitted when no selection is made. A disabled configuration can still retain non-empty selections. You cannot determine whether the configuration is enabled solely based on whether the array is empty.
	Bindings []*CreateManagedAgentResponseBodyDataContextConfigurationBindings `json:"bindings,omitempty" xml:"bindings,omitempty" type:"Repeated"`
	// Read-only. When the configuration object is returned, this field indicates the enabling status of the configuration. A value of false indicates that the configuration is disabled. The absence of the entire configuration object indicates that the response does not provide the configuration, which cannot be treated as false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataContextConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataContextConfiguration) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataContextConfiguration) GetBindings() []*CreateManagedAgentResponseBodyDataContextConfigurationBindings {
	return s.Bindings
}

func (s *CreateManagedAgentResponseBodyDataContextConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataContextConfiguration) SetBindings(v []*CreateManagedAgentResponseBodyDataContextConfigurationBindings) *CreateManagedAgentResponseBodyDataContextConfiguration {
	s.Bindings = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataContextConfiguration) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataContextConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataContextConfiguration) Validate() error {
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataContextConfigurationBindings struct {
	// The context resource ID in the current workspace.
	//
	// example:
	//
	// kb-example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Read-only string. Currently known types are KNOWLEDGE_BASE (knowledge base) and DATA_CONNECTOR (data connector). The response may be extended with new types. The client should retain unknown values and should not fail to parse the entire response due to unknown types. New request types are subject to the actual support scope of the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataContextConfigurationBindings) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataContextConfigurationBindings) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataContextConfigurationBindings) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateManagedAgentResponseBodyDataContextConfigurationBindings) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataContextConfigurationBindings) SetResourceId(v string) *CreateManagedAgentResponseBodyDataContextConfigurationBindings {
	s.ResourceId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataContextConfigurationBindings) SetType(v string) *CreateManagedAgentResponseBodyDataContextConfigurationBindings {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataContextConfigurationBindings) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataEffectiveContextConfiguration struct {
	// Read-only resource references. This field may be omitted when no selection is made. A disabled configuration can still retain non-empty selections. You cannot determine whether the configuration is enabled solely based on whether the array is empty.
	Bindings []*CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings `json:"bindings,omitempty" xml:"bindings,omitempty" type:"Repeated"`
	// Read-only. When the configuration object is returned, this field indicates the enabling status of the configuration. A value of false indicates that the configuration is disabled. The absence of the entire configuration object indicates that the response does not provide the configuration, which cannot be treated as false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) GetBindings() []*CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings {
	return s.Bindings
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) SetBindings(v []*CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration {
	s.Bindings = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) SetEnabled(v bool) *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfiguration) Validate() error {
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings struct {
	// The context resource ID in the current workspace.
	//
	// example:
	//
	// kb-example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Read-only string. Currently known types are KNOWLEDGE_BASE (knowledge base) and DATA_CONNECTOR (data connector). The response may be extended with new types. The client should retain unknown values and should not fail to parse the entire response due to unknown types. New request types are subject to the actual support scope of the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) SetResourceId(v string) *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings {
	s.ResourceId = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) SetType(v string) *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataEffectiveContextConfigurationBindings) Validate() error {
	return dara.Validate(s)
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
	// The ID of the credential.
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
	// The runtime framework configuration.
	Configuration *CreateManagedAgentResponseBodyDataHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// Returned only in queries. The server calculates this based on the complete instance distribution, which is not affected by the pagination and filtering of the instance list. It is empty when the image tag is missing, pagination is incomplete, or it cannot be confirmed, which is not equivalent to 0.
	//
	// example:
	//
	// 2
	OutdatedInstanceCount *int64 `json:"outdatedInstanceCount,omitempty" xml:"outdatedInstanceCount,omitempty"`
	// The runtime framework type.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Returned only in queries. Currently returns the pool target image tag. It is empty when it cannot be parsed, and cannot be replaced by the agent configuration version.
	//
	// example:
	//
	// 202611031754
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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

func (s *CreateManagedAgentResponseBodyDataHarness) GetOutdatedInstanceCount() *int64 {
	return s.OutdatedInstanceCount
}

func (s *CreateManagedAgentResponseBodyDataHarness) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataHarness) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetConfiguration(v *CreateManagedAgentResponseBodyDataHarnessConfiguration) *CreateManagedAgentResponseBodyDataHarness {
	s.Configuration = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetOutdatedInstanceCount(v int64) *CreateManagedAgentResponseBodyDataHarness {
	s.OutdatedInstanceCount = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetType(v string) *CreateManagedAgentResponseBodyDataHarness {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataHarness) SetVersion(v string) *CreateManagedAgentResponseBodyDataHarness {
	s.Version = &v
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
	// The connector account key record ID bound to qodercli, which is not the original service account key.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The connector account name bound to qodercli.
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
	// The model token quota configuration and usage status for the current period. This field is empty if no quota is configured.
	Quota *CreateManagedAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The thinking level. If not configured, the value defaults to default. Empty strings, null values, or omitted fields are handled according to the default behavior of the model and Harness. Valid values: off, low, medium, high, xhigh, and max. If this parameter is not passed or is null during creation, default is used. When updating and submitting the model configuration, if the thinking level is not passed or is null, it reverts to default. Except for off, which is converted to none, the value is passed to the runtime framework as-is without validating model support or automatically downgrading. If the model does not support the specified level, adjust the level accordingly.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// high
	ThinkingLevel *string `json:"thinkingLevel,omitempty" xml:"thinkingLevel,omitempty"`
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

func (s *CreateManagedAgentResponseBodyDataModel) GetThinkingLevel() *string {
	return s.ThinkingLevel
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

func (s *CreateManagedAgentResponseBodyDataModel) SetThinkingLevel(v string) *CreateManagedAgentResponseBodyDataModel {
	s.ThinkingLevel = &v
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
	// Indicates whether the quota is exceeded in the current period. This field is read-only and returned by the backend.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The quota statistical period. A value of day indicates daily, and a value of month indicates monthly.
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
	// The maximum number of tokens allowed to be consumed in a single period.
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
	// Specifies whether to allow Internet access.
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
	// The name of the OSS bucket. This field is required for each mount item and is validated by the backend.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path in the container, such as `/data/train`. Mount targets cannot be duplicated or form parent-child overwrites, and must avoid platform-reserved directories.
	//
	// example:
	//
	// /mnt/oss/datasets
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty relative directory prefix in the configured bucket.
	//
	// example:
	//
	// datasets
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. Default value: false.
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

type CreateManagedAgentResponseBodyDataPromptVariables struct {
	// The default value.
	//
	// example:
	//
	// default
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// The variable description.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The variable name.
	//
	// example:
	//
	// brand
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataPromptVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataPromptVariables) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) GetDescription() *string {
	return s.Description
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) SetDefaultValue(v string) *CreateManagedAgentResponseBodyDataPromptVariables {
	s.DefaultValue = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) SetDescription(v string) *CreateManagedAgentResponseBodyDataPromptVariables {
	s.Description = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) SetName(v string) *CreateManagedAgentResponseBodyDataPromptVariables {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataPromptVariables) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The sandbox auto-scaling and session configuration.
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

type CreateManagedAgentResponseBodyDataRuntimeHpa struct {
	// Indicates whether elastic auto-scaling is enabled in the current configuration.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions per instance in the current configuration.
	//
	// example:
	//
	// 5
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of instances in the current configuration.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of instances in the current configuration.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The maximum session idle time in the current configuration, in seconds.
	//
	// example:
	//
	// 3600
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
	// The HTTP header name used for session affinity. This field takes effect when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The type of the session policy.
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
	// The effective runtime version. This field is read-only.
	//
	// example:
	//
	// 1.0.0
	AppliedVersion *string `json:"appliedVersion,omitempty" xml:"appliedVersion,omitempty"`
	// Indicates whether the skill is from a fixed template. This field is read-only. Template items cannot be removed.
	//
	// example:
	//
	// false
	FromTemplate *bool `json:"fromTemplate,omitempty" xml:"fromTemplate,omitempty"`
	// The name of the skill.
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
	// The source type of the skill.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The current target version for legacy client compatibility. New clients use resolvedVersion and determine the effective runtime version based on appliedVersion.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version selector for the skill reference. Supports LABEL for tags and VERSION for specific versions. Defaults to LABEL/latest when omitted. Not applicable for STATIC sources.
	VersionSelector *CreateManagedAgentResponseBodyDataSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s CreateManagedAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetFromTemplate() *bool {
	return s.FromTemplate
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetResolvedVersion() *string {
	return s.ResolvedVersion
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataSkills) GetVersionSelector() *CreateManagedAgentResponseBodyDataSkillsVersionSelector {
	return s.VersionSelector
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetAppliedVersion(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.AppliedVersion = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetFromTemplate(v bool) *CreateManagedAgentResponseBodyDataSkills {
	s.FromTemplate = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetName(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetResolvedVersion(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.ResolvedVersion = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetSourceType(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.SourceType = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetVersion(v string) *CreateManagedAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) SetVersionSelector(v *CreateManagedAgentResponseBodyDataSkillsVersionSelector) *CreateManagedAgentResponseBodyDataSkills {
	s.VersionSelector = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentResponseBodyDataSkillsVersionSelector struct {
	// The type of the version selector.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the selector. For LABEL, specify the tag name, such as latest. For VERSION, specify the exact version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentResponseBodyDataSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentResponseBodyDataSkillsVersionSelector) SetType(v string) *CreateManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkillsVersionSelector) SetValue(v string) *CreateManagedAgentResponseBodyDataSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentResponseBodyDataSubAgents struct {
	// The instruction for the sub-agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// Review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The name of the sub-agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of skills for the sub-agent.
	Skills []*CreateManagedAgentResponseBodyDataSubAgentsSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
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

func (s *CreateManagedAgentResponseBodyDataSubAgents) GetSkills() []*CreateManagedAgentResponseBodyDataSubAgentsSkills {
	return s.Skills
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) SetInstruction(v string) *CreateManagedAgentResponseBodyDataSubAgents {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) SetName(v string) *CreateManagedAgentResponseBodyDataSubAgents {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) SetSkills(v []*CreateManagedAgentResponseBodyDataSubAgentsSkills) *CreateManagedAgentResponseBodyDataSubAgents {
	s.Skills = v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgents) Validate() error {
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

type CreateManagedAgentResponseBodyDataSubAgentsSkills struct {
	// The name of the skill.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the skill.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateManagedAgentResponseBodyDataSubAgentsSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponseBodyDataSubAgentsSkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponseBodyDataSubAgentsSkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentResponseBodyDataSubAgentsSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentResponseBodyDataSubAgentsSkills) SetName(v string) *CreateManagedAgentResponseBodyDataSubAgentsSkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgentsSkills) SetVersion(v string) *CreateManagedAgentResponseBodyDataSubAgentsSkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentResponseBodyDataSubAgentsSkills) Validate() error {
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
	// The name of the tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the tool.
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
