// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribePreCheckResultRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribePreCheckResultRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribePreCheckResultRequest
	GetTaskId() *string
}

type DescribePreCheckResultRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga76p6****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the precheck task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4561
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribePreCheckResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePreCheckResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreCheckResultRequest) SetDrdsInstanceId(v string) *DescribePreCheckResultRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetRegionId(v string) *DescribePreCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckResultRequest) SetTaskId(v string) *DescribePreCheckResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribePreCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
