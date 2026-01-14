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
	// The ID of the acceleration region.
	//
	// You can call the [ListAvailableBusiRegions](https://help.aliyun.com/document_detail/261190.html) operation to query the most recent acceleration region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth that you want to allocate to the acceleration region. Unit: **Mbit/s**.
	//
	// You must allocate at least 2 Mbit/s of bandwidth to the acceleration region.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) lines.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines.
	//
	// Valid values if you are allowed to use single-ISP bandwidth:
	//
	// 	- **ChinaTelecom**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaTelecom_L2**
	//
	// 	- **ChinaUnicom_L2**
	//
	// 	- **ChinaMobile_L2**
	//
	// >
	//
	// 	- If the bandwidth metering method of the GA instance is **pay-by-data-transfer**, this parameter is required.
	//
	// 	- If the acceleration region is China (Hong Kong) and a basic bandwidth plan whose bandwidth type is Premium is associated with the GA instance, the default value of IspType is BGP_PRO.
	//
	// 	- The supported single-ISP type varies based on the acceleration region.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	// The region ID of the basic GA instance. Set the value to **cn-hangzhou**.
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
