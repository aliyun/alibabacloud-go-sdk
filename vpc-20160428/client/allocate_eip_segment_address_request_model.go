// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipSegmentAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *AllocateEipSegmentAddressRequest
	GetBandwidth() *string
	SetClientToken(v string) *AllocateEipSegmentAddressRequest
	GetClientToken() *string
	SetEipMask(v string) *AllocateEipSegmentAddressRequest
	GetEipMask() *string
	SetInternetChargeType(v string) *AllocateEipSegmentAddressRequest
	GetInternetChargeType() *string
	SetIsp(v string) *AllocateEipSegmentAddressRequest
	GetIsp() *string
	SetNetmode(v string) *AllocateEipSegmentAddressRequest
	GetNetmode() *string
	SetOwnerAccount(v string) *AllocateEipSegmentAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateEipSegmentAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AllocateEipSegmentAddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AllocateEipSegmentAddressRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AllocateEipSegmentAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateEipSegmentAddressRequest
	GetResourceOwnerId() *int64
	SetZone(v string) *AllocateEipSegmentAddressRequest
	GetZone() *string
}

type AllocateEipSegmentAddressRequest struct {
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// - If **InternetChargeType*	- is set to **PayByBandwidth**, valid values of **Bandwidth*	- are **1*	- to **500**.
	//
	// - If **InternetChargeType*	- is set to **PayByTraffic**, valid values of **Bandwidth*	- are **1*	- to **200**.
	//
	// Default value: **5*	- Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subnet mask of the contiguous EIPs. Valid values:
	//
	// - **28**: The system allocates 16 contiguous EIPs per call.
	//
	// - **27**: The system allocates 32 contiguous EIPs per call.
	//
	// - **26**: The system allocates 64 contiguous EIPs per call.
	//
	// - **25**: The system allocates 128 contiguous EIPs per call.
	//
	// - **24**: The system allocates 256 contiguous EIPs per call.
	//
	// > Due to IP address reservation, the actual number of contiguous EIPs may be 1, 3, or 4 fewer than expected.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28
	EipMask *string `json:"EipMask,omitempty" xml:"EipMask,omitempty"`
	// The billable methods of the contiguous EIPs. Valid values:
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// - **PayByTraffic**: pay-by-data-transfer.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The line type. Valid values:
	//
	// - **BGP*	- (default): BGP (multi-ISP) line. All regions support BGP (multi-ISP) EIPs.
	//
	// - **BGP_PRO**: BGP (multi-ISP) premium line. Only Hong Kong (China), Singapore, Tokyo (Japan), Kuala Lumpur (Malaysia), Manila (Philippines), Jakarta (Indonesia), and Bangkok (Thailand) regions support BGP (multi-ISP) premium EIPs.
	//
	// For more information about BGP (multi-ISP) lines and BGP (multi-ISP) premium lines, see [EIP line types](https://help.aliyun.com/document_detail/32321.html).
	//
	// If you are a whitelist user of single-ISP bandwidth, you can also select the following types:
	//
	// - **ChinaTelecom**: China Telecom
	//
	// - **ChinaUnicom**: China Unicom
	//
	// - **ChinaMobile**: China Mobile
	//
	// - **ChinaTelecom_L2**: China Telecom L2
	//
	// - **ChinaUnicom_L2**: China Unicom L2
	//
	// - **ChinaMobile_L2**: China Mobile L2
	//
	// If you are an Alibaba Finance Cloud user, this parameter is required. Set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The network type. Set the value to **public**, which specifies the public network.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the contiguous EIPs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone of the contiguous EIP group.
	//
	// example:
	//
	// cn-hangzhou-a
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s AllocateEipSegmentAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipSegmentAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipSegmentAddressRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *AllocateEipSegmentAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateEipSegmentAddressRequest) GetEipMask() *string {
	return s.EipMask
}

func (s *AllocateEipSegmentAddressRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateEipSegmentAddressRequest) GetIsp() *string {
	return s.Isp
}

func (s *AllocateEipSegmentAddressRequest) GetNetmode() *string {
	return s.Netmode
}

func (s *AllocateEipSegmentAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateEipSegmentAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateEipSegmentAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateEipSegmentAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateEipSegmentAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateEipSegmentAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateEipSegmentAddressRequest) GetZone() *string {
	return s.Zone
}

func (s *AllocateEipSegmentAddressRequest) SetBandwidth(v string) *AllocateEipSegmentAddressRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetClientToken(v string) *AllocateEipSegmentAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetEipMask(v string) *AllocateEipSegmentAddressRequest {
	s.EipMask = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetInternetChargeType(v string) *AllocateEipSegmentAddressRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetIsp(v string) *AllocateEipSegmentAddressRequest {
	s.Isp = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetNetmode(v string) *AllocateEipSegmentAddressRequest {
	s.Netmode = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetOwnerAccount(v string) *AllocateEipSegmentAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetOwnerId(v int64) *AllocateEipSegmentAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetRegionId(v string) *AllocateEipSegmentAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetResourceGroupId(v string) *AllocateEipSegmentAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetResourceOwnerAccount(v string) *AllocateEipSegmentAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetResourceOwnerId(v int64) *AllocateEipSegmentAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) SetZone(v string) *AllocateEipSegmentAddressRequest {
	s.Zone = &v
	return s
}

func (s *AllocateEipSegmentAddressRequest) Validate() error {
	return dara.Validate(s)
}
