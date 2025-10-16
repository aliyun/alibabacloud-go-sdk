// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcZoneResponseBody
	GetRequestId() *string
	SetZoneList(v []*DescribeVpcZoneResponseBodyZoneList) *DescribeVpcZoneResponseBody
	GetZoneList() []*DescribeVpcZoneResponseBodyZoneList
}

type DescribeVpcZoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 694DFBF3-C060-529F-92D0-7FC7E0DA1E21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones.
	ZoneList []*DescribeVpcZoneResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeVpcZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcZoneResponseBody) GetZoneList() []*DescribeVpcZoneResponseBodyZoneList {
	return s.ZoneList
}

func (s *DescribeVpcZoneResponseBody) SetRequestId(v string) *DescribeVpcZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcZoneResponseBody) SetZoneList(v []*DescribeVpcZoneResponseBodyZoneList) *DescribeVpcZoneResponseBody {
	s.ZoneList = v
	return s
}

func (s *DescribeVpcZoneResponseBody) Validate() error {
	if s.ZoneList != nil {
		for _, item := range s.ZoneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcZoneResponseBodyZoneList struct {
	// The name of the zone.
	//
	// example:
	//
	// Hangzhou Zone B
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone type. Default value: AvailabilityZone. This value indicates Alibaba Cloud zones.
	//
	// example:
	//
	// AvailabilityZone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeVpcZoneResponseBodyZoneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcZoneResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponseBodyZoneList) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeVpcZoneResponseBodyZoneList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcZoneResponseBodyZoneList) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetLocalName(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.LocalName = &v
	return s
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetZoneId(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetZoneType(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.ZoneType = &v
	return s
}

func (s *DescribeVpcZoneResponseBodyZoneList) Validate() error {
	return dara.Validate(s)
}
