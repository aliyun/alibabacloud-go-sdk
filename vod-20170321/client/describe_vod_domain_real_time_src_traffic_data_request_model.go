// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeSrcTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeSrcTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeSrcTrafficDataRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeSrcTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeSrcTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeSrcTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
