// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainWebsocketHttpCodeDataRequest struct {
	// The accelerated domain name. You can specify multiple accelerated domain names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2018-03-01T06:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for a query. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP).
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
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-01T05:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
