// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDSPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstType(v string) *DescribeRDSPerformanceRequest
	GetDbInstType() *string
	SetDrdsInstanceId(v string) *DescribeRDSPerformanceRequest
	GetDrdsInstanceId() *string
	SetEndTime(v int64) *DescribeRDSPerformanceRequest
	GetEndTime() *int64
	SetKeys(v string) *DescribeRDSPerformanceRequest
	GetKeys() *string
	SetRdsInstanceId(v string) *DescribeRDSPerformanceRequest
	GetRdsInstanceId() *string
	SetStartTime(v int64) *DescribeRDSPerformanceRequest
	GetStartTime() *int64
}

type DescribeRDSPerformanceRequest struct {
	// The type of the database engine.
	//
	// example:
	//
	// mysql
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the Distributed Relational Database Service (DRDS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time of the query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// example:
	//
	// 1603209690000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics for a query at a time. Separate multiple metric parameters with commas (,).
	//
	// >  For more information about the details of performance monitoring metrics, see [Storage monitoring](https://help.aliyun.com/document_detail/186705.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_MemCpuUsage
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The ID of the storage-layer ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-************
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The start time of the query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// example:
	//
	// 1603123290000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRDSPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDSPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeRDSPerformanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRDSPerformanceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRDSPerformanceRequest) GetKeys() *string {
	return s.Keys
}

func (s *DescribeRDSPerformanceRequest) GetRdsInstanceId() *string {
	return s.RdsInstanceId
}

func (s *DescribeRDSPerformanceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeRDSPerformanceRequest) SetDbInstType(v string) *DescribeRDSPerformanceRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetDrdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetEndTime(v int64) *DescribeRDSPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetKeys(v string) *DescribeRDSPerformanceRequest {
	s.Keys = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetRdsInstanceId(v string) *DescribeRDSPerformanceRequest {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) SetStartTime(v int64) *DescribeRDSPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRDSPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
