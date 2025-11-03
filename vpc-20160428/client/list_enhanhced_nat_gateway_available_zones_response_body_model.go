// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnhanhcedNatGatewayAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEnhanhcedNatGatewayAvailableZonesResponseBody
	GetRequestId() *string
	SetZones(v []*ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) *ListEnhanhcedNatGatewayAvailableZonesResponseBody
	GetZones() []*ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones
}

type ListEnhanhcedNatGatewayAvailableZonesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AE6F367-52EA-535D-9A3D-EF23D70527C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	Zones []*ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBody) GetZones() []*ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones {
	return s.Zones
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBody) SetRequestId(v string) *ListEnhanhcedNatGatewayAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBody) SetZones(v []*ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) *ListEnhanhcedNatGatewayAvailableZonesResponseBody {
	s.Zones = v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBody) Validate() error {
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

type ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones struct {
	// The name of the zone.
	//
	// example:
	//
	// Dubai Zone A
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the zone where the instance is deployed.
	//
	// example:
	//
	// me-east-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) GetLocalName() *string {
	return s.LocalName
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) SetLocalName(v string) *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) SetZoneId(v string) *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
