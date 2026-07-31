// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *BindAccountShrinkRequest
	GetAccountName() *string
	SetDBClusterId(v string) *BindAccountShrinkRequest
	GetDBClusterId() *string
	SetRamUser(v string) *BindAccountShrinkRequest
	GetRamUser() *string
	SetRamUserListShrink(v string) *BindAccountShrinkRequest
	GetRamUserListShrink() *string
}

type BindAccountShrinkRequest struct {
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
	RamUserListShrink *string `json:"RamUserList,omitempty" xml:"RamUserList,omitempty"`
}

func (s BindAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAccountShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BindAccountShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindAccountShrinkRequest) GetRamUser() *string {
	return s.RamUser
}

func (s *BindAccountShrinkRequest) GetRamUserListShrink() *string {
	return s.RamUserListShrink
}

func (s *BindAccountShrinkRequest) SetAccountName(v string) *BindAccountShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *BindAccountShrinkRequest) SetDBClusterId(v string) *BindAccountShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindAccountShrinkRequest) SetRamUser(v string) *BindAccountShrinkRequest {
	s.RamUser = &v
	return s
}

func (s *BindAccountShrinkRequest) SetRamUserListShrink(v string) *BindAccountShrinkRequest {
	s.RamUserListShrink = &v
	return s
}

func (s *BindAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
