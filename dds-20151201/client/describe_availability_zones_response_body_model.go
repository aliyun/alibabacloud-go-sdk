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
	// The details of the zones in which you can create ApsaraDB for MongoDB instances.
	AvailableZones []*DescribeAvailabilityZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type DescribeAvailabilityZonesResponseBodyAvailableZones struct {
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the latest available regions.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// The return value of the ZoneName parameter is in the language that is specified by the **AcceptLanguage*	- parameter. For example, if the value of the ZoneId parameter in the response is **cn-hangzhou-h**, the following values are returned for the ZoneName parameter:
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **zh**, **H*	- is returned for the ZoneName parameter.
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **en**, **Hangzhou Zone H*	- is returned for the ZoneName parameter.
	//
	// example:
	//
	// Hangzhou Zone H
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
