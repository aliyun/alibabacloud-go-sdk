// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTeamResponseBody
	GetCode() *string
	SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody
	GetData() *GetTeamResponseBodyData
	SetHttpStatusCode(v int32) *GetTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTeamResponseBody
	GetSuccess() *bool
}

type GetTeamResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTeamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBody) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTeamResponseBody) GetData() *GetTeamResponseBodyData {
	return s.Data
}

func (s *GetTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTeamResponseBody) SetCode(v string) *GetTeamResponseBody {
	s.Code = &v
	return s
}

func (s *GetTeamResponseBody) SetData(v *GetTeamResponseBodyData) *GetTeamResponseBody {
	s.Data = v
	return s
}

func (s *GetTeamResponseBody) SetHttpStatusCode(v int32) *GetTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTeamResponseBody) SetMessage(v string) *GetTeamResponseBody {
	s.Message = &v
	return s
}

func (s *GetTeamResponseBody) SetRequestId(v string) *GetTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTeamResponseBody) SetSuccess(v bool) *GetTeamResponseBody {
	s.Success = &v
	return s
}

func (s *GetTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTeamResponseBodyData struct {
	Agents []*GetTeamResponseBodyDataAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 负责智能客服业务的团队
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// tm-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string                         `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Users     []*GetTeamResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyData) GetAgents() []*GetTeamResponseBodyDataAgents {
	return s.Agents
}

func (s *GetTeamResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyData) GetTeamId() *string {
	return s.TeamId
}

func (s *GetTeamResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyData) GetUsers() []*GetTeamResponseBodyDataUsers {
	return s.Users
}

func (s *GetTeamResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyData) SetAgents(v []*GetTeamResponseBodyDataAgents) *GetTeamResponseBodyData {
	s.Agents = v
	return s
}

func (s *GetTeamResponseBodyData) SetCreatedAt(v string) *GetTeamResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetDescription(v string) *GetTeamResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyData) SetName(v string) *GetTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyData) SetRegionId(v string) *GetTeamResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTeamResponseBodyData) SetStatus(v string) *GetTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyData) SetTeamId(v string) *GetTeamResponseBodyData {
	s.TeamId = &v
	return s
}

func (s *GetTeamResponseBodyData) SetUpdatedAt(v string) *GetTeamResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyData) SetUsers(v []*GetTeamResponseBodyDataUsers) *GetTeamResponseBodyData {
	s.Users = v
	return s
}

func (s *GetTeamResponseBodyData) SetWorkspaceId(v string) *GetTeamResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyData) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTeamResponseBodyDataAgents struct {
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// example:
	//
	// 处理售后咨询的智能体
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	EffectiveSpecVersion *int64 `json:"effectiveSpecVersion,omitempty" xml:"effectiveSpecVersion,omitempty"`
	// example:
	//
	// 2
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// example:
	//
	// agent-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyDataAgents) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataAgents) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *GetTeamResponseBodyDataAgents) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetTeamResponseBodyDataAgents) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyDataAgents) GetDeployType() *string {
	return s.DeployType
}

func (s *GetTeamResponseBodyDataAgents) GetDescription() *string {
	return s.Description
}

func (s *GetTeamResponseBodyDataAgents) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *GetTeamResponseBodyDataAgents) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *GetTeamResponseBodyDataAgents) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyDataAgents) GetRuntime() *string {
	return s.Runtime
}

func (s *GetTeamResponseBodyDataAgents) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyDataAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *GetTeamResponseBodyDataAgents) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyDataAgents) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyDataAgents) SetAgentId(v string) *GetTeamResponseBodyDataAgents {
	s.AgentId = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetCreateMode(v string) *GetTeamResponseBodyDataAgents {
	s.CreateMode = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetCreatedAt(v string) *GetTeamResponseBodyDataAgents {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetDeployType(v string) *GetTeamResponseBodyDataAgents {
	s.DeployType = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetDescription(v string) *GetTeamResponseBodyDataAgents {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetEffectiveSpecVersion(v int64) *GetTeamResponseBodyDataAgents {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetLatestSpecVersion(v int64) *GetTeamResponseBodyDataAgents {
	s.LatestSpecVersion = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetName(v string) *GetTeamResponseBodyDataAgents {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetRuntime(v string) *GetTeamResponseBodyDataAgents {
	s.Runtime = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetStatus(v string) *GetTeamResponseBodyDataAgents {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetTeamRole(v string) *GetTeamResponseBodyDataAgents {
	s.TeamRole = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetUpdatedAt(v string) *GetTeamResponseBodyDataAgents {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) SetWorkspaceId(v string) *GetTeamResponseBodyDataAgents {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyDataAgents) Validate() error {
	return dara.Validate(s)
}

type GetTeamResponseBodyDataUsers struct {
	// example:
	//
	// password
	AuthMethod *string `json:"authMethod,omitempty" xml:"authMethod,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 张三
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// Example@2026
	InitialPassword *string `json:"initialPassword,omitempty" xml:"initialPassword,omitempty"`
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 智能体运营组成员
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// usr-123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTeamResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyDataUsers) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *GetTeamResponseBodyDataUsers) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTeamResponseBodyDataUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTeamResponseBodyDataUsers) GetEmail() *string {
	return s.Email
}

func (s *GetTeamResponseBodyDataUsers) GetInitialPassword() *string {
	return s.InitialPassword
}

func (s *GetTeamResponseBodyDataUsers) GetName() *string {
	return s.Name
}

func (s *GetTeamResponseBodyDataUsers) GetNote() *string {
	return s.Note
}

func (s *GetTeamResponseBodyDataUsers) GetStatus() *string {
	return s.Status
}

func (s *GetTeamResponseBodyDataUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *GetTeamResponseBodyDataUsers) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetTeamResponseBodyDataUsers) GetUserId() *string {
	return s.UserId
}

func (s *GetTeamResponseBodyDataUsers) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTeamResponseBodyDataUsers) SetAuthMethod(v string) *GetTeamResponseBodyDataUsers {
	s.AuthMethod = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetCreatedAt(v string) *GetTeamResponseBodyDataUsers {
	s.CreatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetDisplayName(v string) *GetTeamResponseBodyDataUsers {
	s.DisplayName = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetEmail(v string) *GetTeamResponseBodyDataUsers {
	s.Email = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetInitialPassword(v string) *GetTeamResponseBodyDataUsers {
	s.InitialPassword = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetName(v string) *GetTeamResponseBodyDataUsers {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetNote(v string) *GetTeamResponseBodyDataUsers {
	s.Note = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetStatus(v string) *GetTeamResponseBodyDataUsers {
	s.Status = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetTeamRole(v string) *GetTeamResponseBodyDataUsers {
	s.TeamRole = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetUpdatedAt(v string) *GetTeamResponseBodyDataUsers {
	s.UpdatedAt = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetUserId(v string) *GetTeamResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) SetWorkspaceId(v string) *GetTeamResponseBodyDataUsers {
	s.WorkspaceId = &v
	return s
}

func (s *GetTeamResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
