// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerName(v string) *DescribeServiceLogRequest
	GetContainerName() *string
	SetEndTime(v string) *DescribeServiceLogRequest
	GetEndTime() *string
	SetInstanceName(v string) *DescribeServiceLogRequest
	GetInstanceName() *string
	SetIp(v string) *DescribeServiceLogRequest
	GetIp() *string
	SetKeyword(v string) *DescribeServiceLogRequest
	GetKeyword() *string
	SetPageNum(v int64) *DescribeServiceLogRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeServiceLogRequest
	GetPageSize() *int64
	SetPrevious(v bool) *DescribeServiceLogRequest
	GetPrevious() *bool
	SetStartTime(v string) *DescribeServiceLogRequest
	GetStartTime() *string
}

type DescribeServiceLogRequest struct {
	// The name of the container that runs the service.
	//
	// example:
	//
	// worker0
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The end of the time range to query. The time must be in UTC.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the instance that runs the service. For more information about how to query the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// echo-da290ac8-7fckm
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address of the instance whose logs you want to query. For more information about how to query the IP address of an instance, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// 10.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The keyword that you use to query the logs of the service.
	//
	// example:
	//
	// key
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 500.
	//
	// example:
	//
	// 500
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query the logs that are generated before the instance last restarts. This parameter is available only if the instance restarts.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Previous *bool `json:"Previous,omitempty" xml:"Previous,omitempty"`
	// The beginning of the time range to query. The time must be in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServiceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeServiceLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeServiceLogRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeServiceLogRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeServiceLogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeServiceLogRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeServiceLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeServiceLogRequest) GetPrevious() *bool {
	return s.Previous
}

func (s *DescribeServiceLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeServiceLogRequest) SetContainerName(v string) *DescribeServiceLogRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeServiceLogRequest) SetEndTime(v string) *DescribeServiceLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServiceLogRequest) SetInstanceName(v string) *DescribeServiceLogRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeServiceLogRequest) SetIp(v string) *DescribeServiceLogRequest {
	s.Ip = &v
	return s
}

func (s *DescribeServiceLogRequest) SetKeyword(v string) *DescribeServiceLogRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageNum(v int64) *DescribeServiceLogRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageSize(v int64) *DescribeServiceLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPrevious(v bool) *DescribeServiceLogRequest {
	s.Previous = &v
	return s
}

func (s *DescribeServiceLogRequest) SetStartTime(v string) *DescribeServiceLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeServiceLogRequest) Validate() error {
	return dara.Validate(s)
}
