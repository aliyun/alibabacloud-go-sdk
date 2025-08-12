// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyntheticProbeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *DescribeSyntheticProbeListRequest
	GetCity() *string
	SetIdcProbe(v bool) *DescribeSyntheticProbeListRequest
	GetIdcProbe() *bool
	SetIpv4(v bool) *DescribeSyntheticProbeListRequest
	GetIpv4() *bool
	SetIpv6(v bool) *DescribeSyntheticProbeListRequest
	GetIpv6() *bool
	SetIsp(v string) *DescribeSyntheticProbeListRequest
	GetIsp() *string
	SetLmProbe(v bool) *DescribeSyntheticProbeListRequest
	GetLmProbe() *bool
	SetMbProbe(v bool) *DescribeSyntheticProbeListRequest
	GetMbProbe() *bool
	SetRegionId(v string) *DescribeSyntheticProbeListRequest
	GetRegionId() *string
	SetViewAll(v bool) *DescribeSyntheticProbeListRequest
	GetViewAll() *bool
}

type DescribeSyntheticProbeListRequest struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// false
	IdcProbe *bool `json:"IdcProbe,omitempty" xml:"IdcProbe,omitempty"`
	// example:
	//
	// false
	Ipv4 *bool `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	// example:
	//
	// false
	Ipv6 *bool   `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	Isp  *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// false
	LmProbe *bool `json:"LmProbe,omitempty" xml:"LmProbe,omitempty"`
	// example:
	//
	// false
	MbProbe  *bool   `json:"MbProbe,omitempty" xml:"MbProbe,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	ViewAll *bool `json:"ViewAll,omitempty" xml:"ViewAll,omitempty"`
}

func (s DescribeSyntheticProbeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyntheticProbeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyntheticProbeListRequest) GetCity() *string {
	return s.City
}

func (s *DescribeSyntheticProbeListRequest) GetIdcProbe() *bool {
	return s.IdcProbe
}

func (s *DescribeSyntheticProbeListRequest) GetIpv4() *bool {
	return s.Ipv4
}

func (s *DescribeSyntheticProbeListRequest) GetIpv6() *bool {
	return s.Ipv6
}

func (s *DescribeSyntheticProbeListRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSyntheticProbeListRequest) GetLmProbe() *bool {
	return s.LmProbe
}

func (s *DescribeSyntheticProbeListRequest) GetMbProbe() *bool {
	return s.MbProbe
}

func (s *DescribeSyntheticProbeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyntheticProbeListRequest) GetViewAll() *bool {
	return s.ViewAll
}

func (s *DescribeSyntheticProbeListRequest) SetCity(v string) *DescribeSyntheticProbeListRequest {
	s.City = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetIdcProbe(v bool) *DescribeSyntheticProbeListRequest {
	s.IdcProbe = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetIpv4(v bool) *DescribeSyntheticProbeListRequest {
	s.Ipv4 = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetIpv6(v bool) *DescribeSyntheticProbeListRequest {
	s.Ipv6 = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetIsp(v string) *DescribeSyntheticProbeListRequest {
	s.Isp = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetLmProbe(v bool) *DescribeSyntheticProbeListRequest {
	s.LmProbe = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetMbProbe(v bool) *DescribeSyntheticProbeListRequest {
	s.MbProbe = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetRegionId(v string) *DescribeSyntheticProbeListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) SetViewAll(v bool) *DescribeSyntheticProbeListRequest {
	s.ViewAll = &v
	return s
}

func (s *DescribeSyntheticProbeListRequest) Validate() error {
	return dara.Validate(s)
}
