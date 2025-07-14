// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppStackTask interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *AppStackTask
	GetCreateTime() *int64
	SetEndTime(v int64) *AppStackTask
	GetEndTime() *int64
	SetInstanceId(v string) *AppStackTask
	GetInstanceId() *string
	SetStackId(v string) *AppStackTask
	GetStackId() *string
	SetStartTime(v int64) *AppStackTask
	GetStartTime() *int64
	SetStatus(v string) *AppStackTask
	GetStatus() *string
	SetSteps(v []*AppStackTaskSteps) *AppStackTask
	GetSteps() []*AppStackTaskSteps
	SetTaskId(v string) *AppStackTask
	GetTaskId() *string
}

type AppStackTask struct {
	// example:
	//
	// 1706518652
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1706518652
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// i-789y
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// palworld
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// example:
	//
	// 1706518652
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// WAIT
	Status *string              `json:"Status,omitempty" xml:"Status,omitempty"`
	Steps  []*AppStackTaskSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// example:
	//
	// t-789y-deploy
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AppStackTask) String() string {
	return dara.Prettify(s)
}

func (s AppStackTask) GoString() string {
	return s.String()
}

func (s *AppStackTask) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AppStackTask) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AppStackTask) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppStackTask) GetStackId() *string {
	return s.StackId
}

func (s *AppStackTask) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AppStackTask) GetStatus() *string {
	return s.Status
}

func (s *AppStackTask) GetSteps() []*AppStackTaskSteps {
	return s.Steps
}

func (s *AppStackTask) GetTaskId() *string {
	return s.TaskId
}

func (s *AppStackTask) SetCreateTime(v int64) *AppStackTask {
	s.CreateTime = &v
	return s
}

func (s *AppStackTask) SetEndTime(v int64) *AppStackTask {
	s.EndTime = &v
	return s
}

func (s *AppStackTask) SetInstanceId(v string) *AppStackTask {
	s.InstanceId = &v
	return s
}

func (s *AppStackTask) SetStackId(v string) *AppStackTask {
	s.StackId = &v
	return s
}

func (s *AppStackTask) SetStartTime(v int64) *AppStackTask {
	s.StartTime = &v
	return s
}

func (s *AppStackTask) SetStatus(v string) *AppStackTask {
	s.Status = &v
	return s
}

func (s *AppStackTask) SetSteps(v []*AppStackTaskSteps) *AppStackTask {
	s.Steps = v
	return s
}

func (s *AppStackTask) SetTaskId(v string) *AppStackTask {
	s.TaskId = &v
	return s
}

func (s *AppStackTask) Validate() error {
	return dara.Validate(s)
}

type AppStackTaskSteps struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1706518652
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// VPC_CREATE_NETWORK
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 初始化 VPC 网络
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1706518652
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// WAIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AppStackTaskSteps) String() string {
	return dara.Prettify(s)
}

func (s AppStackTaskSteps) GoString() string {
	return s.String()
}

func (s *AppStackTaskSteps) GetCode() *string {
	return s.Code
}

func (s *AppStackTaskSteps) GetDuration() *int64 {
	return s.Duration
}

func (s *AppStackTaskSteps) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AppStackTaskSteps) GetId() *string {
	return s.Id
}

func (s *AppStackTaskSteps) GetMessage() *string {
	return s.Message
}

func (s *AppStackTaskSteps) GetName() *string {
	return s.Name
}

func (s *AppStackTaskSteps) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AppStackTaskSteps) GetStatus() *string {
	return s.Status
}

func (s *AppStackTaskSteps) SetCode(v string) *AppStackTaskSteps {
	s.Code = &v
	return s
}

func (s *AppStackTaskSteps) SetDuration(v int64) *AppStackTaskSteps {
	s.Duration = &v
	return s
}

func (s *AppStackTaskSteps) SetEndTime(v int64) *AppStackTaskSteps {
	s.EndTime = &v
	return s
}

func (s *AppStackTaskSteps) SetId(v string) *AppStackTaskSteps {
	s.Id = &v
	return s
}

func (s *AppStackTaskSteps) SetMessage(v string) *AppStackTaskSteps {
	s.Message = &v
	return s
}

func (s *AppStackTaskSteps) SetName(v string) *AppStackTaskSteps {
	s.Name = &v
	return s
}

func (s *AppStackTaskSteps) SetStartTime(v int64) *AppStackTaskSteps {
	s.StartTime = &v
	return s
}

func (s *AppStackTaskSteps) SetStatus(v string) *AppStackTaskSteps {
	s.Status = &v
	return s
}

func (s *AppStackTaskSteps) Validate() error {
	return dara.Validate(s)
}
