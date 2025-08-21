// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVsDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVsDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVsDomainBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVsDomainBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeVsDomainBpsDataRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-10-02T02:30:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// example:
	//
	// guangdong
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-12-26T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVsDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVsDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVsDomainBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainBpsDataRequest) SetDomainName(v string) *DescribeVsDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetEndTime(v string) *DescribeVsDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetInterval(v string) *DescribeVsDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetIspNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeVsDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetOwnerId(v int64) *DescribeVsDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) SetStartTime(v string) *DescribeVsDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
