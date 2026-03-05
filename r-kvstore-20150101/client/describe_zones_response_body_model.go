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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.KVStoreZone != nil {
		for _, item := range s.KVStoreZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesKVStoreZone struct {
	Disabled      *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	IsRds         *bool   `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SwitchNetwork *bool   `json:"SwitchNetwork,omitempty" xml:"SwitchNetwork,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName      *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
