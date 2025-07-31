// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMongoDBLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeMongoDBLogConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeMongoDBLogConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMongoDBLogConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeMongoDBLogConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMongoDBLogConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeMongoDBLogConfigRequest struct {
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/61939.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMongoDBLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMongoDBLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeMongoDBLogConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMongoDBLogConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMongoDBLogConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMongoDBLogConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMongoDBLogConfigRequest) SetDBInstanceId(v string) *DescribeMongoDBLogConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetOwnerAccount(v string) *DescribeMongoDBLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetOwnerId(v int64) *DescribeMongoDBLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeMongoDBLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) SetResourceOwnerId(v int64) *DescribeMongoDBLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMongoDBLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
