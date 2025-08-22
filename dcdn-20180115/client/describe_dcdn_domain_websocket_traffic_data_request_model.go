// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainWebsocketTrafficDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you do not specify a value for this parameter, all accelerated domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for a query. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Operation Description**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the ISP.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
