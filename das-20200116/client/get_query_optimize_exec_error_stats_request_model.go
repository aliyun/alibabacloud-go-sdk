// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetAsc() *string
	SetDbNames(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetDbNames() *string
	SetEngine(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetEngine() *string
	SetInstanceIds(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetInstanceIds() *string
	SetKeywords(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetKeywords() *string
	SetLogicalOperator(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetLogicalOperator() *string
	SetOrderBy(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetOrderBy() *string
	SetPageNo(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetPageNo() *string
	SetPageSize(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetPageSize() *string
	SetRegion(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetRegion() *string
	SetTime(v string) *GetQueryOptimizeExecErrorStatsRequest
	GetTime() *string
}

type GetQueryOptimizeExecErrorStatsRequest struct {
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
	// The field by which to sort the returned entries. Only error_count is supported, which specifies that the entries are sorted based on the number of failed executions.
	//
	// example:
	//
	// error_count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
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
	// The time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642953600000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetAsc() *string {
	return s.Asc
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQueryOptimizeExecErrorStatsRequest) GetTime() *string {
	return s.Time
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetAsc(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Asc = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetDbNames(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.DbNames = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetEngine(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetInstanceIds(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetKeywords(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Keywords = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetLogicalOperator(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetOrderBy(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetPageNo(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetPageSize(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetRegion(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Region = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetTime(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) Validate() error {
	return dara.Validate(s)
}
