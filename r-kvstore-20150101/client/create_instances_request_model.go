// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateInstancesRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateInstancesRequest
	GetAutoRenew() *string
	SetBusinessInfo(v string) *CreateInstancesRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *CreateInstancesRequest
	GetCouponNo() *string
	SetEngineVersion(v string) *CreateInstancesRequest
	GetEngineVersion() *string
	SetInstances(v string) *CreateInstancesRequest
	GetInstances() *string
	SetOwnerAccount(v string) *CreateInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstancesRequest
	GetOwnerId() *int64
	SetRebuildInstance(v bool) *CreateInstancesRequest
	GetRebuildInstance() *bool
	SetResourceGroupId(v string) *CreateInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstancesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateInstancesRequest
	GetSecurityToken() *string
	SetToken(v string) *CreateInstancesRequest
	GetToken() *string
}

type CreateInstancesRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// \\	- **true**: Enables automatic payment. This is the default value.
	//
	// \\	- **false**: Disables automatic payment. You must go to the console to complete the payment. In the top navigation bar, choose **Expenses*	- > **Renewal Management**. In the navigation pane on the left, click **Or*er Management*	- > **My Or*ers**, find the or*er, and then complete the payment.
	//
	// \\> This parameter is valid only when **ChargeType*	- is set to **PrePaid*	- in **Instances**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// \\	- **true**: Enables auto-renewal.
	//
	// \\	- **false**: Disables auto-renewal. This is the default value.
	//
	// \\> This parameter is valid only when **ChargeType*	- is set to **PrePaid*	- in **Instances**.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Additional business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code. The default value is `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The Redis-compatible engine version for the instance. Valid values: **4.0*	- and **5.0**. The default value is **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The configurations of the new instances, specified in JSON format. For more information, see the details of the Instances parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{
	//
	//     "RegionId": "cn-hangzhou",
	//
	//     "izNo": "cn-hangzhou-b",
	//
	//     "quantity": 1,
	//
	//     "instanceType": "Redis",
	//
	//     "instanceClass": "redis.logic.sharding.1g.2db.0rodb.4proxy.default",
	//
	//     "EngineVersion": "5.0",
	//
	//     "ChargeType":"PrePaid",
	//
	//     "Period":"1",
	//
	//     "networkType": "VPC" ,
	//
	//     "vpcId": "vpc-2zex6u1nu32k3ux35oxxx",
	//
	//     "vSwitchId": "vsw-2zesk464e647104kw3xxx"
	//
	// }]
	Instances    *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to rebuild an instance from the recycle bin. Valid values:
	//
	// \\	- **true**: Rebuilds the instance.
	//
	// \\	- **false**: Does not rebuild the instance. This is the default value.
	//
	// \\> This parameter is valid only when **SrcDBInstanceId*	- is specified in **Instances**.
	//
	// example:
	//
	// false
	RebuildInstance *bool `json:"RebuildInstance,omitempty" xml:"RebuildInstance,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// A client-generated token to ensure request idempotence. The value must be unique across requests, case-sensitive, and up to 64 ASCII characters long.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateInstancesRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateInstancesRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateInstancesRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateInstancesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateInstancesRequest) GetInstances() *string {
	return s.Instances
}

func (s *CreateInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstancesRequest) GetRebuildInstance() *bool {
	return s.RebuildInstance
}

func (s *CreateInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateInstancesRequest) GetToken() *string {
	return s.Token
}

func (s *CreateInstancesRequest) SetAutoPay(v bool) *CreateInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenew(v string) *CreateInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstancesRequest) SetBusinessInfo(v string) *CreateInstancesRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateInstancesRequest) SetCouponNo(v string) *CreateInstancesRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateInstancesRequest) SetEngineVersion(v string) *CreateInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateInstancesRequest) SetInstances(v string) *CreateInstancesRequest {
	s.Instances = &v
	return s
}

func (s *CreateInstancesRequest) SetOwnerAccount(v string) *CreateInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstancesRequest) SetOwnerId(v int64) *CreateInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstancesRequest) SetRebuildInstance(v bool) *CreateInstancesRequest {
	s.RebuildInstance = &v
	return s
}

func (s *CreateInstancesRequest) SetResourceGroupId(v string) *CreateInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstancesRequest) SetResourceOwnerAccount(v string) *CreateInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstancesRequest) SetResourceOwnerId(v int64) *CreateInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstancesRequest) SetSecurityToken(v string) *CreateInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateInstancesRequest) SetToken(v string) *CreateInstancesRequest {
	s.Token = &v
	return s
}

func (s *CreateInstancesRequest) Validate() error {
	return dara.Validate(s)
}
