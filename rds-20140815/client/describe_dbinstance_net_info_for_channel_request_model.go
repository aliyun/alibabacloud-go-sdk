// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoForChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetDBInstanceId() *string
	SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetDBInstanceNetRWSplitType() *string
	SetFlag(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetFlag() *string
	SetOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceNetInfoForChannelRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1k8s41l2o52****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the endpoint to query. Valid values:
	//
	// 	- **0**: a regular endpoint. This is the default value.
	//
	// 	- **1**: a read/write splitting endpoint that is assigned after the shared proxy feature is enabled.
	//
	// example:
	//
	// 0
	DBInstanceNetRWSplitType *string `json:"DBInstanceNetRWSplitType,omitempty" xml:"DBInstanceNetRWSplitType,omitempty"`
	// An internal parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// None
	Flag                 *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetDBInstanceNetRWSplitType() *string {
	return s.DBInstanceNetRWSplitType
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetFlag() *string {
	return s.Flag
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetClientToken(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.DBInstanceNetRWSplitType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetFlag(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.Flag = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) Validate() error {
	return dara.Validate(s)
}
