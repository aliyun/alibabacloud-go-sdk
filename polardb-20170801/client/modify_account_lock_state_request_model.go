// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountLockStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountLockState(v string) *ModifyAccountLockStateRequest
	GetAccountLockState() *string
	SetAccountName(v string) *ModifyAccountLockStateRequest
	GetAccountName() *string
	SetAccountPasswordValidTime(v string) *ModifyAccountLockStateRequest
	GetAccountPasswordValidTime() *string
	SetDBClusterId(v string) *ModifyAccountLockStateRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyAccountLockStateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountLockStateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountLockStateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountLockStateRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountLockStateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Lock
	AccountLockState *string `json:"AccountLockState,omitempty" xml:"AccountLockState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// your_account_name
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2026-09-17 10:00:00
	AccountPasswordValidTime *string `json:"AccountPasswordValidTime,omitempty" xml:"AccountPasswordValidTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-***
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountLockStateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountLockStateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountLockStateRequest) GetAccountLockState() *string {
	return s.AccountLockState
}

func (s *ModifyAccountLockStateRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountLockStateRequest) GetAccountPasswordValidTime() *string {
	return s.AccountPasswordValidTime
}

func (s *ModifyAccountLockStateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountLockStateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountLockStateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountLockStateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountLockStateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountLockStateRequest) SetAccountLockState(v string) *ModifyAccountLockStateRequest {
	s.AccountLockState = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetAccountName(v string) *ModifyAccountLockStateRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetAccountPasswordValidTime(v string) *ModifyAccountLockStateRequest {
	s.AccountPasswordValidTime = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetDBClusterId(v string) *ModifyAccountLockStateRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetOwnerAccount(v string) *ModifyAccountLockStateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetOwnerId(v int64) *ModifyAccountLockStateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetResourceOwnerAccount(v string) *ModifyAccountLockStateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountLockStateRequest) SetResourceOwnerId(v int64) *ModifyAccountLockStateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountLockStateRequest) Validate() error {
	return dara.Validate(s)
}
