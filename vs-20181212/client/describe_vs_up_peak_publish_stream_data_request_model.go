// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUpPeakPublishStreamDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsUpPeakPublishStreamDataRequest
	GetDomainName() *string
	SetDomainSwitch(v string) *DescribeVsUpPeakPublishStreamDataRequest
	GetDomainSwitch() *string
	SetEndTime(v string) *DescribeVsUpPeakPublishStreamDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsUpPeakPublishStreamDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsUpPeakPublishStreamDataRequest
	GetStartTime() *string
}

type DescribeVsUpPeakPublishStreamDataRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// on
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) GetDomainSwitch() *string {
	return s.DomainSwitch
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetDomainName(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetDomainSwitch(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetEndTime(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetOwnerId(v int64) *DescribeVsUpPeakPublishStreamDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) SetStartTime(v string) *DescribeVsUpPeakPublishStreamDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataRequest) Validate() error {
	return dara.Validate(s)
}
