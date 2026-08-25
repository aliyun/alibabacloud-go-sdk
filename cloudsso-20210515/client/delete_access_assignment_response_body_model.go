// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessAssignmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessAssignmentResponseBody
	GetRequestId() *string
	SetTask(v *DeleteAccessAssignmentResponseBodyTask) *DeleteAccessAssignmentResponseBody
	GetTask() *DeleteAccessAssignmentResponseBodyTask
}

type DeleteAccessAssignmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5C9D0CF4-5CE8-5CE6-932A-826EF4ADD007
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task *DeleteAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s DeleteAccessAssignmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessAssignmentResponseBody) GetTask() *DeleteAccessAssignmentResponseBodyTask {
	return s.Task
}

func (s *DeleteAccessAssignmentResponseBody) SetRequestId(v string) *DeleteAccessAssignmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBody) SetTask(v *DeleteAccessAssignmentResponseBodyTask) *DeleteAccessAssignmentResponseBody {
	s.Task = v
	return s
}

func (s *DeleteAccessAssignmentResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAccessAssignmentResponseBodyTask struct {
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
	// InProgress
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
	// The task type. The value is fixed as DeleteAccessAssignment, which indicates that access permissions on an account in your resource directory are removed.
	//
	// example:
	//
	// DeleteAccessAssignment
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeleteAccessAssignmentResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessAssignmentResponseBodyTask) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTargetId() *string {
	return s.TargetId
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTargetName() *string {
	return s.TargetName
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTargetPath() *string {
	return s.TargetPath
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTargetType() *string {
	return s.TargetType
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteAccessAssignmentResponseBodyTask) GetTaskType() *string {
	return s.TaskType
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetAccessConfigurationId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetAccessConfigurationName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetOriginTargetId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.OriginTargetId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetPrincipalType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetStatus(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.Status = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetPath(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetPathName(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTargetType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTaskId(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) SetTaskType(v string) *DeleteAccessAssignmentResponseBodyTask {
	s.TaskType = &v
	return s
}

func (s *DeleteAccessAssignmentResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
