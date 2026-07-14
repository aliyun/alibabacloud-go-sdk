// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetAccountRequest
	GetAccountName() *string
	SetDBInstanceId(v string) *GetAccountRequest
	GetDBInstanceId() *string
}

type GetAccountRequest struct {
	// The name of the initial account. The name must meet the following requirements:
	//
	// - The name can contain lowercase letters, digits, and underscores (_).
	//
	// - The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// - The name cannot start with gp.
	//
	// - The name must be 2 to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// >You can specify up to 30 instance IDs for batch operations. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s GetAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetAccountRequest) SetAccountName(v string) *GetAccountRequest {
	s.AccountName = &v
	return s
}

func (s *GetAccountRequest) SetDBInstanceId(v string) *GetAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetAccountRequest) Validate() error {
	return dara.Validate(s)
}
