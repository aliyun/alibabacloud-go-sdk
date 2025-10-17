// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CheckAccountNameZonalRequest
	GetAccountName() *string
	SetDBClusterId(v string) *CheckAccountNameZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckAccountNameZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckAccountNameZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckAccountNameZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckAccountNameZonalRequest
	GetResourceOwnerId() *int64
}

type CheckAccountNameZonalRequest struct {
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
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckAccountNameZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameZonalRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountNameZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CheckAccountNameZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckAccountNameZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckAccountNameZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckAccountNameZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckAccountNameZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckAccountNameZonalRequest) SetAccountName(v string) *CheckAccountNameZonalRequest {
	s.AccountName = &v
	return s
}

func (s *CheckAccountNameZonalRequest) SetDBClusterId(v string) *CheckAccountNameZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckAccountNameZonalRequest) SetOwnerAccount(v string) *CheckAccountNameZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckAccountNameZonalRequest) SetOwnerId(v int64) *CheckAccountNameZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckAccountNameZonalRequest) SetResourceOwnerAccount(v string) *CheckAccountNameZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckAccountNameZonalRequest) SetResourceOwnerId(v int64) *CheckAccountNameZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckAccountNameZonalRequest) Validate() error {
	return dara.Validate(s)
}
