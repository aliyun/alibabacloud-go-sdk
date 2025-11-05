// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSchemaPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *RevokeSchemaPermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *RevokeSchemaPermissionRequest
	GetPrivileges() []*string
	SetSchemaName(v string) *RevokeSchemaPermissionRequest
	GetSchemaName() *string
	SetUserName(v string) *RevokeSchemaPermissionRequest
	GetUserName() *string
}

type RevokeSchemaPermissionRequest struct {
	// example:
	//
	// test_db
	DatabaseName *string   `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Privileges   []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// example:
	//
	// my_schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// example:
	//
	// p4_1234xxxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RevokeSchemaPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSchemaPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeSchemaPermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *RevokeSchemaPermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *RevokeSchemaPermissionRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *RevokeSchemaPermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *RevokeSchemaPermissionRequest) SetDatabaseName(v string) *RevokeSchemaPermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *RevokeSchemaPermissionRequest) SetPrivileges(v []*string) *RevokeSchemaPermissionRequest {
	s.Privileges = v
	return s
}

func (s *RevokeSchemaPermissionRequest) SetSchemaName(v string) *RevokeSchemaPermissionRequest {
	s.SchemaName = &v
	return s
}

func (s *RevokeSchemaPermissionRequest) SetUserName(v string) *RevokeSchemaPermissionRequest {
	s.UserName = &v
	return s
}

func (s *RevokeSchemaPermissionRequest) Validate() error {
	return dara.Validate(s)
}
