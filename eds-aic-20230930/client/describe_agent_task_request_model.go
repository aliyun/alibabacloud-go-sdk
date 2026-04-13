// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *DescribeAgentTaskRequest
	GetTaskIds() []*string
}

type DescribeAgentTaskRequest struct {
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgentTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeAgentTaskRequest) SetTaskIds(v []*string) *DescribeAgentTaskRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
