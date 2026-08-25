// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListTasksResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody
	GetTasks() []*ListTasksResponseBodyTasks
	SetTotalCounts(v int32) *ListTasksResponseBody
	GetTotalCounts() *int32
}

type ListTasksResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// > This parameter is returned only when the value of `IsTruncated` is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C0DA2DFC-EB18-59EF-BD82-C30862EBA3A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTasks() []*ListTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListTasksResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListTasksResponseBody) SetIsTruncated(v bool) *ListTasksResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListTasksResponseBody) SetMaxResults(v int32) *ListTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCounts(v int32) *ListTasksResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyTasks struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2021-11-09T05:50:50Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task failure.
	//
	// > This parameter is returned only when the value of `Status` is `Failed`.
	//
	// example:
	//
	// No Permission.
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the CloudSSO identity.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	//
	// example:
	//
	// Alice
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// - User
	//
	// - Group
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2021-11-09T05:50:50Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
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
	// The ID of the task object.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	//
	// example:
	//
	// dev-test
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object.
	//
	// The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// t-sh5k4gesm6twlrqb****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
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

func (s ListTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListTasksResponseBodyTasks) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *ListTasksResponseBodyTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTasksResponseBodyTasks) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListTasksResponseBodyTasks) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ListTasksResponseBodyTasks) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListTasksResponseBodyTasks) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListTasksResponseBodyTasks) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListTasksResponseBodyTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListTasksResponseBodyTasks) GetTargetId() *string {
	return s.TargetId
}

func (s *ListTasksResponseBodyTasks) GetTargetName() *string {
	return s.TargetName
}

func (s *ListTasksResponseBodyTasks) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListTasksResponseBodyTasks) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *ListTasksResponseBodyTasks) GetTargetType() *string {
	return s.TargetType
}

func (s *ListTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksResponseBodyTasks) SetAccessConfigurationId(v string) *ListTasksResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetAccessConfigurationName(v string) *ListTasksResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetEndTime(v string) *ListTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetFailureReason(v string) *ListTasksResponseBodyTasks {
	s.FailureReason = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetOriginTargetId(v string) *ListTasksResponseBodyTasks {
	s.OriginTargetId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalId(v string) *ListTasksResponseBodyTasks {
	s.PrincipalId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalName(v string) *ListTasksResponseBodyTasks {
	s.PrincipalName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetPrincipalType(v string) *ListTasksResponseBodyTasks {
	s.PrincipalType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStartTime(v string) *ListTasksResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetId(v string) *ListTasksResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetName(v string) *ListTasksResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetPath(v string) *ListTasksResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetPathName(v string) *ListTasksResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTargetType(v string) *ListTasksResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskType(v string) *ListTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
