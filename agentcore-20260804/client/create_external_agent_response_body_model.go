// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateExternalAgentResponseBody
	GetCode() *string
	SetData(v *CreateExternalAgentResponseBodyData) *CreateExternalAgentResponseBody
	GetData() *CreateExternalAgentResponseBodyData
	SetHttpStatusCode(v int32) *CreateExternalAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateExternalAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExternalAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExternalAgentResponseBody
	GetSuccess() *bool
}

type CreateExternalAgentResponseBody struct {
	// The business status code. The value is SUCCESS when the request succeeds.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information about the created external agent.
	Data *CreateExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value is 200 when the request succeeds.
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
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateExternalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateExternalAgentResponseBody) GetData() *CreateExternalAgentResponseBodyData {
	return s.Data
}

func (s *CreateExternalAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateExternalAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExternalAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExternalAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExternalAgentResponseBody) SetCode(v string) *CreateExternalAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExternalAgentResponseBody) SetData(v *CreateExternalAgentResponseBodyData) *CreateExternalAgentResponseBody {
	s.Data = v
	return s
}

func (s *CreateExternalAgentResponseBody) SetHttpStatusCode(v int32) *CreateExternalAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExternalAgentResponseBody) SetMessage(v string) *CreateExternalAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExternalAgentResponseBody) SetRequestId(v string) *CreateExternalAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExternalAgentResponseBody) SetSuccess(v bool) *CreateExternalAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExternalAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentResponseBodyData struct {
	// The external agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
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
	// The description of the external agent.
	//
	// example:
	//
	// A code review agent running in the user environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The runtime status information reported by the external agent.
	ExternalAgentStatus *CreateExternalAgentResponseBodyDataExternalAgentStatus `json:"externalAgentStatus,omitempty" xml:"externalAgentStatus,omitempty" type:"Struct"`
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
	// The processing status of the latest specification version. Valid values:
	//
	// - pending: Pending.
	//
	// - processing: Processing.
	//
	// - waiting_retry: Waiting for retry.
	//
	// - succeeded: Succeeded.
	//
	// - failed: Failed.
	//
	// - superseded: Superseded by a newer version.
	//
	// example:
	//
	// pending
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The model configuration. Available only when modelSource is set to PLATFORM.
	Model *CreateExternalAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The model configuration source. PLATFORM indicates that the platform parses and delivers the model configuration. RUNTIME indicates that the external runtime manages the model independently, and the model parameter cannot be specified at the same time. Valid values:
	//
	// - PLATFORM: platform model.
	//
	// - RUNTIME: runtime model.
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The runtime type reported by the external agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The list of skill configurations.
	Skills []*CreateExternalAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The external agent status. Valid values:
	//
	// - Creating: Being created.
	//
	// - Running: Running.
	//
	// - Failed: Failed.
	//
	// - Updating: Being updated.
	//
	// - Deleting: Being deleted.
	//
	// - Deleted: Deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The agent template configuration.
	Template *CreateExternalAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*CreateExternalAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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

func (s CreateExternalAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateExternalAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateExternalAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateExternalAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateExternalAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateExternalAgentResponseBodyData) GetExternalAgentStatus() *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	return s.ExternalAgentStatus
}

func (s *CreateExternalAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateExternalAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *CreateExternalAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *CreateExternalAgentResponseBodyData) GetModel() *CreateExternalAgentResponseBodyDataModel {
	return s.Model
}

func (s *CreateExternalAgentResponseBodyData) GetModelSource() *string {
	return s.ModelSource
}

func (s *CreateExternalAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExternalAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateExternalAgentResponseBodyData) GetSkills() []*CreateExternalAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *CreateExternalAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateExternalAgentResponseBodyData) GetTemplate() *CreateExternalAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *CreateExternalAgentResponseBodyData) GetTools() []*CreateExternalAgentResponseBodyDataTools {
	return s.Tools
}

