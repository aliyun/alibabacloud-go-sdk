// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeUpBpsPeakDataRequest
	GetDomainName() *string
	SetDomainSwitch(v string) *DescribeUpBpsPeakDataRequest
	GetDomainSwitch() *string
	SetEndTime(v string) *DescribeUpBpsPeakDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeUpBpsPeakDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUpBpsPeakDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeUpBpsPeakDataRequest
	GetStartTime() *string
}

type DescribeUpBpsPeakDataRequest struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to collect statistics at the domain name level. Valid values:
	//
	// 	- on: collects statistics at the domain name level.
	//
	// 	- off: collects statistics at the user level. This is the default value.
	//
	// example:
	//
	// on
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	// The end time of stream pulling. The time must be in UTC. The end time must be later than the start time. The maximum time range that can be specified by StartTime and EndTime is 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T21:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of stream pulling. The time must be in UTC. The maximum time range that can be specified by StartTime and EndTime is 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUpBpsPeakDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUpBpsPeakDataRequest) GetDomainSwitch() *string {
	return s.DomainSwitch
}

func (s *DescribeUpBpsPeakDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUpBpsPeakDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpBpsPeakDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUpBpsPeakDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUpBpsPeakDataRequest) SetDomainName(v string) *DescribeUpBpsPeakDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetDomainSwitch(v string) *DescribeUpBpsPeakDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetEndTime(v string) *DescribeUpBpsPeakDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetOwnerId(v int64) *DescribeUpBpsPeakDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetRegionId(v string) *DescribeUpBpsPeakDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetStartTime(v string) *DescribeUpBpsPeakDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) Validate() error {
	return dara.Validate(s)
}
