// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessAssignmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAccessAssignmentResponseBody
	GetRequestId() *string
	SetTask(v *CreateAccessAssignmentResponseBodyTask) *CreateAccessAssignmentResponseBody
	GetTask() *CreateAccessAssignmentResponseBodyTask
}

type CreateAccessAssignmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4726AA56-E138-5C99-85E4-F493536D042F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried task.
	Task *CreateAccessAssignmentResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s CreateAccessAssignmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessAssignmentResponseBody) GetTask() *CreateAccessAssignmentResponseBodyTask {
	return s.Task
}

func (s *CreateAccessAssignmentResponseBody) SetRequestId(v string) *CreateAccessAssignmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBody) SetTask(v *CreateAccessAssignmentResponseBodyTask) *CreateAccessAssignmentResponseBody {
	s.Task = v
	return s
}

func (s *CreateAccessAssignmentResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessAssignmentResponseBodyTask struct {
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
	// The ID of the job.
	//
	// example:
	//
	// t-sh6tceylhvgejpip****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as CreateAccessAssignment, which indicates that access permissions on an account in your resource directory are assigned.
	//
	// example:
	//
	// CreateAccessAssignment
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateAccessAssignmentResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessAssignmentResponseBodyTask) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponseBodyTask) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *CreateAccessAssignmentResponseBodyTask) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *CreateAccessAssignmentResponseBodyTask) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *CreateAccessAssignmentResponseBodyTask) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateAccessAssignmentResponseBodyTask) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *CreateAccessAssignmentResponseBodyTask) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateAccessAssignmentResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTargetName() *string {
	return s.TargetName
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAccessAssignmentResponseBodyTask) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateAccessAssignmentResponseBodyTask) SetAccessConfigurationId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetAccessConfigurationName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetOriginTargetId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.OriginTargetId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetPrincipalType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.PrincipalType = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetStatus(v string) *CreateAccessAssignmentResponseBodyTask {
	s.Status = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetPath(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetPath = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetPathName(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetPathName = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTargetType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTaskId(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) SetTaskType(v string) *CreateAccessAssignmentResponseBodyTask {
	s.TaskType = &v
	return s
}

func (s *CreateAccessAssignmentResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
