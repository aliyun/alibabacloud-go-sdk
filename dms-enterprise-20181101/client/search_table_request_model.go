// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *SearchTableRequest
	GetDbType() *string
	SetEnvType(v string) *SearchTableRequest
	GetEnvType() *string
	SetPageNumber(v int32) *SearchTableRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTableRequest
	GetPageSize() *int32
	SetReturnGuid(v bool) *SearchTableRequest
	GetReturnGuid() *bool
	SetSearchKey(v string) *SearchTableRequest
	GetSearchKey() *string
	SetSearchRange(v string) *SearchTableRequest
	GetSearchRange() *string
	SetSearchTarget(v string) *SearchTableRequest
	GetSearchTarget() *string
	SetTid(v int64) *SearchTableRequest
	GetTid() *int64
}

type SearchTableRequest struct {
	// The type of database. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// 	- **DRDS**
	//
	// 	- **OceanBase**
	//
	// 	- **Mongo**
	//
	// 	- **Redis**
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which databases belong. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// PRODUCT
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return the GUID of each table.
	//
	// example:
	//
	// false
	ReturnGuid *bool `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
	// The keyword that is used to query tables.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The scope of tables that you want to query. Valid values:
	//
	// 	- **HAS_PERMSSION**: the tables on which the current account has permissions.
	//
	// 	- **OWNER**: the tables owned by the current account.
	//
	// 	- **MY_FOCUS**: the tables that the current account follows.
	//
	// 	- **UNKNOWN**: all tables.
	//
	// example:
	//
	// OWNER
	SearchRange *string `json:"SearchRange,omitempty" xml:"SearchRange,omitempty"`
	// The type of table that you want to query. Valid values:
	//
	// 	- **TABLE**: physical and logical tables
	//
	// 	- **SINGLE_TABLE**: physical tables
	//
	// 	- **LOGIC_TABLE**: logical tables
	//
	// example:
	//
	// LOGIC_TABLE
	SearchTarget *string `json:"SearchTarget,omitempty" xml:"SearchTarget,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SearchTableRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTableRequest) GoString() string {
	return s.String()
}

func (s *SearchTableRequest) GetDbType() *string {
	return s.DbType
}

func (s *SearchTableRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *SearchTableRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTableRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTableRequest) GetReturnGuid() *bool {
	return s.ReturnGuid
}

func (s *SearchTableRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchTableRequest) GetSearchRange() *string {
	return s.SearchRange
}

func (s *SearchTableRequest) GetSearchTarget() *string {
	return s.SearchTarget
}

func (s *SearchTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SearchTableRequest) SetDbType(v string) *SearchTableRequest {
	s.DbType = &v
	return s
}

func (s *SearchTableRequest) SetEnvType(v string) *SearchTableRequest {
	s.EnvType = &v
	return s
}

func (s *SearchTableRequest) SetPageNumber(v int32) *SearchTableRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTableRequest) SetPageSize(v int32) *SearchTableRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTableRequest) SetReturnGuid(v bool) *SearchTableRequest {
	s.ReturnGuid = &v
	return s
}

func (s *SearchTableRequest) SetSearchKey(v string) *SearchTableRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchTableRequest) SetSearchRange(v string) *SearchTableRequest {
	s.SearchRange = &v
	return s
}

func (s *SearchTableRequest) SetSearchTarget(v string) *SearchTableRequest {
	s.SearchTarget = &v
	return s
}

func (s *SearchTableRequest) SetTid(v int64) *SearchTableRequest {
	s.Tid = &v
	return s
}

func (s *SearchTableRequest) Validate() error {
	return dara.Validate(s)
}
