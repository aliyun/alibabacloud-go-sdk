// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakOfLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeUpBpsPeakOfLineRequest
	GetDomainName() *string
	SetDomainSwitch(v string) *DescribeUpBpsPeakOfLineRequest
	GetDomainSwitch() *string
	SetEndTime(v string) *DescribeUpBpsPeakOfLineRequest
	GetEndTime() *string
	SetLine(v string) *DescribeUpBpsPeakOfLineRequest
	GetLine() *string
	SetOwnerId(v int64) *DescribeUpBpsPeakOfLineRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUpBpsPeakOfLineRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeUpBpsPeakOfLineRequest
	GetStartTime() *string
}

type DescribeUpBpsPeakOfLineRequest struct {
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
	// 	- off (default): collects statistics at the user level.
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
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The leased line. Valid values: integers from 0 to 9.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Line     *string `json:"Line,omitempty" xml:"Line,omitempty"`
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

func (s DescribeUpBpsPeakOfLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUpBpsPeakOfLineRequest) GetDomainSwitch() *string {
	return s.DomainSwitch
}

func (s *DescribeUpBpsPeakOfLineRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUpBpsPeakOfLineRequest) GetLine() *string {
	return s.Line
}

func (s *DescribeUpBpsPeakOfLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpBpsPeakOfLineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUpBpsPeakOfLineRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUpBpsPeakOfLineRequest) SetDomainName(v string) *DescribeUpBpsPeakOfLineRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetDomainSwitch(v string) *DescribeUpBpsPeakOfLineRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetEndTime(v string) *DescribeUpBpsPeakOfLineRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetLine(v string) *DescribeUpBpsPeakOfLineRequest {
	s.Line = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetOwnerId(v int64) *DescribeUpBpsPeakOfLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetRegionId(v string) *DescribeUpBpsPeakOfLineRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetStartTime(v string) *DescribeUpBpsPeakOfLineRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) Validate() error {
	return dara.Validate(s)
}
