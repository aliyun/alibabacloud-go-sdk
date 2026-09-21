// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlInsightStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *DescribeSqlInsightStatisticRequest
	GetAsc() *bool
	SetConsoleContext(v string) *DescribeSqlInsightStatisticRequest
	GetConsoleContext() *string
	SetDbName(v string) *DescribeSqlInsightStatisticRequest
	GetDbName() *string
	SetDoFillTrend(v bool) *DescribeSqlInsightStatisticRequest
	GetDoFillTrend() *bool
	SetEndTime(v int64) *DescribeSqlInsightStatisticRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSqlInsightStatisticRequestFilters) *DescribeSqlInsightStatisticRequest
	GetFilters() []*DescribeSqlInsightStatisticRequestFilters
	SetInstanceId(v string) *DescribeSqlInsightStatisticRequest
	GetInstanceId() *string
	SetJobId(v string) *DescribeSqlInsightStatisticRequest
	GetJobId() *string
	SetKeyword(v string) *DescribeSqlInsightStatisticRequest
	GetKeyword() *string
	SetNodeId(v string) *DescribeSqlInsightStatisticRequest
	GetNodeId() *string
	SetOrderBy(v string) *DescribeSqlInsightStatisticRequest
	GetOrderBy() *string
	SetPageNo(v int32) *DescribeSqlInsightStatisticRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeSqlInsightStatisticRequest
	GetPageSize() *int32
	SetRole(v string) *DescribeSqlInsightStatisticRequest
	GetRole() *string
	SetSqlType(v string) *DescribeSqlInsightStatisticRequest
	GetSqlType() *string
	SetStartTime(v int64) *DescribeSqlInsightStatisticRequest
	GetStartTime() *int64
	SetTemplateId(v string) *DescribeSqlInsightStatisticRequest
	GetTemplateId() *string
	SetType(v string) *DescribeSqlInsightStatisticRequest
	GetType() *string
}

