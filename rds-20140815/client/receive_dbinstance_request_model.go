// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ReceiveDBInstanceRequest
	GetDBInstanceId() *string
	SetGuardDBInstanceId(v string) *ReceiveDBInstanceRequest
	GetGuardDBInstanceId() *string
	SetOwnerAccount(v string) *ReceiveDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReceiveDBInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReceiveDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReceiveDBInstanceRequest
	GetResourceOwnerId() *int64
}

type ReceiveDBInstanceRequest struct {
	// The ID of the primary instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the ID of the primary instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the disaster recovery instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the ID of the disaster recovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-tr2whku*****
	GuardDBInstanceId    *string `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReceiveDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReceiveDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReceiveDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReceiveDBInstanceRequest) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *ReceiveDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReceiveDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReceiveDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReceiveDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReceiveDBInstanceRequest) SetDBInstanceId(v string) *ReceiveDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReceiveDBInstanceRequest) SetGuardDBInstanceId(v string) *ReceiveDBInstanceRequest {
	s.GuardDBInstanceId = &v
	return s
}

func (s *ReceiveDBInstanceRequest) SetOwnerAccount(v string) *ReceiveDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReceiveDBInstanceRequest) SetOwnerId(v int64) *ReceiveDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReceiveDBInstanceRequest) SetResourceOwnerAccount(v string) *ReceiveDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReceiveDBInstanceRequest) SetResourceOwnerId(v int64) *ReceiveDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReceiveDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
