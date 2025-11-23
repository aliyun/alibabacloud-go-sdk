// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserOwnedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *ListUserOwnedResourcesRequest
	GetDatabaseName() *string
	SetDbType(v string) *ListUserOwnedResourcesRequest
	GetDbType() *string
	SetEnvType(v string) *ListUserOwnedResourcesRequest
	GetEnvType() *string
	SetLogic(v bool) *ListUserOwnedResourcesRequest
	GetLogic() *bool
	SetOwnerType(v string) *ListUserOwnedResourcesRequest
	GetOwnerType() *string
	SetPageNumber(v int32) *ListUserOwnedResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserOwnedResourcesRequest
	GetPageSize() *int32
	SetTid(v int64) *ListUserOwnedResourcesRequest
	GetTid() *int64
	SetUserId(v string) *ListUserOwnedResourcesRequest
	GetUserId() *string
}

type ListUserOwnedResourcesRequest struct {
	// The database name.
	//
	// example:
	//
	// db_name
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the database instance. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- **product**: production environment.
	//
	// 	- **dev**: development environment.
	//
	// 	- **pre**: pre-release environment.
	//
	// 	- **test**: test environment.
	//
	// 	- **sit**: system integration testing (SIT) environment.
	//
	// 	- **uat**: user acceptance testing (UAT) environment.
	//
	// 	- **pet**: stress testing environment.
	//
	// 	- **stag**: staging environment.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true.**: The database is a logical database
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The type of the owner. Valid values:
	//
	// 	- INSTANCE: an owner of an instance.
	//
	// 	- DATABASE: an owner of a physical database.
	//
	// 	- TABLE: an owner of a physical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserOwnedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserOwnedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserOwnedResourcesRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListUserOwnedResourcesRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListUserOwnedResourcesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUserOwnedResourcesRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListUserOwnedResourcesRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ListUserOwnedResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserOwnedResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserOwnedResourcesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListUserOwnedResourcesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserOwnedResourcesRequest) SetDatabaseName(v string) *ListUserOwnedResourcesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetDbType(v string) *ListUserOwnedResourcesRequest {
	s.DbType = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetEnvType(v string) *ListUserOwnedResourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetLogic(v bool) *ListUserOwnedResourcesRequest {
	s.Logic = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetOwnerType(v string) *ListUserOwnedResourcesRequest {
	s.OwnerType = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetPageNumber(v int32) *ListUserOwnedResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetPageSize(v int32) *ListUserOwnedResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetTid(v int64) *ListUserOwnedResourcesRequest {
	s.Tid = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) SetUserId(v string) *ListUserOwnedResourcesRequest {
	s.UserId = &v
	return s
}

func (s *ListUserOwnedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
