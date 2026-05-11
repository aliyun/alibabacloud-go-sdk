// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticalDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTimeLeftRange(v int64) *DescribeStatisticalDataRequest
	GetBeginTimeLeftRange() *int64
	SetBeginTimeRightRange(v int64) *DescribeStatisticalDataRequest
	GetBeginTimeRightRange() *int64
	SetInstanceId(v string) *DescribeStatisticalDataRequest
	GetInstanceId() *string
	SetTimeUnit(v string) *DescribeStatisticalDataRequest
	GetTimeUnit() *string
}

type DescribeStatisticalDataRequest struct {
	// example:
	//
	// 1582283640000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1582298040000
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Day/Hour
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeStatisticalDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataRequest) GetBeginTimeLeftRange() *int64 {
	return s.BeginTimeLeftRange
}

func (s *DescribeStatisticalDataRequest) GetBeginTimeRightRange() *int64 {
	return s.BeginTimeRightRange
}

func (s *DescribeStatisticalDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeStatisticalDataRequest) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeRightRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetInstanceId(v string) *DescribeStatisticalDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetTimeUnit(v string) *DescribeStatisticalDataRequest {
	s.TimeUnit = &v
	return s
}

func (s *DescribeStatisticalDataRequest) Validate() error {
	return dara.Validate(s)
}
