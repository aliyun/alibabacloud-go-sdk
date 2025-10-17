// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkAnalyzeLogTask interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SparkAnalyzeLogTask
	GetDBClusterId() *string
	SetResult(v *LogAnalyzeResult) *SparkAnalyzeLogTask
	GetResult() *LogAnalyzeResult
	SetRuleMatched(v bool) *SparkAnalyzeLogTask
	GetRuleMatched() *bool
	SetStartedTimeInMillis(v int64) *SparkAnalyzeLogTask
	GetStartedTimeInMillis() *int64
	SetSubmittedTimeInMillis(v int64) *SparkAnalyzeLogTask
	GetSubmittedTimeInMillis() *int64
	SetTaskErrMsg(v string) *SparkAnalyzeLogTask
	GetTaskErrMsg() *string
	SetTaskId(v int64) *SparkAnalyzeLogTask
	GetTaskId() *int64
	SetTaskState(v string) *SparkAnalyzeLogTask
	GetTaskState() *string
	SetTerminatedTimeInMillis(v int64) *SparkAnalyzeLogTask
	GetTerminatedTimeInMillis() *int64
	SetUserId(v int64) *SparkAnalyzeLogTask
	GetUserId() *int64
}

type SparkAnalyzeLogTask struct {
	// example:
	//
	// amv-adbxxxxx
	DBClusterId *string           `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Result      *LogAnalyzeResult `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	RuleMatched *bool `json:"RuleMatched,omitempty" xml:"RuleMatched,omitempty"`
	// example:
	//
	// 1672123543000
	StartedTimeInMillis *int64 `json:"StartedTimeInMillis,omitempty" xml:"StartedTimeInMillis,omitempty"`
	// example:
	//
	// 1672123543000
	SubmittedTimeInMillis *int64 `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
	// example:
	//
	// Driver log not found
	TaskErrMsg *string `json:"TaskErrMsg,omitempty" xml:"TaskErrMsg,omitempty"`
	// example:
	//
	// 10
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// WAITING
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// example:
	//
	// 1672123543000
	TerminatedTimeInMillis *int64 `json:"TerminatedTimeInMillis,omitempty" xml:"TerminatedTimeInMillis,omitempty"`
	// example:
	//
	// 13719918xxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SparkAnalyzeLogTask) String() string {
	return dara.Prettify(s)
}

func (s SparkAnalyzeLogTask) GoString() string {
	return s.String()
}

func (s *SparkAnalyzeLogTask) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SparkAnalyzeLogTask) GetResult() *LogAnalyzeResult {
	return s.Result
}

func (s *SparkAnalyzeLogTask) GetRuleMatched() *bool {
	return s.RuleMatched
}

func (s *SparkAnalyzeLogTask) GetStartedTimeInMillis() *int64 {
	return s.StartedTimeInMillis
}

func (s *SparkAnalyzeLogTask) GetSubmittedTimeInMillis() *int64 {
	return s.SubmittedTimeInMillis
}

func (s *SparkAnalyzeLogTask) GetTaskErrMsg() *string {
	return s.TaskErrMsg
}

func (s *SparkAnalyzeLogTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SparkAnalyzeLogTask) GetTaskState() *string {
	return s.TaskState
}

func (s *SparkAnalyzeLogTask) GetTerminatedTimeInMillis() *int64 {
	return s.TerminatedTimeInMillis
}

func (s *SparkAnalyzeLogTask) GetUserId() *int64 {
	return s.UserId
}

func (s *SparkAnalyzeLogTask) SetDBClusterId(v string) *SparkAnalyzeLogTask {
	s.DBClusterId = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetResult(v *LogAnalyzeResult) *SparkAnalyzeLogTask {
	s.Result = v
	return s
}

func (s *SparkAnalyzeLogTask) SetRuleMatched(v bool) *SparkAnalyzeLogTask {
	s.RuleMatched = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetStartedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.StartedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetSubmittedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.SubmittedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskErrMsg(v string) *SparkAnalyzeLogTask {
	s.TaskErrMsg = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskId(v int64) *SparkAnalyzeLogTask {
	s.TaskId = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskState(v string) *SparkAnalyzeLogTask {
	s.TaskState = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTerminatedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.TerminatedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetUserId(v int64) *SparkAnalyzeLogTask {
	s.UserId = &v
	return s
}

func (s *SparkAnalyzeLogTask) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
