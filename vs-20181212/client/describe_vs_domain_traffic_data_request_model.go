// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVsDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVsDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVsDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVsDomainTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeVsDomainTrafficDataRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-09-21T02:50:42Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-08-18T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVsDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVsDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVsDomainTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainTrafficDataRequest) SetDomainName(v string) *DescribeVsDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetEndTime(v string) *DescribeVsDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetInterval(v string) *DescribeVsDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeVsDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeVsDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeVsDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) SetStartTime(v string) *DescribeVsDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
