// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersForPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetTotalCount() *int32
	SetVirtualBorderRouterForPhysicalConnectionSet(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetVirtualBorderRouterForPhysicalConnectionSet() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7C5AE8B3-A2D8-428D-A2FF-93A225C0821E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about VBRs.
	VirtualBorderRouterForPhysicalConnectionSet *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet `json:"VirtualBorderRouterForPhysicalConnectionSet,omitempty" xml:"VirtualBorderRouterForPhysicalConnectionSet,omitempty" type:"Struct"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetVirtualBorderRouterForPhysicalConnectionSet() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet {
	return s.VirtualBorderRouterForPhysicalConnectionSet
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetRequestId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetTotalCount(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetVirtualBorderRouterForPhysicalConnectionSet(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.VirtualBorderRouterForPhysicalConnectionSet = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet struct {
	VirtualBorderRouterForPhysicalConnectionType []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType `json:"VirtualBorderRouterForPhysicalConnectionType,omitempty" xml:"VirtualBorderRouterForPhysicalConnectionType,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) GetVirtualBorderRouterForPhysicalConnectionType() []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	return s.VirtualBorderRouterForPhysicalConnectionType
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) SetVirtualBorderRouterForPhysicalConnectionType(v []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet {
	s.VirtualBorderRouterForPhysicalConnectionType = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) Validate() error {
	return dara.Validate(s)
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType struct {
	// The time when the VBR was first activated.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	ActivationTime *string `json:"ActivationTime,omitempty" xml:"ActivationTime,omitempty"`
	// The bandwidth of the VBR that is associated with the Express Connect circuit. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The status of the bandwidth. Valid values:
	//
	// 	- **Active**
	//
	// 	- **Inactive**
	//
	// example:
	//
	// Active
	BandwidthStatus *string `json:"BandwidthStatus,omitempty" xml:"BandwidthStatus,omitempty"`
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// example:
	//
	// longtel0**
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The time when the VBR was created.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the ECC instance.
	//
	// example:
	//
	// ecc-sjghe****
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The IPv4 address of the gateway device on the Alibaba Cloud side.
	//
	// example:
	//
	// 192.168.XX.X
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// The IPv6 address of the gateway device on the Alibaba Cloud side.
	//
	// example:
	//
	// ipv6bw-uf6hcyzu65v98v3du****
	LocalIpv6GatewayIp *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	// The status of the VBR associated with the Express Connect circuit. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	PConnVbrBussinessStatus *string `json:"PConnVbrBussinessStatus,omitempty" xml:"PConnVbrBussinessStatus,omitempty"`
	// The billing method of the VBR. Valid values:
	//
	// 	- **PrePaid**: subscription. If you choose this billing method, make sure that your Apsara Stack account supports balance payments or credit payments.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PConnVbrChargeType *string `json:"PConnVbrChargeType,omitempty" xml:"PConnVbrChargeType,omitempty"`
	// The time when the VBR associated with the Express Connect circuit expires.
	//
	// example:
	//
	// 2021-06-10T12:20:55
	PConnVbrExpireTime *string `json:"PConnVbrExpireTime,omitempty" xml:"PConnVbrExpireTime,omitempty"`
	// The IPv4 address of the gateway device on the user side.
	//
	// example:
	//
	// 162.62.XX.XX
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The IPv6 address of the gateway device on the user side.
	//
	// This parameter is required when you create a VBR for the owner of the Express Connect circuit. You can ignore this parameter when you create a VBR for another Alibaba Cloud account.
	//
	// example:
	//
	// 2001:XXXX:3c4d:0015:0000:0000:0000:1a2b
	PeerIpv6GatewayIp *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	// The subnet mask for the IPv6 addresses of the gateway devices on the Alibaba Cloud side and on the user side.
	//
	// The two IPv6 addresses must fall within the same subnet.
	//
	// example:
	//
	// 2408:4004:cc:400::/56
	PeeringIpv6SubnetMask *string `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	// The subnet mask of the IPv4 addresses configured on the user side and Alibaba Cloud side.
	//
	// The two IPv4 addresses must fall within the same subnet.
	//
	// example:
	//
	// 255.255.255.0
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// The last time when the status of the VBR changed from Terminated to Active.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	RecoveryTime *string `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	// The status of the VBR. Valid values:
	//
	// 	- **unconfirmed**
	//
	// 	- **active**
	//
	// 	- **terminating**
	//
	// 	- **terminated**
	//
	// 	- **recovering**
	//
	// 	- **deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last time when the VBR was disabled.
	//
	// example:
	//
	// 2021-06-07T12:20:55
	TerminationTime *string `json:"TerminationTime,omitempty" xml:"TerminationTime,omitempty"`
	// The VBR type.
	//
	// example:
	//
	// pconnVBR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The VBR ID.
	//
	// example:
	//
	// vbr-bp16ksp61j7e0tk****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// If the owner of the VBR is the same as that of the Express Connect circuit, this parameter is empty.
	//
	// example:
	//
	// 253460731706911258
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The VLAN ID of the VBR.
	//
	// example:
	//
	// 1678
	VlanId *int32 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetActivationTime() *string {
	return s.ActivationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetBandwidthStatus() *string {
	return s.BandwidthStatus
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetEccId() *string {
	return s.EccId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPConnVbrBussinessStatus() *string {
	return s.PConnVbrBussinessStatus
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPConnVbrChargeType() *string {
	return s.PConnVbrChargeType
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPConnVbrExpireTime() *string {
	return s.PConnVbrExpireTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetRecoveryTime() *string {
	return s.RecoveryTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetStatus() *string {
	return s.Status
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetTerminationTime() *string {
	return s.TerminationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetType() *string {
	return s.Type
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVbrOwnerUid() *int64 {
	return s.VbrOwnerUid
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVlanId() *int32 {
	return s.VlanId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetActivationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.ActivationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetBandwidth(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetBandwidthStatus(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.BandwidthStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetCircuitCode(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.CircuitCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetCreationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetEccId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.EccId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetEnableIpv6(v bool) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetLocalGatewayIp(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.LocalGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetLocalIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPConnVbrBussinessStatus(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PConnVbrBussinessStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPConnVbrChargeType(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PConnVbrChargeType = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPConnVbrExpireTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PConnVbrExpireTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPeerGatewayIp(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PeerGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPeerIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPeeringIpv6SubnetMask(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetPeeringSubnetMask(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.PeeringSubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetRecoveryTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.RecoveryTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetStatus(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.Status = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetTerminationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.TerminationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetType(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.Type = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVbrId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VbrId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVbrOwnerUid(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VbrOwnerUid = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVlanId(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VlanId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) Validate() error {
	return dara.Validate(s)
}
