// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceDbMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsInstanceDbMonitorRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsInstanceDbMonitorRequest
	GetDrdsInstanceId() *string
	SetEndTime(v int64) *DescribeDrdsInstanceDbMonitorRequest
	GetEndTime() *int64
	SetKey(v string) *DescribeDrdsInstanceDbMonitorRequest
	GetKey() *string
	SetRegionId(v string) *DescribeDrdsInstanceDbMonitorRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest
	GetStartTime() *int64
}

type DescribeDrdsInstanceDbMonitorRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the Distributed Relational Database Service (DRDS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1603166400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics for a query at a time. Separate multiple metric parameters with commas (,).
	//
	// >  For more information about the details of performance monitoring metrics, see [Database monitoring](https://help.aliyun.com/document_detail/186704.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1603162800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstanceDbMonitorRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDbName(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetKey(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceDbMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceDbMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorRequest) Validate() error {
	return dara.Validate(s)
}
