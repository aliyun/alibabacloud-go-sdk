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
	// The database name.
	//
	// example:
	//
	// db_demo
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// The list of permissions.
	Privileges []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// The schema name.
	//
	// example:
	//
	// my_schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The user ID. For more information, see [Hologres account overview](https://help.aliyun.com/document_detail/130393.html).
	//
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
