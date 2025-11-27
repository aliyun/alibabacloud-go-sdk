// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// The available regions and zones.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	RDSRegion []*DescribeRegionsResponseBodyRegionsRDSRegion `json:"RDSRegion,omitempty" xml:"RDSRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRDSRegion() []*DescribeRegionsResponseBodyRegionsRDSRegion {
	return s.RDSRegion
}

func (s *DescribeRegionsResponseBodyRegions) SetRDSRegion(v []*DescribeRegionsResponseBodyRegionsRDSRegion) *DescribeRegionsResponseBodyRegions {
	s.RDSRegion = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	if s.RDSRegion != nil {
		for _, item := range s.RDSRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsRDSRegion struct {
	// The region name. The return value of this parameter is in the language that is specified by the **AcceptLanguage*	- parameter. For example, if the value of the RegionId parameter in the response is cn-hangzhou, the following values are returned for the LocalName parameter:
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **zh-CN**, the value  1()is returned for the LocalName parameter.
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **en-US**, the value China (Hangzhou) is returned for the LocalName parameter.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint that is used to connect to Alibaba Cloud services in the region. For more information, see [Endpoints](https://help.aliyun.com/document_detail/610370.html).
	//
	// example:
	//
	// rds.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name. The return value of this parameter is in the language that is specified by the **AcceptLanguage*	- parameter. For example, if the value of the ZoneId parameter in the response is cn-hangzhou-j, the following values are returned for the ZoneName parameter:
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **zh-CN**, the value   J is returned for the ZoneName parameter.
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **en-US**, the value Hangzhou Zone J is returned for the ZoneName parameter.
	//
	// example:
	//
	// Hangzhou Zone H
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetZoneName(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.ZoneName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) Validate() error {
	return dara.Validate(s)
}
