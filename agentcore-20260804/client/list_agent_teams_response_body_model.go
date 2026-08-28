// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentTeamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentTeamsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentTeamsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListAgentTeamsResponseBodyItems) *ListAgentTeamsResponseBody
	GetItems() []*ListAgentTeamsResponseBodyItems
	SetMaxResults(v int32) *ListAgentTeamsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAgentTeamsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAgentTeamsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentTeamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentTeamsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgentTeamsResponseBody
	GetTotalCount() *int64
}

type ListAgentTeamsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of agent-team membership relationships.
	Items []*ListAgentTeamsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of records returned on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor used to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// dXNlci1vZmZzZXQ6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAgentTeamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentTeamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentTeamsResponseBody) GetItems() []*ListAgentTeamsResponseBodyItems {
	return s.Items
}

func (s *ListAgentTeamsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentTeamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentTeamsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentTeamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentTeamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentTeamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentTeamsResponseBody) SetCode(v string) *ListAgentTeamsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetHttpStatusCode(v int32) *ListAgentTeamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetItems(v []*ListAgentTeamsResponseBodyItems) *ListAgentTeamsResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentTeamsResponseBody) SetMaxResults(v int32) *ListAgentTeamsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetMessage(v string) *ListAgentTeamsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetNextToken(v string) *ListAgentTeamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetRequestId(v string) *ListAgentTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetSuccess(v bool) *ListAgentTeamsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentTeamsResponseBody) SetTotalCount(v int64) *ListAgentTeamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentTeamsResponseBody) Validate() error {
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

type ListAgentTeamsResponseBodyItems struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The team ID.
	//
	// example:
	//
	// team-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The team name.
	//
	// example:
	//
	// Default Team
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// The role of the agent in the team.
	//
	// example:
	//
	// MEMBER
	TeamRole *string `json:"teamRole,omitempty" xml:"teamRole,omitempty"`
}

func (s ListAgentTeamsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentTeamsResponseBodyItems) GetTeamId() *string {
	return s.TeamId
}

func (s *ListAgentTeamsResponseBodyItems) GetTeamName() *string {
	return s.TeamName
}

func (s *ListAgentTeamsResponseBodyItems) GetTeamRole() *string {
	return s.TeamRole
}

func (s *ListAgentTeamsResponseBodyItems) SetAgentId(v string) *ListAgentTeamsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListAgentTeamsResponseBodyItems) SetTeamId(v string) *ListAgentTeamsResponseBodyItems {
	s.TeamId = &v
	return s
}

func (s *ListAgentTeamsResponseBodyItems) SetTeamName(v string) *ListAgentTeamsResponseBodyItems {
	s.TeamName = &v
	return s
}

func (s *ListAgentTeamsResponseBodyItems) SetTeamRole(v string) *ListAgentTeamsResponseBodyItems {
	s.TeamRole = &v
	return s
}

func (s *ListAgentTeamsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
