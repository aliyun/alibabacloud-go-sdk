// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskInstanceEvent interface {
	dara.Model
	String() string
	GoString() string
	SetGmtEndTime(v string) *TaskInstanceEvent
	GetGmtEndTime() *string
	SetGmtStartTime(v string) *TaskInstanceEvent
	GetGmtStartTime() *string
	SetMessage(v string) *TaskInstanceEvent
	GetMessage() *string
	SetPodName(v string) *TaskInstanceEvent
	GetPodName() *string
	SetStatus(v string) *TaskInstanceEvent
	GetStatus() *string
	SetWorkloadType(v string) *TaskInstanceEvent
	GetWorkloadType() *string
}

type TaskInstanceEvent struct {
	// The end time.
	//
	// example:
	//
	// 2023-06-22T00:02:00Z
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The message content.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the pod involved in the task.
	//
	// example:
	//
	// test
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The status.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-product type involved in the task.
	//
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s TaskInstanceEvent) String() string {
	return dara.Prettify(s)
}

func (s TaskInstanceEvent) GoString() string {
	return s.String()
}

func (s *TaskInstanceEvent) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *TaskInstanceEvent) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *TaskInstanceEvent) GetMessage() *string {
	return s.Message
}

func (s *TaskInstanceEvent) GetPodName() *string {
	return s.PodName
}

func (s *TaskInstanceEvent) GetStatus() *string {
	return s.Status
}

func (s *TaskInstanceEvent) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *TaskInstanceEvent) SetGmtEndTime(v string) *TaskInstanceEvent {
	s.GmtEndTime = &v
	return s
}

func (s *TaskInstanceEvent) SetGmtStartTime(v string) *TaskInstanceEvent {
	s.GmtStartTime = &v
	return s
}

func (s *TaskInstanceEvent) SetMessage(v string) *TaskInstanceEvent {
	s.Message = &v
	return s
}

func (s *TaskInstanceEvent) SetPodName(v string) *TaskInstanceEvent {
	s.PodName = &v
	return s
}

func (s *TaskInstanceEvent) SetStatus(v string) *TaskInstanceEvent {
	s.Status = &v
	return s
}

func (s *TaskInstanceEvent) SetWorkloadType(v string) *TaskInstanceEvent {
	s.WorkloadType = &v
	return s
}

func (s *TaskInstanceEvent) Validate() error {
	return dara.Validate(s)
}
