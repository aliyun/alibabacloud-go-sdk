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
	// This parameter is required.
	//
	// example:
	//
	// 1
	AreaId *int32 `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MD5
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40
	DeadTime *int32 `json:"DeadTime,omitempty" xml:"DeadTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	HelloTime     *int32  `json:"HelloTime,omitempty" xml:"HelloTime,omitempty"`
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	// example:
	//
	// 5
	Md5Key *string `json:"Md5Key,omitempty" xml:"Md5Key,omitempty"`
	// example:
	//
	// 7
	Md5KeyId        *int32  `json:"Md5KeyId,omitempty" xml:"Md5KeyId,omitempty"`
	Networks        *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
	OspfCost        *int32  `json:"OspfCost,omitempty" xml:"OspfCost,omitempty"`
	OspfNetworkType *string `json:"OspfNetworkType,omitempty" xml:"OspfNetworkType,omitempty"`
	// example:
	//
	// duuf****
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RedistributeProtocol *string `json:"RedistributeProtocol,omitempty" xml:"RedistributeProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sag-3manef62evrfr6****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
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
