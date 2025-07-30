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
	// 	- **true*	- (default).
	//
	// 	- **false**. If automatic payment is disabled, you must perform the following steps to complete the payment in the Tair (Redis OSS-compatible) console: In the top navigation bar, choose **Expenses*	- > **Renewal Management**. In the left-side navigation pane, click **Orders**. On the **Orders*	- page, find the order and complete the payment.
	//
	// >  This parameter is valid only if the value of the **ChargeType*	- field in the **Instances*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  This parameter is available only if **ChargeType*	- in the **Instances*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The database engine version of the instance. Valid values: **4.0*	- and **5.0**.
	//
	// >  The default value is **5.0**.
	//
	// Valid values:
	//
	// 	- 2.8
	//
	// 	- 4.0
	//
	// 	- 5.0
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The JSON-formatted configurations of the instance. For more information, see the "Additional description of the Instances parameter" section.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{	"RegionId": "cn-hangzhou",	"izNo": "cn-hangzhou-b",	"quantity": 2,	"instanceType": "Redis",	"instanceClass": "redis.master.small.default",	"EngineVersion": "5.0",	"ChargeType": "PostPaid"}]
	Instances    *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to restore the source instance from the recycle bin. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter is valid only if the **SrcDBInstanceId*	- field in the **Instances*	- parameter is specified.
	//
	// example:
	//
	// false
	RebuildInstance *bool `json:"RebuildInstance,omitempty" xml:"RebuildInstance,omitempty"`
	// The ID of the resource group to which to assign the instance.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token is case-sensitive. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
