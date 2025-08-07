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
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node ID.
	//
	// >  Queries the HA failover records of the Node `DBNodeId` . You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as node IDs.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// example:
	//
	// 2020-09-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The log type.
	//
	// This parameter is required.
	//
	// example:
	//
	// HaSwitchLogList
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2020-05-01T00:00Z
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
