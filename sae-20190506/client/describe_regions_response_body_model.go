// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeRegionsResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeRegionsResponseBody
	GetMessage() *string
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DDE85827-B0B3-4E56-86E8-17C42009****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
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
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRegion() []*DescribeRegionsResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The region name. Valid values:
	//
	// 	- **China (Hangzhou)**
	//
	// 	- **China (Shanghai)**
	//
	// 	- **China (Beijing)**
	//
	// 	- **China (Zhangjiakou)**
	//
	// 	- **China (Shenzhen)**
	//
	// 	- **China (Guangzhou)**
	//
	// 	- **China (Hong Kong)**
	//
	// 	- **Singapore**
	//
	// 	- **US (Silicon Valley)**
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The recommended zones.
	RecommendZones *DescribeRegionsResponseBodyRegionsRegionRecommendZones `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Struct"`
	// The endpoint for the region. Valid values:
	//
	// 	- **sae.cn-hangzhou.aliyuncs.com**
	//
	// 	- **sae.cn-shanghai.aliyuncs.com**
	//
	// 	- **sae.cn-beijing.aliyuncs.com**
	//
	// 	- **sae.cn-zhangjiakou.aliyuncs.com**
	//
	// 	- **sae.cn-shenzhen.aliyuncs.com**
	//
	// 	- **sae.cn-guangzhou.aliyuncs.com**
	//
	// 	- **sae.cn-hongkong.aliyuncs.com**
	//
	// 	- **sae.ap-southeast-1.aliyuncs.com**
	//
	// 	- **sae.us-west-1.aliyuncs.com**
	//
	// example:
	//
	// sae.cn-shanghai.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID. Valid values:
	//
	// 	- **cn-hangzhou**: the ID of the China (Hangzhou) region
	//
	// 	- **cn-shanghai**: the ID of the China (Shanghai) region
	//
	// 	- **cn-beijing**: the ID of the China (Beijing) region
	//
	// 	- **cn-zhangjiakou**: the ID of the China (Zhangjiakou) region
	//
	// 	- **cn-shenzhen**: the ID of the China (Shenzhen) region
	//
	// 	- **cn-guangzhou**: the ID of the China (Guangzhou) region
	//
	// 	- **cn-hongkong**: the ID of the China (Hong Kong) region
	//
	// 	- **ap-southeast-1**: the ID of the Singapore region
	//
	// 	- **us-west-1**: the ID of the US (Silicon Valley) region
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRecommendZones() *DescribeRegionsResponseBodyRegionsRegionRecommendZones {
	return s.RecommendZones
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRecommendZones(v *DescribeRegionsResponseBodyRegionsRegionRecommendZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.RecommendZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionsRegionRecommendZones struct {
	RecommendZone []*string `json:"RecommendZone,omitempty" xml:"RecommendZone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionRecommendZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionRecommendZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionRecommendZones) GetRecommendZone() []*string {
	return s.RecommendZone
}

func (s *DescribeRegionsResponseBodyRegionsRegionRecommendZones) SetRecommendZone(v []*string) *DescribeRegionsResponseBodyRegionsRegionRecommendZones {
	s.RecommendZone = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionRecommendZones) Validate() error {
	return dara.Validate(s)
}
