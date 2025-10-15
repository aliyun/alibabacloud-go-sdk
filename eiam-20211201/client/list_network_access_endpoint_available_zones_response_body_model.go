// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListNetworkAccessEndpointAvailableZonesResponseBody
	GetRequestId() *string
	SetZones(v []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones) *ListNetworkAccessEndpointAvailableZonesResponseBody
	GetZones() []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones
}

type ListNetworkAccessEndpointAvailableZonesResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) GetZones() []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones {
	return s.Zones
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) SetZones(v []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones) *ListNetworkAccessEndpointAvailableZonesResponseBody {
	s.Zones = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) Validate() error {
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

type ListNetworkAccessEndpointAvailableZonesResponseBodyZones struct {
	// 可用区名称。
	//
	// example:
	//
	// 华东1（杭州）可用区J
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 可用区ID。
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) GetLocalName() *string {
	return s.LocalName
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) SetLocalName(v string) *ListNetworkAccessEndpointAvailableZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) SetZoneId(v string) *ListNetworkAccessEndpointAvailableZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
