// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DeleteAccountRequest
	GetAccount() *string
	SetDBInstanceId(v string) *DeleteAccountRequest
	GetDBInstanceId() *string
	SetProduct(v string) *DeleteAccountRequest
	GetProduct() *string
	SetRegionId(v string) *DeleteAccountRequest
	GetRegionId() *string
}

type DeleteAccountRequest struct {
	// The destination database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The code of the cloud service.
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

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAccount() *string {
	return s.Account
}

func (s *DeleteAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteAccountRequest) GetProduct() *string {
	return s.Product
}

func (s *DeleteAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAccountRequest) SetAccount(v string) *DeleteAccountRequest {
	s.Account = &v
	return s
}

func (s *DeleteAccountRequest) SetDBInstanceId(v string) *DeleteAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAccountRequest) SetProduct(v string) *DeleteAccountRequest {
	s.Product = &v
	return s
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
