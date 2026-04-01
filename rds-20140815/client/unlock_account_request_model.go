// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *UnlockAccountRequest
	GetAccountName() *string
	SetDBInstanceId(v string) *UnlockAccountRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UnlockAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnlockAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnlockAccountRequest
	GetResourceOwnerId() *int64
}

type UnlockAccountRequest struct {
	// The account that you want to unlock. You can unlock a single account at a time.
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

func (s UnlockAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockAccountRequest) GoString() string {
	return s.String()
}

func (s *UnlockAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *UnlockAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnlockAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnlockAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnlockAccountRequest) SetAccountName(v string) *UnlockAccountRequest {
	s.AccountName = &v
	return s
}

func (s *UnlockAccountRequest) SetDBInstanceId(v string) *UnlockAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnlockAccountRequest) SetOwnerId(v int64) *UnlockAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockAccountRequest) SetResourceOwnerAccount(v string) *UnlockAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockAccountRequest) SetResourceOwnerId(v int64) *UnlockAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockAccountRequest) Validate() error {
	return dara.Validate(s)
}
