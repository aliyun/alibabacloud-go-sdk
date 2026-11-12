// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayOspfRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v int32) *UpdateSmartAccessGatewayOspfRouteRequest
	GetAreaId() *int32
	SetAuthenticationType(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetAuthenticationType() *string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayOspfRouteRequest
	GetCrossAccount() *bool
	SetDeadTime(v int32) *UpdateSmartAccessGatewayOspfRouteRequest
	GetDeadTime() *int32
	SetHelloTime(v int32) *UpdateSmartAccessGatewayOspfRouteRequest
	GetHelloTime() *int32
	SetInterfaceName(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetInterfaceName() *string
	SetMd5Key(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetMd5Key() *string
	SetMd5KeyId(v int32) *UpdateSmartAccessGatewayOspfRouteRequest
	GetMd5KeyId() *int32
	SetNetworks(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetNetworks() *string
	SetOspfCost(v int32) *UpdateSmartAccessGatewayOspfRouteRequest
	GetOspfCost() *int32
	SetOspfNetworkType(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetOspfNetworkType() *string
	SetPassword(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetPassword() *string
	SetRedistributeProtocol(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetRedistributeProtocol() *string
	SetRegionId(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetResourceUid() *string
	SetRouterId(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetRouterId() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayOspfRouteRequest
	GetSagSn() *string
}

type UpdateSmartAccessGatewayOspfRouteRequest struct {
	// The OSPF area ID.
	//
	// Valid values: **1*	- to **2147483647**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AreaId *int32 `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The authentication type. Valid values:
	//
	// - **NONE**: disables authentication.
	//
	// - **CLEARTEXT**: uses plaintext authentication.
	//
	// - **MD5**: uses MD5 authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// Specifies whether to query only cross-account Smart Access Gateway instances. Valid values:
	//
	// - **false*	- (default): no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The OSPF dead interval. Unit: seconds.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40
	DeadTime *int32 `json:"DeadTime,omitempty" xml:"DeadTime,omitempty"`
	// The interval at which Hello packets are sent. Unit: seconds.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	HelloTime *int32 `json:"HelloTime,omitempty" xml:"HelloTime,omitempty"`
	// The port name.
	//
	// > This parameter does not take effect.
	//
	// example:
	//
	// 2
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	// The value of the MD5 key.
	//
	// Valid values: **1*	- to **47**.
	//
	// > This parameter is required only when **AuthenticationType*	- is set to **MD5**.
	//
	// example:
	//
	// 5
	Md5Key *string `json:"Md5Key,omitempty" xml:"Md5Key,omitempty"`
	// The ID of the MD5 key.
	//
	// Valid values: **1*	- to **2147483647**.
	//
	// > This parameter is required only when **AuthenticationType*	- is set to **MD5**.
	//
	// example:
	//
	// 7
	Md5KeyId *int32 `json:"Md5KeyId,omitempty" xml:"Md5KeyId,omitempty"`
	// The CIDR block.
	//
	// > This parameter does not take effect.
	//
	// example:
	//
	// 192.168.10.0/24
	Networks *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
	// The OSPF route cost.
	//
	// > This parameter does not take effect.
	//
	// example:
	//
	// 11
	OspfCost *int32 `json:"OspfCost,omitempty" xml:"OspfCost,omitempty"`
	// The OSPF area type.
	//
	// > This parameter does not take effect.
	//
	// example:
	//
	// NSSA
	OspfNetworkType *string `json:"OspfNetworkType,omitempty" xml:"OspfNetworkType,omitempty"`
	// The password for plaintext authentication.
	//
	// The password must be 1 to 8 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// > This parameter is required only when **AuthenticationType*	- is set to **CLEARTEXT**.
	//
	// example:
	//
	// duuf****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The route redistribution protocol.
	//
	// > This parameter does not take effect.
	//
	// example:
	//
	// OSPF
	RedistributeProtocol *string `json:"RedistributeProtocol,omitempty" xml:"RedistributeProtocol,omitempty"`
	// The region ID of the Smart Access Gateway instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to obtain region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The OSPF router ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The Smart Access Gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-3manef62evrfr6****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the Smart Access Gateway device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dk****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s UpdateSmartAccessGatewayOspfRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayOspfRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetAreaId() *int32 {
	return s.AreaId
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetDeadTime() *int32 {
	return s.DeadTime
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetHelloTime() *int32 {
	return s.HelloTime
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetInterfaceName() *string {
	return s.InterfaceName
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetMd5Key() *string {
	return s.Md5Key
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetMd5KeyId() *int32 {
	return s.Md5KeyId
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetNetworks() *string {
	return s.Networks
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetOspfCost() *int32 {
	return s.OspfCost
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetOspfNetworkType() *string {
	return s.OspfNetworkType
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetRedistributeProtocol() *string {
	return s.RedistributeProtocol
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetAreaId(v int32) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.AreaId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetAuthenticationType(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.AuthenticationType = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetDeadTime(v int32) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.DeadTime = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetHelloTime(v int32) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.HelloTime = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetInterfaceName(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.InterfaceName = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetMd5Key(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.Md5Key = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetMd5KeyId(v int32) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.Md5KeyId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetNetworks(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.Networks = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetOspfCost(v int32) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.OspfCost = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetOspfNetworkType(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.OspfNetworkType = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetPassword(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.Password = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetRedistributeProtocol(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.RedistributeProtocol = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetRegionId(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetRouterId(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.RouterId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) SetSagSn(v string) *UpdateSmartAccessGatewayOspfRouteRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteRequest) Validate() error {
	return dara.Validate(s)
}
