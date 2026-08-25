// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *ListTasksRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *ListTasksRequest
	GetDirectoryId() *string
	SetFilter(v string) *ListTasksRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTasksRequest
	GetNextToken() *string
	SetPrincipalId(v string) *ListTasksRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *ListTasksRequest
	GetPrincipalType() *string
	SetStatus(v string) *ListTasksRequest
	GetStatus() *string
	SetTargetId(v string) *ListTasksRequest
	GetTargetId() *string
	SetTargetType(v string) *ListTasksRequest
	GetTargetType() *string
	SetTaskType(v string) *ListTasksRequest
	GetTaskType() *string
}

type ListTasksRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// The condition is not case-sensitive. The condition must be in the StartTime ge YYYY-MM-DDTHH:mm:SSZ format. You must set YYYY-MM-DDTHH:mm:SSZ to a value that is no more than 7 days from the current time. ge indicates Greater Than or Equals.
	//
	// For example, if you set the Filter parameter to StartTime ge 2021-03-15T01:12:23Z, the operation queries the tasks from 2021-03-15T01:12:23 GMT.
	//
	// > If you do not specify this parameter, the operation queries the tasks within the previous 24 hours by default.
	//
	// example:
	//
	// StartTime ge 2021-03-15T01:12:23Z
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the CloudSSO identity. The ID can be used to filter access permissions.
	//
	// - If you set `PrincipalType` to `User`, set `PrincipalId` to the ID of the CloudSSO user.
	//
	// - If you set `PrincipalType` to `Group`, set `PrincipalId` to the ID of the CloudSSO group.
	//
	// > You can use the type to filter access permissions only if you specify both `PrincipalId` and `PrincipalType`.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. The type can be used to filter access permissions. Valid values:
	//
	// - User
	//
	// - Group
	//
	// > You can use the type to filter access permissions only if you specify both `PrincipalId` and `PrincipalType`.
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task. The ID can be used to filter tasks. Valid values:
	//
	// - InProgress: The task is running.
	//
	// - Success: The task is successful.
	//
	// - Failed: The task failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object. The ID can be used to filter access permissions.
	//
	// > You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. The type can be used to filter access permissions.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// > You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the task. The type can be used to filter tasks. Valid values:
	//
	// - ProvisionAccessConfiguration: An access configuration is provisioned.
	//
	// - DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	//
	// - CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	//
	// - DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	//
	// example:
	//
	// CreateAccessAssignment
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListTasksRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTasksRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListTasksRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListTasksRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksRequest) SetAccessConfigurationId(v string) *ListTasksRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListTasksRequest) SetDirectoryId(v string) *ListTasksRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListTasksRequest) SetFilter(v string) *ListTasksRequest {
	s.Filter = &v
	return s
}

func (s *ListTasksRequest) SetMaxResults(v int32) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetPrincipalId(v string) *ListTasksRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListTasksRequest) SetPrincipalType(v string) *ListTasksRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTargetId(v string) *ListTasksRequest {
	s.TargetId = &v
	return s
}

func (s *ListTasksRequest) SetTargetType(v string) *ListTasksRequest {
	s.TargetType = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
