// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateNodeRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateNodeRequest
	GetAccountPassword() *string
	SetAutoPay(v bool) *CreateNodeRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *CreateNodeRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *CreateNodeRequest
	GetClientToken() *string
	SetCouponNo(v string) *CreateNodeRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *CreateNodeRequest
	GetDBInstanceId() *string
	SetNodeClass(v string) *CreateNodeRequest
	GetNodeClass() *string
	SetNodeStorage(v int32) *CreateNodeRequest
	GetNodeStorage() *int32
	SetNodeType(v string) *CreateNodeRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *CreateNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNodeRequest
	GetOwnerId() *int64
	SetReadonlyReplicas(v int32) *CreateNodeRequest
	GetReadonlyReplicas() *int32
	SetResourceOwnerAccount(v string) *CreateNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNodeRequest
	GetResourceOwnerId() *int64
	SetSearchDBInstanceClass(v string) *CreateNodeRequest
	GetSearchDBInstanceClass() *string
	SetSearchNodeCount(v int32) *CreateNodeRequest
	GetSearchNodeCount() *int32
	SetSearchStorage(v int32) *CreateNodeRequest
	GetSearchStorage() *int32
	SetShardDirect(v bool) *CreateNodeRequest
	GetShardDirect() *bool
}

type CreateNodeRequest struct {
	// The username. Valid values:
	//
	// - The username must start with a lowercase letter.
	//
	// - The username can contain lowercase letters, digits, and underscores (_).
	//
	// - The username must be 4 to 16 characters in length.
	//
	// > - ApsaraDB for MongoDB does not support using keywords as usernames.
	//
	// > - The permissions of the account are fixed to read-only.
	//
	// > - When you apply for a direct connection endpoint for the shard node for the first time, you must set the username and password. You do not need to set them again for subsequent applications.
	//
	// example:
	//
	// ceshi
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the account. Valid values:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include !@#$%^&*()_+-=.
	//
	// - The password must be 8 to 32 characters in length.
	//
	// > ApsaraDB for MongoDB does not support resetting the password of the account for the shard node.
	//
	// example:
	//
	// 123+abc
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. This is the default value. Make sure that your account has sufficient balance.
	//
	// <props="china">- **false**: disables automatic payment. You can log on to the ApsaraDB for MongoDB console, choose **Expenses*	- > **Expenses and Costs*	- in the upper-right corner, choose **Subscription Orders*	- > **My Orders*	- in the left-side navigation pane, and then pay for the order on the **Orders*	- tab.
	//
	// <props="intl">- **false**: disables automatic payment. You can log on to the ApsaraDB for MongoDB console, choose **Expenses*	- > **Expenses and Costs*	- in the upper-right corner, click **Order Management*	- in the left-side navigation pane, and then pay for the order on the **Orders*	- page.
	//
	// > This parameter is required when the billing method is subscription.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The additional parameter that contains business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null**: uses a coupon. This is the default value.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: does not use a coupon.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance ID of the sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11501cd7b5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance type of the shard node or mongos node. For more information about instance types, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.shard.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The disk space of the node. Unit: GB.
	//
	// The valid values vary based on the storage type. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// > This parameter is required when the node type is **Shard**.
	//
	// example:
	//
	// 20
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The node type. Valid values:
	//
	// - **shard**: shard node.
	//
	// - **mongos**: mongos node.
	//
	// This parameter is required.
	//
	// example:
	//
	// shard
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes in the shard node.
	//
	// Valid values: **0*	- to **5**. Default value: **0**.
	//
	// > This parameter is supported only on the Alibaba Cloud China site.
	//
	// example:
	//
	// 5
	ReadonlyReplicas      *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchDBInstanceClass *string `json:"SearchDBInstanceClass,omitempty" xml:"SearchDBInstanceClass,omitempty"`
	SearchNodeCount       *int32  `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
	SearchStorage         *int32  `json:"SearchStorage,omitempty" xml:"SearchStorage,omitempty"`
	// Specifies whether to apply for a direct connection endpoint for the shard node. Valid values:
	//
	// - **true**: applies for a direct connection endpoint for the shard node.
	//
	// - **false**: does not apply for a direct connection endpoint for the shard node. This is the default value.
	//
	// example:
	//
	// false
	ShardDirect *bool `json:"ShardDirect,omitempty" xml:"ShardDirect,omitempty"`
}

func (s CreateNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateNodeRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNodeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNodeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateNodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateNodeRequest) GetNodeClass() *string {
	return s.NodeClass
}

func (s *CreateNodeRequest) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *CreateNodeRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNodeRequest) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *CreateNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNodeRequest) GetSearchDBInstanceClass() *string {
	return s.SearchDBInstanceClass
}

func (s *CreateNodeRequest) GetSearchNodeCount() *int32 {
	return s.SearchNodeCount
}

func (s *CreateNodeRequest) GetSearchStorage() *int32 {
	return s.SearchStorage
}

func (s *CreateNodeRequest) GetShardDirect() *bool {
	return s.ShardDirect
}

func (s *CreateNodeRequest) SetAccountName(v string) *CreateNodeRequest {
	s.AccountName = &v
	return s
}

func (s *CreateNodeRequest) SetAccountPassword(v string) *CreateNodeRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateNodeRequest) SetAutoPay(v bool) *CreateNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNodeRequest) SetBusinessInfo(v string) *CreateNodeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateNodeRequest) SetClientToken(v string) *CreateNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNodeRequest) SetCouponNo(v string) *CreateNodeRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateNodeRequest) SetDBInstanceId(v string) *CreateNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNodeRequest) SetNodeClass(v string) *CreateNodeRequest {
	s.NodeClass = &v
	return s
}

func (s *CreateNodeRequest) SetNodeStorage(v int32) *CreateNodeRequest {
	s.NodeStorage = &v
	return s
}

func (s *CreateNodeRequest) SetNodeType(v string) *CreateNodeRequest {
	s.NodeType = &v
	return s
}

func (s *CreateNodeRequest) SetOwnerAccount(v string) *CreateNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNodeRequest) SetOwnerId(v int64) *CreateNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNodeRequest) SetReadonlyReplicas(v int32) *CreateNodeRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *CreateNodeRequest) SetResourceOwnerAccount(v string) *CreateNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNodeRequest) SetResourceOwnerId(v int64) *CreateNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNodeRequest) SetSearchDBInstanceClass(v string) *CreateNodeRequest {
	s.SearchDBInstanceClass = &v
	return s
}

func (s *CreateNodeRequest) SetSearchNodeCount(v int32) *CreateNodeRequest {
	s.SearchNodeCount = &v
	return s
}

func (s *CreateNodeRequest) SetSearchStorage(v int32) *CreateNodeRequest {
	s.SearchStorage = &v
	return s
}

func (s *CreateNodeRequest) SetShardDirect(v bool) *CreateNodeRequest {
	s.ShardDirect = &v
	return s
}

func (s *CreateNodeRequest) Validate() error {
	return dara.Validate(s)
}
