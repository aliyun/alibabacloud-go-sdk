// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExternalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateExternalAgentResponseBody
	GetCode() *string
	SetData(v *UpdateExternalAgentResponseBodyData) *UpdateExternalAgentResponseBody
	GetData() *UpdateExternalAgentResponseBodyData
	SetHttpStatusCode(v int32) *UpdateExternalAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateExternalAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateExternalAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateExternalAgentResponseBody
	GetSuccess() *bool
}

type UpdateExternalAgentResponseBody struct {
	// The business status code. The value is SUCCESS when the request succeeds.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the updated external agent.
	Data *UpdateExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateExternalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateExternalAgentResponseBody) GetData() *UpdateExternalAgentResponseBodyData {
	return s.Data
}

func (s *UpdateExternalAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateExternalAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExternalAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExternalAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateExternalAgentResponseBody) SetCode(v string) *UpdateExternalAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExternalAgentResponseBody) SetData(v *UpdateExternalAgentResponseBodyData) *UpdateExternalAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateExternalAgentResponseBody) SetHttpStatusCode(v int32) *UpdateExternalAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateExternalAgentResponseBody) SetMessage(v string) *UpdateExternalAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExternalAgentResponseBody) SetRequestId(v string) *UpdateExternalAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExternalAgentResponseBody) SetSuccess(v bool) *UpdateExternalAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateExternalAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExternalAgentResponseBodyData struct {
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
	ExternalAgentStatus *UpdateExternalAgentResponseBodyDataExternalAgentStatus `json:"externalAgentStatus,omitempty" xml:"externalAgentStatus,omitempty" type:"Struct"`
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
	Model *UpdateExternalAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
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
	Skills []*UpdateExternalAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
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
	Template *UpdateExternalAgentResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*UpdateExternalAgentResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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

func (s UpdateExternalAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateExternalAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *UpdateExternalAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateExternalAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateExternalAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateExternalAgentResponseBodyData) GetExternalAgentStatus() *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	return s.ExternalAgentStatus
}

func (s *UpdateExternalAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateExternalAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *UpdateExternalAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *UpdateExternalAgentResponseBodyData) GetModel() *UpdateExternalAgentResponseBodyDataModel {
	return s.Model
}

func (s *UpdateExternalAgentResponseBodyData) GetModelSource() *string {
	return s.ModelSource
}

func (s *UpdateExternalAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateExternalAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *UpdateExternalAgentResponseBodyData) GetSkills() []*UpdateExternalAgentResponseBodyDataSkills {
	return s.Skills
}

func (s *UpdateExternalAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateExternalAgentResponseBodyData) GetTemplate() *UpdateExternalAgentResponseBodyDataTemplate {
	return s.Template
}

func (s *UpdateExternalAgentResponseBodyData) GetTools() []*UpdateExternalAgentResponseBodyDataTools {
	return s.Tools
}

