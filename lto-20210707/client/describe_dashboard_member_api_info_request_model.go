// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberApiInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *DescribeDashboardMemberApiInfoRequest
	GetBizChainId() *string
	SetEndTime(v int64) *DescribeDashboardMemberApiInfoRequest
	GetEndTime() *int64
	SetRegionId(v string) *DescribeDashboardMemberApiInfoRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeDashboardMemberApiInfoRequest
	GetStartTime() *int64
}

type DescribeDashboardMemberApiInfoRequest struct {
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	// This parameter is required.
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDashboardMemberApiInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberApiInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberApiInfoRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *DescribeDashboardMemberApiInfoRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDashboardMemberApiInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDashboardMemberApiInfoRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDashboardMemberApiInfoRequest) SetBizChainId(v string) *DescribeDashboardMemberApiInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoRequest) SetEndTime(v int64) *DescribeDashboardMemberApiInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoRequest) SetRegionId(v string) *DescribeDashboardMemberApiInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoRequest) SetStartTime(v int64) *DescribeDashboardMemberApiInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoRequest) Validate() error {
	return dara.Validate(s)
}
