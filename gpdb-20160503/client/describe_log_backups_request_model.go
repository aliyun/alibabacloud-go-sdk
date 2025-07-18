// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLogBackupsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeLogBackupsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeLogBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogBackupsRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeLogBackupsRequest
	GetStartTime() *string
}

type DescribeLogBackupsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-12T03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-12T02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLogBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLogBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLogBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLogBackupsRequest) SetDBInstanceId(v string) *DescribeLogBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetEndTime(v string) *DescribeLogBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageNumber(v int32) *DescribeLogBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageSize(v int32) *DescribeLogBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetStartTime(v string) *DescribeLogBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) Validate() error {
	return dara.Validate(s)
}
