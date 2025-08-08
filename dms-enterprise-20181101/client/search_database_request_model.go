// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *SearchDatabaseRequest
	GetDbType() *string
	SetEnvType(v string) *SearchDatabaseRequest
	GetEnvType() *string
	SetPageNumber(v int32) *SearchDatabaseRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchDatabaseRequest
	GetPageSize() *int32
	SetRealLoginUserUid(v string) *SearchDatabaseRequest
	GetRealLoginUserUid() *string
	SetSearchKey(v string) *SearchDatabaseRequest
	GetSearchKey() *string
	SetSearchRange(v string) *SearchDatabaseRequest
	GetSearchRange() *string
	SetSearchTarget(v string) *SearchDatabaseRequest
	GetSearchTarget() *string
	SetTid(v int64) *SearchDatabaseRequest
	GetTid() *int64
}

type SearchDatabaseRequest struct {
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The environment type of the database. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
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
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The keyword that is used to search for databases.
	//
	// example:
	//
	// testdb
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The query range based on permissions. Valid values:
	//
	// 	- **HAS_PERMSSION**: searches for databases on which the current user has permissions.
	//
	// 	- **OWNER**: searches for databases owned by the current user.
	//
	// 	- **MY_FOCUS**: searches for databases that the current user follows.
	//
	// 	- **UNKNOWN**: searches for all databases.
	//
	// example:
	//
	// HAS_PERMSSION
	SearchRange *string `json:"SearchRange,omitempty" xml:"SearchRange,omitempty"`
	// The category of the database. Valid values:
	//
	// 	- **DB**: single database or logical database.
	//
	// 	- **SINGLE_DB**: single database.
	//
	// 	- **LOGIC_DB**: logical database.
	//
	// example:
	//
	// SINGLE_DB
	SearchTarget *string `json:"SearchTarget,omitempty" xml:"SearchTarget,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SearchDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseRequest) GoString() string {
	return s.String()
}

func (s *SearchDatabaseRequest) GetDbType() *string {
	return s.DbType
}

func (s *SearchDatabaseRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *SearchDatabaseRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchDatabaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchDatabaseRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *SearchDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchDatabaseRequest) GetSearchRange() *string {
	return s.SearchRange
}

func (s *SearchDatabaseRequest) GetSearchTarget() *string {
	return s.SearchTarget
}

func (s *SearchDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SearchDatabaseRequest) SetDbType(v string) *SearchDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *SearchDatabaseRequest) SetEnvType(v string) *SearchDatabaseRequest {
	s.EnvType = &v
	return s
}

func (s *SearchDatabaseRequest) SetPageNumber(v int32) *SearchDatabaseRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDatabaseRequest) SetPageSize(v int32) *SearchDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDatabaseRequest) SetRealLoginUserUid(v string) *SearchDatabaseRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchKey(v string) *SearchDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchRange(v string) *SearchDatabaseRequest {
	s.SearchRange = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchTarget(v string) *SearchDatabaseRequest {
	s.SearchTarget = &v
	return s
}

func (s *SearchDatabaseRequest) SetTid(v int64) *SearchDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *SearchDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
