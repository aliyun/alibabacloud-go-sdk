// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteDnsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsAddress(v []*string) *ModifyOfficeSiteDnsInfoRequest
	GetDnsAddress() []*string
	SetOfficeSiteId(v string) *ModifyOfficeSiteDnsInfoRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifyOfficeSiteDnsInfoRequest
	GetRegionId() *string
}

type ModifyOfficeSiteDnsInfoRequest struct {
	// The IP addresses of the custom DNS servers. Up to 2 IP addresses can be specified.
	DnsAddress []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-778418****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID of the instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteDnsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteDnsInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteDnsInfoRequest) GetDnsAddress() []*string {
	return s.DnsAddress
}

func (s *ModifyOfficeSiteDnsInfoRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteDnsInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteDnsInfoRequest) SetDnsAddress(v []*string) *ModifyOfficeSiteDnsInfoRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyOfficeSiteDnsInfoRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteDnsInfoRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteDnsInfoRequest) SetRegionId(v string) *ModifyOfficeSiteDnsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteDnsInfoRequest) Validate() error {
	return dara.Validate(s)
}
