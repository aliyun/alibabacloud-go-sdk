// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppliedAdvicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceType(v string) *DescribeAppliedAdvicesRequest
	GetAdviceType() *string
	SetDBClusterId(v string) *DescribeAppliedAdvicesRequest
	GetDBClusterId() *string
	SetEndTime(v int64) *DescribeAppliedAdvicesRequest
	GetEndTime() *int64
	SetKeyword(v string) *DescribeAppliedAdvicesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeAppliedAdvicesRequest
	GetLang() *string
	SetOrder(v string) *DescribeAppliedAdvicesRequest
	GetOrder() *string
	SetPageNumber(v int64) *DescribeAppliedAdvicesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAppliedAdvicesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeAppliedAdvicesRequest
	GetRegionId() *string
	SetSchemaTableName(v string) *DescribeAppliedAdvicesRequest
	GetSchemaTableName() *string
	SetStartTime(v int64) *DescribeAppliedAdvicesRequest
	GetStartTime() *int64
}

type DescribeAppliedAdvicesRequest struct {
	// The type of the advice. Valid values:
	//
	// - **INDEX**: index optimization
	//
	// - **TIERING**: hot/cold data optimization
	//
	// example:
	//
	// INDEX
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end date of the query. The date is in the `yyyyMMdd` format.
	//
	// example:
	//
	// 20220824
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword for the query. Fuzzy match by table name is supported.
	//
	// example:
	//
	// you_table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the query results. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// - **ja**: Japanese
	//
	// - **zh-tw**: Traditional Chinese
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort the query results. The value is a JSON string. Example: `[{"Field":"SchemaName","Type":"Asc"}]`. Fields:
	//
	// - `Field`: The field by which to sort the results. Valid values:
	//
	//   - `SchemaName`: the database name
	//
	//   - `TableName`: the table name
	//
	//   - `JobStatus`: the status of the build job for the table
	//
	//   - `SubmitTime`: the time when the advice was submitted
	//
	//   - `Benefit`: the estimated benefit
	//
	// - `Type`: The sort order. Valid values:
	//
	//   - `Asc`: ascending
	//
	//   - `Desc`: descending
	//
	// > If you do not set this parameter, the query results are sorted by advice submission time in descending order.
	//
	// example:
	//
	// [{\\"Field\\":\\"AdviceType\\",\\"Type\\":\\"Desc\\"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database and table. Format: **database.table**.
	//
	// example:
	//
	// tpch.lineitem
	SchemaTableName *string `json:"SchemaTableName,omitempty" xml:"SchemaTableName,omitempty"`
	// The start date of the query. The date is in the `yyyyMMdd` format.
	//
	// example:
	//
	// 20220811
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAppliedAdvicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppliedAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesRequest) GetAdviceType() *string {
	return s.AdviceType
}

func (s *DescribeAppliedAdvicesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAppliedAdvicesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAppliedAdvicesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeAppliedAdvicesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAppliedAdvicesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeAppliedAdvicesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAppliedAdvicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAppliedAdvicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppliedAdvicesRequest) GetSchemaTableName() *string {
	return s.SchemaTableName
}

func (s *DescribeAppliedAdvicesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAppliedAdvicesRequest) SetAdviceType(v string) *DescribeAppliedAdvicesRequest {
	s.AdviceType = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetDBClusterId(v string) *DescribeAppliedAdvicesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetEndTime(v int64) *DescribeAppliedAdvicesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetKeyword(v string) *DescribeAppliedAdvicesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetLang(v string) *DescribeAppliedAdvicesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetOrder(v string) *DescribeAppliedAdvicesRequest {
	s.Order = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetPageNumber(v int64) *DescribeAppliedAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetPageSize(v int64) *DescribeAppliedAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetRegionId(v string) *DescribeAppliedAdvicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetSchemaTableName(v string) *DescribeAppliedAdvicesRequest {
	s.SchemaTableName = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetStartTime(v int64) *DescribeAppliedAdvicesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) Validate() error {
	return dara.Validate(s)
}
