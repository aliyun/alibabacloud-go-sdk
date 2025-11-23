// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *ListAuthorizedUsersForDatabaseRequest
	GetDbId() *string
	SetLogic(v bool) *ListAuthorizedUsersForDatabaseRequest
	GetLogic() *bool
	SetPageNumber(v string) *ListAuthorizedUsersForDatabaseRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedUsersForDatabaseRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedUsersForDatabaseRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedUsersForDatabaseRequest
	GetTid() *int64
}

type ListAuthorizedUsersForDatabaseRequest struct {
	// The database ID. The database can be a physical database or a logical database.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 135***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used for the search.
	//
	// example:
	//
	// poc_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAuthorizedUsersForDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetDbId() *string {
	return s.DbId
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetDbId(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.DbId = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetLogic(v bool) *ListAuthorizedUsersForDatabaseRequest {
	s.Logic = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetPageNumber(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetPageSize(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetSearchKey(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetTid(v int64) *ListAuthorizedUsersForDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
