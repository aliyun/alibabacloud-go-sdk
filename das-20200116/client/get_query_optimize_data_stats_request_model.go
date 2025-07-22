// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v string) *GetQueryOptimizeDataStatsRequest
	GetAsc() *string
	SetDbNames(v string) *GetQueryOptimizeDataStatsRequest
	GetDbNames() *string
	SetEngine(v string) *GetQueryOptimizeDataStatsRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeDataStatsRequest
	GetInstanceIds() *string
	SetKeywords(v string) *GetQueryOptimizeDataStatsRequest
	GetKeywords() *string
	SetLogicalOperator(v string) *GetQueryOptimizeDataStatsRequest
	GetLogicalOperator() *string
	SetOnlyOptimizedSql(v string) *GetQueryOptimizeDataStatsRequest
	GetOnlyOptimizedSql() *string
	SetOrderBy(v string) *GetQueryOptimizeDataStatsRequest
	GetOrderBy() *string
	SetPageNo(v string) *GetQueryOptimizeDataStatsRequest
	GetPageNo() *string
	SetPageSize(v string) *GetQueryOptimizeDataStatsRequest
	GetPageSize() *string
	SetRegion(v string) *GetQueryOptimizeDataStatsRequest
	GetRegion() *string
	SetRules(v string) *GetQueryOptimizeDataStatsRequest
	GetRules() *string
	SetSqlIds(v string) *GetQueryOptimizeDataStatsRequest
	GetSqlIds() *string
	SetTagNames(v string) *GetQueryOptimizeDataStatsRequest
	GetTagNames() *string
	SetTime(v string) *GetQueryOptimizeDataStatsRequest
	GetTime() *string
	SetUser(v string) *GetQueryOptimizeDataStatsRequest
	GetUser() *string
}

type GetQueryOptimizeDataStatsRequest struct {
	// Specifies whether to sort the returned entries in ascending order. Default value: **true**. Valid values:
	//
	// 	- **true**: sorts the returned entries in ascending order.
	//
	// 	- **false**: does not sort the returned entries in ascending order.
	//
	// example:
	//
	// true
	Asc *string `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The name of the database to be queried.
	//
	// example:
	//
	// testdb01
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PostgreSQL**
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
	// rm-2ze8g2am97624****
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
	// 	- **true**: queries only SQL templates that need to be optimized.
	//
	// 	- **false**: does not query only SQL statements that need to be optimized.
	//
	// example:
	//
	// false
	OnlyOptimizedSql *string `json:"OnlyOptimizedSql,omitempty" xml:"OnlyOptimizedSql,omitempty"`
	// The field by which to sort the returned entries. Default value: **count**. Valid values:
	//
	// 	- **count**: the number of executions.
	//
	// 	- **maxQueryTime**: the longest execution time.
	//
	// 	- **avgQueryTime**: the average execution time.
	//
	// 	- **maxLockTime**: the longest lock wait time.
	//
	// 	- **avgLockTime**: the longest lock wait time.
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
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// >  Set this parameter to **cn-china*	- for the instances that are created in the regions in the Chinese mainland.
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
	// The SQL template ID. You can query the ID of a template by calling the [GetQueryOptimizeExecErrorStats](https://help.aliyun.com/document_detail/405235.html) operation.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlIds *string `json:"SqlIds,omitempty" xml:"SqlIds,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	TagNames *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	// The time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642953600000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The account of the database to be queried.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetQueryOptimizeDataStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsRequest) GetAsc() *string {
	return s.Asc
}

func (s *GetQueryOptimizeDataStatsRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *GetQueryOptimizeDataStatsRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeDataStatsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeDataStatsRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *GetQueryOptimizeDataStatsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *GetQueryOptimizeDataStatsRequest) GetOnlyOptimizedSql() *string {
	return s.OnlyOptimizedSql
}

func (s *GetQueryOptimizeDataStatsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetQueryOptimizeDataStatsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetQueryOptimizeDataStatsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetQueryOptimizeDataStatsRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeDataStatsRequest) GetRules() *string {
	return s.Rules
}

func (s *GetQueryOptimizeDataStatsRequest) GetSqlIds() *string {
	return s.SqlIds
}

func (s *GetQueryOptimizeDataStatsRequest) GetTagNames() *string {
	return s.TagNames
}

func (s *GetQueryOptimizeDataStatsRequest) GetTime() *string {
	return s.Time
}

func (s *GetQueryOptimizeDataStatsRequest) GetUser() *string {
	return s.User
}

func (s *GetQueryOptimizeDataStatsRequest) SetAsc(v string) *GetQueryOptimizeDataStatsRequest {
	s.Asc = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetDbNames(v string) *GetQueryOptimizeDataStatsRequest {
	s.DbNames = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetEngine(v string) *GetQueryOptimizeDataStatsRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetInstanceIds(v string) *GetQueryOptimizeDataStatsRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetKeywords(v string) *GetQueryOptimizeDataStatsRequest {
	s.Keywords = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetLogicalOperator(v string) *GetQueryOptimizeDataStatsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetOnlyOptimizedSql(v string) *GetQueryOptimizeDataStatsRequest {
	s.OnlyOptimizedSql = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetOrderBy(v string) *GetQueryOptimizeDataStatsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetPageNo(v string) *GetQueryOptimizeDataStatsRequest {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetPageSize(v string) *GetQueryOptimizeDataStatsRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetRegion(v string) *GetQueryOptimizeDataStatsRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetRules(v string) *GetQueryOptimizeDataStatsRequest {
	s.Rules = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetSqlIds(v string) *GetQueryOptimizeDataStatsRequest {
	s.SqlIds = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetTagNames(v string) *GetQueryOptimizeDataStatsRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetTime(v string) *GetQueryOptimizeDataStatsRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetUser(v string) *GetQueryOptimizeDataStatsRequest {
	s.User = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) Validate() error {
	return dara.Validate(s)
}
