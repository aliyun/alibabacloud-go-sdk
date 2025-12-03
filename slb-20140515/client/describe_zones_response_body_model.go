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
	// The request ID.
	//
	// example:
	//
	// A48D35FF-440A-4BC0-A4A2-A9BF69B7E43A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones.
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
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetZone() []*DescribeZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZone struct {
	// The name of the zone.
	//
	// example:
	//
	// The list of secondary zones.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The secondary zones.
	SlaveZones *DescribeZonesResponseBodyZonesZoneSlaveZones `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Struct"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeZonesResponseBodyZonesZone) GetSlaveZones() *DescribeZonesResponseBodyZonesZoneSlaveZones {
	return s.SlaveZones
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetSlaveZones(v *DescribeZonesResponseBodyZonesZoneSlaveZones) *DescribeZonesResponseBodyZonesZone {
	s.SlaveZones = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) Validate() error {
	if s.SlaveZones != nil {
		if err := s.SlaveZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneSlaveZones struct {
	SlaveZone []*DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone `json:"SlaveZone,omitempty" xml:"SlaveZone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZones) GetSlaveZone() []*DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone {
	return s.SlaveZone
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZones) SetSlaveZone(v []*DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) *DescribeZonesResponseBodyZonesZoneSlaveZones {
	s.SlaveZone = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZones) Validate() error {
	if s.SlaveZone != nil {
		for _, item := range s.SlaveZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone struct {
	// The name of the secondary zone.
	//
	// example:
	//
	// Queries zones in a specified region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the secondary zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) Validate() error {
	return dara.Validate(s)
}
