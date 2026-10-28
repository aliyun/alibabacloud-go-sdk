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
	// The sort order. Default value: **false*	- (descending order). Valid values:
	//
	// - **true**: ascending order.
	//
	// - **false**: descending order.
	//
	// example:
	//
	// false
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// null
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The name of the database used to filter the results.
	//
	// > In the aggregation storage links of some instances, you can specify multiple database names separated by commas. In other links, you can specify only a single database name.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Specifies whether to populate time series trend data for each statistical result, which corresponds to the **Trend*	- field in the response. Default value: **true**. Valid values:
	//
	// - **true**: populates trend data.
	//
	// - **false**: does not populate trend data.
	//
	// > If this parameter is set to true, the system additionally queries data for each time slice for each result, which significantly increases the query overhead. If the trend filling capability is not enabled for the instance, this parameter does not take effect.
	//
	// example:
	//
	// true
	DoFillTrend *bool `json:"DoFillTrend,omitempty" xml:"DoFillTrend,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in milliseconds. The system rounds up the value to the nearest minute.
	//
	// > The time span between **EndTime*	- and **StartTime*	- cannot exceed 7 days. If **EndTime*	- is earlier than the time when the SQL Explorer feature was enabled for the instance, an error indicating that the query time is earlier than the available time is returned.
	//
	// > Due to delays in data aggregation, the actual effective value is truncated to a few minutes before the current time. Therefore, data from the most recent few minutes may not be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1718600000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of structured filter conditions. The conditions are passed as key-value pairs. The POP format is **Filters.N.Key*	- and **Filters.N.Value**. A maximum of 100 pairs can be specified. **Key*	- is not case-sensitive. Entries with an empty **Value*	- are ignored. Valid values of **Key*	- that take effect in this operation:
	//
	// - **keyWord**: the keywords. The keywords are split into multiple words by whitespace characters and take effect together with the **Keyword*	- parameter.
	//
	// - **hostAddress**: the access source addresses. Separate multiple addresses with commas. This key takes effect together with **TemplateId*	- when **Type*	- is set to **OriginHost**.
	//
	// - **accountName**: the database usernames. Separate multiple usernames with commas.
	//
	// - **dbName**: the database names. Separate multiple names with commas. This key takes effect together with the **DbName*	- parameter.
	//
	// - **sqlType**: the SQL statement types. Separate multiple types with commas. This key takes effect together with the **SqlType*	- parameter.
	//
	// - **sqlId**: the SQL template IDs. Separate multiple IDs with commas. This key takes effect together with **TemplateId*	- when **Type*	- is set to **SQL**.
	//
	// - **insRole**: the primary/secondary role. Valid values are **master*	- and **slave**. This value is case-sensitive.
	//
	// > If you specify a **Key*	- other than the preceding values, the key is ignored.
	//
	// > This parameter takes effect only in the aggregation storage links of some instances. In other links, this parameter is completely ignored.
	Filters []*DescribeSqlInsightStatisticRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The database instance ID.
	//
	// > This operation supports ApsaraDB RDS for MySQL, PolarDB for MySQL, PolarDB-X, ApsaraDB RDS for PostgreSQL, PolarDB for PostgreSQL, ApsaraDB RDS for SQL Server, and Lindorm instances that have the SQL Explorer feature enabled. MongoDB and Redis instances are not supported. If you invoke this operation for unsupported instances, a fault indicating that the operation is not supported is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1q6f9z5xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The idempotent ID for querying SQL Explorer data. You do not need to specify this parameter for common queries.
	//
	// example:
	//
	// 9a4f5c4494dbd6713185d87a97aa53e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The keywords used to retrieve SQL templates by fuzzy match based on the SQL template content. Separate multiple keywords with spaces. The system first performs keyword match to retrieve up to 1,000 corresponding SQL templates, and then performs aggregation statistics based on these templates. If no matching templates are found, an empty list is returned.
	//
	// > This parameter does not take effect when **Type*	- is set to **SQL*	- and **TemplateId*	- is specified.
	//
	// example:
	//
	// t_order
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the node in the instance. This parameter is used to narrow the statistics scope to the specified node. Only a single node ID is supported. Multiple node IDs separated by commas are not supported.
	//
	// > This parameter is required only for instances that consist of multiple nodes, such as PolarDB-X and Lindorm instances. You can ignore this parameter for single-node instances.
	//
	// example:
	//
	// pi-bp1xxxxxxxxxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field by which to sort the results. If this parameter is not specified or an unsupported value is specified, the results are sorted by **rt*	- (total response time). Valid values:
	//
	// - Response time: **rt**, **avgRt**, **maxRt**, **minRt**, and **rtRate**.
	//
	// - Executions: **count**, **countRate**, and **timestamp**.
	//
	// - Scanned rows: **rowsExamined*	- and **avgRowsExamined**.
	//
	// - Returned rows: **rowsReturned**, **totalRowsReturned**, **avgRowsReturned**, **maxRowsReturned**, **minRowsReturned**, **maxRowReturned**, and **minRowReturned**.
	//
	// - Logical reads: **logicalRead**, **totalLogicalRead**, **avgLogicalRead**, **maxLogicalRead**, and **minLogicalRead**.
	//
	// - Physical reads: **physicalRead**, **totalPhysicalRead**, **avgPhysicalRead**, **maxPhysicalRead**, and **minPhysicalRead**.
	//
	// - Logical writes (valid only for SQL Server instances): **writes**, **totalWrites**, **avgWrites**, **maxWrites**, and **minWrites**.
	//
	// - CPU time (valid only for SQL Server instances): **totalCpuTime**, **avgCpuTime**, **maxCpuTime**, and **minCpuTime**.
	//
	// - PolarDB-X compute nodes (valid only when **Role*	- is set to **polarx_cn**): **scnt**, **avgScnt**, **rows**, **avgRows**, **frows**, and **avgFrows**.
	//
	// - Affected rows (valid only for Lindorm instances): **totalAffectRows*	- and **avgAffectRows**.
	//
	// > **timestamp*	- indicates that the results are sorted by data timestamp. The timestamp is a UNIX timestamp in milliseconds.
	//
	// > Only the first letter is case-insensitive. The remaining characters must exactly match the preceding values. For example, **AvgRt*	- is valid, but **avgrt*	- is invalid.
	//
	// example:
	//
	// rt
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. If this parameter is not specified, the default value 1 is used.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of statistical entries to return on each page. If this parameter is not specified, the default value 10 is used. The maximum value is 2000.
	//
	// > If the value exceeds 2000, an InvalidParams error is returned instead of being truncated.
	//
	// > When statistics are aggregated by access source or database user (when **Type*	- is set to **FullRequestOrigin*	- or **FullRequestUser**), a value greater than 100 is reset to 10 in some aggregation storage links.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role of the instance node. The value is not case-sensitive. If this parameter is not specified, the role is automatically parsed based on **InstanceId**. Valid values:
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
	// The SQL statement type used to filter the results. Valid values:
	//
	// - **select**
	//
	// - **insert**
	//
	// - **update**
	//
	// - **delete**
	//
	// > The values must be in lowercase. In the aggregation storage links of some instances, you can specify multiple values separated by commas. In other links, you can specify only a single value.
	//
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in milliseconds. The system rounds down the value to the nearest minute.
	//
	// > The start time cannot be more than 30 days before the current time. If the specified start time is earlier than the time when the SQL Explorer feature was enabled for the instance, the start time is automatically adjusted to the enabling time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1718000000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The identifier of the statistics object. The meaning varies based on the value of **Type**. Valid values:
	//
	// - When **Type*	- is set to **SQL**: specify the SQL template ID, which is the **SqlId*	- in the response.
	//
	// - When **Type*	- is set to **OriginHost**: specify the access source address.
	//
	// - When **Type*	- is set to **User**: specify the database username.
	//
	// > This parameter does not take effect when **Type*	- is empty, **FullRequestOrigin**, or **FullRequestUser**.
	//
	// > When **Type*	- is set to **SQL**, you can specify multiple template IDs separated by commas. In this case, **Keyword*	- does not take effect.
	//
	// example:
	//
	// 9f8e7d6c5b4a3210
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The aggregation or filter dimension for statistics. The value is not case-sensitive. If this parameter is not specified, the statistics are aggregated by SQL template by default. Valid values:
	//
	// - **FullRequestOrigin**: aggregates statistics by access source address.
	//
	// - **FullRequestUser**: aggregates statistics by database user.
	//
	// - **SQL**: filters statistics by SQL template. You must also specify **TemplateId*	- as the SQL template ID.
	//
	// - **OriginHost**: filters statistics by access source. You must also specify **TemplateId*	- as the source address.
	//
	// - **User**: filters statistics by database user. You must also specify **TemplateId*	- as the username.
	//
	// > If you specify a value other than the preceding values, an InvalidParams error is returned.
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
	// The key of the conditional filter. The key is not case-sensitive. For active values, see the metric description of the **Filters*	- parameter.
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
