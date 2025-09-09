// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsPerformanceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeRdsPerformanceSummaryRequest
	GetDrdsInstanceId() *string
	SetRdsInstanceId(v []*string) *DescribeRdsPerformanceSummaryRequest
	GetRdsInstanceId() []*string
	SetRegionId(v string) *DescribeRdsPerformanceSummaryRequest
	GetRegionId() *string
}

type DescribeRdsPerformanceSummaryRequest struct {
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-*****************
	RdsInstanceId []*string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty" type:"Repeated"`
	// The ID of the region where the streaming domain resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRdsPerformanceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRdsPerformanceSummaryRequest) GetRdsInstanceId() []*string {
	return s.RdsInstanceId
}

func (s *DescribeRdsPerformanceSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsPerformanceSummaryRequest) SetDrdsInstanceId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRdsInstanceId(v []*string) *DescribeRdsPerformanceSummaryRequest {
	s.RdsInstanceId = v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) SetRegionId(v string) *DescribeRdsPerformanceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
