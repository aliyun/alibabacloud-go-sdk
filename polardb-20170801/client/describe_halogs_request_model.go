// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHALogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeHALogsRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *DescribeHALogsRequest
	GetDBNodeId() *string
	SetEndTime(v string) *DescribeHALogsRequest
	GetEndTime() *string
	SetLogType(v string) *DescribeHALogsRequest
	GetLogType() *string
	SetPageNumber(v int32) *DescribeHALogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHALogsRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeHALogsRequest
	GetStartTime() *string
}

type DescribeHALogsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node ID.
	//
	// > 这是一个optional 字段，需要增加一个条件If specified,If specified, queries the high availability (HA) switchover records of `DBNodeId`. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view the detailed information about all clusters under your account, including node IDs.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time follows the `YYYY-MM-DDThh:mm:ssZ` format (UTC time).
	//
	// example:
	//
	// 2025-05-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The log type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HaSwitchLogList
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 5 to 50. Default value: 10.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. The time follows the `YYYY-MM-DDThh:mm:ssZ` format (UTC time).
	//
	// example:
	//
	// 2025-05-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeHALogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHALogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHALogsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeHALogsRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeHALogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHALogsRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeHALogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHALogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHALogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHALogsRequest) SetDBClusterId(v string) *DescribeHALogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeHALogsRequest) SetDBNodeId(v string) *DescribeHALogsRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeHALogsRequest) SetEndTime(v string) *DescribeHALogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHALogsRequest) SetLogType(v string) *DescribeHALogsRequest {
	s.LogType = &v
	return s
}

func (s *DescribeHALogsRequest) SetPageNumber(v int32) *DescribeHALogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHALogsRequest) SetPageSize(v int32) *DescribeHALogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHALogsRequest) SetStartTime(v string) *DescribeHALogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeHALogsRequest) Validate() error {
	return dara.Validate(s)
}
