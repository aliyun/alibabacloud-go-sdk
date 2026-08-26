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
	// The region code. Valid values:
	//
	// - **CN**: the Chinese mainland.
	//
	// - **OverSeas**: outside the Chinese mainland.
	//
	// - **AP1**: Asia-Pacific 1.
	//
	// - **AP2**: Asia-Pacific 2.
	//
	// - **AP3**: Asia-Pacific 3.
	//
	// - **NA**: North America.
	//
	// - **SA**: South America.
	//
	// - **EU**: Europe.
	//
	// - **MEAA**: Middle East and Africa.
	//
	// - **all**: all regions.
	//
	// > If this parameter is not specified, the default value is the Chinese mainland. Regions outside the Chinese mainland: - Asia-Pacific 1: Hong Kong (China), Macao (China), Taiwan (China), Japan, and Southeast Asian countries except Vietnam and Indonesia. - Asia-Pacific 2: Indonesia, South Korea, and Vietnam. - Asia-Pacific 3: Australia and New Zealand. North America: the United States and Canada. - South America: Brazil. - Europe: Ukraine, the United Kingdom, France, the Netherlands, Spain, Italy, Sweden, and Germany. - Middle East and Africa: South Africa, Oman, the United Arab Emirates, and Kuwait.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The protocol of the data to retrieve. Valid values:
	//
	// - **http**: HTTP.
	//
	// - **https**: HTTPS.
	//
	// - **quic**: QUIC.
	//
	// - **all*	- (default): all of the preceding protocols.
	//
	// example:
	//
	// all
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	// The streaming domain.
	//
	// - You can specify a single domain name or multiple domain names. Separate multiple domain names with commas (,).
	//
	// - If this parameter is empty, the merged data of all streaming domains is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// The end time must be later than the start time, and the difference between the end time and the start time cannot exceed **31*	- days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data type of the usage data to query. Valid values:
	//
	// - **bps**: playback bandwidth.
	//
	// - **traf**: traffic.
	//
	// - **req_traf**: when Type is set to push, this indicates stream ingest traffic. When Type is set to push_proxy, this indicates relay traffic.
	//
	// - **req_bps**: when Type is set to push, this indicates stream ingest bandwidth. When Type is set to push_proxy, this indicates relay bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// traf
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Forces retrieval of data at the specified time granularity, in seconds. Valid values: **300*	- (5 minutes), **3600*	- (1 hour), and **86400*	- (1 day).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of usage data to retrieve.
	//
	// When **Field*	- is set to **bps*	- or **traf**, valid values:
	//
	// - **rts**: RTS bandwidth or traffic.
	//
	// - **quic**: QUIC bandwidth or traffic.
	//
	// When **Field*	- is set to **req_traf*	- or **req_bps**, valid values:
	//
	// - **push**: stream ingest bandwidth or traffic.
	//
	// - **push_proxy**: relay bandwidth or traffic.
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
