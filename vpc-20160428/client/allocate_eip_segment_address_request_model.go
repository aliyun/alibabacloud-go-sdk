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
	// The maximum bandwidth of the contiguous EIP group. Unit: Mbit/s.
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**: **1*	- to **500**.****
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**: **1*	- to **200**.****
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PrePaid**: **1*	- to **1000**.****
	//
	// Default value: **5**. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subnet mask of the contiguous EIP group. Valid values:
	//
	// 	- **28**: applies for 16 contiguous EIPs in each call.
	//
	// 	- **27**: applies for 32 contiguous EIPs in each call.
	//
	// 	- **26**: applies for 64 contiguous EIPs in each call.
	//
	// 	- **25**: applies for 128 contiguous EIPs in each call.
	//
	// 	- **24**: applies for 256 contiguous EIPs in each call.
	//
	// >  Some IP address are reserved for specific purposes. Therefore, the actual number of the contiguous EIPs may be one, three, or four less than the expected number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28
	EipMask *string `json:"EipMask,omitempty" xml:"EipMask,omitempty"`
	// The metering method of contiguous EIPs. Valid values:
	//
	// 	- **PayByBandwidth*	- (default)
	//
	// 	- **PayByTraffic**
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) line The BGP (Multi-ISP) line is supported in all regions.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro line BGP (Multi-ISP) Pro line is supported only in the China (Hong Kong), Singapore, Japan (Tokyo), Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// For more information about the BGP (Multi-ISP) line and BGP (Multi-ISP) Pro line, see [EIP line types](https://help.aliyun.com/document_detail/32321.html).
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
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
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the parameter to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The network type. Set the value to **public**, which specifies the public network type.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the contiguous EIP group resides.
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
