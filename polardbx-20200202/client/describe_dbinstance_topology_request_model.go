// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceTopologyRequest
	GetDBInstanceName() *string
	SetEndTime(v string) *DescribeDBInstanceTopologyRequest
	GetEndTime() *string
	SetMinuteSimple(v bool) *DescribeDBInstanceTopologyRequest
	GetMinuteSimple() *bool
	SetRegionId(v string) *DescribeDBInstanceTopologyRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDBInstanceTopologyRequest
	GetStartTime() *string
}

type DescribeDBInstanceTopologyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The end time of the time range for querying historical instances in the topology. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-10-22T10:30:45Z 04:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query the historical topology.
	//
	// example:
	//
	// true
	MinuteSimple *bool `json:"MinuteSimple,omitempty" xml:"MinuteSimple,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the time range for querying historical instances in the topology. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-10-21T10:30:45Z 04:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstanceTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceTopologyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstanceTopologyRequest) GetMinuteSimple() *bool {
	return s.MinuteSimple
}

func (s *DescribeDBInstanceTopologyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceTopologyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceTopologyRequest) SetDBInstanceName(v string) *DescribeDBInstanceTopologyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetEndTime(v string) *DescribeDBInstanceTopologyRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetMinuteSimple(v bool) *DescribeDBInstanceTopologyRequest {
	s.MinuteSimple = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetRegionId(v string) *DescribeDBInstanceTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) SetStartTime(v string) *DescribeDBInstanceTopologyRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyRequest) Validate() error {
	return dara.Validate(s)
}
