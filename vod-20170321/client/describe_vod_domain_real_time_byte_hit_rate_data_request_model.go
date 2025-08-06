// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeByteHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeByteHitRateDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeByteHitRateDataRequest struct {
	// You can specify multiple domain names and separate them with commas (,). You can specify at most 100 domain names in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2020-05-15T09:15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-05-15T09:13:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeByteHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeByteHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
