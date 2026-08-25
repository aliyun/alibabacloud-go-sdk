// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCrossAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddCrossAccountRequest
	GetAlias() *string
	SetCrossAccountRoleName(v string) *AddCrossAccountRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *AddCrossAccountRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *AddCrossAccountRequest
	GetCrossAccountUserId() *int64
}

type AddCrossAccountRequest struct {
	// The alias. The maximum length is 32 characters. This parameter is not required for cross-account backups that are configured based on a resource directory.
	//
	// example:
	//
	// Source account 1
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the RAM role for the account to back up. This parameter is used when you configure a cross-account backup by assuming a RAM role.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The type of cross-account backup. Valid values:
	//
	// - **CROSS_ACCOUNT**: Configures a cross-account backup by assuming a RAM role.
	//
	// - **CROSS_ACCOUNT_BY_RD**: Configures a cross-account backup based on a resource directory.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The UID of the account to back up.
	//
	// example:
	//
	// 1589753xxxxxx625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s AddCrossAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCrossAccountRequest) GoString() string {
	return s.String()
}

func (s *AddCrossAccountRequest) GetAlias() *string {
	return s.Alias
}

func (s *AddCrossAccountRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *AddCrossAccountRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *AddCrossAccountRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *AddCrossAccountRequest) SetAlias(v string) *AddCrossAccountRequest {
	s.Alias = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountRoleName(v string) *AddCrossAccountRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountType(v string) *AddCrossAccountRequest {
	s.CrossAccountType = &v
	return s
}

func (s *AddCrossAccountRequest) SetCrossAccountUserId(v int64) *AddCrossAccountRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *AddCrossAccountRequest) Validate() error {
	return dara.Validate(s)
}
