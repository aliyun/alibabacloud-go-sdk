// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *LockAccountRequest
	GetAccountName() *string
	SetDBInstanceId(v string) *LockAccountRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *LockAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *LockAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LockAccountRequest
	GetResourceOwnerId() *int64
}

type LockAccountRequest struct {
	// The account that you want to lock. You can lock only a single account at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bpxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s LockAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s LockAccountRequest) GoString() string {
	return s.String()
}

func (s *LockAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *LockAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *LockAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LockAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LockAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LockAccountRequest) SetAccountName(v string) *LockAccountRequest {
	s.AccountName = &v
	return s
}

func (s *LockAccountRequest) SetDBInstanceId(v string) *LockAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *LockAccountRequest) SetOwnerId(v int64) *LockAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *LockAccountRequest) SetResourceOwnerAccount(v string) *LockAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockAccountRequest) SetResourceOwnerId(v int64) *LockAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockAccountRequest) Validate() error {
	return dara.Validate(s)
}
