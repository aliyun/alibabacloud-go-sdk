// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetExternalAgentResponseBody
	GetCode() *string
	SetData(v *GetExternalAgentResponseBodyData) *GetExternalAgentResponseBody
	GetData() *GetExternalAgentResponseBodyData
	SetHttpStatusCode(v int32) *GetExternalAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetExternalAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExternalAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExternalAgentResponseBody
	GetSuccess() *bool
}

type GetExternalAgentResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the managed agent.
	Data *GetExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message about the request processing result.
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetExternalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetExternalAgentResponseBody) GetData() *GetExternalAgentResponseBodyData {
	return s.Data
}

func (s *GetExternalAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetExternalAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExternalAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExternalAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExternalAgentResponseBody) SetCode(v string) *GetExternalAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetExternalAgentResponseBody) SetData(v *GetExternalAgentResponseBodyData) *GetExternalAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetExternalAgentResponseBody) SetHttpStatusCode(v int32) *GetExternalAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetExternalAgentResponseBody) SetMessage(v string) *GetExternalAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetExternalAgentResponseBody) SetRequestId(v string) *GetExternalAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExternalAgentResponseBody) SetSuccess(v bool) *GetExternalAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetExternalAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentResponseBodyData struct {
	// The managed agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The explicitly configured skills.
	ConfiguredSkills []*GetExternalAgentResponseBodyDataConfiguredSkills `json:"configuredSkills,omitempty" xml:"configuredSkills,omitempty" type:"Repeated"`
	// The creation mode.
	//
	// example:
	//
	// CUSTOM
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
	// SELF_HOSTED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the managed agent.
	//
	// example:
	//
	// Code review agent running in the user environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The runtime status information reported by the managed agent.
	ExternalAgentStatus *GetExternalAgentResponseBodyDataExternalAgentStatus `json:"externalAgentStatus,omitempty" xml:"externalAgentStatus,omitempty" type:"Struct"`
	// The agent instruction used to guide the behavior of the agent.
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
	// The processing status of the latest specification version. Valid values: pending, processing, waiting_retry, succeeded, failed, and superseded.
	//
	// example:
	//
	// pending
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The model configuration. This parameter is available only when modelSource is set to PLATFORM.
	Model *GetExternalAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The source of the model configuration. A value of PLATFORM indicates that the platform parses and delivers the model configuration. A value of RUNTIME indicates that the external runtime manages the model, and the model parameter cannot be passed at the same time. Valid values: PLATFORM and RUNTIME.
	//
	// example:
	//
	// PLATFORM
	ModelSource *string `json:"modelSource,omitempty" xml:"modelSource,omitempty"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-external-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime type reported by the managed agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The list of skill configurations.
	Skills []*GetExternalAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The status of the managed agent. Valid values: Creating, Running, Failed, Updating, Deleting, and Deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The agent template configuration.
	Template *GetExternalAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*GetExternalAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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

func (s GetExternalAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetExternalAgentResponseBodyData) GetConfiguredSkills() []*GetExternalAgentResponseBodyDataConfiguredSkills {
	return s.ConfiguredSkills
}

func (s *GetExternalAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetExternalAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetExternalAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *GetExternalAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetExternalAgentResponseBodyData) GetExternalAgentStatus() *GetExternalAgentResponseBodyDataExternalAgentStatus {
	return s.ExternalAgentStatus
}

func (s *GetExternalAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *GetExternalAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *GetExternalAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *GetExternalAgentResponseBodyData) GetModel() *GetExternalAgentResponseBodyDataModel {
	return s.Model
}

func (s *GetExternalAgentResponseBodyData) GetModelSource() *string {
	return s.ModelSource
}

func (s *GetExternalAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExternalAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *GetExternalAgentResponseBodyData) GetSkills() []*GetExternalAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *GetExternalAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetExternalAgentResponseBodyData) GetTemplate() *GetExternalAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *GetExternalAgentResponseBodyData) GetTools() []*GetExternalAgentResponseBodyDataTools {
	return s.Tools
}

func (s *GetExternalAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetExternalAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExternalAgentResponseBodyData) SetAgentId(v string) *GetExternalAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetConfiguredSkills(v []*GetExternalAgentResponseBodyDataConfiguredSkills) *GetExternalAgentResponseBodyData {
	s.ConfiguredSkills = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetCreateMode(v string) *GetExternalAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetCreatedAt(v string) *GetExternalAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetDeployType(v string) *GetExternalAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetDescription(v string) *GetExternalAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetExternalAgentStatus(v *GetExternalAgentResponseBodyDataExternalAgentStatus) *GetExternalAgentResponseBodyData {
	s.ExternalAgentStatus = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetInstruction(v string) *GetExternalAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetLatestSpecVersion(v int64) *GetExternalAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetLatestVersionStatus(v string) *GetExternalAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetModel(v *GetExternalAgentResponseBodyDataModel) *GetExternalAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetModelSource(v string) *GetExternalAgentResponseBodyData {
	s.ModelSource = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetName(v string) *GetExternalAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetRegionId(v string) *GetExternalAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetRuntime(v string) *GetExternalAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetSkills(v []*GetExternalAgentResponseBodyDataSkills) *GetExternalAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetStatus(v string) *GetExternalAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetTemplate(v *GetExternalAgentResponseBodyDataTemplate) *GetExternalAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetTools(v []*GetExternalAgentResponseBodyDataTools) *GetExternalAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetUpdatedAt(v string) *GetExternalAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetWorkspaceId(v string) *GetExternalAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetExternalAgentResponseBodyData) Validate() error {
	if s.ConfiguredSkills != nil {
		for _, item := range s.ConfiguredSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExternalAgentStatus != nil {
		if err := s.ExternalAgentStatus.Validate(); err != nil {
			return err
		}
	}
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

type GetExternalAgentResponseBodyDataConfiguredSkills struct {
	// The name of the skill in the Workspace AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill source type. A value of REFERENCE indicates that the skill references AI Registry. A value of STATIC indicates that the skill is statically provided with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// A legacy compatibility field. For new requests, use sourceType and versionSelector.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The skill reference version selector. It supports LABEL for tags and VERSION for specific versions. If omitted, LABEL/latest is used. Selectors are not provided for STATIC sources.
	VersionSelector *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s GetExternalAgentResponseBodyDataConfiguredSkills) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataConfiguredSkills) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) GetVersion() *string {
	return s.Version
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) GetVersionSelector() *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector {
	return s.VersionSelector
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) SetName(v string) *GetExternalAgentResponseBodyDataConfiguredSkills {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) SetSourceType(v string) *GetExternalAgentResponseBodyDataConfiguredSkills {
	s.SourceType = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) SetVersion(v string) *GetExternalAgentResponseBodyDataConfiguredSkills {
	s.Version = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) SetVersionSelector(v *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) *GetExternalAgentResponseBodyDataConfiguredSkills {
	s.VersionSelector = v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector struct {
	// The version selector type. A value of LABEL indicates selection by tag, and a value of VERSION indicates selection by specific version.
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

func (s GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) SetType(v string) *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) SetValue(v string) *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataConfiguredSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataExternalAgentStatus struct {
	// The heartbeat status. ONLINE indicates that the latest heartbeat is within the configured timeout threshold. STALE indicates that the heartbeat has timed out. UNKNOWN indicates that the heartbeat is missing or the format is invalid. Valid values: ONLINE, STALE, and UNKNOWN.
	//
	// example:
	//
	// ONLINE
	HeartbeatStatus *string `json:"heartbeatStatus,omitempty" xml:"heartbeatStatus,omitempty"`
	// The last active time of the managed agent, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The time of the last heartbeat from the managed agent, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastHeartbeat *string `json:"lastHeartbeat,omitempty" xml:"lastHeartbeat,omitempty"`
	// The local IP address reported by the managed agent.
	//
	// example:
	//
	// 10.0.0.42
	LocalIP *string `json:"localIP,omitempty" xml:"localIP,omitempty"`
	// The runtime type reported by the managed agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
}

func (s GetExternalAgentResponseBodyDataExternalAgentStatus) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataExternalAgentStatus) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) GetHeartbeatStatus() *string {
	return s.HeartbeatStatus
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) GetLastActiveAt() *string {
	return s.LastActiveAt
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) GetLastHeartbeat() *string {
	return s.LastHeartbeat
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) GetLocalIP() *string {
	return s.LocalIP
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) GetRuntime() *string {
	return s.Runtime
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) SetHeartbeatStatus(v string) *GetExternalAgentResponseBodyDataExternalAgentStatus {
	s.HeartbeatStatus = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) SetLastActiveAt(v string) *GetExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastActiveAt = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) SetLastHeartbeat(v string) *GetExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastHeartbeat = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) SetLocalIP(v string) *GetExternalAgentResponseBodyDataExternalAgentStatus {
	s.LocalIP = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) SetRuntime(v string) *GetExternalAgentResponseBodyDataExternalAgentStatus {
	s.Runtime = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataExternalAgentStatus) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataModel struct {
	// The model connection ID.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// The name of the upstream model.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration and usage status in the current cycle. This parameter is empty if no quota is configured.
	Quota *GetExternalAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetExternalAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *GetExternalAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *GetExternalAgentResponseBodyDataModel) GetQuota() *GetExternalAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *GetExternalAgentResponseBodyDataModel) SetModelConnectionId(v string) *GetExternalAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModel) SetModelName(v string) *GetExternalAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModel) SetQuota(v *GetExternalAgentResponseBodyDataModelQuota) *GetExternalAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *GetExternalAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentResponseBodyDataModelQuota struct {
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
	// Indicates whether the quota is exceeded in the current cycle. This parameter is read-only and returned by the backend.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The quota statistical period. A value of day indicates daily statistics, and a value of month indicates monthly statistics.
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The gateway quota rule status. This parameter is read-only and returned by the backend.
	//
	// example:
	//
	// ACTIVE
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// The maximum number of tokens allowed to be consumed in a single cycle.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The number of tokens consumed in the current cycle. This parameter is read-only and returned by the backend.
	//
	// example:
	//
	// 12345
	UsedAmount *int64 `json:"usedAmount,omitempty" xml:"usedAmount,omitempty"`
}

func (s GetExternalAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *GetExternalAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetEnabled(v bool) *GetExternalAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetLimitType(v string) *GetExternalAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *GetExternalAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetPeriodType(v string) *GetExternalAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *GetExternalAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *GetExternalAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *GetExternalAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModelQuota) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataSkills struct {
	// The version applied at runtime. This field is read-only.
	//
	// example:
	//
	// 1.0.0
	AppliedVersion *string `json:"appliedVersion,omitempty" xml:"appliedVersion,omitempty"`
	// Indicates whether the item is from a fixed template. This field is read-only. Items from a template cannot be removed.
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
	// The skill source type. A value of REFERENCE indicates that the skill references AI Registry. A value of STATIC indicates that the skill is statically provided with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The current target version for compatibility with legacy clients. New clients use resolvedVersion and determine the effective runtime version based on appliedVersion.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The skill reference version selector. It supports LABEL for tags and VERSION for specific versions. If omitted, LABEL/latest is used. Selectors are not provided for STATIC sources.
	VersionSelector *GetExternalAgentResponseBodyDataSkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s GetExternalAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataSkills) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *GetExternalAgentResponseBodyDataSkills) GetFromTemplate() *bool {
	return s.FromTemplate
}

func (s *GetExternalAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyDataSkills) GetResolvedVersion() *string {
	return s.ResolvedVersion
}

func (s *GetExternalAgentResponseBodyDataSkills) GetSourceType() *string {
	return s.SourceType
}

func (s *GetExternalAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *GetExternalAgentResponseBodyDataSkills) GetVersionSelector() *GetExternalAgentResponseBodyDataSkillsVersionSelector {
	return s.VersionSelector
}

func (s *GetExternalAgentResponseBodyDataSkills) SetAppliedVersion(v string) *GetExternalAgentResponseBodyDataSkills {
	s.AppliedVersion = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetFromTemplate(v bool) *GetExternalAgentResponseBodyDataSkills {
	s.FromTemplate = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetName(v string) *GetExternalAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetResolvedVersion(v string) *GetExternalAgentResponseBodyDataSkills {
	s.ResolvedVersion = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetSourceType(v string) *GetExternalAgentResponseBodyDataSkills {
	s.SourceType = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetVersion(v string) *GetExternalAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetVersionSelector(v *GetExternalAgentResponseBodyDataSkillsVersionSelector) *GetExternalAgentResponseBodyDataSkills {
	s.VersionSelector = v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentResponseBodyDataSkillsVersionSelector struct {
	// The version selector type. A value of LABEL indicates selection by tag, and a value of VERSION indicates selection by specific version.
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

func (s GetExternalAgentResponseBodyDataSkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataSkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataSkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *GetExternalAgentResponseBodyDataSkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *GetExternalAgentResponseBodyDataSkillsVersionSelector) SetType(v string) *GetExternalAgentResponseBodyDataSkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkillsVersionSelector) SetValue(v string) *GetExternalAgentResponseBodyDataSkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *GetExternalAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s GetExternalAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataTemplate) GetAiRegistry() *GetExternalAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *GetExternalAgentResponseBodyDataTemplate) SetAiRegistry(v *GetExternalAgentResponseBodyDataTemplateAiRegistry) *GetExternalAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *GetExternalAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalAgentResponseBodyDataTemplateAiRegistry struct {
	// The name of the template in the AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in the AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetExternalAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *GetExternalAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *GetExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *GetExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataTools struct {
	// The name of the tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tool type. Valid values: MCP.
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetExternalAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *GetExternalAgentResponseBodyDataTools) SetName(v string) *GetExternalAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataTools) SetType(v string) *GetExternalAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
