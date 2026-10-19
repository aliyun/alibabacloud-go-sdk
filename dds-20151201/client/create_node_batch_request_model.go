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
	// The name of the account. Valid values:
	//
	// - The name must start with a lowercase letter.
	//
	// - The name can contain lowercase letters, digits, and underscores (_).
	//
	// - The name must be 4 to 16 characters in length.
	//
	// > - ApsaraDB for MongoDB does not support using keywords as account names.
	//
	// > - The permissions of the account are fixed to read-only.
	//
	// > - When you apply for a direct connection endpoint for a shard node for the first time, you must set the account name and password. You do not need to set them again for subsequent applications.
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
	// > ApsaraDB for MongoDB does not support resetting the password of the account for a shard node.
	//
	// example:
	//
	// PassWord123
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: (Default) Automatic payment is enabled. Make sure that your account has sufficient balance.
	//
	// - **false**: Manual payment is enabled. For more information, see [Manually renew a subscription instance](https://help.aliyun.com/document_detail/85052.html).
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
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null**: (Default) A coupon is used.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: No coupon is used.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp18b0934e70****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The source of the request. Valid values:
	//
	// - **OpenApi**: OpenAPI.
	//
	// - **mongo_buy**: The console.
	//
	// example:
	//
	// OpenApi
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The specifications of the Mongos or shard nodes to be added. For more information about the specifications, see [Instance specifications](https://help.aliyun.com/document_detail/57141.html).
	//
	// > By default, a sharded cluster instance supports up to 32 Mongos nodes and 32 shard nodes.
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
	// Specifies whether to apply for a direct connection endpoint for the shard node. Valid values:
	//
	// - **true**: Applies for a direct connection endpoint for the shard node.
	//
	// - **false**: (Default) Does not apply for a direct connection endpoint for the shard node.
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
