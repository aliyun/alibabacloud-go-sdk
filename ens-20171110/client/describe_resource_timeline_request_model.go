// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *DescribeResourceTimelineRequest
	GetBeginTime() *string
	SetEndTime(v string) *DescribeResourceTimelineRequest
	GetEndTime() *string
	SetUuid(v string) *DescribeResourceTimelineRequest
	GetUuid() *string
}

type DescribeResourceTimelineRequest struct {
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Uuid      *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeResourceTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeResourceTimelineRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeResourceTimelineRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeResourceTimelineRequest) SetBeginTime(v string) *DescribeResourceTimelineRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeResourceTimelineRequest) SetEndTime(v string) *DescribeResourceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceTimelineRequest) SetUuid(v string) *DescribeResourceTimelineRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeResourceTimelineRequest) Validate() error {
	return dara.Validate(s)
}
