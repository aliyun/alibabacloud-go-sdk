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
	SetAutoRenew(v bool) *CreateRouterInterfaceRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateRouterInterfaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouterInterfaceRequest
	GetDescription() *string
	SetFastLinkMode(v bool) *CreateRouterInterfaceRequest
	GetFastLinkMode() *bool
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
	SetResourceGroupId(v string) *CreateRouterInterfaceRequest
	GetResourceGroupId() *string
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
	SetTags(v []*CreateRouterInterfaceRequestTags) *CreateRouterInterfaceRequest
	GetTags() []*CreateRouterInterfaceRequestTags
}

type CreateRouterInterfaceRequest struct {
	// The ID of the access point to which the VBR belongs.
	//
	// You can call the [DescribeAccessPoints](https://help.aliyun.com/document_detail/36062.html) operation to obtain the IDs of access points.
	//
	// >  This parameter is required if the VBR is connected to an Express Connect circuit.
	//
	// example:
	//
	// ap-cn-hangzhou-yh-ts-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **false*	- (default): The automatic payment is disabled. If you select this option, you must go to the Order Center to complete the payment after an order is generated.
	//
	// 	- **true**: The automatic payment is enabled. Payments are automatically complete after an order is generated.
	//
	// >  This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the router interface.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abcabc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the VBR that is created in the Fast Link mode is uplinked to the router interface. The Fast Link mode helps automatically connect router interfaces that are created for the VBR and its peer VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >
	//
	// 	- This parameter takes effect only if **RouterType*	- is set to **VBR*	- and **OppositeRouterType*	- is set to **VRouter**.
	//
	// 	- If **FastLinkMode*	- is set to **true**, **Role*	- must be set to **InitiatingSide**. In this case, **AccessPointId**, **OppositeRouterType**, **OpppsiteRouterId**, and **OppositeInterfaceOwnerId*	- are required.
	//
	// example:
	//
	// false
	FastLinkMode *bool `json:"FastLinkMode,omitempty" xml:"FastLinkMode,omitempty"`
	// The source IP address that is used to perform health checks. The source IP address must be an idle IP address of the local virtual private cloud (VPC).
	//
	// >  You can set this parameter when an Express Connect circuit is used.
	//
	// example:
	//
	// 192.168.0.6
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The destination IP address that is used to perform health checks.
	//
	// >  This parameter is required if you specify **HealthCheckSourceIp**
	//
	// example:
	//
	// 192.168.0.8
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// The billing method of the router interface. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the router interface.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the access point to which the peer belongs.
	//
	// >  This parameter is required if the peer router interface is associated with a VBR. The specified value cannot be changed after the router interface is created.
	//
	// example:
	//
	// ap-cn-shanghai-nt-aligroup-C
	OppositeAccessPointId *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	// The ID of the peer router interface.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzd****
	OppositeInterfaceId *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the peer router interface belongs.
	//
	// example:
	//
	// 253460731706911258
	OppositeInterfaceOwnerId *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// The ID of the region in which the acceptor is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	// The ID of the peer router.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	OppositeRouterId *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	// The type of router that is associated with the peer router interface. Valid values:
	//
	// 	- **VRouter**
	//
	// 	- **VBR**
	//
	// example:
	//
	// VRouter
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- Valid values when PricingCycle is set to Month: **1 to 9**.
	//
	// 	- Valid values when PricingCycle is set to Year: **1 to 3**.
	//
	// >  This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// example:
	//
	// 3
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the subscription. Valid values:
	//
	// 	- **Month*	- (default)
	//
	// 	- **Year**
	//
	// >  This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the region to which the router interface belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource group, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the router interface. Valid values:
	//
	// 	- **InitiatingSide**: requester
	//
	// 	- **AcceptingSide**: acceptor
	//
	// This parameter is required.
	//
	// example:
	//
	// InitiatingSide
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the router that is associated with the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-m5ebm6g9ptc9mly1c****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The type of router that is associated with the router interface. Valid values:
	//
	// 	- **VRouter**
	//
	// 	- **VBR**
	//
	// This parameter is required.
	//
	// example:
	//
	// VRouter
	RouterType *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	// The specification of the router interface and the corresponding bandwidth. Valid values:
	//
	// 	- **Mini.2**: 2 Mbit/s
	//
	// 	- **Mini.5**: 5 Mbit/s
	//
	// 	- **Small.1**: 10 Mbit/s
	//
	// 	- **Small.2**: 20 Mbit/s
	//
	// 	- **Small.5**: 50 Mbit/s
	//
	// 	- **Middle.1**: 100 Mbit/s
	//
	// 	- **Middle.2**: 200 Mbit/s
	//
	// 	- **Middle.5**: 500 Mbit/s
	//
	// 	- **Large.1**: 1,000 Mbit/s
	//
	// 	- **Large.2**: 2,000 Mbit/s
	//
	// 	- **Large.5**: 5,000 Mbit/s
	//
	// 	- **Xlarge.1**: 10,000 Mbit/s
	//
	// >  If **Role*	- is set to **AcceptingSide**, set **Spec*	- to **Negative**. This indicates that you do not need to specify the specification when you create an acceptor router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mini.2
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tag to add to the resource.
	Tags []*CreateRouterInterfaceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *CreateRouterInterfaceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRouterInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouterInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouterInterfaceRequest) GetFastLinkMode() *bool {
	return s.FastLinkMode
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

func (s *CreateRouterInterfaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *CreateRouterInterfaceRequest) GetTags() []*CreateRouterInterfaceRequestTags {
	return s.Tags
}

func (s *CreateRouterInterfaceRequest) SetAccessPointId(v string) *CreateRouterInterfaceRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetAutoPay(v bool) *CreateRouterInterfaceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetAutoRenew(v bool) *CreateRouterInterfaceRequest {
	s.AutoRenew = &v
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

func (s *CreateRouterInterfaceRequest) SetFastLinkMode(v bool) *CreateRouterInterfaceRequest {
	s.FastLinkMode = &v
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

func (s *CreateRouterInterfaceRequest) SetResourceGroupId(v string) *CreateRouterInterfaceRequest {
	s.ResourceGroupId = &v
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

func (s *CreateRouterInterfaceRequest) SetTags(v []*CreateRouterInterfaceRequestTags) *CreateRouterInterfaceRequest {
	s.Tags = v
	return s
}

func (s *CreateRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRouterInterfaceRequestTags struct {
	// The tag key to add to the resource. You must enter at least one tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRouterInterfaceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateRouterInterfaceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateRouterInterfaceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateRouterInterfaceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateRouterInterfaceRequestTags) SetKey(v string) *CreateRouterInterfaceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateRouterInterfaceRequestTags) SetValue(v string) *CreateRouterInterfaceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateRouterInterfaceRequestTags) Validate() error {
	return dara.Validate(s)
}
