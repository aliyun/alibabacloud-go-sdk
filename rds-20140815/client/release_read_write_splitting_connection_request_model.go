// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseReadWriteSplittingConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ReleaseReadWriteSplittingConnectionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ReleaseReadWriteSplittingConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseReadWriteSplittingConnectionRequest
	GetOwnerId() *int64
	SetRWAddressType(v string) *ReleaseReadWriteSplittingConnectionRequest
	GetRWAddressType() *string
	SetResourceOwnerAccount(v string) *ReleaseReadWriteSplittingConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseReadWriteSplittingConnectionRequest
	GetResourceOwnerId() *int64
}

type ReleaseReadWriteSplittingConnectionRequest struct {
	// The ID of the primary instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RWAddressType        *string `json:"RWAddressType,omitempty" xml:"RWAddressType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseReadWriteSplittingConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseReadWriteSplittingConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetRWAddressType() *string {
	return s.RWAddressType
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseReadWriteSplittingConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetDBInstanceId(v string) *ReleaseReadWriteSplittingConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetOwnerAccount(v string) *ReleaseReadWriteSplittingConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetOwnerId(v int64) *ReleaseReadWriteSplittingConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetRWAddressType(v string) *ReleaseReadWriteSplittingConnectionRequest {
	s.RWAddressType = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseReadWriteSplittingConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) SetResourceOwnerId(v int64) *ReleaseReadWriteSplittingConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionRequest) Validate() error {
	return dara.Validate(s)
}
