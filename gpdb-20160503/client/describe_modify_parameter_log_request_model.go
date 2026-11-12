// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeModifyParameterLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeModifyParameterLogRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeModifyParameterLogRequest
	GetStartTime() *string
}

type DescribeModifyParameterLogRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2020-05-05T11:22:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-02T11:22:22Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModifyParameterLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModifyParameterLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModifyParameterLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModifyParameterLogRequest) SetDBInstanceId(v string) *DescribeModifyParameterLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) Validate() error {
	return dara.Validate(s)
}
