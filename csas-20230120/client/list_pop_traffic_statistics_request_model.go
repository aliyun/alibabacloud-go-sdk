// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPopTrafficStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListPopTrafficStatisticsRequest
	GetEndTime() *string
	SetRegion(v string) *ListPopTrafficStatisticsRequest
	GetRegion() *string
	SetStartTime(v string) *ListPopTrafficStatisticsRequest
	GetStartTime() *string
}

type ListPopTrafficStatisticsRequest struct {
	// example:
	//
	// 1681293719
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1681035708
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListPopTrafficStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPopTrafficStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListPopTrafficStatisticsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListPopTrafficStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListPopTrafficStatisticsRequest) SetEndTime(v string) *ListPopTrafficStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPopTrafficStatisticsRequest) SetRegion(v string) *ListPopTrafficStatisticsRequest {
	s.Region = &v
	return s
}

func (s *ListPopTrafficStatisticsRequest) SetStartTime(v string) *ListPopTrafficStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPopTrafficStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
