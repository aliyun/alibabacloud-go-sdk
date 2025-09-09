// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegeRequest
	GetAccountName() *string
	SetDbPrivilege(v []*ModifyAccountPrivilegeRequestDbPrivilege) *ModifyAccountPrivilegeRequest
	GetDbPrivilege() []*ModifyAccountPrivilegeRequestDbPrivilege
	SetDrdsInstanceId(v string) *ModifyAccountPrivilegeRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *ModifyAccountPrivilegeRequest
	GetRegionId() *string
}

type ModifyAccountPrivilegeRequest struct {
	// The username of the account that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// account_sec
	AccountName *string                                     `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbPrivilege []*ModifyAccountPrivilegeRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbgaen89****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegeRequest) GetDbPrivilege() []*ModifyAccountPrivilegeRequestDbPrivilege {
	return s.DbPrivilege
}

func (s *ModifyAccountPrivilegeRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyAccountPrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegeRequest) SetAccountName(v string) *ModifyAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDbPrivilege(v []*ModifyAccountPrivilegeRequestDbPrivilege) *ModifyAccountPrivilegeRequest {
	s.DbPrivilege = v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetDrdsInstanceId(v string) *ModifyAccountPrivilegeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) SetRegionId(v string) *ModifyAccountPrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAccountPrivilegeRequestDbPrivilege struct {
	// The name of the database that you want to manage by using the account to modify.
	//
	// example:
	//
	// test123
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The permissions that you want to grant to the account.
	//
	// example:
	//
	// ReadWrite
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegeRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) GetDbName() *string {
	return s.DbName
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) GetPrivilege() *string {
	return s.Privilege
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetDbName(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) SetPrivilege(v string) *ModifyAccountPrivilegeRequestDbPrivilege {
	s.Privilege = &v
	return s
}

func (s *ModifyAccountPrivilegeRequestDbPrivilege) Validate() error {
	return dara.Validate(s)
}
