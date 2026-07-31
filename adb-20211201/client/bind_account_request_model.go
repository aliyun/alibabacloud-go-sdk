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
	SetRamUserList(v []*string) *BindAccountRequest
	GetRamUserList() []*string
}

type BindAccountRequest struct {
	// A standard database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// ID of the cluster. Applies to Enterprise Edition, Basic Edition, or Data Lakehouse Edition clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz99d9nh532****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// ID of the Alibaba Cloud RAM user to bind.
	//
	// example:
	//
	// 1444832459****
	RamUser *string `json:"RamUser,omitempty" xml:"RamUser,omitempty"`
	// List of Alibaba Cloud RAM user IDs to bind. You can bind only one RAM user at a time. If you specify this parameter, the RamUser parameter is ignored.
	RamUserList []*string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty" type:"Repeated"`
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

func (s *BindAccountRequest) GetRamUserList() []*string {
	return s.RamUserList
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

func (s *BindAccountRequest) SetRamUserList(v []*string) *BindAccountRequest {
	s.RamUserList = v
	return s
}

func (s *BindAccountRequest) Validate() error {
	return dara.Validate(s)
}
