// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyAccountAuthorityRequest
	GetAccount() *string
	SetDBInstanceId(v string) *ModifyAccountAuthorityRequest
	GetDBInstanceId() *string
	SetDmlAuthSetting(v *ModifyAccountAuthorityRequestDmlAuthSetting) *ModifyAccountAuthorityRequest
	GetDmlAuthSetting() *ModifyAccountAuthorityRequestDmlAuthSetting
	SetRegionId(v string) *ModifyAccountAuthorityRequest
	GetRegionId() *string
}

type ModifyAccountAuthorityRequest struct {
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
	// The information about permissions.
	//
	// This parameter is required.
	DmlAuthSetting *ModifyAccountAuthorityRequestDmlAuthSetting `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyAccountAuthorityRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountAuthorityRequest) GetDmlAuthSetting() *ModifyAccountAuthorityRequestDmlAuthSetting {
	return s.DmlAuthSetting
}

func (s *ModifyAccountAuthorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountAuthorityRequest) SetAccount(v string) *ModifyAccountAuthorityRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDBInstanceId(v string) *ModifyAccountAuthorityRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDmlAuthSetting(v *ModifyAccountAuthorityRequestDmlAuthSetting) *ModifyAccountAuthorityRequest {
	s.DmlAuthSetting = v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetRegionId(v string) *ModifyAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) Validate() error {
	if s.DmlAuthSetting != nil {
		if err := s.DmlAuthSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAccountAuthorityRequestDmlAuthSetting struct {
	// The databases on which you want to grant permissions. Separate multiple databases with commas (,).
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The dictionaries on which you want to grant permissions. Separate multiple dictionaries with commas (,).
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Specifies whether to grant the DDL permissions to the database account. Valid values:
	//
	// 	- **true**: The account has the permissions to execute DDL statements.
	//
	// 	- **false**: The account does not have the permissions to execute DDL statements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant the DML permissions to the database account. Valid values:
	//
	// 	- **0**: The account has the permissions to read data from the database, write data to the database, and modify the settings of the database.
	//
	// 	- **1**: The account only has the permissions to read data from the database.
	//
	// 	- **2**: The account only has the permissions to read data from the database and modify the settings of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
}

func (s ModifyAccountAuthorityRequestDmlAuthSetting) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityRequestDmlAuthSetting) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) GetAllowDatabases() []*string {
	return s.AllowDatabases
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) GetAllowDictionaries() []*string {
	return s.AllowDictionaries
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) GetDmlAuthority() *int32 {
	return s.DmlAuthority
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetAllowDatabases(v []*string) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.AllowDatabases = v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetAllowDictionaries(v []*string) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.AllowDictionaries = v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetDdlAuthority(v bool) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.DdlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) SetDmlAuthority(v int32) *ModifyAccountAuthorityRequestDmlAuthSetting {
	s.DmlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequestDmlAuthSetting) Validate() error {
	return dara.Validate(s)
}
