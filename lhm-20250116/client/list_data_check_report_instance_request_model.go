// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *ListDataCheckReportInstanceRequest
	GetTaskId() *int64
}

type ListDataCheckReportInstanceRequest struct {
	// The check task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListDataCheckReportInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportInstanceRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDataCheckReportInstanceRequest) SetTaskId(v int64) *ListDataCheckReportInstanceRequest {
	s.TaskId = &v
	return s
}

func (s *ListDataCheckReportInstanceRequest) Validate() error {
	return dara.Validate(s)
}
