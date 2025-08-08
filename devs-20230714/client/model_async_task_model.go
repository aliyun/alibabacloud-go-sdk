// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelAsyncTask interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModelAsyncTask
	GetErrCode() *string
	SetErrMsg(v string) *ModelAsyncTask
	GetErrMsg() *string
	SetFinished(v bool) *ModelAsyncTask
	GetFinished() *bool
	SetFinishedTime(v int64) *ModelAsyncTask
	GetFinishedTime() *int64
	SetResult(v interface{}) *ModelAsyncTask
	GetResult() interface{}
	SetStartTime(v int64) *ModelAsyncTask
	GetStartTime() *int64
	SetTaskType(v string) *ModelAsyncTask
	GetTaskType() *string
	SetUpdateTime(v int64) *ModelAsyncTask
	GetUpdateTime() *int64
}

type ModelAsyncTask struct {
	ErrCode      *string     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg       *string     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Finished     *bool       `json:"finished,omitempty" xml:"finished,omitempty"`
	FinishedTime *int64      `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Result       interface{} `json:"result,omitempty" xml:"result,omitempty"`
	StartTime    *int64      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskType     *string     `json:"taskType,omitempty" xml:"taskType,omitempty"`
	UpdateTime   *int64      `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ModelAsyncTask) String() string {
	return dara.Prettify(s)
}

func (s ModelAsyncTask) GoString() string {
	return s.String()
}

func (s *ModelAsyncTask) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelAsyncTask) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ModelAsyncTask) GetFinished() *bool {
	return s.Finished
}

func (s *ModelAsyncTask) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ModelAsyncTask) GetResult() interface{} {
	return s.Result
}

func (s *ModelAsyncTask) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelAsyncTask) GetTaskType() *string {
	return s.TaskType
}

func (s *ModelAsyncTask) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ModelAsyncTask) SetErrCode(v string) *ModelAsyncTask {
	s.ErrCode = &v
	return s
}

func (s *ModelAsyncTask) SetErrMsg(v string) *ModelAsyncTask {
	s.ErrMsg = &v
	return s
}

func (s *ModelAsyncTask) SetFinished(v bool) *ModelAsyncTask {
	s.Finished = &v
	return s
}

func (s *ModelAsyncTask) SetFinishedTime(v int64) *ModelAsyncTask {
	s.FinishedTime = &v
	return s
}

func (s *ModelAsyncTask) SetResult(v interface{}) *ModelAsyncTask {
	s.Result = v
	return s
}

func (s *ModelAsyncTask) SetStartTime(v int64) *ModelAsyncTask {
	s.StartTime = &v
	return s
}

func (s *ModelAsyncTask) SetTaskType(v string) *ModelAsyncTask {
	s.TaskType = &v
	return s
}

func (s *ModelAsyncTask) SetUpdateTime(v int64) *ModelAsyncTask {
	s.UpdateTime = &v
	return s
}

func (s *ModelAsyncTask) Validate() error {
	return dara.Validate(s)
}
