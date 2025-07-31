// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateNodeBatchRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateNodeBatchRequest
	GetAccountPassword() *string
	SetAutoPay(v bool) *CreateNodeBatchRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *CreateNodeBatchRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *CreateNodeBatchRequest
	GetClientToken() *string
	SetCouponNo(v string) *CreateNodeBatchRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *CreateNodeBatchRequest
	GetDBInstanceId() *string
	SetFromApp(v string) *CreateNodeBatchRequest
	GetFromApp() *string
	SetNodesInfo(v string) *CreateNodeBatchRequest
	GetNodesInfo() *string
	SetOwnerAccount(v string) *CreateNodeBatchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNodeBatchRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateNodeBatchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNodeBatchRequest
	GetResourceOwnerId() *int64
	SetShardDirect(v bool) *CreateNodeBatchRequest
	GetShardDirect() *bool
}

type CreateNodeBatchRequest struct {
	// The username of the account. The username must meet the following requirements:
	//
	// - The username starts with a lowercase letter.
	//
	// - The username contains lowercase letters, digits, and underscores (_).
	//
	// - The username is 4 to 16 characters in length.
	//
	// > - Keywords cannot be used as account usernames.
	//
	// > - The permissions of this account are fixed at read-only.
	//
	// > - The username and password are required to be set only when you apply for an endpoint for the shard node for the first time.
	//
	// example:
	//
	// ceshi
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the account. The password must meet the following requirements:
	//
	// - The password contains at least three of the following character types: uppercase letters, lowercase letters, digits, and specific special characters.
	//
	// - These special characters include ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password is 8 to 32 characters in length.
	//
	// > The account password of the shard node cannot be reset.
	//
	// example:
	//
	// 123+abc
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// - **false**: disables automatic payment. In this case, you must manually pay for the instance. You can perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, choose **Expenses*	- > **Orders**. On the Orders page, find the order and complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use coupons. Default value: null. Valid values:
	//
	// 	- **default*	- or **null**: uses coupons.
	//
	// 	- **youhuiquan_promotion_option_id_for_blank**: does not use coupons.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the instance for which you want to add nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp18b0934e70****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The source of the request. Valid values:
	//
	// - **OpenApi**: ApsaraDB for MongoDB API
	//
	// - **mongo_buy**: ApsaraDB for MongoDB console
	//
	// example:
	//
	// OpenApi
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The specifications of the mongos or shard node that you want to add. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// > Up to 32 mongos or shard nodes are supported for each sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Shards":[{"DBInstanceClass":"mdb.shard.4x.large.d","Storage":20}]}
	NodesInfo            *string `json:"NodesInfo,omitempty" xml:"NodesInfo,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to apply for an endpoint for the shard node. Default value: false. Valid values:
	//
	// - **true**: applies for an endpoint for the shard node.
	//
	// - **false**: does not apply for an endpoint for the shard node.
	//
	// example:
	//
	// false
	ShardDirect *bool `json:"ShardDirect,omitempty" xml:"ShardDirect,omitempty"`
}

func (s CreateNodeBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeBatchRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateNodeBatchRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateNodeBatchRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNodeBatchRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateNodeBatchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNodeBatchRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateNodeBatchRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateNodeBatchRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *CreateNodeBatchRequest) GetNodesInfo() *string {
	return s.NodesInfo
}

func (s *CreateNodeBatchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNodeBatchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNodeBatchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNodeBatchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNodeBatchRequest) GetShardDirect() *bool {
	return s.ShardDirect
}

func (s *CreateNodeBatchRequest) SetAccountName(v string) *CreateNodeBatchRequest {
	s.AccountName = &v
	return s
}

func (s *CreateNodeBatchRequest) SetAccountPassword(v string) *CreateNodeBatchRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateNodeBatchRequest) SetAutoPay(v bool) *CreateNodeBatchRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNodeBatchRequest) SetBusinessInfo(v string) *CreateNodeBatchRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateNodeBatchRequest) SetClientToken(v string) *CreateNodeBatchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNodeBatchRequest) SetCouponNo(v string) *CreateNodeBatchRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateNodeBatchRequest) SetDBInstanceId(v string) *CreateNodeBatchRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNodeBatchRequest) SetFromApp(v string) *CreateNodeBatchRequest {
	s.FromApp = &v
	return s
}

func (s *CreateNodeBatchRequest) SetNodesInfo(v string) *CreateNodeBatchRequest {
	s.NodesInfo = &v
	return s
}

func (s *CreateNodeBatchRequest) SetOwnerAccount(v string) *CreateNodeBatchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNodeBatchRequest) SetOwnerId(v int64) *CreateNodeBatchRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNodeBatchRequest) SetResourceOwnerAccount(v string) *CreateNodeBatchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNodeBatchRequest) SetResourceOwnerId(v int64) *CreateNodeBatchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNodeBatchRequest) SetShardDirect(v bool) *CreateNodeBatchRequest {
	s.ShardDirect = &v
	return s
}

func (s *CreateNodeBatchRequest) Validate() error {
	return dara.Validate(s)
}
