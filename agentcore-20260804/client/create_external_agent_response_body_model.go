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
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information about the external agent after creation.
	Data *CreateExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request processing.
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
	// The runtime result corresponding to the currently effective specification.
	EffectiveResult *CreateExternalAgentResponseBodyDataEffectiveResult `json:"effectiveResult,omitempty" xml:"effectiveResult,omitempty" type:"Struct"`
	// The currently effective specification version number.
	//
	// example:
	//
	// 1
	EffectiveSpecVersion *int64 `json:"effectiveSpecVersion,omitempty" xml:"effectiveSpecVersion,omitempty"`
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
	// The model configuration. This parameter is available only when modelSource is set to PLATFORM.
	Model *CreateExternalAgentResponseBodyDataModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The source of the model configuration. Valid values:
	//
	// - PLATFORM: The platform parses and delivers the model configuration. You can specify the model parameter.
	//
	// - RUNTIME: The external runtime manages the model on its own. You cannot specify the model parameter at the same time.
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
	// The status of the external agent. Valid values:
	//
	// - Creating: The agent is being created.
	//
	// - Running: The agent is running.
	//
	// - Failed: The agent creation failed.
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

func (s *CreateExternalAgentResponseBodyData) GetEffectiveResult() *CreateExternalAgentResponseBodyDataEffectiveResult {
	return s.EffectiveResult
}

func (s *CreateExternalAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
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

func (s *CreateExternalAgentResponseBodyData) SetEffectiveResult(v *CreateExternalAgentResponseBodyDataEffectiveResult) *CreateExternalAgentResponseBodyData {
	s.EffectiveResult = v
	return s
}

func (s *CreateExternalAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *CreateExternalAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
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

type CreateExternalAgentResponseBodyDataEffectiveResult struct {
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

func (s CreateExternalAgentResponseBodyDataEffectiveResult) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentResponseBodyDataEffectiveResult) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetMatrixUserId() *string {
	return s.MatrixUserId
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetPersonalRoomId() *string {
	return s.PersonalRoomId
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetRuntimeAcceptStatus() *string {
	return s.RuntimeAcceptStatus
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetRuntimeRequestVersion() *int64 {
	return s.RuntimeRequestVersion
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) GetWorkspacePrefix() *string {
	return s.WorkspacePrefix
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetMatrixUserId(v string) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.MatrixUserId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetPersonalRoomId(v string) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.PersonalRoomId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetRuntimeAcceptStatus(v string) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeAcceptStatus = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetRuntimeId(v string) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetRuntimeRequestVersion(v int64) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.RuntimeRequestVersion = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) SetWorkspacePrefix(v string) *CreateExternalAgentResponseBodyDataEffectiveResult {
	s.WorkspacePrefix = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataEffectiveResult) Validate() error {
	return dara.Validate(s)
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
	// The most recent active time of the external agent in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The most recent heartbeat time of the external agent in RFC 3339 format.
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

func (s *CreateExternalAgentResponseBodyDataModel) SetModelConnectionId(v string) *CreateExternalAgentResponseBodyDataModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModel) SetModelName(v string) *CreateExternalAgentResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *CreateExternalAgentResponseBodyDataModel) Validate() error {
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
