// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceZoneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAccessInstanceZoneListResponseBody
	GetRequestId() *string
	SetZoneList(v []*string) *DescribeAccessInstanceZoneListResponseBody
	GetZoneList() []*string
	SetZones(v []*DescribeAccessInstanceZoneListResponseBodyZones) *DescribeAccessInstanceZoneListResponseBody
	GetZones() []*DescribeAccessInstanceZoneListResponseBodyZones
}

type DescribeAccessInstanceZoneListResponseBody struct {
	// example:
	//
	// 31306819-C4BC-56F3-BBE6-*****
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneList  []*string                                          `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
	Zones     []*DescribeAccessInstanceZoneListResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeAccessInstanceZoneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceZoneListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceZoneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessInstanceZoneListResponseBody) GetZoneList() []*string {
	return s.ZoneList
}

func (s *DescribeAccessInstanceZoneListResponseBody) GetZones() []*DescribeAccessInstanceZoneListResponseBodyZones {
	return s.Zones
}

func (s *DescribeAccessInstanceZoneListResponseBody) SetRequestId(v string) *DescribeAccessInstanceZoneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessInstanceZoneListResponseBody) SetZoneList(v []*string) *DescribeAccessInstanceZoneListResponseBody {
	s.ZoneList = v
	return s
}

func (s *DescribeAccessInstanceZoneListResponseBody) SetZones(v []*DescribeAccessInstanceZoneListResponseBodyZones) *DescribeAccessInstanceZoneListResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeAccessInstanceZoneListResponseBody) Validate() error {
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

type DescribeAccessInstanceZoneListResponseBodyZones struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAccessInstanceZoneListResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceZoneListResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceZoneListResponseBodyZones) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeAccessInstanceZoneListResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAccessInstanceZoneListResponseBodyZones) SetLocalName(v string) *DescribeAccessInstanceZoneListResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeAccessInstanceZoneListResponseBodyZones) SetZoneId(v string) *DescribeAccessInstanceZoneListResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeAccessInstanceZoneListResponseBodyZones) Validate() error {
	return dara.Validate(s)
}
