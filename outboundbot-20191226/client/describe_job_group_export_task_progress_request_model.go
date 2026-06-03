// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupExportTaskProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeJobGroupExportTaskProgressRequest
	GetInstanceId() *string
	SetTaskId(v string) *DescribeJobGroupExportTaskProgressRequest
	GetTaskId() *string
}

type DescribeJobGroupExportTaskProgressRequest struct {
	// example:
	//
	// b3dbfb82-1ae6-4e73-b717-f494727d2af3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// e4e2a770-b97b-465a-80d8-06dca008c503
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeJobGroupExportTaskProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupExportTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupExportTaskProgressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJobGroupExportTaskProgressRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeJobGroupExportTaskProgressRequest) SetInstanceId(v string) *DescribeJobGroupExportTaskProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressRequest) SetTaskId(v string) *DescribeJobGroupExportTaskProgressRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressRequest) Validate() error {
	return dara.Validate(s)
}
