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
	SetShardDirect(v bool) *CreateNodeRequest
	GetShardDirect() *bool
}

type CreateNodeRequest struct {
	// The username of the account. The username must meet the following requirements:
	//
	// 	- The username starts with a lowercase letter.
	//
	// 	- The username can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The username must be 4 to 16 characters in length.
	//
	// >
	//
	// 	- Keywords cannot be used as accounts.
	//
	// 	- This account is granted the read-only permissions.
	//
	// 	- The username and password need to be set if you apply for an endpoint for the shard node for the first time.
	//
	// example:
	//
	// ceshi
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the account. The password must meet the following requirements:
	//
	// 	- The password contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- These special characters include ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password is 8 to 32 characters in length.
	//
	// >  ApsaraDB for MongoDB does not allow you to reset the password of an account.
	//
	// example:
	//
	// 123+abc
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. You can perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, choose **Expenses*	- > Orders. On the **Orders*	- page, find the order that you want to pay for and complete the payment.
	//
	// >  This parameter is required only when the billing method of the instance is subscription.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The coupon code. Default value: **youhuiquan_promotion_option_id_for_blank**.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11501cd7b5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance type of the shard or mongos node. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.shard.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The disk capacity of the node. Unit: GB.
	//
	// Valid values: **10*	- to **2000**. The value must be a multiple of 10.
	//
	// >  This parameter is required only when the NodeType parameter is set to **shard**.
	//
	// example:
	//
	// 10
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **shard**: shard node
	//
	// 	- **mongos**: mongos node
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
	// Valid values: **0**, 1, 2, 3, 4, and **5**. Default value: **0**.
	//
	// >  This parameter is available only for ApsaraDB for MongoDB instances that are purchased on the China site (aliyun.com).
	//
	// example:
	//
	// 5
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to apply for an endpoint for the shard node. Valid values:
	//
	// 	- **true**: applies for an endpoint for the shard node.
	//
	// 	- **false*	- (default): does not apply for an endpoint for the shard node.
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

func (s *CreateNodeRequest) SetShardDirect(v bool) *CreateNodeRequest {
	s.ShardDirect = &v
	return s
}

func (s *CreateNodeRequest) Validate() error {
	return dara.Validate(s)
}
