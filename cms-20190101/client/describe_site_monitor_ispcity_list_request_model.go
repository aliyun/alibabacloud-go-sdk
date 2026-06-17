// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorISPCityListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *DescribeSiteMonitorISPCityListRequest
	GetCity() *string
	SetIPV4(v bool) *DescribeSiteMonitorISPCityListRequest
	GetIPV4() *bool
	SetIPV6(v bool) *DescribeSiteMonitorISPCityListRequest
	GetIPV6() *bool
	SetIsp(v string) *DescribeSiteMonitorISPCityListRequest
	GetIsp() *string
	SetRegionId(v string) *DescribeSiteMonitorISPCityListRequest
	GetRegionId() *string
	SetViewAll(v bool) *DescribeSiteMonitorISPCityListRequest
	GetViewAll() *bool
}

type DescribeSiteMonitorISPCityListRequest struct {
	// The name or ID of the city where the carrier detection point resides.
	//
	// > Fuzzy match is supported for city names.
	//
	// example:
	//
	// Guiyang
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Specifies whether to query IPv4 probes. Valid values:
	//
	// - true (default): Queries IPv4 probes.
	//
	// - false: Does not query IPv4 probes.
	//
	// example:
	//
	// true
	IPV4 *bool `json:"IPV4,omitempty" xml:"IPV4,omitempty"`
	// Specifies whether to query IPv6 probes. Valid values:
	//
	// - true (default): Queries IPv6 probes.
	//
	// - false: Does not query IPv6 probes.
	//
	// example:
	//
	// true
	IPV6 *bool `json:"IPV6,omitempty" xml:"IPV6,omitempty"`
	// The name or ID of the carrier detection point.
	//
	// > Fuzzy match is supported for carrier names.
	//
	// example:
	//
	// China Unicom
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return all detection points. Valid values:
	//
	// - true (default): Returns all detection points.
	//
	// - false: Returns only available detection points.
	//
	// example:
	//
	// true
	ViewAll *bool `json:"ViewAll,omitempty" xml:"ViewAll,omitempty"`
}

func (s DescribeSiteMonitorISPCityListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListRequest) GetCity() *string {
	return s.City
}

func (s *DescribeSiteMonitorISPCityListRequest) GetIPV4() *bool {
	return s.IPV4
}

func (s *DescribeSiteMonitorISPCityListRequest) GetIPV6() *bool {
	return s.IPV6
}

func (s *DescribeSiteMonitorISPCityListRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSiteMonitorISPCityListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorISPCityListRequest) GetViewAll() *bool {
	return s.ViewAll
}

func (s *DescribeSiteMonitorISPCityListRequest) SetCity(v string) *DescribeSiteMonitorISPCityListRequest {
	s.City = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) SetIPV4(v bool) *DescribeSiteMonitorISPCityListRequest {
	s.IPV4 = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) SetIPV6(v bool) *DescribeSiteMonitorISPCityListRequest {
	s.IPV6 = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) SetIsp(v string) *DescribeSiteMonitorISPCityListRequest {
	s.Isp = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) SetRegionId(v string) *DescribeSiteMonitorISPCityListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) SetViewAll(v bool) *DescribeSiteMonitorISPCityListRequest {
	s.ViewAll = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListRequest) Validate() error {
	return dara.Validate(s)
}
