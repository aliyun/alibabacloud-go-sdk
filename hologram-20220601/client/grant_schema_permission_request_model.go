// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSchemaPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *GrantSchemaPermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *GrantSchemaPermissionRequest
	GetPrivileges() []*string
	SetSchemaName(v string) *GrantSchemaPermissionRequest
	GetSchemaName() *string
	SetUserName(v string) *GrantSchemaPermissionRequest
	GetUserName() *string
}

type GrantSchemaPermissionRequest struct {
	// example:
	//
	// db_demo
	DatabaseName *string   `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Privileges   []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// example:
	//
	// my_schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// example:
	//
	// p4_134xxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GrantSchemaPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantSchemaPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantSchemaPermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GrantSchemaPermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *GrantSchemaPermissionRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GrantSchemaPermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *GrantSchemaPermissionRequest) SetDatabaseName(v string) *GrantSchemaPermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GrantSchemaPermissionRequest) SetPrivileges(v []*string) *GrantSchemaPermissionRequest {
	s.Privileges = v
	return s
}

func (s *GrantSchemaPermissionRequest) SetSchemaName(v string) *GrantSchemaPermissionRequest {
	s.SchemaName = &v
	return s
}

func (s *GrantSchemaPermissionRequest) SetUserName(v string) *GrantSchemaPermissionRequest {
	s.UserName = &v
	return s
}

func (s *GrantSchemaPermissionRequest) Validate() error {
	return dara.Validate(s)
}
