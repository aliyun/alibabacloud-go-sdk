// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameAvailableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CheckAccountNameAvailableRequest
	GetAccountName() *string
	SetClientToken(v string) *CheckAccountNameAvailableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CheckAccountNameAvailableRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CheckAccountNameAvailableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckAccountNameAvailableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckAccountNameAvailableRequest
	GetResourceOwnerAccount() *string
}

type CheckAccountNameAvailableRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s CheckAccountNameAvailableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameAvailableRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountNameAvailableRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CheckAccountNameAvailableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckAccountNameAvailableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckAccountNameAvailableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckAccountNameAvailableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckAccountNameAvailableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckAccountNameAvailableRequest) SetAccountName(v string) *CheckAccountNameAvailableRequest {
	s.AccountName = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetClientToken(v string) *CheckAccountNameAvailableRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetDBInstanceId(v string) *CheckAccountNameAvailableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetOwnerAccount(v string) *CheckAccountNameAvailableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetOwnerId(v int64) *CheckAccountNameAvailableRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetResourceOwnerAccount(v string) *CheckAccountNameAvailableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) Validate() error {
	return dara.Validate(s)
}
