// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *GrantUserPermissionRequest
	GetDbId() *string
	SetDsType(v string) *GrantUserPermissionRequest
	GetDsType() *string
	SetExpireDate(v string) *GrantUserPermissionRequest
	GetExpireDate() *string
	SetInstanceId(v int64) *GrantUserPermissionRequest
	GetInstanceId() *int64
	SetLogic(v bool) *GrantUserPermissionRequest
	GetLogic() *bool
	SetPermTypes(v string) *GrantUserPermissionRequest
	GetPermTypes() *string
	SetTableId(v string) *GrantUserPermissionRequest
	GetTableId() *string
	SetTableName(v string) *GrantUserPermissionRequest
	GetTableName() *string
	SetTid(v int64) *GrantUserPermissionRequest
	GetTid() *int64
	SetUserId(v string) *GrantUserPermissionRequest
	GetUserId() *string
}

type GrantUserPermissionRequest struct {
	// The ID of the database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// >  The value of the DatabaseId parameter is that of the DbId parameter.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The permissions on a specific type of object that you want to grant to the user. Valid values:
	//
	// 	- INSTANCE: permissions on instances
	//
	// 	- DATABASE: permissions on physical databases
	//
	// 	- LOGIC_DATABASE: permissions on logical databases
	//
	// 	- TABLE: permissions on physical tables
	//
	// 	- LOGIC_TABLE: permissions on logical tables
	//
	// This parameter is required.
	//
	// example:
	//
	// DATABASE
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The time when the permissions expire.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-12-12 00:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the instance. You must specify this parameter if you grant permissions on an instance to the user. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the ID of the instance.
	//
	// example:
	//
	// 174****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the database is a logical database. You must specify this parameter if you grant permissions on a database to the user. Valid values:
	//
	// 	- true: The database is a logical database.
	//
	// 	- false: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The permission type. Separate multiple permission types with commas (,). Valid values:
	//
	// 	- **QUERY**: the query permissions
	//
	// 	- **EXPORT**: the export permissions
	//
	// 	- **CORRECT**: the change permissions
	//
	// 	- **LOGIN**: the logon permissions
	//
	// 	- **PERF**: the query permissions on the performance details of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// QUERY
	PermTypes *string `json:"PermTypes,omitempty" xml:"PermTypes,omitempty"`
	// The ID of the table. You must specify this parameter if you grant permissions on a table to the user. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table ID.
	//
	// example:
	//
	// 132***
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table. You must specify this parameter if you grant permissions on a table to the user.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the user. You can call the [GetUser](https://help.aliyun.com/document_detail/147098.html) or [ListUsers](https://help.aliyun.com/document_detail/141938.html) operation to query the ID of the user.
	//
	// >  The user ID is different from the ID of your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantUserPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionRequest) GetDbId() *string {
	return s.DbId
}

func (s *GrantUserPermissionRequest) GetDsType() *string {
	return s.DsType
}

func (s *GrantUserPermissionRequest) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GrantUserPermissionRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GrantUserPermissionRequest) GetLogic() *bool {
	return s.Logic
}

func (s *GrantUserPermissionRequest) GetPermTypes() *string {
	return s.PermTypes
}

func (s *GrantUserPermissionRequest) GetTableId() *string {
	return s.TableId
}

func (s *GrantUserPermissionRequest) GetTableName() *string {
	return s.TableName
}

func (s *GrantUserPermissionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GrantUserPermissionRequest) GetUserId() *string {
	return s.UserId
}

func (s *GrantUserPermissionRequest) SetDbId(v string) *GrantUserPermissionRequest {
	s.DbId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetDsType(v string) *GrantUserPermissionRequest {
	s.DsType = &v
	return s
}

func (s *GrantUserPermissionRequest) SetExpireDate(v string) *GrantUserPermissionRequest {
	s.ExpireDate = &v
	return s
}

func (s *GrantUserPermissionRequest) SetInstanceId(v int64) *GrantUserPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetLogic(v bool) *GrantUserPermissionRequest {
	s.Logic = &v
	return s
}

func (s *GrantUserPermissionRequest) SetPermTypes(v string) *GrantUserPermissionRequest {
	s.PermTypes = &v
	return s
}

func (s *GrantUserPermissionRequest) SetTableId(v string) *GrantUserPermissionRequest {
	s.TableId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetTableName(v string) *GrantUserPermissionRequest {
	s.TableName = &v
	return s
}

func (s *GrantUserPermissionRequest) SetTid(v int64) *GrantUserPermissionRequest {
	s.Tid = &v
	return s
}

func (s *GrantUserPermissionRequest) SetUserId(v string) *GrantUserPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GrantUserPermissionRequest) Validate() error {
	return dara.Validate(s)
}
