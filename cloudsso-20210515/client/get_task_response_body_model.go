// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody
	GetTask() *GetTaskResponseBodyTask
}

type GetTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 923CA5E8-57BF-5E15-8BA6-E75A966B7E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task *GetTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetTask() *GetTaskResponseBodyTask {
	return s.Task
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTask struct {
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
	// 2021-11-05T02:58:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task failure.
	//
	// >  This parameter is returned only when the value of `Status` is `Failed`.
	//
	// example:
	//
	// No Permission.
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// example:
	//
	// 17xxxxxxxxxxxx73
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
	// 	- User
	//
	// 	- Group
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2021-11-05T02:58:07Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- InProgress: The task is running.
	//
	// 	- Success: The task is successful.
	//
	// 	- Failed: The task failed.
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
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-shfqw1u1edszvxw5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// 	- ProvisionAccessConfiguration: An access configuration is provisioned.
	//
	// 	- DeprovisionAccessConfiguration: An access configuration is de-provisioned.
	//
	// 	- CreateAccessAssignment: Access permissions on an account in the resource directory are assigned.
	//
	// 	- DeleteAccessAssignment: Access permissions on an account in the resource directory are removed.
	//
	// example:
	//
	// DeleteAccessAssignment
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *GetTaskResponseBodyTask) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *GetTaskResponseBodyTask) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskResponseBodyTask) GetFailureReason() *string {
	return s.FailureReason
}

func (s *GetTaskResponseBodyTask) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *GetTaskResponseBodyTask) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetTaskResponseBodyTask) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetTaskResponseBodyTask) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetTaskResponseBodyTask) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyTask) GetTargetId() *string {
	return s.TargetId
}

func (s *GetTaskResponseBodyTask) GetTargetName() *string {
	return s.TargetName
}

func (s *GetTaskResponseBodyTask) GetTargetPath() *string {
	return s.TargetPath
}

func (s *GetTaskResponseBodyTask) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *GetTaskResponseBodyTask) GetTargetType() *string {
	return s.TargetType
}

func (s *GetTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBodyTask) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskResponseBodyTask) SetAccessConfigurationId(v string) *GetTaskResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetAccessConfigurationName(v string) *GetTaskResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetEndTime(v string) *GetTaskResponseBodyTask {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetFailureReason(v string) *GetTaskResponseBodyTask {
	s.FailureReason = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetOriginTargetId(v string) *GetTaskResponseBodyTask {
	s.OriginTargetId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalId(v string) *GetTaskResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalName(v string) *GetTaskResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPrincipalType(v string) *GetTaskResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStartTime(v string) *GetTaskResponseBodyTask {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetId(v string) *GetTaskResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetName(v string) *GetTaskResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetPath(v string) *GetTaskResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetPathName(v string) *GetTaskResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTargetType(v string) *GetTaskResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskType(v string) *GetTaskResponseBodyTask {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
