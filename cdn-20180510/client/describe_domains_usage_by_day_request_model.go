// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsUsageByDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainsUsageByDayRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainsUsageByDayRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainsUsageByDayRequest
	GetStartTime() *string
}

type DescribeDomainsUsageByDayRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// If you do not specify an accelerated domain name, the data of all accelerated domain names that belong to your account is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-23T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-22T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainsUsageByDayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainsUsageByDayRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainsUsageByDayRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainsUsageByDayRequest) SetDomainName(v string) *DescribeDomainsUsageByDayRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetEndTime(v string) *DescribeDomainsUsageByDayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetStartTime(v string) *DescribeDomainsUsageByDayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) Validate() error {
	return dara.Validate(s)
}
