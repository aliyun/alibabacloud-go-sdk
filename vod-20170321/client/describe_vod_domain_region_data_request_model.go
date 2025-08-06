// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRegionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRegionDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRegionDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainRegionDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRegionDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRegionDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRegionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRegionDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRegionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRegionDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRegionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRegionDataRequest) SetDomainName(v string) *DescribeVodDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRegionDataRequest) SetEndTime(v string) *DescribeVodDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRegionDataRequest) SetOwnerId(v int64) *DescribeVodDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRegionDataRequest) SetStartTime(v string) *DescribeVodDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRegionDataRequest) Validate() error {
	return dara.Validate(s)
}
