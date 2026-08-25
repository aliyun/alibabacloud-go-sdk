// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUsersResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody
	GetItems() []*ListUsersResponseBodyItems
	SetMaxResults(v int32) *ListUsersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListUsersResponseBody
	GetTotalCount() *int64
}

type ListUsersResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                        `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListUsersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// dXNlci1vZmZzZXQ6MTA
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

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUsersResponseBody) GetItems() []*ListUsersResponseBodyItems {
	return s.Items
}

func (s *ListUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetHttpStatusCode(v int32) *ListUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUsersResponseBody) SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody {
	s.Items = v
	return s
}

func (s *ListUsersResponseBody) SetMaxResults(v int32) *ListUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetNextToken(v string) *ListUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
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

type ListUsersResponseBodyItems struct {
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
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
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 智能体运营组成员
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
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
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListUsersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyItems) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *ListUsersResponseBodyItems) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ListUsersResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListUsersResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyItems) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListUsersResponseBodyItems) GetNote() *string {
	return s.Note
}

func (s *ListUsersResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUsersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListUsersResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListUsersResponseBodyItems) SetAgentCoreUserId(v string) *ListUsersResponseBodyItems {
	s.AgentCoreUserId = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetAuthMethod(v string) *ListUsersResponseBodyItems {
	s.AuthMethod = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetCreatedAt(v string) *ListUsersResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetDisplayName(v string) *ListUsersResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetEmail(v string) *ListUsersResponseBodyItems {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetName(v string) *ListUsersResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetNote(v string) *ListUsersResponseBodyItems {
	s.Note = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetRegionId(v string) *ListUsersResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetStatus(v string) *ListUsersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetUpdatedAt(v string) *ListUsersResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetWorkspaceId(v string) *ListUsersResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListUsersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
