// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVirtualBorderRoutersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetTotalCount() *int32
	SetVirtualBorderRouterSet(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) *DescribeVirtualBorderRoutersResponseBody
	GetVirtualBorderRouterSet() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet
}

type DescribeVirtualBorderRoutersResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount             *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VirtualBorderRouterSet *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet `json:"VirtualBorderRouterSet,omitempty" xml:"VirtualBorderRouterSet,omitempty" type:"Struct"`
}

func (s DescribeVirtualBorderRoutersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetVirtualBorderRouterSet() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet {
	return s.VirtualBorderRouterSet
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetPageNumber(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetPageSize(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetRequestId(v string) *DescribeVirtualBorderRoutersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetTotalCount(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetVirtualBorderRouterSet(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) *DescribeVirtualBorderRoutersResponseBody {
	s.VirtualBorderRouterSet = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) Validate() error {
	if s.VirtualBorderRouterSet != nil {
		if err := s.VirtualBorderRouterSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet struct {
	VirtualBorderRouterType []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType `json:"VirtualBorderRouterType,omitempty" xml:"VirtualBorderRouterType,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) GetVirtualBorderRouterType() []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	return s.VirtualBorderRouterType
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) SetVirtualBorderRouterType(v []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet {
	s.VirtualBorderRouterType = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) Validate() error {
	if s.VirtualBorderRouterType != nil {
		for _, item := range s.VirtualBorderRouterType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType struct {
	AccessPointId                    *string                                                                                                             `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	ActivationTime                   *string                                                                                                             `json:"ActivationTime,omitempty" xml:"ActivationTime,omitempty"`
	AssociatedCens                   *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens                `json:"AssociatedCens,omitempty" xml:"AssociatedCens,omitempty" type:"Struct"`
	AssociatedPhysicalConnections    *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections `json:"AssociatedPhysicalConnections,omitempty" xml:"AssociatedPhysicalConnections,omitempty" type:"Struct"`
	Bandwidth                        *int32                                                                                                              `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CircuitCode                      *string                                                                                                             `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CloudBoxInstanceId               *string                                                                                                             `json:"CloudBoxInstanceId,omitempty" xml:"CloudBoxInstanceId,omitempty"`
	CreationTime                     *string                                                                                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                      *string                                                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DetectMultiplier                 *int64                                                                                                              `json:"DetectMultiplier,omitempty" xml:"DetectMultiplier,omitempty"`
	EccId                            *string                                                                                                             `json:"EccId,omitempty" xml:"EccId,omitempty"`
	EcrAttatchStatus                 *string                                                                                                             `json:"EcrAttatchStatus,omitempty" xml:"EcrAttatchStatus,omitempty"`
	EcrId                            *string                                                                                                             `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	EcrOwnerId                       *string                                                                                                             `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	EnableIpv6                       *bool                                                                                                               `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	LocalGatewayIp                   *string                                                                                                             `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	LocalIpv6GatewayIp               *string                                                                                                             `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	MinRxInterval                    *int64                                                                                                              `json:"MinRxInterval,omitempty" xml:"MinRxInterval,omitempty"`
	MinTxInterval                    *int64                                                                                                              `json:"MinTxInterval,omitempty" xml:"MinTxInterval,omitempty"`
	Mtu                              *int32                                                                                                              `json:"Mtu,omitempty" xml:"Mtu,omitempty"`
	Name                             *string                                                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PConnVbrChargeType               *string                                                                                                             `json:"PConnVbrChargeType,omitempty" xml:"PConnVbrChargeType,omitempty"`
	PConnVbrExpireTime               *string                                                                                                             `json:"PConnVbrExpireTime,omitempty" xml:"PConnVbrExpireTime,omitempty"`
	PeerGatewayIp                    *string                                                                                                             `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	PeerIpv6GatewayIp                *string                                                                                                             `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	PeeringIpv6SubnetMask            *string                                                                                                             `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	PeeringSubnetMask                *string                                                                                                             `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	PhysicalConnectionBusinessStatus *string                                                                                                             `json:"PhysicalConnectionBusinessStatus,omitempty" xml:"PhysicalConnectionBusinessStatus,omitempty"`
	PhysicalConnectionId             *string                                                                                                             `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	PhysicalConnectionOwnerUid       *string                                                                                                             `json:"PhysicalConnectionOwnerUid,omitempty" xml:"PhysicalConnectionOwnerUid,omitempty"`
	PhysicalConnectionStatus         *string                                                                                                             `json:"PhysicalConnectionStatus,omitempty" xml:"PhysicalConnectionStatus,omitempty"`
	RecoveryTime                     *string                                                                                                             `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	ResourceGroupId                  *string                                                                                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RouteTableId                     *string                                                                                                             `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	SitelinkEnable                   *bool                                                                                                               `json:"SitelinkEnable,omitempty" xml:"SitelinkEnable,omitempty"`
	Status                           *string                                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                             *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags                          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TerminationTime                  *string                                                                                                             `json:"TerminationTime,omitempty" xml:"TerminationTime,omitempty"`
	Type                             *string                                                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	VbrId                            *string                                                                                                             `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VlanId                           *int32                                                                                                              `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	VlanInterfaceId                  *string                                                                                                             `json:"VlanInterfaceId,omitempty" xml:"VlanInterfaceId,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetActivationTime() *string {
	return s.ActivationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetAssociatedCens() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens {
	return s.AssociatedCens
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetAssociatedPhysicalConnections() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections {
	return s.AssociatedPhysicalConnections
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetCloudBoxInstanceId() *string {
	return s.CloudBoxInstanceId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetDescription() *string {
	return s.Description
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetDetectMultiplier() *int64 {
	return s.DetectMultiplier
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetEccId() *string {
	return s.EccId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetEcrAttatchStatus() *string {
	return s.EcrAttatchStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetEcrOwnerId() *string {
	return s.EcrOwnerId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetMinRxInterval() *int64 {
	return s.MinRxInterval
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetMinTxInterval() *int64 {
	return s.MinTxInterval
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetMtu() *int32 {
	return s.Mtu
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetName() *string {
	return s.Name
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPConnVbrChargeType() *string {
	return s.PConnVbrChargeType
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPConnVbrExpireTime() *string {
	return s.PConnVbrExpireTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionBusinessStatus() *string {
	return s.PhysicalConnectionBusinessStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionOwnerUid() *string {
	return s.PhysicalConnectionOwnerUid
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionStatus() *string {
	return s.PhysicalConnectionStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetRecoveryTime() *string {
	return s.RecoveryTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetSitelinkEnable() *bool {
	return s.SitelinkEnable
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetStatus() *string {
	return s.Status
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetTags() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags {
	return s.Tags
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetTerminationTime() *string {
	return s.TerminationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetType() *string {
	return s.Type
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVlanId() *int32 {
	return s.VlanId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVlanInterfaceId() *string {
	return s.VlanInterfaceId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetAccessPointId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.AccessPointId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetActivationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.ActivationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetAssociatedCens(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.AssociatedCens = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetAssociatedPhysicalConnections(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.AssociatedPhysicalConnections = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetBandwidth(v int32) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetCircuitCode(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.CircuitCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetCloudBoxInstanceId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.CloudBoxInstanceId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetCreationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetDescription(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Description = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetDetectMultiplier(v int64) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.DetectMultiplier = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetEccId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.EccId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetEcrAttatchStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.EcrAttatchStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetEcrId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.EcrId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetEcrOwnerId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.EcrOwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetEnableIpv6(v bool) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetLocalGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.LocalGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetLocalIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetMinRxInterval(v int64) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.MinRxInterval = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetMinTxInterval(v int64) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.MinTxInterval = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetMtu(v int32) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Mtu = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetName(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Name = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPConnVbrChargeType(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PConnVbrChargeType = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPConnVbrExpireTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PConnVbrExpireTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeerGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeerGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeerIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeeringIpv6SubnetMask(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeeringSubnetMask(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeeringSubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionBusinessStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionBusinessStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionOwnerUid(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionOwnerUid = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetRecoveryTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.RecoveryTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetResourceGroupId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetRouteTableId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetSitelinkEnable(v bool) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.SitelinkEnable = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Status = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetTags(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Tags = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetTerminationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.TerminationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetType(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Type = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVbrId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VbrId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVlanId(v int32) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VlanId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVlanInterfaceId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VlanInterfaceId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) Validate() error {
	if s.AssociatedCens != nil {
		if err := s.AssociatedCens.Validate(); err != nil {
			return err
		}
	}
	if s.AssociatedPhysicalConnections != nil {
		if err := s.AssociatedPhysicalConnections.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens struct {
	AssociatedCen []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen `json:"AssociatedCen,omitempty" xml:"AssociatedCen,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) GetAssociatedCen() []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen {
	return s.AssociatedCen
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) SetAssociatedCen(v []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens {
	s.AssociatedCen = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCens) Validate() error {
	if s.AssociatedCen != nil {
		for _, item := range s.AssociatedCen {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen struct {
	CenId      *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	CenStatus  *string `json:"CenStatus,omitempty" xml:"CenStatus,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) GetCenStatus() *string {
	return s.CenStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) SetCenId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen {
	s.CenId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) SetCenOwnerId(v int64) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen {
	s.CenOwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) SetCenStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen {
	s.CenStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedCensAssociatedCen) Validate() error {
	return dara.Validate(s)
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections struct {
	AssociatedPhysicalConnection []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection `json:"AssociatedPhysicalConnection,omitempty" xml:"AssociatedPhysicalConnection,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) GetAssociatedPhysicalConnection() []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	return s.AssociatedPhysicalConnection
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) SetAssociatedPhysicalConnection(v []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections {
	s.AssociatedPhysicalConnection = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnections) Validate() error {
	if s.AssociatedPhysicalConnection != nil {
		for _, item := range s.AssociatedPhysicalConnection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection struct {
	CircuitCode                      *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	EnableIpv6                       *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	LocalGatewayIp                   *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	LocalIpv6GatewayIp               *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	PeerGatewayIp                    *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	PeerIpv6GatewayIp                *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	PeeringIpv6SubnetMask            *string `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	PeeringSubnetMask                *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	PhysicalConnectionBusinessStatus *string `json:"PhysicalConnectionBusinessStatus,omitempty" xml:"PhysicalConnectionBusinessStatus,omitempty"`
	PhysicalConnectionId             *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	PhysicalConnectionOwnerUid       *string `json:"PhysicalConnectionOwnerUid,omitempty" xml:"PhysicalConnectionOwnerUid,omitempty"`
	PhysicalConnectionStatus         *string `json:"PhysicalConnectionStatus,omitempty" xml:"PhysicalConnectionStatus,omitempty"`
	Status                           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VlanId                           *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	VlanInterfaceId                  *string `json:"VlanInterfaceId,omitempty" xml:"VlanInterfaceId,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPhysicalConnectionBusinessStatus() *string {
	return s.PhysicalConnectionBusinessStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPhysicalConnectionOwnerUid() *string {
	return s.PhysicalConnectionOwnerUid
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetPhysicalConnectionStatus() *string {
	return s.PhysicalConnectionStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetStatus() *string {
	return s.Status
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetVlanId() *string {
	return s.VlanId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) GetVlanInterfaceId() *string {
	return s.VlanInterfaceId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetCircuitCode(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.CircuitCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetEnableIpv6(v bool) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetLocalGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.LocalGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetLocalIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPeerGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PeerGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPeerIpv6GatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPeeringIpv6SubnetMask(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPeeringSubnetMask(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PeeringSubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPhysicalConnectionBusinessStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PhysicalConnectionBusinessStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPhysicalConnectionOwnerUid(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PhysicalConnectionOwnerUid = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetPhysicalConnectionStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.PhysicalConnectionStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.Status = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetVlanId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.VlanId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) SetVlanInterfaceId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection {
	s.VlanInterfaceId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeAssociatedPhysicalConnectionsAssociatedPhysicalConnection) Validate() error {
	return dara.Validate(s)
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags struct {
	Tags []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) GetTags() []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags {
	return s.Tags
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) SetTags(v []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags {
	s.Tags = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTags) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) SetKey(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags {
	s.Key = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) SetValue(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags {
	s.Value = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterTypeTagsTags) Validate() error {
	return dara.Validate(s)
}
