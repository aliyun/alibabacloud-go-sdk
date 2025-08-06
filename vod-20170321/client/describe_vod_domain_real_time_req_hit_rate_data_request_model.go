// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeReqHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeReqHitRateDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeReqHitRateDataRequest struct {
	// You can specify multiple domain names and separate them with commas (,). You can specify a maximum of 100 domain names in each call. If you specify multiple domain names, merged data is returned.
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
	// 2024-01-02T11:26:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-01-02T11:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeReqHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeReqHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
