// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDashboardMemberDeviceInfoRequest
	GetEndTime() *int64
	SetRegionId(v string) *DescribeDashboardMemberDeviceInfoRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeDashboardMemberDeviceInfoRequest
	GetStartTime() *int64
}

type DescribeDashboardMemberDeviceInfoRequest struct {
	// This parameter is required.
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDashboardMemberDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberDeviceInfoRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDashboardMemberDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDashboardMemberDeviceInfoRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDashboardMemberDeviceInfoRequest) SetEndTime(v int64) *DescribeDashboardMemberDeviceInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoRequest) SetRegionId(v string) *DescribeDashboardMemberDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoRequest) SetStartTime(v int64) *DescribeDashboardMemberDeviceInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
