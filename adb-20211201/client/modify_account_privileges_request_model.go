// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegesRequest
	GetAccountName() *string
	SetAccountPrivileges(v []*ModifyAccountPrivilegesRequestAccountPrivileges) *ModifyAccountPrivilegesRequest
	GetAccountPrivileges() []*ModifyAccountPrivilegesRequestAccountPrivileges
	SetDBClusterId(v string) *ModifyAccountPrivilegesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ModifyAccountPrivilegesRequest
	GetRegionId() *string
}

type ModifyAccountPrivilegesRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that you want to grant to the database account.
	//
	// This parameter is required.
	AccountPrivileges []*ModifyAccountPrivilegesRequestAccountPrivileges `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegesRequest) GetAccountPrivileges() []*ModifyAccountPrivilegesRequestAccountPrivileges {
	return s.AccountPrivileges
}

func (s *ModifyAccountPrivilegesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPrivilegesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegesRequest) SetAccountName(v string) *ModifyAccountPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetAccountPrivileges(v []*ModifyAccountPrivilegesRequestAccountPrivileges) *ModifyAccountPrivilegesRequest {
	s.AccountPrivileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetRegionId(v string) *ModifyAccountPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAccountPrivilegesRequestAccountPrivileges struct {
	// The objects on which you want to grant permissions, including databases, tables, and columns.
	PrivilegeObject *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level that you want to assign to the database account. You can call the `DescribeEnabledPrivileges` operation to query the permission level that can be assigned to the database account.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The permissions that you want to grant to the database account.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivilegeObject() *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	return s.PrivilegeObject
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivileges() []*string {
	return s.Privileges
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeObject(v *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeObject = v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeType(v string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeType = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivileges(v []*string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.Privileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) Validate() error {
	return dara.Validate(s)
}

type ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject struct {
	// The columns on which you want to grant permissions. This parameter must be specified when the PrivilegeType parameter is set to Column.
	//
	// example:
	//
	// column1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The databases on which you want to grant permissions. This parameter must be specified when the PrivilegeType parameter is set to Database, Table, or Column.
	//
	// example:
	//
	// tsdb1
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The tables on which you want to grant permissions. This parameter must be specified when the PrivilegeType parameter is set to Table or Column.
	//
	// example:
	//
	// table1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetColumn() *string {
	return s.Column
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetDatabase() *string {
	return s.Database
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetTable() *string {
	return s.Table
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetColumn(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Column = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetDatabase(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Database = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetTable(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Table = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) Validate() error {
	return dara.Validate(s)
}
