// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreateRouterInterfaceRequest
	GetAccessPointId() *string
	SetAutoPay(v bool) *CreateRouterInterfaceRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreateRouterInterfaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouterInterfaceRequest
	GetDescription() *string
	SetHealthCheckSourceIp(v string) *CreateRouterInterfaceRequest
	GetHealthCheckSourceIp() *string
	SetHealthCheckTargetIp(v string) *CreateRouterInterfaceRequest
	GetHealthCheckTargetIp() *string
	SetInstanceChargeType(v string) *CreateRouterInterfaceRequest
	GetInstanceChargeType() *string
	SetName(v string) *CreateRouterInterfaceRequest
	GetName() *string
	SetOppositeAccessPointId(v string) *CreateRouterInterfaceRequest
	GetOppositeAccessPointId() *string
	SetOppositeInterfaceId(v string) *CreateRouterInterfaceRequest
	GetOppositeInterfaceId() *string
	SetOppositeInterfaceOwnerId(v string) *CreateRouterInterfaceRequest
	GetOppositeInterfaceOwnerId() *string
	SetOppositeRegionId(v string) *CreateRouterInterfaceRequest
	GetOppositeRegionId() *string
	SetOppositeRouterId(v string) *CreateRouterInterfaceRequest
	GetOppositeRouterId() *string
	SetOppositeRouterType(v string) *CreateRouterInterfaceRequest
	GetOppositeRouterType() *string
	SetOwnerAccount(v string) *CreateRouterInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRouterInterfaceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateRouterInterfaceRequest
	GetPeriod() *int32
	SetPricingCycle(v string) *CreateRouterInterfaceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *CreateRouterInterfaceRequest
	GetRole() *string
	SetRouterId(v string) *CreateRouterInterfaceRequest
	GetRouterId() *string
	SetRouterType(v string) *CreateRouterInterfaceRequest
	GetRouterType() *string
	SetSpec(v string) *CreateRouterInterfaceRequest
	GetSpec() *string
	SetUserCidr(v string) *CreateRouterInterfaceRequest
	GetUserCidr() *string
}

type CreateRouterInterfaceRequest struct {
	// The access point ID.
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values are `true` and `false`. The default value is `true`.
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the router interface. The description must be 2 to 256 characters long, must start with a letter, and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source IP address that is used for the health check.
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The destination IP address that is used for the health check.
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// The billing method of the instance. Set the value to `PrePaid`. This parameter is required if you also specify `PricingCycle`.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the router interface. The name must be 2 to 128 characters long and start with a letter. It can contain letters, digits, underscores (_), and hyphens (-).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the peer access point.
	OppositeAccessPointId *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	// The ID of the peer router interface.
	OppositeInterfaceId *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	// The ID of the account to which the peer router interface belongs.
	OppositeInterfaceOwnerId *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// The ID of the peer region.
	//
	// This parameter is required.
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	// The ID of the peer router. This parameter is available only when the local and peer router interfaces belong to the same account.
	OppositeRouterId *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	// The type of the peer router. Valid values:
	//
	// - **VRouter**
	//
	// - **VBR**
	//
	// Default value: **VRouter**.
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. This parameter is required when `InstanceChargeType` is set to `PrePaid` and `PricingCycle` is set to `Month` or `Year`. Valid values:
	//
	// - If `PricingCycle` is set to `Month`, the valid values are 1 to 9.
	//
	// - If `PricingCycle` is set to `Year`, the valid values are 1 to 3.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle. This parameter is required if `InstanceChargeType` is set to `PrePaid`. Valid values are `Month` and `Year`.
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID of the router interface.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the router interface in the peering connection. Valid values:
	//
	// - **InitiatingSide**: The router interface is the initiator.
	//
	// - **AcceptingSide**: The router interface is the acceptor.
	//
	// This parameter is required.
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The router ID.
	//
	// This parameter is required.
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The router type. Valid values:
	//
	// - **VRouter**
	//
	// - **VBR**
	//
	// This parameter is required.
	RouterType *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	// The specification of the router interface. Valid values:
	//
	// - **Mini.2**
	//
	// - **Mini.5**
	//
	// - **Small.1**
	//
	// - **Small.2**
	//
	// - **Small.5**
	//
	// - **Middle.1**
	//
	// - **Middle.2**
	//
	// - **Middle.5**
	//
	// - **Large.1**
	//
	// - **Large.2**
	//
	// - **Large.5**
	//
	// - **Xlarge.1**
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The CIDR block of the user. This parameter is required when you create a router interface for a virtual border router (VBR) that is in the same region as the Express Connect circuit, or when both `RouterType` and `OppositeRouterType` are set to `VBR`.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s CreateRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateRouterInterfaceRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreateRouterInterfaceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateRouterInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouterInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouterInterfaceRequest) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *CreateRouterInterfaceRequest) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *CreateRouterInterfaceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRouterInterfaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouterInterfaceRequest) GetOppositeAccessPointId() *string {
	return s.OppositeAccessPointId
}

func (s *CreateRouterInterfaceRequest) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *CreateRouterInterfaceRequest) GetOppositeInterfaceOwnerId() *string {
	return s.OppositeInterfaceOwnerId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *CreateRouterInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRouterInterfaceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateRouterInterfaceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRouterInterfaceRequest) GetRole() *string {
	return s.Role
}

func (s *CreateRouterInterfaceRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *CreateRouterInterfaceRequest) GetRouterType() *string {
	return s.RouterType
}

func (s *CreateRouterInterfaceRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateRouterInterfaceRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreateRouterInterfaceRequest) SetAccessPointId(v string) *CreateRouterInterfaceRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetAutoPay(v bool) *CreateRouterInterfaceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetClientToken(v string) *CreateRouterInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetDescription(v string) *CreateRouterInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetHealthCheckSourceIp(v string) *CreateRouterInterfaceRequest {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetHealthCheckTargetIp(v string) *CreateRouterInterfaceRequest {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetInstanceChargeType(v string) *CreateRouterInterfaceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetName(v string) *CreateRouterInterfaceRequest {
	s.Name = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeAccessPointId(v string) *CreateRouterInterfaceRequest {
	s.OppositeAccessPointId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeInterfaceId(v string) *CreateRouterInterfaceRequest {
	s.OppositeInterfaceId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeInterfaceOwnerId(v string) *CreateRouterInterfaceRequest {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRegionId(v string) *CreateRouterInterfaceRequest {
	s.OppositeRegionId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRouterId(v string) *CreateRouterInterfaceRequest {
	s.OppositeRouterId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRouterType(v string) *CreateRouterInterfaceRequest {
	s.OppositeRouterType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOwnerAccount(v string) *CreateRouterInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOwnerId(v int64) *CreateRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetPeriod(v int32) *CreateRouterInterfaceRequest {
	s.Period = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetPricingCycle(v string) *CreateRouterInterfaceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRegionId(v string) *CreateRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetResourceOwnerAccount(v string) *CreateRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetResourceOwnerId(v int64) *CreateRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRole(v string) *CreateRouterInterfaceRequest {
	s.Role = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRouterId(v string) *CreateRouterInterfaceRequest {
	s.RouterId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRouterType(v string) *CreateRouterInterfaceRequest {
	s.RouterType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetSpec(v string) *CreateRouterInterfaceRequest {
	s.Spec = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetUserCidr(v string) *CreateRouterInterfaceRequest {
	s.UserCidr = &v
	return s
}

func (s *CreateRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
