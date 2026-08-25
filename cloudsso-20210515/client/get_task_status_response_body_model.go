// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskStatusResponseBody
	GetRequestId() *string
	SetTaskStatus(v *GetTaskStatusResponseBodyTaskStatus) *GetTaskStatusResponseBody
	GetTaskStatus() *GetTaskStatusResponseBodyTaskStatus
}

type GetTaskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 005F4623-AE53-504D-830F-44825F7DC211
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status information about the task.
	TaskStatus *GetTaskStatusResponseBodyTaskStatus `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" type:"Struct"`
}

func (s GetTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatusResponseBody) GetTaskStatus() *GetTaskStatusResponseBodyTaskStatus {
	return s.TaskStatus
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTaskStatus(v *GetTaskStatusResponseBodyTaskStatus) *GetTaskStatusResponseBody {
	s.TaskStatus = v
	return s
}

func (s *GetTaskStatusResponseBody) Validate() error {
	if s.TaskStatus != nil {
		if err := s.TaskStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskStatusResponseBodyTaskStatus struct {
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

func (s GetTaskStatusResponseBodyTaskStatus) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBodyTaskStatus) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetFailureReason() *string {
	return s.FailureReason
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetStatus() *string {
	return s.Status
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskStatusResponseBodyTaskStatus) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetEndTime(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.EndTime = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetFailureReason(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.FailureReason = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetStartTime(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.StartTime = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetStatus(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetTaskId(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) SetTaskType(v string) *GetTaskStatusResponseBodyTaskStatus {
	s.TaskType = &v
	return s
}

func (s *GetTaskStatusResponseBodyTaskStatus) Validate() error {
	return dara.Validate(s)
}
