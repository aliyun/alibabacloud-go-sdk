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
	// The list of users.
	Items []*ListUsersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page that takes effect for this query.
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
	// The pagination token for the next page. This value is empty if no more pages are available.
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
	// The total number of users that match the query conditions.
	//
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
	// The user ID.
	//
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// The authentication method of the user. A value of password indicates local password authentication in the workspace. Values of dingtalk and feishu indicate that the user is synchronized and authenticated by the corresponding external identity provider.
	//
	// example:
	//
	// password
	AuthMethod *string `json:"authMethod,omitempty" xml:"authMethod,omitempty"`
	// The creation time in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The display name of the user. The value is 1 to 32 characters in length.
	//
	// example:
	//
	// John Smith
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user. The value can be up to 256 characters in length.
	//
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The username. The value must be unique within the workspace and can contain only lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or digit. The value is 1 to 32 characters in length.
	//
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The remarks of the user. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// Agent operations team member
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The user status. Valid values: Creating, Active, Updating, Deleting, Failed, DeleteFailed.
	//
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time of the last modification in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
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
