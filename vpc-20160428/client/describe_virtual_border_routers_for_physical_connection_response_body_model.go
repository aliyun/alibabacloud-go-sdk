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
	TotalCount                                  *int32                                                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	if s.VirtualBorderRouterForPhysicalConnectionSet != nil {
		if err := s.VirtualBorderRouterForPhysicalConnectionSet.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.VirtualBorderRouterForPhysicalConnectionType != nil {
		for _, item := range s.VirtualBorderRouterForPhysicalConnectionType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType struct {
	ActivationTime          *string `json:"ActivationTime,omitempty" xml:"ActivationTime,omitempty"`
	Bandwidth               *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthStatus         *string `json:"BandwidthStatus,omitempty" xml:"BandwidthStatus,omitempty"`
	CircuitCode             *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EccId                   *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	EnableIpv6              *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	LocalGatewayIp          *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	LocalIpv6GatewayIp      *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	PConnVbrBussinessStatus *string `json:"PConnVbrBussinessStatus,omitempty" xml:"PConnVbrBussinessStatus,omitempty"`
	PConnVbrChargeType      *string `json:"PConnVbrChargeType,omitempty" xml:"PConnVbrChargeType,omitempty"`
	PConnVbrExpireTime      *string `json:"PConnVbrExpireTime,omitempty" xml:"PConnVbrExpireTime,omitempty"`
	PeerGatewayIp           *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	PeerIpv6GatewayIp       *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	PeeringIpv6SubnetMask   *string `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	PeeringSubnetMask       *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	RecoveryTime            *string `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminationTime         *string `json:"TerminationTime,omitempty" xml:"TerminationTime,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VbrId                   *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VbrName                 *string `json:"VbrName,omitempty" xml:"VbrName,omitempty"`
	VbrOwnerUid             *int64  `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	VlanId                  *int32  `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
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

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVbrName() *string {
	return s.VbrName
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

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVbrName(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VbrName = &v
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
