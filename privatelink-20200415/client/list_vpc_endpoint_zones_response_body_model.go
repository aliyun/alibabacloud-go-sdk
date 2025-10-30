// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointZonesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointZonesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointZonesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVpcEndpointZonesResponseBody
	GetTotalCount() *int32
	SetZones(v []*ListVpcEndpointZonesResponseBodyZones) *ListVpcEndpointZonesResponseBody
	GetZones() []*ListVpcEndpointZonesResponseBodyZones
}

type ListVpcEndpointZonesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the zones.
	Zones []*ListVpcEndpointZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointZonesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpcEndpointZonesResponseBody) GetZones() []*ListVpcEndpointZonesResponseBodyZones {
	return s.Zones
}

func (s *ListVpcEndpointZonesResponseBody) SetMaxResults(v int32) *ListVpcEndpointZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetNextToken(v string) *ListVpcEndpointZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetRequestId(v string) *ListVpcEndpointZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetTotalCount(v int32) *ListVpcEndpointZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetZones(v []*ListVpcEndpointZonesResponseBodyZones) *ListVpcEndpointZonesResponseBody {
	s.Zones = v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointZonesResponseBodyZones struct {
	// The ID of the endpoint ENI.
	//
	// example:
	//
	// eni-hp3c8qj1tyct8aqy****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The IP address of the endpoint ENI.
	//
	// example:
	//
	// 192.168.2.23
	EniIp *string `json:"EniIp,omitempty" xml:"EniIp,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the zone. The system automatically creates an endpoint elastic network interface (ENI) in the vSwitch.
	//
	// example:
	//
	// vsw-hjkshjvdkdvd****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The domain name of the zone.
	//
	// After the endpoint in the zone is connected to the endpoint service, you can access the service resources of the endpoint service by using the domain name of the zone.
	//
	// example:
	//
	// ep-hp3f033dp24c5yc9****-cn-huhehaote.epsrv-hp3itcpowf37m3d5****.cn-huhehaote-a.privatelink.aliyuncs.com
	ZoneDomain *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-huhehaote-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ZoneIpv6Address *string `json:"ZoneIpv6Address,omitempty" xml:"ZoneIpv6Address,omitempty"`
	// The state of the zone. Valid values:
	//
	// 	- **Creating**: The zone is being created.
	//
	// 	- **Wait**: The zone is to be connected.
	//
	// 	- **Connected**: The zone is connected.
	//
	// 	- **Deleting**: The zone is being deleted.
	//
	// 	- **Disconnecting**: The zone is being disconnected.
	//
	// 	- **Disconnected**: The zone is disconnected.
	//
	// 	- **Connecting**: The zone is being connected.
	//
	// example:
	//
	// Wait
	ZoneStatus *string `json:"ZoneStatus,omitempty" xml:"ZoneStatus,omitempty"`
}

func (s ListVpcEndpointZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetEniId() *string {
	return s.EniId
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetEniIp() *string {
	return s.EniIp
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetZoneDomain() *string {
	return s.ZoneDomain
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetZoneIpv6Address() *string {
	return s.ZoneIpv6Address
}

func (s *ListVpcEndpointZonesResponseBodyZones) GetZoneStatus() *string {
	return s.ZoneStatus
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetEniId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetEniIp(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.EniIp = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetRegionId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetVSwitchId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneDomain(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneDomain = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneIpv6Address(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneIpv6Address = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneStatus(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneStatus = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