func (s *CreateExternalAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateExternalAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateExternalAgentResponseBodyData) SetAgentId(v string) *CreateExternalAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetCreateMode(v string) *CreateExternalAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetCreatedAt(v string) *CreateExternalAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetDeployType(v string) *CreateExternalAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetDescription(v string) *CreateExternalAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetExternalAgentStatus(v *CreateExternalAgentResponseBodyDataExternalAgentStatus) *CreateExternalAgentResponseBodyData {
	s.ExternalAgentStatus = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetInstruction(v string) *CreateExternalAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetLatestSpecVersion(v int64) *CreateExternalAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetLatestVersionStatus(v string) *CreateExternalAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetModel(v *CreateExternalAgentResponseBodyDataModel) *CreateExternalAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetModelSource(v string) *CreateExternalAgentResponseBodyData {
	s.ModelSource = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetName(v string) *CreateExternalAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetRegionId(v string) *CreateExternalAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetRuntime(v string) *CreateExternalAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetSkills(v []*CreateExternalAgentResponseBodyDataSkills) *CreateExternalAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetStatus(v string) *CreateExternalAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetTemplate(v *CreateExternalAgentResponseBodyDataTemplate) *CreateExternalAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetTools(v []*CreateExternalAgentResponseBodyDataTools) *CreateExternalAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetUpdatedAt(v string) *CreateExternalAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetWorkspaceId(v string) *CreateExternalAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyData) Validate() error {
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

type CreateExternalAgentResponseBodyDataExternalAgentStatus struct {
	// The heartbeat status. ONLINE indicates that the most recent heartbeat has not exceeded the configured timeout threshold. STALE indicates that the heartbeat has timed out. UNKNOWN indicates that the heartbeat is missing or has an invalid format. Valid values:
	//
	// - ONLINE: Online.
	//
	// - STALE: Heartbeat expired.
	//
	// - UNKNOWN: Unknown.
	//
	// example:
	//
	// ONLINE
	HeartbeatStatus *string `json:"heartbeatStatus,omitempty" xml:"heartbeatStatus,omitempty"`
	// The time when the external agent was last active, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The time of the last heartbeat from the external agent, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastHeartbeat *string `json:"lastHeartbeat,omitempty" xml:"lastHeartbeat,omitempty"`
	// The local IP address reported by the external agent.
	//
	// example:
	//
	// 10.0.0.42
	LocalIP *string `json:"localIP,omitempty" xml:"localIP,omitempty"`
	// The runtime type reported by the external agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
}

func (s CreateExternalAgentResponseBodyDataExternalAgentStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataExternalAgentStatus) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) GetHeartbeatStatus() *string {
	return s.HeartbeatStatus
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) GetLastActiveAt() *string {
	return s.LastActiveAt
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) GetLastHeartbeat() *string {
	return s.LastHeartbeat
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) GetLocalIP() *string {
	return s.LocalIP
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) SetHeartbeatStatus(v string) *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	s.HeartbeatStatus = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) SetLastActiveAt(v string) *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastActiveAt = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) SetLastHeartbeat(v string) *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastHeartbeat = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) SetLocalIP(v string) *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LocalIP = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) SetRuntime(v string) *CreateExternalAgentResponseBodyDataExternalAgentStatus {
	s.Runtime = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataExternalAgentStatus) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentResponseBodyDataModel struct {
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
	// The model token quota configuration and the quota usage status in the current cycle. This field is empty if no quota is configured.
	Quota *CreateExternalAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s CreateExternalAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateExternalAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateExternalAgentResponseBodyDataModel) GetQuota() *CreateExternalAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *CreateExternalAgentResponseBodyDataModel) SetModelConnectionId(v string) *CreateExternalAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModel) SetModelName(v string) *CreateExternalAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModel) SetQuota(v *CreateExternalAgentResponseBodyDataModelQuota) *CreateExternalAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentResponseBodyDataModelQuota struct {
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
	// Indicates whether the quota has been exceeded in the current cycle. This is a read-only field returned by the backend.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The quota statistical period. day indicates a daily period. month indicates a monthly period.
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

func (s CreateExternalAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetEnabled(v bool) *CreateExternalAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetLimitType(v string) *CreateExternalAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *CreateExternalAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetPeriodType(v string) *CreateExternalAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *CreateExternalAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *CreateExternalAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *CreateExternalAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModelQuota) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentResponseBodyDataSkills struct {
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

func (s CreateExternalAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateExternalAgentResponseBodyDataSkills) SetName(v string) *CreateExternalAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataSkills) SetVersion(v string) *CreateExternalAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentResponseBodyDataTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *CreateExternalAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s CreateExternalAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataTemplate) GetAiRegistry() *CreateExternalAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *CreateExternalAgentResponseBodyDataTemplate) SetAiRegistry(v *CreateExternalAgentResponseBodyDataTemplateAiRegistry) *CreateExternalAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *CreateExternalAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalAgentResponseBodyDataTemplateAiRegistry struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateExternalAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *CreateExternalAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *CreateExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *CreateExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateExternalAgentResponseBodyDataTools struct {
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

func (s CreateExternalAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *CreateExternalAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *CreateExternalAgentResponseBodyDataTools) SetName(v string) *CreateExternalAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataTools) SetType(v string) *CreateExternalAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
