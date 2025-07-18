// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountPasswordRequest
	GetAccountPassword() *string
	SetDBInstanceId(v string) *ResetAccountPasswordRequest
	GetDBInstanceId() *string
}

type ResetAccountPasswordRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount_1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Testaccount_1
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Before you call this operation, make sure that the following requirements are met:
	//
	// 	- The instance is in the running state.
	//
	// 	- The instance is not locked.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-t4nf48vf15713****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountPasswordRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