type DescribeSqlInsightStatisticRequest struct {
	// The sort direction. Default value: **false*	- (descending). Valid values:
	//
	// - **true**: ascending.
	//
	// - **false**: descending.
	//
	// example:
	//
	// false
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database name used for filtering.
	//
	// > In certain aggregation storage pipelines, you can specify multiple database names separated by commas. In other pipelines, only a single database name is supported.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Specifies whether to populate time series trend data for each statistical entry, which corresponds to the **Trend*	- field in the response. Default value: **true**. Valid values:
	//
	// - **true**: Populates trend data.
	//
	// - **false**: Does not populate trend data.
	//
	// > Enabling this option triggers additional queries for each time slice per entry, which significantly increases query overhead. If the trend filling capability is not enabled for the instance, this parameter does not take effect.
	//
	// example:
	//
	// true
	DoFillTrend *bool `json:"DoFillTrend,omitempty" xml:"DoFillTrend,omitempty"`
	// The end time of the query. Specify a UNIX timestamp in milliseconds. The system rounds up to the nearest minute.
	//
	// > The span between this value and **StartTime*	- must not exceed 7 days. If **EndTime*	- is earlier than the time when SQL Explorer was enabled for the instance, an error indicating that the query time is earlier than the available time is returned.
	//
	// > Because data aggregation involves latency, the actual effective value is trimmed to a few minutes before the current time. Data from the most recent minutes may not be available.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1718600000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of structured filter conditions, specified as Key/Value pairs. The POP format is **Filters.N.Key*	- and **Filters.N.Value**, with a maximum of 100 pairs. **Key*	- is case-insensitive. Entries with an empty **Value*	- are ignored. Valid values of **Key*	- for this operation:
	//
	// - **keyWord**: the keyword. The value is split by whitespace into multiple words and takes effect together with the **Keyword*	- parameter.
	//
	// - **hostAddress**: the access source address. Separate multiple values with commas. This takes effect together with **TemplateId*	- when **Type*	- is set to **OriginHost**.
	//
	// - **accountName**: the database username. Separate multiple values with commas.
	//
	// - **dbName**: the database name. Separate multiple values with commas. This takes effect together with the **DbName*	- parameter.
	//
	// - **sqlType**: the SQL type. Separate multiple values with commas. This takes effect together with the **SqlType*	- parameter.
	//
	// - **sqlId**: the SQL template ID. Separate multiple values with commas. This takes effect together with **TemplateId*	- when **Type*	- is set to **SQL**.
	//
	// - **insRole**: the primary/secondary role. Valid values: **master*	- and **slave**. These values are case-sensitive.
	//
	// > Any **Key*	- value other than the preceding values is ignored.
	//
	// > This parameter takes effect only in certain aggregation storage pipelines. In other pipelines, this parameter is entirely ignored.
	Filters []*DescribeSqlInsightStatisticRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The database instance ID.
	//
	// > This operation supports RDS for MySQL, PolarDB for MySQL, PolarDB-X, RDS for PostgreSQL, PolarDB for PostgreSQL, RDS for SQL Server, and Lindorm instances that have SQL Explorer enabled. MongoDB and Redis instances are not supported. Calling this operation for unsupported instances returns an error indicating that the operation is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1q6f9z5xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The idempotency ID for the SQL Explorer data query. This parameter is not required for regular queries.
	//
	// example:
	//
	// 9a4f5c4494dbd6713185d87a97aa53e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The keyword for fuzzy retrieve on SQL template content. Separate multiple keywords with spaces. The system first performs keyword match to find the corresponding SQL templates (up to 1,000 templates), and then performs aggregation statistics based on these templates. If no templates match, an empty list is returned.
	//
	// > This parameter does not take effect when **Type*	- is set to **SQL*	- and **TemplateId*	- is specified.
	//
	// example:
	//
	// t_order
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The node ID of the instance. This parameter narrows the statistical scope to the specified node. Only a single node ID is supported. You cannot specify multiple node IDs separated by commas.
	//
	// > This parameter is required only for instances that consist of multiple nodes, such as PolarDB-X and Lindorm instances. You can ignore this parameter for single-node instances.
	//
	// example:
	//
	// pi-bp1xxxxxxxxxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field used for sorting. If this parameter is not specified or an unsupported value is specified, the results are sorted by **rt*	- (total response time). Valid values:
	//
	// - Response time: **rt**, **avgRt**, **maxRt**, **minRt**, **rtRate**.
	//
	// - Executions: **count**, **countRate**, **timestamp**.
	//
	// - Scan rows: **rowsExamined**, **avgRowsExamined**.
	//
	// - Returned rows: **rowsReturned**, **totalRowsReturned**, **avgRowsReturned**, **maxRowsReturned**, **minRowsReturned**, **maxRowReturned**, **minRowReturned**.
	//
	// - Logical reads: **logicalRead**, **totalLogicalRead**, **avgLogicalRead**, **maxLogicalRead**, **minLogicalRead**.
	//
	// - Physical reads: **physicalRead**, **totalPhysicalRead**, **avgPhysicalRead**, **maxPhysicalRead**, **minPhysicalRead**.
	//
	// - Logical writes (valid only for SQL Server instances): **writes**, **totalWrites**, **avgWrites**, **maxWrites**, **minWrites**.
	//
	// - CPU time (valid only for SQL Server instances): **totalCpuTime**, **avgCpuTime**, **maxCpuTime**, **minCpuTime**.
	//
	// - PolarDB-X compute node metrics (valid only when **Role*	- is set to **polarx_cn**): **scnt**, **avgScnt**, **rows**, **avgRows**, **frows**, **avgFrows**.
	//
	// - Affected rows (valid only for Lindorm instances): **totalAffectRows**, **avgAffectRows**.
	//
	// > **timestamp*	- sorts by the data timestamp, which is a millisecond-level UNIX timestamp.
	//
	// > Only the first letter is case-insensitive. The remaining characters must exactly match the preceding values. For example, **AvgRt*	- is valid but **avgrt*	- is not.
	//
	// example:
	//
	// rt
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of statistical entries per page. Default value: 10. Maximum value: 2000.
	//
	// > A value greater than 2000 returns an InvalidParams error instead of being trimmed.
	//
	// > When aggregating by access source or database user (**Type*	- is set to **FullRequestOrigin*	- or **FullRequestUser**), a value greater than 100 may be reset to 10 in certain aggregation storage pipelines.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role of the instance node. The value is case-insensitive. If this parameter is not specified, the role is automatically resolved from **InstanceId**. Valid values:
	//
	// - **polarx_cn**: PolarDB-X compute node.
	//
	// - **polarx_dn**: PolarDB-X storage node.
	//
	// > This value affects the scope of returned fields. For example, **Scnt**, **Rows**, and **Frows*	- are returned only when the value is **polarx_cn**.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The SQL type used for filtering. Valid values:
	//
	// - **select**
	//
	// - **insert**
	//
	// - **update**
	//
	// - **delete**
	//
	// > Values are lowercase. In certain aggregation storage pipelines, you can specify multiple values separated by commas. In other pipelines, only a single value is supported.
	//
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. Specify a UNIX timestamp in milliseconds. The system rounds down to the nearest minute.
	//
	// > The value must be within the last 30 days. If the value is earlier than the time when SQL Explorer was enabled for the instance, it is automatically adjusted to the time when SQL Explorer was enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1718000000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The identifier of the statistical object. The meaning varies depending on the value of **Type**. Valid values:
	//
	// - When **Type*	- is set to **SQL**: the SQL template ID, which corresponds to **SqlId*	- in the response.
	//
	// - When **Type*	- is set to **OriginHost**: the access source address.
	//
	// - When **Type*	- is set to **User**: the database username.
	//
	// > This parameter does not take effect when **Type*	- is not specified, or is set to **FullRequestOrigin*	- or **FullRequestUser**.
	//
	// > When **Type*	- is set to **SQL**, you can specify multiple template IDs separated by commas. In this case, the **Keyword*	- parameter does not take effect.
	//
	// example:
	//
	// 9f8e7d6c5b4a3210
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The aggregation or filter dimension for statistics. The value is case-insensitive. If this parameter is not specified, statistics are aggregated by SQL template by default. Valid values:
	//
	// - **FullRequestOrigin**: Aggregates by access source address.
	//
	// - **FullRequestUser**: Aggregates by database user.
	//
	// - **SQL**: Filters by SQL template. You must also specify **TemplateId*	- as the SQL template ID.
	//
	// - **OriginHost**: Filters by access source. You must also specify **TemplateId*	- as the source address.
	//
	// - **User**: Filters by database user. You must also specify **TemplateId*	- as the username.
	//
	// > Specifying a value other than the preceding values returns an InvalidParams error.
	//
	// example:
	//
	// SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSqlInsightStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticRequest) GetAsc() *bool {
	return s.Asc
}

func (s *DescribeSqlInsightStatisticRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DescribeSqlInsightStatisticRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeSqlInsightStatisticRequest) GetDoFillTrend() *bool {
	return s.DoFillTrend
}

func (s *DescribeSqlInsightStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSqlInsightStatisticRequest) GetFilters() []*DescribeSqlInsightStatisticRequestFilters {
	return s.Filters
}

func (s *DescribeSqlInsightStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlInsightStatisticRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSqlInsightStatisticRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeSqlInsightStatisticRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSqlInsightStatisticRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeSqlInsightStatisticRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeSqlInsightStatisticRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlInsightStatisticRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeSqlInsightStatisticRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlInsightStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSqlInsightStatisticRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSqlInsightStatisticRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSqlInsightStatisticRequest) SetAsc(v bool) *DescribeSqlInsightStatisticRequest {
	s.Asc = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetConsoleContext(v string) *DescribeSqlInsightStatisticRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetDbName(v string) *DescribeSqlInsightStatisticRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetDoFillTrend(v bool) *DescribeSqlInsightStatisticRequest {
	s.DoFillTrend = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetEndTime(v int64) *DescribeSqlInsightStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetFilters(v []*DescribeSqlInsightStatisticRequestFilters) *DescribeSqlInsightStatisticRequest {
	s.Filters = v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetInstanceId(v string) *DescribeSqlInsightStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetJobId(v string) *DescribeSqlInsightStatisticRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetKeyword(v string) *DescribeSqlInsightStatisticRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetNodeId(v string) *DescribeSqlInsightStatisticRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetOrderBy(v string) *DescribeSqlInsightStatisticRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetPageNo(v int32) *DescribeSqlInsightStatisticRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetPageSize(v int32) *DescribeSqlInsightStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetRole(v string) *DescribeSqlInsightStatisticRequest {
	s.Role = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetSqlType(v string) *DescribeSqlInsightStatisticRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetStartTime(v int64) *DescribeSqlInsightStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetTemplateId(v string) *DescribeSqlInsightStatisticRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) SetType(v string) *DescribeSqlInsightStatisticRequest {
	s.Type = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlInsightStatisticRequestFilters struct {
	// The key of the filter condition. The value is case-insensitive. For valid values, see the description of the **Filters*	- parameter.
	//
	// example:
	//
	// sqlType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition. If the value is null or an empty string, the filter condition is ignored. Separate multiple values with commas. The specific upper limit depends on the corresponding key.
	//
	// example:
	//
	// select
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlInsightStatisticRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSqlInsightStatisticRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSqlInsightStatisticRequestFilters) SetKey(v string) *DescribeSqlInsightStatisticRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequestFilters) SetValue(v string) *DescribeSqlInsightStatisticRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSqlInsightStatisticRequestFilters) Validate() error {
	return dara.Validate(s)
}
