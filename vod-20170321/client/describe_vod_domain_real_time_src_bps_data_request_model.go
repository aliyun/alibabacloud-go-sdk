// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeSrcBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeSrcBpsDataRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
