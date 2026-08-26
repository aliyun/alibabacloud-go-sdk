// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpPeakPublishStreamDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeUpPeakPublishStreamDataRequest
	GetDomainName() *string
	SetDomainSwitch(v string) *DescribeUpPeakPublishStreamDataRequest
	GetDomainSwitch() *string
	SetEndTime(v string) *DescribeUpPeakPublishStreamDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeUpPeakPublishStreamDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUpPeakPublishStreamDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeUpPeakPublishStreamDataRequest
	GetStartTime() *string
}

type DescribeUpPeakPublishStreamDataRequest struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain name statistics switch. Valid values:
	//
	// - on: enables domain name-level statistics.
	//
	// - off: user-level data. User-level data is collected by default.
	//
	// example:
	//
	// on
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	// The end time of stream pulling in UTC. The interval between StartTime and EndTime must be within 30 days, and EndTime must be later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of stream pulling in UTC. The interval between StartTime and EndTime must be within 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUpPeakPublishStreamDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetDomainSwitch() *string {
	return s.DomainSwitch
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUpPeakPublishStreamDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetDomainName(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetDomainSwitch(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetEndTime(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetOwnerId(v int64) *DescribeUpPeakPublishStreamDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetRegionId(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetStartTime(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) Validate() error {
	return dara.Validate(s)
}
