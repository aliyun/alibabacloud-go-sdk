// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteOpenSearchAccountRequest
	GetAccountName() *string
	SetDBInstanceName(v string) *DeleteOpenSearchAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteOpenSearchAccountRequest
	GetRegionId() *string
}

type DeleteOpenSearchAccountRequest struct {
	// The name of the account to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOpenSearchAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteOpenSearchAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteOpenSearchAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOpenSearchAccountRequest) SetAccountName(v string) *DeleteOpenSearchAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteOpenSearchAccountRequest) SetDBInstanceName(v string) *DeleteOpenSearchAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteOpenSearchAccountRequest) SetRegionId(v string) *DeleteOpenSearchAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOpenSearchAccountRequest) Validate() error {
	return dara.Validate(s)
}
