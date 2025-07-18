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
	SetDBInstanceId(v string) *DeleteAccountRequest
	GetDBInstanceId() *string
}

type DeleteAccountRequest struct {
	// The name of the database account. The privileged account cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

func (s *DeleteAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBInstanceId(v string) *DeleteAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
