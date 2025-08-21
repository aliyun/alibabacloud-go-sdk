// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRegionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainRegionDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainRegionDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainRegionDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainRegionDataRequest
	GetStartTime() *string
}

type DescribeVsDomainRegionDataRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-10-31T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-10-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainRegionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainRegionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainRegionDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainRegionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainRegionDataRequest) SetDomainName(v string) *DescribeVsDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetEndTime(v string) *DescribeVsDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetOwnerId(v int64) *DescribeVsDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) SetStartTime(v string) *DescribeVsDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataRequest) Validate() error {
	return dara.Validate(s)
}
