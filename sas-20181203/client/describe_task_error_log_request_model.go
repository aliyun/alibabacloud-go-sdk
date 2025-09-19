// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskErrorLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTaskId(v string) *DescribeTaskErrorLogRequest
	GetBuildTaskId() *string
}

type DescribeTaskErrorLogRequest struct {
	// The ID of the task.
	//
	// >  You can call the DescribeImageFixTask operation to query the IDs of tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ivf-6e520160-205d-4801-b8e9-9e7e****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
}

func (s DescribeTaskErrorLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskErrorLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskErrorLogRequest) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *DescribeTaskErrorLogRequest) SetBuildTaskId(v string) *DescribeTaskErrorLogRequest {
	s.BuildTaskId = &v
	return s
}

func (s *DescribeTaskErrorLogRequest) Validate() error {
	return dara.Validate(s)
}
