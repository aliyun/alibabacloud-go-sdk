// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLogAnalyzeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *GetSparkLogAnalyzeTaskRequest
	GetTaskId() *int64
}

type GetSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark log analysis task. You can call the ListSparkLogAnalyzeTasks operation to query the IDs of all Spark log analysis tasks that are submitted in the current cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSparkLogAnalyzeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetSparkLogAnalyzeTaskRequest) SetTaskId(v int64) *GetSparkLogAnalyzeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetSparkLogAnalyzeTaskRequest) Validate() error {
	return dara.Validate(s)
}
