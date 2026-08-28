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
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The external agent details.
	Data *GetExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
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
	// The external agent description.
	//
	// example:
	//
	// A code review agent running in the user environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The runtime result corresponding to the currently effective specification.
	EffectiveResult *GetExternalAgentResponseBodyDataEffectiveResult `json:"effectiveResult,omitempty" xml:"effectiveResult,omitempty" type:"Struct"`
	// The currently effective specification version number.
	//
	// example:
	//
	// 1
	EffectiveSpecVersion *int64 `json:"effectiveSpecVersion,omitempty" xml:"effectiveSpecVersion,omitempty"`
	// The runtime status information reported by the external agent.
	ExternalAgentStatus *GetExternalAgentResponseBodyDataExternalAgentStatus `json:"externalAgentStatus,omitempty" xml:"externalAgentStatus,omitempty" type:"Struct"`
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
	// - pending: Pending processing.
	//
	// - processing: Being processed.
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
	Model *GetExternalAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The model configuration source. Valid values:
	//
	// - PLATFORM: The model configuration is parsed and delivered by the platform.
	//
	// - RUNTIME: The model is managed by the external runtime. The model parameter cannot be specified at the same time.
	//
	// example:
	//
	// PLATFORM
	ModelSource *string `json:"modelSource,omitempty" xml:"modelSource,omitempty"`
	// The external agent name.
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
	Skills []*GetExternalAgentResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The external agent status. Valid values:
	//
	// - Creating: The agent is being created.
	//
	// - Running: The agent is running.
	//
	// - Failed: The agent has failed.
	//
	// - Updating: The agent is being updated.
	//
	// - Deleting: The agent is being deleted.
	//
	// - Deleted: The agent has been deleted.
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

func (s *GetExternalAgentResponseBodyData) GetEffectiveResult() *GetExternalAgentResponseBodyDataEffectiveResult {
	return s.EffectiveResult
}

func (s *GetExternalAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
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

func (s *GetExternalAgentResponseBodyData) SetEffectiveResult(v *GetExternalAgentResponseBodyDataEffectiveResult) *GetExternalAgentResponseBodyData {
	s.EffectiveResult = v
	return s
}

func (s *GetExternalAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *GetExternalAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
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
	if s.EffectiveResult != nil {
		if err := s.EffectiveResult.Validate(); err != nil {
			return err
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

type GetExternalAgentResponseBodyDataEffectiveResult struct {
	// The user ID of the agent in Matrix.
	//
	// example:
	//
	// @agent-1:matrix.example.com
	MatrixUserId *string `json:"matrixUserId,omitempty" xml:"matrixUserId,omitempty"`
	// The Matrix personal room ID of the agent.
	//
	// example:
	//
	// !room:matrix.example.com
	PersonalRoomId *string `json:"personalRoomId,omitempty" xml:"personalRoomId,omitempty"`
	// The acceptance status of the runtime for the current request version.
	//
	// example:
	//
	// ACCEPTED
	RuntimeAcceptStatus *string `json:"runtimeAcceptStatus,omitempty" xml:"runtimeAcceptStatus,omitempty"`
	// The runtime instance ID.
	//
	// example:
	//
	// runtime-123
	RuntimeId *string `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
	// The runtime request version number.
	//
	// example:
	//
	// 5
	RuntimeRequestVersion *int64 `json:"runtimeRequestVersion,omitempty" xml:"runtimeRequestVersion,omitempty"`
	// The storage prefix of the agent in the workspace.
	//
	// example:
	//
	// agents/agent-1
	WorkspacePrefix *string `json:"workspacePrefix,omitempty" xml:"workspacePrefix,omitempty"`
}

func (s GetExternalAgentResponseBodyDataEffectiveResult) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataEffectiveResult) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetMatrixUserId() *string {
	return s.MatrixUserId
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetPersonalRoomId() *string {
	return s.PersonalRoomId
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetRuntimeAcceptStatus() *string {
	return s.RuntimeAcceptStatus
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetRuntimeRequestVersion() *int64 {
	return s.RuntimeRequestVersion
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) GetWorkspacePrefix() *string {
	return s.WorkspacePrefix
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetMatrixUserId(v string) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.MatrixUserId = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetPersonalRoomId(v string) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.PersonalRoomId = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetRuntimeAcceptStatus(v string) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeAcceptStatus = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetRuntimeId(v string) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeId = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetRuntimeRequestVersion(v int64) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeRequestVersion = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) SetWorkspacePrefix(v string) *GetExternalAgentResponseBodyDataEffectiveResult {
	s.WorkspacePrefix = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataEffectiveResult) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataExternalAgentStatus struct {
	// The heartbeat status. Valid values:
	//
	// - ONLINE: The latest heartbeat has not exceeded the configured timeout threshold.
	//
	// - STALE: The heartbeat has timed out.
	//
	// - UNKNOWN: The heartbeat is missing or has an invalid format.
	//
	// example:
	//
	// ONLINE
	HeartbeatStatus *string `json:"heartbeatStatus,omitempty" xml:"heartbeatStatus,omitempty"`
	// The time when the external agent was last active in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The time of the last heartbeat from the external agent in RFC 3339 format.
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

func (s *GetExternalAgentResponseBodyDataModel) SetModelConnectionId(v string) *GetExternalAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModel) SetModelName(v string) *GetExternalAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type GetExternalAgentResponseBodyDataSkills struct {
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

func (s GetExternalAgentResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetExternalAgentResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetExternalAgentResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *GetExternalAgentResponseBodyDataSkills) SetName(v string) *GetExternalAgentResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) SetVersion(v string) *GetExternalAgentResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *GetExternalAgentResponseBodyDataSkills) Validate() error {
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
