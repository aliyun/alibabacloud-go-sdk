// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *ListUserPermissionsRequest
	GetDatabaseName() *string
	SetDbType(v string) *ListUserPermissionsRequest
	GetDbType() *string
	SetEnvType(v string) *ListUserPermissionsRequest
	GetEnvType() *string
	SetLogic(v bool) *ListUserPermissionsRequest
	GetLogic() *bool
	SetPageNumber(v int32) *ListUserPermissionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserPermissionsRequest
	GetPageSize() *int32
	SetPermType(v string) *ListUserPermissionsRequest
	GetPermType() *string
	SetSearchKey(v string) *ListUserPermissionsRequest
	GetSearchKey() *string
	SetTid(v int64) *ListUserPermissionsRequest
	GetTid() *int64
	SetUserId(v string) *ListUserPermissionsRequest
	GetUserId() *string
}

type ListUserPermissionsRequest struct {
	// The name of the database.
	//
	// example:
	//
	// db_name
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// polardb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- product: production environment
	//
	// 	- dev: development environment
	//
	// 	- pre: staging environment
	//
	// 	- test: test environment
	//
	// 	- sit: SIT environment
	//
	// 	- uat: user acceptance testing (UAT) environment
	//
	// 	- pet: stress testing environment
	//
	// 	- stag: STAG environment
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- true: The database is a logical database.
	//
	// 	- false: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The permissions on a specific type of resources that you want to query. Valid values:
	//
	// 	- DATABASE: permissions on databases
	//
	// 	- TABLE: permissions on tables
	//
	// 	- COLUMN: permissions on fields
	//
	// 	- INSTANCE: permissions on instances
	//
	// This parameter is required.
	//
	// example:
	//
	// DATABASE
	PermType *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	// The keyword used in the query. For example, if you want to query permissions on an instance, you can specify the endpoint of the instance, such as rm-bp144d5ky4l4r****.
	//
	// example:
	//
	// rm-bp144d5ky4l4r****
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
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

func (s ListUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListUserPermissionsRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListUserPermissionsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUserPermissionsRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListUserPermissionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserPermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserPermissionsRequest) GetPermType() *string {
	return s.PermType
}

func (s *ListUserPermissionsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListUserPermissionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListUserPermissionsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserPermissionsRequest) SetDatabaseName(v string) *ListUserPermissionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListUserPermissionsRequest) SetDbType(v string) *ListUserPermissionsRequest {
	s.DbType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetEnvType(v string) *ListUserPermissionsRequest {
	s.EnvType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetLogic(v bool) *ListUserPermissionsRequest {
	s.Logic = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPageNumber(v int32) *ListUserPermissionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPageSize(v int32) *ListUserPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPermType(v string) *ListUserPermissionsRequest {
	s.PermType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetSearchKey(v string) *ListUserPermissionsRequest {
	s.SearchKey = &v
	return s
}

func (s *ListUserPermissionsRequest) SetTid(v int64) *ListUserPermissionsRequest {
	s.Tid = &v
	return s
}

func (s *ListUserPermissionsRequest) SetUserId(v string) *ListUserPermissionsRequest {
	s.UserId = &v
	return s
}

func (s *ListUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
