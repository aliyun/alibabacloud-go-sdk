// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegionId(v string) *CreateBasicIpSetRequest
	GetAccelerateRegionId() *string
	SetAcceleratorId(v string) *CreateBasicIpSetRequest
	GetAcceleratorId() *string
	SetBandwidth(v int64) *CreateBasicIpSetRequest
	GetBandwidth() *int64
	SetClientToken(v string) *CreateBasicIpSetRequest
	GetClientToken() *string
	SetIspType(v string) *CreateBasicIpSetRequest
	GetIspType() *string
	SetRegionId(v string) *CreateBasicIpSetRequest
	GetRegionId() *string
}

type CreateBasicIpSetRequest struct {
	// The ID of the region to be accelerated.
	//
	// You can invoke the [ListAvailableBusiRegions](https://help.aliyun.com/document_detail/261190.html) operation to query the active acceleration regions for the specified Alibaba Cloud Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The instance ID of the basic Alibaba Cloud Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth of the acceleration area. Unit: **Mbps**.
	//
	// The minimum bandwidth that can be allocated to an acceleration area is 2 Mbps.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of a request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ISP type of the public network in the acceleration region. Valid values:
	//
	// - **BGP*	- (default): BGP (Multi-ISP)
	//
	// - **BGP_PRO**: BGP (Multi-ISP) Pro
	//
	// If you are a whitelist user of single-ISP bandwidth, you can also select the following types:
	//
	// - **ChinaTelecom**: China Telecom (single ISP)
	//
	// - **ChinaUnicom**: China Unicom (single ISP)
	//
	// - **ChinaMobile**: China Shift (single ISP)
	//
	// - **ChinaTelecom_L2**: China Telecom (single ISP)_L2
	//
	// - **ChinaUnicom_L2**: China Unicom (single ISP)_L2
	//
	// - **ChinaMobile_L2**: China Shift (single ISP)_L2
	//
	// > - This parameter is required for basic Alibaba Cloud Global Accelerator (GA) instances that use the **pay-by-traffic*	- billing method.
	//
	// > - If the acceleration region of the basic Alibaba Cloud Global Accelerator (GA) instance is Hong Kong (China) and the instance is attached with a basic bandwidth plan of the BGP (Multi-ISP) Pro type, the default value is BGP (Multi-ISP) Pro.
	//
	// > - The supported single-ISP line types vary by acceleration region.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	// The region ID of the basic Alibaba Cloud Global Accelerator (GA) instance. Set the value to **ap-southeast-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetRequest) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *CreateBasicIpSetRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicIpSetRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateBasicIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicIpSetRequest) GetIspType() *string {
	return s.IspType
}

func (s *CreateBasicIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicIpSetRequest) SetAccelerateRegionId(v string) *CreateBasicIpSetRequest {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetAcceleratorId(v string) *CreateBasicIpSetRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetBandwidth(v int64) *CreateBasicIpSetRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetClientToken(v string) *CreateBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetIspType(v string) *CreateBasicIpSetRequest {
	s.IspType = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetRegionId(v string) *CreateBasicIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicIpSetRequest) Validate() error {
	return dara.Validate(s)
}
