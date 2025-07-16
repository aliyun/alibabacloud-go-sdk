// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeDomainUsageDataRequest
	GetArea() *string
	SetDataProtocol(v string) *DescribeDomainUsageDataRequest
	GetDataProtocol() *string
	SetDomainName(v string) *DescribeDomainUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainUsageDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeDomainUsageDataRequest
	GetField() *string
	SetInterval(v string) *DescribeDomainUsageDataRequest
	GetInterval() *string
	SetServiceType(v string) *DescribeDomainUsageDataRequest
	GetServiceType() *string
	SetStartTime(v string) *DescribeDomainUsageDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeDomainUsageDataRequest
	GetType() *string
}

type DescribeDomainUsageDataRequest struct {
	// The billable region. Valid values:
	//
	// 	- **CN*	- (default): inside the Chinese mainland
	//
	// 	- **OverSeas**: outside the Chinese mainland
	//
	// 	- **AP1**: Asia Pacific 1
	//
	// 	- **AP2**: Asia Pacific 2
	//
	// 	- **AP3**: Asia Pacific 3
	//
	// 	- **NA**: North America
	//
	// 	- **SA**: South America
	//
	// 	- **EU**: Europe
	//
	// 	- **MEAA**: Middle East and Africa
	//
	// 	- **all**: all the preceding billable regions
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The protocol of the data that you want to query. Valid values:
	//
	// 	- **http:*	- HTTP
	//
	// 	- **https:*	- HTTPS
	//
	// 	- **quic**: QUIC
	//
	// 	- **all*	- (default): HTTP, HTTPS, and QUIC
	//
	// example:
	//
	// all
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	// The accelerated domain name. You can specify up to 100 domain names in each request. Separate multiple domain names with commas (,).
	//
	// > If you leave this parameter empty, the usage data of all accelerated domain names in your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the data that you want to query. Valid values:
	//
	// 	- **bps**: bandwidth
	//
	// 	- **traf**: traffic
	//
	// 	- **acc**: requests
	//
	// > If you set this parameter to **acc**, the **Area*	- parameter is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The time granularity of the data entries. Unit: seconds. Valid values: **300*	- (5 minutes), **3600*	- (1 hour), and **86400*	- (1 day).
	//
	// 	- If **Interval*	- is set to **300**, you can query usage data in the last 6 months. The maximum time range per query that can be specified is 3 days.
	//
	// 	- If **Interval*	- is set to **3600*	- or **86400**, you can query usage data of the previous year.
	//
	// 	- If you leave the **Interval*	- parameter empty, the maximum time range that you can query is 1 month. If you specify a time range of 1 to 3 days, the time interval between the entries that are returned is 1 hour. If you specify a time range of at least 4 days, the time interval between the entries that are returned is 1 day.
	//
	// example:
	//
	// 300
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > Data is collected every 5 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of content that you want to query. Valid values:
	//
	// 	- **static**: static content
	//
	// 	- **dynamic**: dynamic content
	//
	// 	- **all*	- (default): both static and dynamic content
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeDomainUsageDataRequest) GetDataProtocol() *string {
	return s.DataProtocol
}

func (s *DescribeDomainUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainUsageDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeDomainUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainUsageDataRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeDomainUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainUsageDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDomainUsageDataRequest) SetArea(v string) *DescribeDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetDataProtocol(v string) *DescribeDomainUsageDataRequest {
	s.DataProtocol = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetDomainName(v string) *DescribeDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetEndTime(v string) *DescribeDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetField(v string) *DescribeDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetInterval(v string) *DescribeDomainUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetServiceType(v string) *DescribeDomainUsageDataRequest {
	s.ServiceType = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetStartTime(v string) *DescribeDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetType(v string) *DescribeDomainUsageDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
