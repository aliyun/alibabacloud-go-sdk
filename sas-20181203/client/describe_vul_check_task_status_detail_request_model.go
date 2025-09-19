// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulCheckTaskStatusDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *DescribeVulCheckTaskStatusDetailRequest
	GetTaskIds() []*string
	SetTypes(v []*string) *DescribeVulCheckTaskStatusDetailRequest
	GetTypes() []*string
	SetUuid(v string) *DescribeVulCheckTaskStatusDetailRequest
	GetUuid() *string
}

type DescribeVulCheckTaskStatusDetailRequest struct {
	// The task IDs.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The types of the vulnerabilities that are detected by the tasks.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// The UUID of the server.
	//
	// example:
	//
	// 5d55af3c-35f3-4d4d-8ccc-8c5443b0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeVulCheckTaskStatusDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulCheckTaskStatusDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulCheckTaskStatusDetailRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeVulCheckTaskStatusDetailRequest) GetTypes() []*string {
	return s.Types
}

func (s *DescribeVulCheckTaskStatusDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeVulCheckTaskStatusDetailRequest) SetTaskIds(v []*string) *DescribeVulCheckTaskStatusDetailRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailRequest) SetTypes(v []*string) *DescribeVulCheckTaskStatusDetailRequest {
	s.Types = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailRequest) SetUuid(v string) *DescribeVulCheckTaskStatusDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailRequest) Validate() error {
	return dara.Validate(s)
}
