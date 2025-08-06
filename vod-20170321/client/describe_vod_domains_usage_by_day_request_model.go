// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainsUsageByDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainsUsageByDayRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainsUsageByDayRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainsUsageByDayRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainsUsageByDayRequest
	GetStartTime() *string
}

type DescribeVodDomainsUsageByDayRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainsUsageByDayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainsUsageByDayRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainsUsageByDayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainsUsageByDayRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainsUsageByDayRequest) SetDomainName(v string) *DescribeVodDomainsUsageByDayRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayRequest) SetEndTime(v string) *DescribeVodDomainsUsageByDayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayRequest) SetOwnerId(v int64) *DescribeVodDomainsUsageByDayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayRequest) SetStartTime(v string) *DescribeVodDomainsUsageByDayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayRequest) Validate() error {
	return dara.Validate(s)
}
