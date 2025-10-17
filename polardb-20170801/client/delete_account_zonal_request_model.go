// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAccountZonalRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DeleteAccountZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteAccountZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAccountZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAccountZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAccountZonalRequest
	GetResourceOwnerId() *int64
}

type DeleteAccountZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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

func (s DeleteAccountZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountZonalRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAccountZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAccountZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAccountZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccountZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccountZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAccountZonalRequest) SetAccountName(v string) *DeleteAccountZonalRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountZonalRequest) SetDBClusterId(v string) *DeleteAccountZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountZonalRequest) SetOwnerAccount(v string) *DeleteAccountZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountZonalRequest) SetOwnerId(v int64) *DeleteAccountZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountZonalRequest) SetResourceOwnerAccount(v string) *DeleteAccountZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountZonalRequest) SetResourceOwnerId(v int64) *DeleteAccountZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountZonalRequest) Validate() error {
	return dara.Validate(s)
}
