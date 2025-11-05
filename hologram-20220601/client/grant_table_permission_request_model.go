// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTablePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllTable(v bool) *GrantTablePermissionRequest
	GetAllTable() *bool
	SetDatabaseName(v string) *GrantTablePermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *GrantTablePermissionRequest
	GetPrivileges() []*string
	SetSchemaName(v string) *GrantTablePermissionRequest
	GetSchemaName() *string
	SetTableName(v string) *GrantTablePermissionRequest
	GetTableName() *string
	SetUserName(v string) *GrantTablePermissionRequest
	GetUserName() *string
}

type GrantTablePermissionRequest struct {
	// example:
	//
	// true
	AllTable *bool `json:"allTable,omitempty" xml:"allTable,omitempty"`
	// example:
	//
	// my_db
	DatabaseName *string   `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Privileges   []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// example:
	//
	// my_schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// example:
	//
	// orders_pay
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// p4_1234xxxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GrantTablePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantTablePermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantTablePermissionRequest) GetAllTable() *bool {
	return s.AllTable
}

func (s *GrantTablePermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GrantTablePermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *GrantTablePermissionRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GrantTablePermissionRequest) GetTableName() *string {
	return s.TableName
}

func (s *GrantTablePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *GrantTablePermissionRequest) SetAllTable(v bool) *GrantTablePermissionRequest {
	s.AllTable = &v
	return s
}

func (s *GrantTablePermissionRequest) SetDatabaseName(v string) *GrantTablePermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GrantTablePermissionRequest) SetPrivileges(v []*string) *GrantTablePermissionRequest {
	s.Privileges = v
	return s
}

func (s *GrantTablePermissionRequest) SetSchemaName(v string) *GrantTablePermissionRequest {
	s.SchemaName = &v
	return s
}

func (s *GrantTablePermissionRequest) SetTableName(v string) *GrantTablePermissionRequest {
	s.TableName = &v
	return s
}

func (s *GrantTablePermissionRequest) SetUserName(v string) *GrantTablePermissionRequest {
	s.UserName = &v
	return s
}

func (s *GrantTablePermissionRequest) Validate() error {
	return dara.Validate(s)
}
