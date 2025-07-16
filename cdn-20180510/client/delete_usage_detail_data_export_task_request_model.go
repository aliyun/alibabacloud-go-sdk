// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsageDetailDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DeleteUsageDetailDataExportTaskRequest
	GetTaskId() *string
}

type DeleteUsageDetailDataExportTaskRequest struct {
	// The ID of the task. You can call the [DescribeUserUsageDataExportTask](https://help.aliyun.com/document_detail/91062.html) operation to query tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteUsageDetailDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteUsageDetailDataExportTaskRequest) SetTaskId(v string) *DeleteUsageDetailDataExportTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteUsageDetailDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
