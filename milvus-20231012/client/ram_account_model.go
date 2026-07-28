// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRamAccount interface {
	dara.Model
	String() string
	GoString() string
	SetBindable(v bool) *RamAccount
	GetBindable() *bool
	SetDisplayName(v string) *RamAccount
	GetDisplayName() *string
	SetUid(v string) *RamAccount
	GetUid() *string
	SetUserName(v string) *RamAccount
	GetUserName() *string
}

type RamAccount struct {
	// Specifies whether the RamAccount can be bound to other resources.
	//
	// example:
	//
	// true
	Bindable *bool `json:"bindable,omitempty" xml:"bindable,omitempty"`
	// The display name for the RamAccount, which appears in the console.
	//
	// example:
	//
	// Test User
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The unique identifier for the RamAccount.
	//
	// example:
	//
	// 1234567890123456
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// The user name for the RamAccount.
	//
	// example:
	//
	// test-user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RamAccount) String() string {
	return dara.Prettify(s)
}

func (s RamAccount) GoString() string {
	return s.String()
}

func (s *RamAccount) GetBindable() *bool {
	return s.Bindable
}

func (s *RamAccount) GetDisplayName() *string {
	return s.DisplayName
}

func (s *RamAccount) GetUid() *string {
	return s.Uid
}

func (s *RamAccount) GetUserName() *string {
	return s.UserName
}

func (s *RamAccount) SetBindable(v bool) *RamAccount {
	s.Bindable = &v
	return s
}

func (s *RamAccount) SetDisplayName(v string) *RamAccount {
	s.DisplayName = &v
	return s
}

func (s *RamAccount) SetUid(v string) *RamAccount {
	s.Uid = &v
	return s
}

func (s *RamAccount) SetUserName(v string) *RamAccount {
	s.UserName = &v
	return s
}

func (s *RamAccount) Validate() error {
	return dara.Validate(s)
}
