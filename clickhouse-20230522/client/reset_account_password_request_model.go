// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ResetAccountPasswordRequest
	GetAccount() *string
	SetDBInstanceId(v string) *ResetAccountPasswordRequest
	GetDBInstanceId() *string
	SetPassword(v string) *ResetAccountPasswordRequest
	GetPassword() *string
	SetProduct(v string) *ResetAccountPasswordRequest
	GetProduct() *string
	SetRegionId(v string) *ResetAccountPasswordRequest
	GetRegionId() *string
}

type ResetAccountPasswordRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: ! @ # $ % ^ & 	- ( ) _ + - =
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service name.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) GetAccount() *string {
	return s.Account
}

func (s *ResetAccountPasswordRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetAccountPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetAccountPasswordRequest) GetProduct() *string {
	return s.Product
}

func (s *ResetAccountPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetAccountPasswordRequest) SetAccount(v string) *ResetAccountPasswordRequest {
	s.Account = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetPassword(v string) *ResetAccountPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetProduct(v string) *ResetAccountPasswordRequest {
	s.Product = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetRegionId(v string) *ResetAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
