// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*DescribePlayEventListResponseBodyEventList) *DescribePlayEventListResponseBody
	GetEventList() []*DescribePlayEventListResponseBodyEventList
	SetPageNo(v int64) *DescribePlayEventListResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayEventListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePlayEventListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePlayEventListResponseBody
	GetTotalCount() *int64
}

type DescribePlayEventListResponseBody struct {
	EventList  []*DescribePlayEventListResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	PageNo     *int64                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlayEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayEventListResponseBody) GetEventList() []*DescribePlayEventListResponseBodyEventList {
	return s.EventList
}

func (s *DescribePlayEventListResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayEventListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayEventListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePlayEventListResponseBody) SetEventList(v []*DescribePlayEventListResponseBodyEventList) *DescribePlayEventListResponseBody {
	s.EventList = v
	return s
}

func (s *DescribePlayEventListResponseBody) SetPageNo(v int64) *DescribePlayEventListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePlayEventListResponseBody) SetPageSize(v int64) *DescribePlayEventListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayEventListResponseBody) SetRequestId(v string) *DescribePlayEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayEventListResponseBody) SetTotalCount(v int64) *DescribePlayEventListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePlayEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayEventListResponseBodyEventList struct {
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EventName   *string  `json:"EventName,omitempty" xml:"EventName,omitempty"`
	Time        *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	Topic       *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DescribePlayEventListResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayEventListResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribePlayEventListResponseBodyEventList) GetDescription() *string {
	return s.Description
}

func (s *DescribePlayEventListResponseBodyEventList) GetDuration() *float32 {
	return s.Duration
}

func (s *DescribePlayEventListResponseBodyEventList) GetEventName() *string {
	return s.EventName
}

func (s *DescribePlayEventListResponseBodyEventList) GetTime() *int64 {
	return s.Time
}

func (s *DescribePlayEventListResponseBodyEventList) GetTopic() *string {
	return s.Topic
}

func (s *DescribePlayEventListResponseBodyEventList) SetDescription(v string) *DescribePlayEventListResponseBodyEventList {
	s.Description = &v
	return s
}

func (s *DescribePlayEventListResponseBodyEventList) SetDuration(v float32) *DescribePlayEventListResponseBodyEventList {
	s.Duration = &v
	return s
}

func (s *DescribePlayEventListResponseBodyEventList) SetEventName(v string) *DescribePlayEventListResponseBodyEventList {
	s.EventName = &v
	return s
}

func (s *DescribePlayEventListResponseBodyEventList) SetTime(v int64) *DescribePlayEventListResponseBodyEventList {
	s.Time = &v
	return s
}

func (s *DescribePlayEventListResponseBodyEventList) SetTopic(v string) *DescribePlayEventListResponseBodyEventList {
	s.Topic = &v
	return s
}

func (s *DescribePlayEventListResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}
