// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShardingNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *AddShardingNodeRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *AddShardingNodeRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *AddShardingNodeRequest
	GetCouponNo() *string
	SetForceTrans(v bool) *AddShardingNodeRequest
	GetForceTrans() *bool
	SetInstanceId(v string) *AddShardingNodeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *AddShardingNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddShardingNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddShardingNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddShardingNodeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *AddShardingNodeRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *AddShardingNodeRequest
	GetShardCount() *int32
	SetSourceBiz(v string) *AddShardingNodeRequest
	GetSourceBiz() *string
	SetVSwitchId(v string) *AddShardingNodeRequest
	GetVSwitchId() *string
}

type AddShardingNodeRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: Automatic payment is enabled. Make sure that your account has sufficient balance.
	//
	// 	- **false**: Automatic payment is disabled. You must manually complete the payment in the console. For more information, see [Manual renewal](https://help.aliyun.com/document_detail/26352.html).
	//
	// > Default value: **true**.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The additional parameter for business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Specifies whether to enable forced transmission. Valid values:
	//
	// - **false*	- (default): Before the Upgrade/Downgrade, the system checks the current minor engine version of the instance. If the minor engine version is outdated, an error is returned. You must upgrade the minor engine version and retry.
	//
	// - **true**: Skips the check and directly executes the Upgrade/Downgrade operation.
	//
	// example:
	//
	// false
	ForceTrans *bool `json:"ForceTrans,omitempty" xml:"ForceTrans,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards to add. Default value: **1**.
	//
	// > The minimum total number of shards for an instance is 2. You can add up to 64 shards in a single operation, and the total number of shards for an instance cannot exceed 256. Make sure that the number of shards does not exceed this limit.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The request source. This parameter is used only for internal maintenance and does not need to be specified.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	// The ID of a different vSwitch within the same VPC. The new data shards are created under the specified vSwitch.
	//
	// If you leave this parameter empty, the new data shards are created under the original vSwitch.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AddShardingNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddShardingNodeRequest) GoString() string {
	return s.String()
}

func (s *AddShardingNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *AddShardingNodeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *AddShardingNodeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *AddShardingNodeRequest) GetForceTrans() *bool {
	return s.ForceTrans
}

func (s *AddShardingNodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddShardingNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddShardingNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddShardingNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddShardingNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddShardingNodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddShardingNodeRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *AddShardingNodeRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *AddShardingNodeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddShardingNodeRequest) SetAutoPay(v bool) *AddShardingNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *AddShardingNodeRequest) SetBusinessInfo(v string) *AddShardingNodeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *AddShardingNodeRequest) SetCouponNo(v string) *AddShardingNodeRequest {
	s.CouponNo = &v
	return s
}

func (s *AddShardingNodeRequest) SetForceTrans(v bool) *AddShardingNodeRequest {
	s.ForceTrans = &v
	return s
}

func (s *AddShardingNodeRequest) SetInstanceId(v string) *AddShardingNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *AddShardingNodeRequest) SetOwnerAccount(v string) *AddShardingNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddShardingNodeRequest) SetOwnerId(v int64) *AddShardingNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddShardingNodeRequest) SetResourceOwnerAccount(v string) *AddShardingNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddShardingNodeRequest) SetResourceOwnerId(v int64) *AddShardingNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddShardingNodeRequest) SetSecurityToken(v string) *AddShardingNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddShardingNodeRequest) SetShardCount(v int32) *AddShardingNodeRequest {
	s.ShardCount = &v
	return s
}

func (s *AddShardingNodeRequest) SetSourceBiz(v string) *AddShardingNodeRequest {
	s.SourceBiz = &v
	return s
}

func (s *AddShardingNodeRequest) SetVSwitchId(v string) *AddShardingNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddShardingNodeRequest) Validate() error {
	return dara.Validate(s)
}
