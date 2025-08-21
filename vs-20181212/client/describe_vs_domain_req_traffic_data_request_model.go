// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainReqTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainReqTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVsDomainReqTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVsDomainReqTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVsDomainReqTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVsDomainReqTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainReqTrafficDataRequest
	GetStartTime() *string
}

type DescribeVsDomainReqTrafficDataRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-10-13T10:00:41Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainReqTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetDomainName(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetEndTime(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetInterval(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetIspNameEn(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetLocationNameEn(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetOwnerId(v int64) *DescribeVsDomainReqTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) SetStartTime(v string) *DescribeVsDomainReqTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
