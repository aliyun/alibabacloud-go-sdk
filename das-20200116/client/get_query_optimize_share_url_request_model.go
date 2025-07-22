// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeShareUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GetQueryOptimizeShareUrlRequest
	GetAsc() *bool
	SetDbNames(v string) *GetQueryOptimizeShareUrlRequest
	GetDbNames() *string
	SetEngine(v string) *GetQueryOptimizeShareUrlRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeShareUrlRequest
	GetInstanceIds() *string
	SetKeywords(v string) *GetQueryOptimizeShareUrlRequest
	GetKeywords() *string
	SetLogicalOperator(v string) *GetQueryOptimizeShareUrlRequest
	GetLogicalOperator() *string
	SetOnlyOptimizedSql(v bool) *GetQueryOptimizeShareUrlRequest
	GetOnlyOptimizedSql() *bool
	SetOrderBy(v string) *GetQueryOptimizeShareUrlRequest
	GetOrderBy() *string
	SetPageNo(v int32) *GetQueryOptimizeShareUrlRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetQueryOptimizeShareUrlRequest
	GetPageSize() *int32
	SetRegion(v string) *GetQueryOptimizeShareUrlRequest
	GetRegion() *string
	SetRules(v string) *GetQueryOptimizeShareUrlRequest
	GetRules() *string
	SetSqlIds(v string) *GetQueryOptimizeShareUrlRequest
	GetSqlIds() *string
	SetTagNames(v string) *GetQueryOptimizeShareUrlRequest
	GetTagNames() *string
	SetTime(v int64) *GetQueryOptimizeShareUrlRequest
	GetTime() *int64
	SetUser(v string) *GetQueryOptimizeShareUrlRequest
	GetUser() *string
}

type GetQueryOptimizeShareUrlRequest struct {
	// Specifies whether to sort the returned entries in ascending order. Default value: **true**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The name of the database to be queried.
	//
	// example:
	//
	// testdb01
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL
	//
	// 	- **PolarDBMySQL**: PolarDB for MySQL
	//
	// 	- **PostgreSQL**: ApsaraDB RDS for PostgreSQL
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The keywords of the SQL template. Separate multiple keywords with spaces.
	//
	// example:
	//
	// select update
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The logical relationship between multiple keywords. Valid values:
	//
	// 	- **or**
	//
	// 	- **and**
	//
	// example:
	//
	// or
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// Specifies whether to query only SQL templates that need to be optimized. Default value: **false**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	OnlyOptimizedSql *bool `json:"OnlyOptimizedSql,omitempty" xml:"OnlyOptimizedSql,omitempty"`
	// The field by which to sort the returned entries. Default value: **count**. Valid values:
	//
	// 	- **count**: the number of executions.
	//
	// 	- **maxQueryTime**: the longest execution duration.
	//
	// 	- **avgQueryTime**: the average execution duration.
	//
	// 	- **maxLockTime**: the longest lock wait duration.
	//
	// 	- **avgLockTime**: the average lock wait duration.
	//
	// 	- **maxRowsExamined**: the largest number of scanned rows.
	//
	// 	- **avgRowsExamined**: the average number of scanned rows.
	//
	// 	- **maxRowsSent**: the largest number of returned rows.
	//
	// 	- **avgRowsSent**: the average number of returned rows.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the instance resides. Valid values:
	//
	// 	- **cn-china**: Chinese mainland
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// This parameter takes effect only if **InstanceIds*	- is left empty. If you leave **InstanceIds*	- empty, the system obtains data from the region set by **Region**. By default, Region is set to **cn-china**. If you specify **InstanceIds**, **Region*	- does not take effect and the system obtains data from the region in which the first specified instance resides.****
	//
	// >  If your instances reside in the regions in the Chinese mainland, set this parameter to **cn-china**.
	//
	// example:
	//
	// cn-china
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The tags that are used to filter SQL templates. Separate multiple tags with commas (,). For more information, see [Query governance](https://help.aliyun.com/document_detail/290038.html).
	//
	// example:
	//
	// DAS_NOT_IMPORTANT
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The SQL template IDs. You can call the [GetQueryOptimizeExecErrorStats](https://help.aliyun.com/document_detail/405261.html) operation to obtain the SQL template IDs.
	//
	// example:
	//
	// 6068ce044e3dc9b903979672fb0b69df,d12515c015fc9f41a0778a9e1de0****
	SqlIds *string `json:"SqlIds,omitempty" xml:"SqlIds,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	TagNames *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	// The date to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642953600000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The account of the database to be queried.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetQueryOptimizeShareUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeShareUrlRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeShareUrlRequest) GetAsc() *bool {
	return s.Asc
}

func (s *GetQueryOptimizeShareUrlRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *GetQueryOptimizeShareUrlRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeShareUrlRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeShareUrlRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *GetQueryOptimizeShareUrlRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *GetQueryOptimizeShareUrlRequest) GetOnlyOptimizedSql() *bool {
	return s.OnlyOptimizedSql
}

func (s *GetQueryOptimizeShareUrlRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetQueryOptimizeShareUrlRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeShareUrlRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeShareUrlRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeShareUrlRequest) GetRules() *string {
	return s.Rules
}

func (s *GetQueryOptimizeShareUrlRequest) GetSqlIds() *string {
	return s.SqlIds
}

func (s *GetQueryOptimizeShareUrlRequest) GetTagNames() *string {
	return s.TagNames
}

func (s *GetQueryOptimizeShareUrlRequest) GetTime() *int64 {
	return s.Time
}

func (s *GetQueryOptimizeShareUrlRequest) GetUser() *string {
	return s.User
}

func (s *GetQueryOptimizeShareUrlRequest) SetAsc(v bool) *GetQueryOptimizeShareUrlRequest {
	s.Asc = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetDbNames(v string) *GetQueryOptimizeShareUrlRequest {
	s.DbNames = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetEngine(v string) *GetQueryOptimizeShareUrlRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetInstanceIds(v string) *GetQueryOptimizeShareUrlRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetKeywords(v string) *GetQueryOptimizeShareUrlRequest {
	s.Keywords = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetLogicalOperator(v string) *GetQueryOptimizeShareUrlRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetOnlyOptimizedSql(v bool) *GetQueryOptimizeShareUrlRequest {
	s.OnlyOptimizedSql = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetOrderBy(v string) *GetQueryOptimizeShareUrlRequest {
	s.OrderBy = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetPageNo(v int32) *GetQueryOptimizeShareUrlRequest {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetPageSize(v int32) *GetQueryOptimizeShareUrlRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetRegion(v string) *GetQueryOptimizeShareUrlRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetRules(v string) *GetQueryOptimizeShareUrlRequest {
	s.Rules = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetSqlIds(v string) *GetQueryOptimizeShareUrlRequest {
	s.SqlIds = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetTagNames(v string) *GetQueryOptimizeShareUrlRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetTime(v int64) *GetQueryOptimizeShareUrlRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) SetUser(v string) *GetQueryOptimizeShareUrlRequest {
	s.User = &v
	return s
}

func (s *GetQueryOptimizeShareUrlRequest) Validate() error {
	return dara.Validate(s)
}
