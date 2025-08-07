// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAccountRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DeleteAccountRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAccountRequest
	GetResourceOwnerId() *int64
}

type DeleteAccountRequest struct {
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
