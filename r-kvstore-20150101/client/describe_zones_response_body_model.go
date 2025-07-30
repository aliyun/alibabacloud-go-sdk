// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() *DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1D42F072-72FE-4DC4-BB8E-64B1D298****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetZones() *DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZones struct {
	KVStoreZone []*DescribeZonesResponseBodyZonesKVStoreZone `json:"KVStoreZone,omitempty" xml:"KVStoreZone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetKVStoreZone() []*DescribeZonesResponseBodyZonesKVStoreZone {
	return s.KVStoreZone
}

func (s *DescribeZonesResponseBodyZones) SetKVStoreZone(v []*DescribeZonesResponseBodyZonesKVStoreZone) *DescribeZonesResponseBodyZones {
	s.KVStoreZone = v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesKVStoreZone struct {
	// Indicates whether Tair (Redis OSS-compatible) instances can be created in the current zone. Valid values:
	//
	// 	- **true**: Tair (Redis OSS-compatible) instances cannot be created in the current zone.
	//
	// 	- **false**: Tair (Redis OSS-compatible) instances can be created in the current zone.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Indicates whether the zone is managed by ApsaraDB RDS. The return value of this parameter is **true*	- in Tair (Redis OSS-compatible).
	//
	// example:
	//
	// true
	IsRds *bool `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the network type of the instance can be changed from the classic network to Virtual Private Cloud (VPC). Valid values:
	//
	// 	- **true**: The network type of the instance can be changed from the classic network to VPC.
	//
	// 	- **false**: The network type of the instance cannot be changed from the classic network to VPC.
	//
	// example:
	//
	// true
	SwitchNetwork *bool `json:"SwitchNetwork,omitempty" xml:"SwitchNetwork,omitempty"`
	// The ID of the zone within the specified region.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone within the specified region.
	//
	// example:
	//
	// Hohhot Zone B
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeZonesResponseBodyZonesKVStoreZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesKVStoreZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetIsRds() *bool {
	return s.IsRds
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetSwitchNetwork() *bool {
	return s.SwitchNetwork
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetDisabled(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.Disabled = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetIsRds(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.IsRds = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetRegionId(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetSwitchNetwork(v bool) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.SwitchNetwork = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) SetZoneName(v string) *DescribeZonesResponseBodyZonesKVStoreZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesKVStoreZone) Validate() error {
	return dara.Validate(s)
}
