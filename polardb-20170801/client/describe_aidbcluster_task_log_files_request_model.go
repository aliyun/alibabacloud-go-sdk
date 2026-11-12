// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskLogFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterTaskLogFilesRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAIDBClusterTaskLogFilesRequest
	GetEndTime() *string
	SetLogType(v string) *DescribeAIDBClusterTaskLogFilesRequest
	GetLogType() *string
	SetPageNumber(v int64) *DescribeAIDBClusterTaskLogFilesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAIDBClusterTaskLogFilesRequest
	GetPageSize() *int64
	SetRelativeDBClusterId(v string) *DescribeAIDBClusterTaskLogFilesRequest
	GetRelativeDBClusterId() *string
	SetReverse(v bool) *DescribeAIDBClusterTaskLogFilesRequest
	GetReverse() *bool
	SetStartTime(v string) *DescribeAIDBClusterTaskLogFilesRequest
	GetStartTime() *string
}

type DescribeAIDBClusterTaskLogFilesRequest struct {
	// The instance ID of the model operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-bp10gr***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-15T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The log type. Valid values:
	//
	// - **sls**
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**. Default value: **100**.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance ID of the associated PolarDB instance.
	//
	// example:
	//
	// pc-bp10ze***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// Specifies whether to sort the results in descending order. Default value: *false*.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The start time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-15T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterTaskLogFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetDBClusterId(v string) *DescribeAIDBClusterTaskLogFilesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetEndTime(v string) *DescribeAIDBClusterTaskLogFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetLogType(v string) *DescribeAIDBClusterTaskLogFilesRequest {
	s.LogType = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetPageNumber(v int64) *DescribeAIDBClusterTaskLogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetPageSize(v int64) *DescribeAIDBClusterTaskLogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetRelativeDBClusterId(v string) *DescribeAIDBClusterTaskLogFilesRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetReverse(v bool) *DescribeAIDBClusterTaskLogFilesRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) SetStartTime(v string) *DescribeAIDBClusterTaskLogFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesRequest) Validate() error {
	return dara.Validate(s)
}
