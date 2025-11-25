// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNameCreatedEndTime(v int64) *DescribePrivateDnsStatisticsRequest
	GetDomainNameCreatedEndTime() *int64
	SetDomainNameCreatedStartTime(v int64) *DescribePrivateDnsStatisticsRequest
	GetDomainNameCreatedStartTime() *int64
}

type DescribePrivateDnsStatisticsRequest struct {
	// example:
	//
	// 1726305596
	DomainNameCreatedEndTime *int64 `json:"DomainNameCreatedEndTime,omitempty" xml:"DomainNameCreatedEndTime,omitempty"`
	// example:
	//
	// 1725864531
	DomainNameCreatedStartTime *int64 `json:"DomainNameCreatedStartTime,omitempty" xml:"DomainNameCreatedStartTime,omitempty"`
}

func (s DescribePrivateDnsStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsStatisticsRequest) GetDomainNameCreatedEndTime() *int64 {
	return s.DomainNameCreatedEndTime
}

func (s *DescribePrivateDnsStatisticsRequest) GetDomainNameCreatedStartTime() *int64 {
	return s.DomainNameCreatedStartTime
}

func (s *DescribePrivateDnsStatisticsRequest) SetDomainNameCreatedEndTime(v int64) *DescribePrivateDnsStatisticsRequest {
	s.DomainNameCreatedEndTime = &v
	return s
}

func (s *DescribePrivateDnsStatisticsRequest) SetDomainNameCreatedStartTime(v int64) *DescribePrivateDnsStatisticsRequest {
	s.DomainNameCreatedStartTime = &v
	return s
}

func (s *DescribePrivateDnsStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
