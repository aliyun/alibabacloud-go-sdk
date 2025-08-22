// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeDcdnDomainUsageDataRequest
	GetArea() *string
	SetDataProtocol(v string) *DescribeDcdnDomainUsageDataRequest
	GetDataProtocol() *string
	SetDomainName(v string) *DescribeDcdnDomainUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainUsageDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeDcdnDomainUsageDataRequest
	GetField() *string
	SetInterval(v string) *DescribeDcdnDomainUsageDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDcdnDomainUsageDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeDcdnDomainUsageDataRequest
	GetType() *string
}

type DescribeDcdnDomainUsageDataRequest struct {
	// The billable region. Valid values:
	//
	// 	- **CN**: Chinese mainland
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
	// Default value: **CN**
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The protocol of the data to query. Valid values:
	//
	// 	- **quic**: Quick UDP Internet Connections (QUIC)
	//
	// 	- **https**: HTTPS
	//
	// 	- **http**: HTTP
	//
	// 	- **all**: all the preceding protocols
	//
	// Default value: **all**
	//
	// example:
	//
	// all
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	// The accelerated domain name. You can specify up to 100 domain names in each request. Separate multiple domain names with commas (,).
	//
	// >  If you do not specify this parameter, the usage data of all accelerated domain names that belong to your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The maximum time range that can be queried is 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of data that you want to query. Valid values:
	//
	// 	- **bps**: bandwidth
	//
	// 	- **traf**: traffic
	//
	// 	- **acc**: requests
	//
	// >  **acc*	- does not support the **Area*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The time interval between the data entries to return. Unit: seconds.
	//
	// The time interval varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The minimum time granularity at which the data is queried is 5 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the requests. Valid values:
	//
	// 	- **static**: static requests
	//
	// 	- **dynamic**: dynamic requests
	//
	// 	- **all**: all requests
	//
	// Default value: **all**
	//
	// example:
	//
	// dynamic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnDomainUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeDcdnDomainUsageDataRequest) GetDataProtocol() *string {
	return s.DataProtocol
}

func (s *DescribeDcdnDomainUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainUsageDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeDcdnDomainUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainUsageDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnDomainUsageDataRequest) SetArea(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetDataProtocol(v string) *DescribeDcdnDomainUsageDataRequest {
	s.DataProtocol = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetDomainName(v string) *DescribeDcdnDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetEndTime(v string) *DescribeDcdnDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetField(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetInterval(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetStartTime(v string) *DescribeDcdnDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetType(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
