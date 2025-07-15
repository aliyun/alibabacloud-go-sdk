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
	SetOwnerId(v int64) *DescribeDomainUsageDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDomainUsageDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDomainUsageDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeDomainUsageDataRequest
	GetType() *string
}

type DescribeDomainUsageDataRequest struct {
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
	// 	- **all**: all regions
	//
	// >  If you do not specify this parameter, the default value CN is used. Alibaba Cloud supports the following countries and regions outside the Chinese mainland: - Asia Pacific 1: Hong Kong (China), Macao (China), Taiwan (China), Japan, and Southeast Asia excluding Vietnam and Indonesia. - Asia Pacific 2: Indonesia, South Korea, and Vietnam. - Asia Pacific 3: Australia and New Zealand. - North America: US and Canada. - South America: Brazil. Europe: Ukraine, UK, France, Netherlands, Spain, Italy, Sweden, and Germany. - Middle East and Africa: South Africa, Oman, UAE, and Kuwait.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The protocol of the data to query. Valid values:
	//
	// 	- **http**: HTTP
	//
	// 	- **https**: HTTPS
	//
	// 	- **quic**: QUIC
	//
	// 	- **all*	- (default): HTTP, HTTPS, and QUIC
	//
	// example:
	//
	// all
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	// The domain name.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// The end time must be later than the start time. The maximum time range that you can specify is **31*	- days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The category of the resource usage data to query. Valid values:
	//
	// 	- **bps**: streaming bandwidth
	//
	// 	- **traf**: streaming traffic
	//
	// 	- **req_traf**: stream ingest traffic if you set Type to push, or stream relay traffic if you set Type to push_proxy
	//
	// 	- **req_bps**: stream ingest bandwidth if you set Type to push, or stream relay bandwidth if you set Type to push_proxy
	//
	// This parameter is required.
	//
	// example:
	//
	// traf
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The time interval between the data entries to return. Unit: seconds. Valid values: **300*	- (5 minutes), **3600*	- (1 hour), and **86400*	- (1 day).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the resource usage data to query.
	//
	// Valid values if you set **Field*	- to **bps*	- or **traf**:
	//
	// 	- **rts**: bandwidth or traffic for Real-Time Streaming (RTS)
	//
	// 	- **quic**: bandwidth or traffic for QUIC
	//
	// 	- **all**: all bandwidth or traffic
	//
	// Valid values if you set **Field*	- to **req_traf*	- or **req_bps**:
	//
	// 	- **push**: stream ingest bandwidth or traffic
	//
	// 	- **push_proxy**: stream relay bandwidth or traffic
	//
	// example:
	//
	// all
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

func (s *DescribeDomainUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDomainUsageDataRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *DescribeDomainUsageDataRequest) SetOwnerId(v int64) *DescribeDomainUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetRegionId(v string) *DescribeDomainUsageDataRequest {
	s.RegionId = &v
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
