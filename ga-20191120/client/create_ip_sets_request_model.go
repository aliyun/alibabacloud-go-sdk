// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegion(v []*CreateIpSetsRequestAccelerateRegion) *CreateIpSetsRequest
	GetAccelerateRegion() []*CreateIpSetsRequestAccelerateRegion
	SetAcceleratorId(v string) *CreateIpSetsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateIpSetsRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateIpSetsRequest
	GetRegionId() *string
}

type CreateIpSetsRequest struct {
	// The information about the acceleration regions.
	//
	// This parameter is required.
	AccelerateRegion []*CreateIpSetsRequestAccelerateRegion `json:"AccelerateRegion,omitempty" xml:"AccelerateRegion,omitempty" type:"Repeated"`
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmqy****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateIpSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *CreateIpSetsRequest) GetAccelerateRegion() []*CreateIpSetsRequestAccelerateRegion {
	return s.AccelerateRegion
}

func (s *CreateIpSetsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateIpSetsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpSetsRequest) SetAccelerateRegion(v []*CreateIpSetsRequestAccelerateRegion) *CreateIpSetsRequest {
	s.AccelerateRegion = v
	return s
}

func (s *CreateIpSetsRequest) SetAcceleratorId(v string) *CreateIpSetsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateIpSetsRequest) SetClientToken(v string) *CreateIpSetsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpSetsRequest) SetRegionId(v string) *CreateIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpSetsRequest) Validate() error {
	if s.AccelerateRegion != nil {
		for _, item := range s.AccelerateRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIpSetsRequestAccelerateRegion struct {
	// The ID of the acceleration region.
	//
	// The number of regions that you can add varies based on the specification of the GA instance. For more information, see [Overview](https://help.aliyun.com/document_detail/153127.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The bandwidth that you want to allocate to the acceleration region. Unit: **Mbit/s**.
	//
	// >	- This parameter is required.
	//
	// >	- You must allocate at least 2 Mbit/s of bandwidth to each acceleration region.
	//
	// >	- The total bandwidth of all acceleration regions cannot exceed the bandwidth limit of your basic bandwidth plan.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IP version used to connect to the GA instance. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// 	- **DUAL_STACK**: IPv4 and IPv6
	//
	// > Only pay-as-you-go standard GA instances support DUAL_STACK.
	//
	// example:
	//
	// IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values:
	//
	// 	- **BGP**: BGP (Multi-ISP) lines.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines
	//
	// > 	- This parameter is required only if the bandwidth metering method of the GA instance is **pay-by-data transfer**.
	//
	// >	- Different acceleration regions support different line types of EIPs.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
}

func (s CreateIpSetsRequestAccelerateRegion) String() string {
	return dara.Prettify(s)
}

func (s CreateIpSetsRequestAccelerateRegion) GoString() string {
	return s.String()
}

func (s *CreateIpSetsRequestAccelerateRegion) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *CreateIpSetsRequestAccelerateRegion) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateIpSetsRequestAccelerateRegion) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateIpSetsRequestAccelerateRegion) GetIspType() *string {
	return s.IspType
}

func (s *CreateIpSetsRequestAccelerateRegion) SetAccelerateRegionId(v string) *CreateIpSetsRequestAccelerateRegion {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetBandwidth(v int32) *CreateIpSetsRequestAccelerateRegion {
	s.Bandwidth = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetIpVersion(v string) *CreateIpSetsRequestAccelerateRegion {
	s.IpVersion = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetIspType(v string) *CreateIpSetsRequestAccelerateRegion {
	s.IspType = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) Validate() error {
	return dara.Validate(s)
}
