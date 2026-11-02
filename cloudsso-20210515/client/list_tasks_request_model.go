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
	// The ID of the access configuration. This parameter is used as a filter condition for the query.
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
	// Format: <Attribute> <Operator> <Value>. The format is not case-sensitive. <Attribute> only supports StartTime. <Operator> only supports ge (greater than or equal to). <Value> must be in the YYYY-MM-DDTHH:mm:ssZ format and can only be a time within the last 7 days.
	//
	// Example: Filter = "StartTime ge 2021-03-15T01:12:23Z". This example queries the list of tasks starting from 2021-03-15T01:12:23Z.
	//
	// > If you do not specify this parameter, the list of tasks created within the last 24 hours is returned by default.
	//
	// example:
	//
	// StartTime ge 2021-03-15T01:12:23Z
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to return the next page of the query results. You do not need to specify NextToken when you call this operation for the first time.
	//
	// When you call this operation for the first time, if the total number of returned entries exceeds the value of MaxResults, the returned data is truncated. Only the number of entries specified by MaxResults is returned, the IsTruncated response parameter is set to true, and a NextToken value is returned. You can use the returned NextToken value to call this operation again while keeping other request parameters unchanged to query the truncated data. You can repeat this process until IsTruncated is set to false, which indicates that all data is returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the CloudSSO identity. This parameter is used as a filter condition for the query. Details:
	//
	// - If PrincipalType is set to User, PrincipalId is the ID of the CloudSSO user.
	//
	// - If PrincipalType is set to Group, PrincipalId is the ID of the CloudSSO group.
	//
	// > PrincipalId and PrincipalType must be specified at the same time to serve as a valid filter condition.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. This parameter is used as a filter condition for the query. Valid values:
	//
	// - User: The CloudSSO user.
	//
	// - Group: The CloudSSO group.
	//
	// > PrincipalId and PrincipalType must be specified at the same time to serve as a valid filter condition.
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The status of the task. This parameter is used as a filter condition for the query. Valid values:
	//
	// - InProgress: The task is running.
	//
	// - Success: The task succeeded.
	//
	// - Failed: The task failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task target. This parameter is used as a filter condition for the query.
	//
	// > TargetId and TargetType must be specified at the same time to serve as a valid filter condition.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task target. This parameter is used as a filter condition for the query. Valid values:
	//
	// - RD-Account: The task target is an account in a resource directory.
	//
	// > TargetId and TargetType must be specified at the same time to serve as a valid filter condition.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the node. This parameter is used as a filter condition for the query. Valid values:
	//
	// - ProvisionAccessConfiguration: Access configuration deployment.
	//
	// - DeprovisionAccessConfiguration: Removes access configuration provisioning.
	//
	// - CreateAccessAssignment: Grants authorization to an account in a resource directory.
	//
	// - DeleteAccessAssignment: Removes authorization from an account in a resource directory.
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
