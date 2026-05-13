// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTaskDetailRequest
	GetInstanceId() *string
	SetTaskId(v string) *DescribeTaskDetailRequest
	GetTaskId() *string
}

type DescribeTaskDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 674546459
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskDetailRequest) SetInstanceId(v string) *DescribeTaskDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTaskDetailRequest) SetTaskId(v string) *DescribeTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
