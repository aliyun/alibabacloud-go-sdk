// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateNodePrivateNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *AllocateNodePrivateNetworkAddressRequest
	GetAccountName() *string
	SetAccountPassword(v string) *AllocateNodePrivateNetworkAddressRequest
	GetAccountPassword() *string
	SetDBInstanceId(v string) *AllocateNodePrivateNetworkAddressRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *AllocateNodePrivateNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *AllocateNodePrivateNetworkAddressRequest
	GetZoneId() *string
}

type AllocateNodePrivateNetworkAddressRequest struct {
	// The username of the account.
	//
	// >
	//
	// 	- The username must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). The username must start with a lowercase letter.
	//
	// 	- You must configure the account and password only when you apply for the endpoint of a shard or Configserver node for the first time. The account and password are required for all shard and Configserver nodes.
	//
	// 	- The permissions of this account are fixed to read-only.
	//
	// example:
	//
	// shardcsaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password for the account.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! # $ % ^ & 	- ( ) _ + - =`
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// example:
	//
	// Test123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The ID of the sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1fa5efaa93****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard or Configserver node.
	//
	// >  You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the ID of the shard or Configserver node.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp124beeb0ac****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the zone to which the instance belongs.
	//
	// >  You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AllocateNodePrivateNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateNodePrivateNetworkAddressRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetAccountName(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.AccountName = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetAccountPassword(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.AccountPassword = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetDBInstanceId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetNodeId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocateNodePrivateNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) SetZoneId(v string) *AllocateNodePrivateNetworkAddressRequest {
	s.ZoneId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
