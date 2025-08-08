// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTask interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentBytes(v string) *ModelTask
	GetCurrentBytes() *string
	SetErrCode(v string) *ModelTask
	GetErrCode() *string
	SetErrMsg(v string) *ModelTask
	GetErrMsg() *string
	SetExtra(v interface{}) *ModelTask
	GetExtra() interface{}
	SetFileSize(v float64) *ModelTask
	GetFileSize() *float64
	SetFinishTime(v float64) *ModelTask
	GetFinishTime() *float64
	SetFinished(v bool) *ModelTask
	GetFinished() *bool
	SetFinishedTime(v float64) *ModelTask
	GetFinishedTime() *float64
	SetId(v string) *ModelTask
	GetId() *string
	SetParams(v string) *ModelTask
	GetParams() *string
	SetResult(v interface{}) *ModelTask
	GetResult() interface{}
	SetSpeed(v string) *ModelTask
	GetSpeed() *string
	SetStartTime(v float64) *ModelTask
	GetStartTime() *float64
	SetStatus(v string) *ModelTask
	GetStatus() *string
	SetTaskId(v string) *ModelTask
	GetTaskId() *string
	SetTaskType(v string) *ModelTask
	GetTaskType() *string
	SetTotal(v float32) *ModelTask
	GetTotal() *float32
	SetTotalBytes(v string) *ModelTask
	GetTotalBytes() *string
	SetUpdateTime(v float64) *ModelTask
	GetUpdateTime() *float64
}

type ModelTask struct {
	CurrentBytes *string     `json:"currentBytes,omitempty" xml:"currentBytes,omitempty"`
	ErrCode      *string     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg       *string     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Extra        interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	// Deprecated
	FileSize *float64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// Deprecated
	FinishTime   *float64    `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Finished     *bool       `json:"finished,omitempty" xml:"finished,omitempty"`
	FinishedTime *float64    `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Id           *string     `json:"id,omitempty" xml:"id,omitempty"`
	Params       *string     `json:"params,omitempty" xml:"params,omitempty"`
	Result       interface{} `json:"result,omitempty" xml:"result,omitempty"`
	Speed        *string     `json:"speed,omitempty" xml:"speed,omitempty"`
	StartTime    *float64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *string     `json:"status,omitempty" xml:"status,omitempty"`
	TaskId       *string     `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskType     *string     `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// Deprecated
	Total      *float32 `json:"total,omitempty" xml:"total,omitempty"`
	TotalBytes *string  `json:"totalBytes,omitempty" xml:"totalBytes,omitempty"`
	UpdateTime *float64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ModelTask) String() string {
	return dara.Prettify(s)
}

func (s ModelTask) GoString() string {
	return s.String()
}

func (s *ModelTask) GetCurrentBytes() *string {
	return s.CurrentBytes
}

func (s *ModelTask) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelTask) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ModelTask) GetExtra() interface{} {
	return s.Extra
}

func (s *ModelTask) GetFileSize() *float64 {
	return s.FileSize
}

func (s *ModelTask) GetFinishTime() *float64 {
	return s.FinishTime
}

func (s *ModelTask) GetFinished() *bool {
	return s.Finished
}

func (s *ModelTask) GetFinishedTime() *float64 {
	return s.FinishedTime
}

func (s *ModelTask) GetId() *string {
	return s.Id
}

func (s *ModelTask) GetParams() *string {
	return s.Params
}

func (s *ModelTask) GetResult() interface{} {
	return s.Result
}

func (s *ModelTask) GetSpeed() *string {
	return s.Speed
}

func (s *ModelTask) GetStartTime() *float64 {
	return s.StartTime
}

func (s *ModelTask) GetStatus() *string {
	return s.Status
}

func (s *ModelTask) GetTaskId() *string {
	return s.TaskId
}

func (s *ModelTask) GetTaskType() *string {
	return s.TaskType
}

func (s *ModelTask) GetTotal() *float32 {
	return s.Total
}

func (s *ModelTask) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *ModelTask) GetUpdateTime() *float64 {
	return s.UpdateTime
}

func (s *ModelTask) SetCurrentBytes(v string) *ModelTask {
	s.CurrentBytes = &v
	return s
}

func (s *ModelTask) SetErrCode(v string) *ModelTask {
	s.ErrCode = &v
	return s
}

func (s *ModelTask) SetErrMsg(v string) *ModelTask {
	s.ErrMsg = &v
	return s
}

func (s *ModelTask) SetExtra(v interface{}) *ModelTask {
	s.Extra = v
	return s
}

func (s *ModelTask) SetFileSize(v float64) *ModelTask {
	s.FileSize = &v
	return s
}

func (s *ModelTask) SetFinishTime(v float64) *ModelTask {
	s.FinishTime = &v
	return s
}

func (s *ModelTask) SetFinished(v bool) *ModelTask {
	s.Finished = &v
	return s
}

func (s *ModelTask) SetFinishedTime(v float64) *ModelTask {
	s.FinishedTime = &v
	return s
}

func (s *ModelTask) SetId(v string) *ModelTask {
	s.Id = &v
	return s
}

func (s *ModelTask) SetParams(v string) *ModelTask {
	s.Params = &v
	return s
}

func (s *ModelTask) SetResult(v interface{}) *ModelTask {
	s.Result = v
	return s
}

func (s *ModelTask) SetSpeed(v string) *ModelTask {
	s.Speed = &v
	return s
}

func (s *ModelTask) SetStartTime(v float64) *ModelTask {
	s.StartTime = &v
	return s
}

func (s *ModelTask) SetStatus(v string) *ModelTask {
	s.Status = &v
	return s
}

func (s *ModelTask) SetTaskId(v string) *ModelTask {
	s.TaskId = &v
	return s
}

func (s *ModelTask) SetTaskType(v string) *ModelTask {
	s.TaskType = &v
	return s
}

func (s *ModelTask) SetTotal(v float32) *ModelTask {
	s.Total = &v
	return s
}

func (s *ModelTask) SetTotalBytes(v string) *ModelTask {
	s.TotalBytes = &v
	return s
}

func (s *ModelTask) SetUpdateTime(v float64) *ModelTask {
	s.UpdateTime = &v
	return s
}

func (s *ModelTask) Validate() error {
	return dara.Validate(s)
}
