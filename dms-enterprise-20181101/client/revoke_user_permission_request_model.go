// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *RevokeUserPermissionRequest
	GetDbId() *string
	SetDsType(v string) *RevokeUserPermissionRequest
	GetDsType() *string
	SetInstanceId(v int64) *RevokeUserPermissionRequest
	GetInstanceId() *int64
	SetLogic(v bool) *RevokeUserPermissionRequest
	GetLogic() *bool
	SetPermTypes(v string) *RevokeUserPermissionRequest
	GetPermTypes() *string
	SetTableId(v string) *RevokeUserPermissionRequest
	GetTableId() *string
	SetTableName(v string) *RevokeUserPermissionRequest
	GetTableName() *string
	SetTid(v int64) *RevokeUserPermissionRequest
	GetTid() *int64
	SetUserAccessId(v string) *RevokeUserPermissionRequest
	GetUserAccessId() *string
	SetUserId(v string) *RevokeUserPermissionRequest
	GetUserId() *string
}

type RevokeUserPermissionRequest struct {
	// The database ID. The database can be a physical database or a logical database.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// example:
	//
	// 1860****
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the object on which you want to revoke permissions from a user. Valid values:
	//
	// 	- **INSTANCE**: instances.
	//
	// 	- **DATABASE**: physical databases.
	//
	// 	- **LOGIC_DATABASE**: logical databases.
	//
	// 	- **TABLE**: physical tables.
	//
	// 	- **LOGIC_TABLE**: logical tables.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATABASE
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The database instance ID. You must specify this parameter if you revoke a permission from the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the ID of the database instance.
	//
	// example:
	//
	// 174****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// >
	//
	// 	- If the database is a logical database, set this parameter to **true**.
	//
	// 	- If the database is a physical database, set this parameter to **false**.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The type of the permissions. Valid values:
	//
	// 	- **QUERY**: query permissions.
	//
	// 	- **EXPORT**: export permissions.
	//
	// 	- **CORRECT**: change permissions.
	//
	// 	- **LOGIN**: logon permissions.
	//
	// 	- **PERF**: query permissions on the performance details of an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// CORRECT
	PermTypes *string `json:"PermTypes,omitempty" xml:"PermTypes,omitempty"`
	// The table ID. You must specify this parameter if you revoke a permission from the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table ID.
	//
	// example:
	//
	// 13****
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table name.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The permission ID. You can call the [ListUserPermission](https://help.aliyun.com/document_detail/146957.html) operation to query the permission ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 774****
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
	// The user ID. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeUserPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionRequest) GetDbId() *string {
	return s.DbId
}

func (s *RevokeUserPermissionRequest) GetDsType() *string {
	return s.DsType
}

func (s *RevokeUserPermissionRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *RevokeUserPermissionRequest) GetLogic() *bool {
	return s.Logic
}

func (s *RevokeUserPermissionRequest) GetPermTypes() *string {
	return s.PermTypes
}

func (s *RevokeUserPermissionRequest) GetTableId() *string {
	return s.TableId
}

func (s *RevokeUserPermissionRequest) GetTableName() *string {
	return s.TableName
}

func (s *RevokeUserPermissionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RevokeUserPermissionRequest) GetUserAccessId() *string {
	return s.UserAccessId
}

func (s *RevokeUserPermissionRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokeUserPermissionRequest) SetDbId(v string) *RevokeUserPermissionRequest {
	s.DbId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetDsType(v string) *RevokeUserPermissionRequest {
	s.DsType = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetInstanceId(v int64) *RevokeUserPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetLogic(v bool) *RevokeUserPermissionRequest {
	s.Logic = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetPermTypes(v string) *RevokeUserPermissionRequest {
	s.PermTypes = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetTableId(v string) *RevokeUserPermissionRequest {
	s.TableId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetTableName(v string) *RevokeUserPermissionRequest {
	s.TableName = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetTid(v int64) *RevokeUserPermissionRequest {
	s.Tid = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetUserAccessId(v string) *RevokeUserPermissionRequest {
	s.UserAccessId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetUserId(v string) *RevokeUserPermissionRequest {
	s.UserId = &v
	return s
}

func (s *RevokeUserPermissionRequest) Validate() error {
	return dara.Validate(s)
}
