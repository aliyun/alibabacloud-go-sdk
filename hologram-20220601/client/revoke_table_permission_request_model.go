// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTablePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllTable(v bool) *RevokeTablePermissionRequest
	GetAllTable() *bool
	SetColumnNames(v []*string) *RevokeTablePermissionRequest
	GetColumnNames() []*string
	SetDatabaseName(v string) *RevokeTablePermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *RevokeTablePermissionRequest
	GetPrivileges() []*string
	SetSchemaName(v string) *RevokeTablePermissionRequest
	GetSchemaName() *string
	SetTableName(v string) *RevokeTablePermissionRequest
	GetTableName() *string
	SetUserName(v string) *RevokeTablePermissionRequest
	GetUserName() *string
}

type RevokeTablePermissionRequest struct {
	// Specifies whether the operation takes effect on all tables. Valid values:
	//
	// - true: Yes.
	//
	// - false or null: No.
	//
	// example:
	//
	// false
	AllTable *bool `json:"allTable,omitempty" xml:"allTable,omitempty"`
	// The list of column names for which permissions are to be revoked.
	ColumnNames []*string `json:"columnNames,omitempty" xml:"columnNames,omitempty" type:"Repeated"`
	// The database name.
	//
	// example:
	//
	// test_db
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// The list of permissions.
	Privileges []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// The schema name.
	//
	// example:
	//
	// my_schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// my_table
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The user ID. For more information, see [Hologres account overview](https://help.aliyun.com/document_detail/130393.html).
	//
	// example:
	//
	// p4_1234xxxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RevokeTablePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTablePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeTablePermissionRequest) GetAllTable() *bool {
	return s.AllTable
}

func (s *RevokeTablePermissionRequest) GetColumnNames() []*string {
	return s.ColumnNames
}

func (s *RevokeTablePermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *RevokeTablePermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *RevokeTablePermissionRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *RevokeTablePermissionRequest) GetTableName() *string {
	return s.TableName
}

func (s *RevokeTablePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *RevokeTablePermissionRequest) SetAllTable(v bool) *RevokeTablePermissionRequest {
	s.AllTable = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetColumnNames(v []*string) *RevokeTablePermissionRequest {
	s.ColumnNames = v
	return s
}

func (s *RevokeTablePermissionRequest) SetDatabaseName(v string) *RevokeTablePermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetPrivileges(v []*string) *RevokeTablePermissionRequest {
	s.Privileges = v
	return s
}

func (s *RevokeTablePermissionRequest) SetSchemaName(v string) *RevokeTablePermissionRequest {
	s.SchemaName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetTableName(v string) *RevokeTablePermissionRequest {
	s.TableName = &v
	return s
}

func (s *RevokeTablePermissionRequest) SetUserName(v string) *RevokeTablePermissionRequest {
	s.UserName = &v
	return s
}

func (s *RevokeTablePermissionRequest) Validate() error {
	return dara.Validate(s)
}