func (s *UpdateExternalAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateExternalAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateExternalAgentResponseBodyData) SetAgentId(v string) *UpdateExternalAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetCreateMode(v string) *UpdateExternalAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetCreatedAt(v string) *UpdateExternalAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetDeployType(v string) *UpdateExternalAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetDescription(v string) *UpdateExternalAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetExternalAgentStatus(v *UpdateExternalAgentResponseBodyDataExternalAgentStatus) *UpdateExternalAgentResponseBodyData {
	s.ExternalAgentStatus = v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetInstruction(v string) *UpdateExternalAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetLatestSpecVersion(v int64) *UpdateExternalAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetLatestVersionStatus(v string) *UpdateExternalAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetModel(v *UpdateExternalAgentResponseBodyDataModel) *UpdateExternalAgentResponseBodyData {
	s.Model = v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetModelSource(v string) *UpdateExternalAgentResponseBodyData {
	s.ModelSource = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetName(v string) *UpdateExternalAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetRegionId(v string) *UpdateExternalAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetRuntime(v string) *UpdateExternalAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetSkills(v []*UpdateExternalAgentResponseBodyDataSkills) *UpdateExternalAgentResponseBodyData {
	s.Skills = v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetStatus(v string) *UpdateExternalAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetTemplate(v *UpdateExternalAgentResponseBodyDataTemplate) *UpdateExternalAgentResponseBodyData {
	s.Template = v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetTools(v []*UpdateExternalAgentResponseBodyDataTools) *UpdateExternalAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetUpdatedAt(v string) *UpdateExternalAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) SetWorkspaceId(v string) *UpdateExternalAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyData) Validate() error {
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

type UpdateExternalAgentResponseBodyDataExternalAgentStatus struct {
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
	// The last active time of the external agent in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The last heartbeat time of the external agent in RFC 3339 format.
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

func (s UpdateExternalAgentResponseBodyDataExternalAgentStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataExternalAgentStatus) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) GetHeartbeatStatus() *string {
	return s.HeartbeatStatus
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) GetLastActiveAt() *string {
	return s.LastActiveAt
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) GetLastHeartbeat() *string {
	return s.LastHeartbeat
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) GetLocalIP() *string {
	return s.LocalIP
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) GetRuntime() *string {
	return s.Runtime
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) SetHeartbeatStatus(v string) *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	s.HeartbeatStatus = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) SetLastActiveAt(v string) *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastActiveAt = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) SetLastHeartbeat(v string) *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LastHeartbeat = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) SetLocalIP(v string) *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	s.LocalIP = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) SetRuntime(v string) *UpdateExternalAgentResponseBodyDataExternalAgentStatus {
	s.Runtime = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataExternalAgentStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentResponseBodyDataModel struct {
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
	Quota *UpdateExternalAgentResponseBodyDataModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s UpdateExternalAgentResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *UpdateExternalAgentResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateExternalAgentResponseBodyDataModel) GetQuota() *UpdateExternalAgentResponseBodyDataModelQuota {
	return s.Quota
}

func (s *UpdateExternalAgentResponseBodyDataModel) SetModelConnectionId(v string) *UpdateExternalAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModel) SetModelName(v string) *UpdateExternalAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModel) SetQuota(v *UpdateExternalAgentResponseBodyDataModelQuota) *UpdateExternalAgentResponseBodyDataModel {
	s.Quota = v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExternalAgentResponseBodyDataModelQuota struct {
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

func (s UpdateExternalAgentResponseBodyDataModelQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataModelQuota) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetEnabled(v bool) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.Enabled = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetLimitType(v string) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.LimitType = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetOverLimit(v bool) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.OverLimit = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetPeriodType(v string) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.PeriodType = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetRuleStatus(v string) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.RuleStatus = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetUsageLimit(v int64) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) SetUsedAmount(v int64) *UpdateExternalAgentResponseBodyDataModelQuota {
	s.UsedAmount = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataModelQuota) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentResponseBodyDataSkills struct {
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

func (s UpdateExternalAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateExternalAgentResponseBodyDataSkills) SetName(v string) *UpdateExternalAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataSkills) SetVersion(v string) *UpdateExternalAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentResponseBodyDataTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *UpdateExternalAgentResponseBodyDataTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s UpdateExternalAgentResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataTemplate) GetAiRegistry() *UpdateExternalAgentResponseBodyDataTemplateAiRegistry {
	return s.AiRegistry
}

func (s *UpdateExternalAgentResponseBodyDataTemplate) SetAiRegistry(v *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) *UpdateExternalAgentResponseBodyDataTemplate {
	s.AiRegistry = v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExternalAgentResponseBodyDataTemplateAiRegistry struct {
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

func (s UpdateExternalAgentResponseBodyDataTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) SetName(v string) *UpdateExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) SetVersion(v string) *UpdateExternalAgentResponseBodyDataTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type UpdateExternalAgentResponseBodyDataTools struct {
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

func (s UpdateExternalAgentResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentResponseBodyDataTools) GetName() *string {
	return s.Name
}

func (s *UpdateExternalAgentResponseBodyDataTools) GetType() *string {
	return s.Type
}

func (s *UpdateExternalAgentResponseBodyDataTools) SetName(v string) *UpdateExternalAgentResponseBodyDataTools {
	s.Name = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataTools) SetType(v string) *UpdateExternalAgentResponseBodyDataTools {
	s.Type = &v
	return s
}

func (s *UpdateExternalAgentResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}
