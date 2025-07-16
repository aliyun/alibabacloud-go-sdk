// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeServiceEventRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeServiceEventRequest
	GetEventType() *string
	SetInstanceName(v string) *DescribeServiceEventRequest
	GetInstanceName() *string
	SetPageNum(v string) *DescribeServiceEventRequest
	GetPageNum() *string
	SetPageSize(v string) *DescribeServiceEventRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeServiceEventRequest
	GetStartTime() *string
}

type DescribeServiceEventRequest struct {
	// The end of the time range to query. By default, the current point in time is the end of the time range to query.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The instance name. For more information about how to obtain the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// echo-test-784xxxx85d-hhnd8
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. The time must be in UTC. The default value is seven days ago.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServiceEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeServiceEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeServiceEventRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeServiceEventRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *DescribeServiceEventRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeServiceEventRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeServiceEventRequest) SetEndTime(v string) *DescribeServiceEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServiceEventRequest) SetEventType(v string) *DescribeServiceEventRequest {
	s.EventType = &v
	return s
}

func (s *DescribeServiceEventRequest) SetInstanceName(v string) *DescribeServiceEventRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeServiceEventRequest) SetPageNum(v string) *DescribeServiceEventRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceEventRequest) SetPageSize(v string) *DescribeServiceEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeServiceEventRequest) SetStartTime(v string) *DescribeServiceEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeServiceEventRequest) Validate() error {
	return dara.Validate(s)
}
