// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseUserPermssionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *ListDatabaseUserPermssionsRequest
	GetDbId() *string
	SetLogic(v bool) *ListDatabaseUserPermssionsRequest
	GetLogic() *bool
	SetPageNumber(v int32) *ListDatabaseUserPermssionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatabaseUserPermssionsRequest
	GetPageSize() *int32
	SetPermType(v string) *ListDatabaseUserPermssionsRequest
	GetPermType() *string
	SetTid(v int64) *ListDatabaseUserPermssionsRequest
	GetTid() *int64
	SetUserName(v string) *ListDatabaseUserPermssionsRequest
	GetUserName() *string
}

type ListDatabaseUserPermssionsRequest struct {
	// The ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database.
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the permission. Valid values:
	//
	// 	- DATABASE: permissions on databases
	//
	// 	- TABLE: permissions on tables
	//
	// 	- COLUMN: permissions on fields
	//
	// This parameter is required.
	//
	// example:
	//
	// DATABASE
	PermType *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// search_user_name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDatabaseUserPermssionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsRequest) GetDbId() *string {
	return s.DbId
}

func (s *ListDatabaseUserPermssionsRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListDatabaseUserPermssionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatabaseUserPermssionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatabaseUserPermssionsRequest) GetPermType() *string {
	return s.PermType
}

func (s *ListDatabaseUserPermssionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDatabaseUserPermssionsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListDatabaseUserPermssionsRequest) SetDbId(v string) *ListDatabaseUserPermssionsRequest {
	s.DbId = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetLogic(v bool) *ListDatabaseUserPermssionsRequest {
	s.Logic = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPageNumber(v int32) *ListDatabaseUserPermssionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPageSize(v int32) *ListDatabaseUserPermssionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPermType(v string) *ListDatabaseUserPermssionsRequest {
	s.PermType = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetTid(v int64) *ListDatabaseUserPermssionsRequest {
	s.Tid = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetUserName(v string) *ListDatabaseUserPermssionsRequest {
	s.UserName = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) Validate() error {
	return dara.Validate(s)
}
