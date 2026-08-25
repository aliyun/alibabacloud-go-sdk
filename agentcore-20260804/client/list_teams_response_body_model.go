// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTeamsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTeamsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListTeamsResponseBodyItems) *ListTeamsResponseBody
	GetItems() []*ListTeamsResponseBodyItems
	SetMaxResults(v int32) *ListTeamsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTeamsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTeamsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTeamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTeamsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTeamsResponseBody
	GetTotalCount() *int64
}

type ListTeamsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                        `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListTeamsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// dGVhbS1vZmZzZXQ6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 42
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTeamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTeamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTeamsResponseBody) GetItems() []*ListTeamsResponseBodyItems {
	return s.Items
}

func (s *ListTeamsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTeamsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTeamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTeamsResponseBody) SetCode(v string) *ListTeamsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTeamsResponseBody) SetHttpStatusCode(v int32) *ListTeamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTeamsResponseBody) SetItems(v []*ListTeamsResponseBodyItems) *ListTeamsResponseBody {
	s.Items = v
	return s
}

func (s *ListTeamsResponseBody) SetMaxResults(v int32) *ListTeamsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsResponseBody) SetMessage(v string) *ListTeamsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTeamsResponseBody) SetNextToken(v string) *ListTeamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamsResponseBody) SetRequestId(v string) *ListTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamsResponseBody) SetSuccess(v bool) *ListTeamsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTeamsResponseBody) SetTotalCount(v int64) *ListTeamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTeamsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTeamsResponseBodyItems struct {
	Agents []*ListTeamsResponseBodyItemsAgents `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
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
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// tm-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string                            `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Users     []*ListTeamsResponseBodyItemsUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTeamsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyItems) GetAgents() []*ListTeamsResponseBodyItemsAgents {
	return s.Agents
}

func (s *ListTeamsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListTeamsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListTeamsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListTeamsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListTeamsResponseBodyItems) GetTeamId() *string {
	return s.TeamId
}

func (s *ListTeamsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListTeamsResponseBodyItems) GetUsers() []*ListTeamsResponseBodyItemsUsers {
	return s.Users
}

func (s *ListTeamsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTeamsResponseBodyItems) SetAgents(v []*ListTeamsResponseBodyItemsAgents) *ListTeamsResponseBodyItems {
	s.Agents = v
	return s
}

func (s *ListTeamsResponseBodyItems) SetCreatedAt(v string) *ListTeamsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetDescription(v string) *ListTeamsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetName(v string) *ListTeamsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetStatus(v string) *ListTeamsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetTeamId(v string) *ListTeamsResponseBodyItems {
	s.TeamId = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetUpdatedAt(v string) *ListTeamsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListTeamsResponseBodyItems) SetUsers(v []*ListTeamsResponseBodyItemsUsers) *ListTeamsResponseBodyItems {
	s.Users = v
	return s
}

func (s *ListTeamsResponseBodyItems) SetWorkspaceId(v string) *ListTeamsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListTeamsResponseBodyItems) Validate() error {
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

type ListTeamsResponseBodyItemsAgents struct {
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// agent-01
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// WORKER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s ListTeamsResponseBodyItemsAgents) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyItemsAgents) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyItemsAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *ListTeamsResponseBodyItemsAgents) GetAgentName() *string {
	return s.AgentName
}

func (s *ListTeamsResponseBodyItemsAgents) GetTeamRole() *string {
	return s.TeamRole
}

func (s *ListTeamsResponseBodyItemsAgents) SetAgentId(v string) *ListTeamsResponseBodyItemsAgents {
	s.AgentId = &v
	return s
}

func (s *ListTeamsResponseBodyItemsAgents) SetAgentName(v string) *ListTeamsResponseBodyItemsAgents {
	s.AgentName = &v
	return s
}

func (s *ListTeamsResponseBodyItemsAgents) SetTeamRole(v string) *ListTeamsResponseBodyItemsAgents {
	s.TeamRole = &v
	return s
}

func (s *ListTeamsResponseBodyItemsAgents) Validate() error {
	return dara.Validate(s)
}

type ListTeamsResponseBodyItemsUsers struct {
	// example:
	//
	// ADMIN
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
	// example:
	//
	// usr-123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// user-01
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListTeamsResponseBodyItemsUsers) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyItemsUsers) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyItemsUsers) GetTeamRole() *string {
	return s.TeamRole
}

func (s *ListTeamsResponseBodyItemsUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListTeamsResponseBodyItemsUsers) GetUserName() *string {
	return s.UserName
}

func (s *ListTeamsResponseBodyItemsUsers) SetTeamRole(v string) *ListTeamsResponseBodyItemsUsers {
	s.TeamRole = &v
	return s
}

func (s *ListTeamsResponseBodyItemsUsers) SetUserId(v string) *ListTeamsResponseBodyItemsUsers {
	s.UserId = &v
	return s
}

func (s *ListTeamsResponseBodyItemsUsers) SetUserName(v string) *ListTeamsResponseBodyItemsUsers {
	s.UserName = &v
	return s
}

func (s *ListTeamsResponseBodyItemsUsers) Validate() error {
	return dara.Validate(s)
}
