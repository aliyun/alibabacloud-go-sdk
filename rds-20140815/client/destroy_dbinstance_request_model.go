// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DestroyDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DestroyDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DestroyDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DestroyDBInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DestroyDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DestroyDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DestroyDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// > You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCziJZNwH****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DestroyDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DestroyDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DestroyDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DestroyDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DestroyDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DestroyDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DestroyDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DestroyDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DestroyDBInstanceRequest) SetClientToken(v string) *DestroyDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DestroyDBInstanceRequest) SetDBInstanceId(v string) *DestroyDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DestroyDBInstanceRequest) SetOwnerAccount(v string) *DestroyDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DestroyDBInstanceRequest) SetOwnerId(v int64) *DestroyDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DestroyDBInstanceRequest) SetResourceOwnerAccount(v string) *DestroyDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DestroyDBInstanceRequest) SetResourceOwnerId(v int64) *DestroyDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DestroyDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
