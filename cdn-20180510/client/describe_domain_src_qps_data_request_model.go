// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainSrcQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainSrcQpsDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainSrcQpsDataRequest
	GetStartTime() *string
}

type DescribeDomainSrcQpsDataRequest struct {
	// The accelerated domain name. You can specify multiple domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries QPS data for all accelerated domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// If you leave this parameter empty, data collected in the last 24 hours is queried.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainSrcQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcQpsDataRequest) SetDomainName(v string) *DescribeDomainSrcQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetEndTime(v string) *DescribeDomainSrcQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetInterval(v string) *DescribeDomainSrcQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetStartTime(v string) *DescribeDomainSrcQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
