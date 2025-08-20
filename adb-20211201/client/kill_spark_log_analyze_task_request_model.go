// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkLogAnalyzeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *KillSparkLogAnalyzeTaskRequest
	GetTaskId() *int64
}

type KillSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark log analysis task. You can call the ListSparkLogAnalyzeTasks operation to query the IDs and states of all analysis tasks in the current cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s KillSparkLogAnalyzeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *KillSparkLogAnalyzeTaskRequest) SetTaskId(v int64) *KillSparkLogAnalyzeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *KillSparkLogAnalyzeTaskRequest) Validate() error {
	return dara.Validate(s)
}
