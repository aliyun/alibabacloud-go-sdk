// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainHttpCodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainHttpCodeDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainHttpCodeDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainHttpCodeDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeVodDomainHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetInterval(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetOwnerId(v int64) *DescribeVodDomainHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeVodDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
