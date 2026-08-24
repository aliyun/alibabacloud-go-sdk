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
	SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() []*DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// The request ID. The request ID is returned regardless of whether the API call is successful.
	//
	// example:
	//
	// 6AA27F1A-A62C-59C3-BCC7-D1DFA4E7EEA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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

func (s *DescribeZonesResponseBody) GetZones() []*DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
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

type DescribeZonesResponseBodyZones struct {
	// The localized name of the zone.
	//
	// example:
	//
	// Hangzhou Zone A
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
