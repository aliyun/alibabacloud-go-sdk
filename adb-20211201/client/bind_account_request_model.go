// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *BindAccountRequest
	GetAccountName() *string
	SetDBClusterId(v string) *BindAccountRequest
	GetDBClusterId() *string
	SetRamUser(v string) *BindAccountRequest
	GetRamUser() *string
}

type BindAccountRequest struct {
	// The standard account of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz99d9nh532****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1444832459****
	RamUser *string `json:"RamUser,omitempty" xml:"RamUser,omitempty"`
}

func (s BindAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAccountRequest) GoString() string {
	return s.String()
}

func (s *BindAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BindAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindAccountRequest) GetRamUser() *string {
	return s.RamUser
}

func (s *BindAccountRequest) SetAccountName(v string) *BindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *BindAccountRequest) SetDBClusterId(v string) *BindAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindAccountRequest) SetRamUser(v string) *BindAccountRequest {
	s.RamUser = &v
	return s
}

func (s *BindAccountRequest) Validate() error {
	return dara.Validate(s)
}
