// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppMaterialTask interface {
	dara.Model
	String() string
	GoString() string
	SetCompleteTime(v string) *AppMaterialTask
	GetCompleteTime() *string
	SetCompleteTimeFormat(v string) *AppMaterialTask
	GetCompleteTimeFormat() *string
	SetFailReason(v string) *AppMaterialTask
	GetFailReason() *string
	SetFinalFileUrls(v []*string) *AppMaterialTask
	GetFinalFileUrls() []*string
	SetStatus(v string) *AppMaterialTask
	GetStatus() *string
	SetSubStatus(v string) *AppMaterialTask
	GetSubStatus() *string
	SetSubmitTime(v string) *AppMaterialTask
	GetSubmitTime() *string
	SetTaskId(v string) *AppMaterialTask
	GetTaskId() *string
	SetTaskParam(v string) *AppMaterialTask
	GetTaskParam() *string
	SetTaskType(v string) *AppMaterialTask
	GetTaskType() *string
}

type AppMaterialTask struct {
	CompleteTime       *string   `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CompleteTimeFormat *string   `json:"CompleteTimeFormat,omitempty" xml:"CompleteTimeFormat,omitempty"`
	FailReason         *string   `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	FinalFileUrls      []*string `json:"FinalFileUrls,omitempty" xml:"FinalFileUrls,omitempty" type:"Repeated"`
	Status             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus          *string   `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	SubmitTime         *string   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	TaskId             *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskParam          *string   `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	TaskType           *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s AppMaterialTask) String() string {
	return dara.Prettify(s)
}

func (s AppMaterialTask) GoString() string {
	return s.String()
}

func (s *AppMaterialTask) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *AppMaterialTask) GetCompleteTimeFormat() *string {
	return s.CompleteTimeFormat
}

func (s *AppMaterialTask) GetFailReason() *string {
	return s.FailReason
}

func (s *AppMaterialTask) GetFinalFileUrls() []*string {
	return s.FinalFileUrls
}

func (s *AppMaterialTask) GetStatus() *string {
	return s.Status
}

func (s *AppMaterialTask) GetSubStatus() *string {
	return s.SubStatus
}

func (s *AppMaterialTask) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *AppMaterialTask) GetTaskId() *string {
	return s.TaskId
}

func (s *AppMaterialTask) GetTaskParam() *string {
	return s.TaskParam
}

func (s *AppMaterialTask) GetTaskType() *string {
	return s.TaskType
}

func (s *AppMaterialTask) SetCompleteTime(v string) *AppMaterialTask {
	s.CompleteTime = &v
	return s
}

func (s *AppMaterialTask) SetCompleteTimeFormat(v string) *AppMaterialTask {
	s.CompleteTimeFormat = &v
	return s
}

func (s *AppMaterialTask) SetFailReason(v string) *AppMaterialTask {
	s.FailReason = &v
	return s
}

func (s *AppMaterialTask) SetFinalFileUrls(v []*string) *AppMaterialTask {
	s.FinalFileUrls = v
	return s
}

func (s *AppMaterialTask) SetStatus(v string) *AppMaterialTask {
	s.Status = &v
	return s
}

func (s *AppMaterialTask) SetSubStatus(v string) *AppMaterialTask {
	s.SubStatus = &v
	return s
}

func (s *AppMaterialTask) SetSubmitTime(v string) *AppMaterialTask {
	s.SubmitTime = &v
	return s
}

func (s *AppMaterialTask) SetTaskId(v string) *AppMaterialTask {
	s.TaskId = &v
	return s
}

func (s *AppMaterialTask) SetTaskParam(v string) *AppMaterialTask {
	s.TaskParam = &v
	return s
}

func (s *AppMaterialTask) SetTaskType(v string) *AppMaterialTask {
	s.TaskType = &v
	return s
}

func (s *AppMaterialTask) Validate() error {
	return dara.Validate(s)
}
