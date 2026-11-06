// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeInspectionTaskAttributeRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeInspectionTaskAttributeRequest
	GetTaskId() *string
}

type DescribeInspectionTaskAttributeRequest struct {
	// The ID of the region where the task is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the inspection task to query. This parameter is required.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInspectionTaskAttributeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTaskAttributeRequest) SetRegionId(v string) *DescribeInspectionTaskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInspectionTaskAttributeRequest) SetTaskId(v string) *DescribeInspectionTaskAttributeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTaskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
