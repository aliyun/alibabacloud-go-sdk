// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailabilityZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeAvailabilityZonesResponseBodyAvailableZones) *DescribeAvailabilityZonesResponseBody
	GetAvailableZones() []*DescribeAvailabilityZonesResponseBodyAvailableZones
	SetRequestId(v string) *DescribeAvailabilityZonesResponseBody
	GetRequestId() *string
}

type DescribeAvailabilityZonesResponseBody struct {
	// The details of the zones where resources are available for purchase.
	AvailableZones []*DescribeAvailabilityZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 19A13A33-0FAD-5120-8AE1-B1636F74DD80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailabilityZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailabilityZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailabilityZonesResponseBody) GetAvailableZones() []*DescribeAvailabilityZonesResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeAvailabilityZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailabilityZonesResponseBody) SetAvailableZones(v []*DescribeAvailabilityZonesResponseBodyAvailableZones) *DescribeAvailabilityZonesResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailabilityZonesResponseBody) SetRequestId(v string) *DescribeAvailabilityZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailabilityZonesResponseBody) Validate() error {
	if s.AvailableZones != nil {
		for _, item := range s.AvailableZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailabilityZonesResponseBodyAvailableZones struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name.
	//
	// The returned language is determined by the **AcceptLanguage*	- parameter. Take the **cn-hangzhou-h*	- zone as an example:
	//
	// - If **AcceptLanguage*	- is set to **zh**: 杭州 可用区H
	//
	// - If **AcceptLanguage*	- is set to **en**: Hangzhou Zone H
	//
	// example:
	//
	// 杭州 可用区H
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeAvailabilityZonesResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailabilityZonesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) SetRegionId(v string) *DescribeAvailabilityZonesResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) SetZoneId(v string) *DescribeAvailabilityZonesResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) SetZoneName(v string) *DescribeAvailabilityZonesResponseBodyAvailableZones {
	s.ZoneName = &v
	return s
}

func (s *DescribeAvailabilityZonesResponseBodyAvailableZones) Validate() error {
	return dara.Validate(s)
}
