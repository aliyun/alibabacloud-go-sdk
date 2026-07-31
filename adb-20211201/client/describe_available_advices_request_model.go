// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableAdvicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceDate(v int64) *DescribeAvailableAdvicesRequest
	GetAdviceDate() *int64
	SetAdviceType(v string) *DescribeAvailableAdvicesRequest
	GetAdviceType() *string
	SetDBClusterId(v string) *DescribeAvailableAdvicesRequest
	GetDBClusterId() *string
	SetKeyword(v string) *DescribeAvailableAdvicesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeAvailableAdvicesRequest
	GetLang() *string
	SetOrder(v string) *DescribeAvailableAdvicesRequest
	GetOrder() *string
	SetPageNumber(v int64) *DescribeAvailableAdvicesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAvailableAdvicesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeAvailableAdvicesRequest
	GetRegionId() *string
	SetSchemaTableName(v string) *DescribeAvailableAdvicesRequest
	GetSchemaTableName() *string
}

type DescribeAvailableAdvicesRequest struct {
	// The date when the advice was generated, in the `yyyyMMdd` format.
	//
	// > Advice is generated daily. To query for advice, specify a date at least one day before the current date. For example, if you query on June 27, 2024, set this parameter to `20240626` or an earlier date.
	//
	// example:
	//
	// 20221124
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The type of advice. Valid values:
	//
	// - **INDEX**: index optimization.
	//
	// - **TIERING**: hot and cold data tiering.
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
	// am-bp198m028ih55****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The keyword for a fuzzy search on table names.
	//
	// example:
	//
	// you_table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The display language for the advice. Valid values:
	//
	// - **zh**: Simplified Chinese (default).
	//
	// - **en**: English.
	//
	// - **ja**: Japanese.
	//
	// - **zh-tw**: Traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies the sort order for the results. The value is a JSON string. Example: `[{"Field":"SchemaName","Type":"Asc"}]`. The JSON string contains the following key-value pairs:
	//
	// - `Field`: the field to sort by. Valid values:
	//
	//   - `SchemaName`: the database name.
	//
	//   - `TableName`: the table name.
	//
	//   - `Benefit`: the expected benefit.
	//
	// - `Type`: the sort order. Valid values:
	//
	//   - `Asc`: ascending order.
	//
	//   - `Desc`: descending order.
	//
	// > By default, results are sorted by expected benefit in descending order.
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
	// The number of entries per page. Valid values:
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
	// A concatenation of the database name and the table name.
	//
	// example:
	//
	// tpch.lineitem
	SchemaTableName *string `json:"SchemaTableName,omitempty" xml:"SchemaTableName,omitempty"`
}

func (s DescribeAvailableAdvicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesRequest) GetAdviceDate() *int64 {
	return s.AdviceDate
}

func (s *DescribeAvailableAdvicesRequest) GetAdviceType() *string {
	return s.AdviceType
}

func (s *DescribeAvailableAdvicesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAvailableAdvicesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeAvailableAdvicesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAvailableAdvicesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeAvailableAdvicesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAvailableAdvicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAvailableAdvicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableAdvicesRequest) GetSchemaTableName() *string {
	return s.SchemaTableName
}

func (s *DescribeAvailableAdvicesRequest) SetAdviceDate(v int64) *DescribeAvailableAdvicesRequest {
	s.AdviceDate = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetAdviceType(v string) *DescribeAvailableAdvicesRequest {
	s.AdviceType = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetDBClusterId(v string) *DescribeAvailableAdvicesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetKeyword(v string) *DescribeAvailableAdvicesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetLang(v string) *DescribeAvailableAdvicesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetOrder(v string) *DescribeAvailableAdvicesRequest {
	s.Order = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetPageNumber(v int64) *DescribeAvailableAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetPageSize(v int64) *DescribeAvailableAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetRegionId(v string) *DescribeAvailableAdvicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetSchemaTableName(v string) *DescribeAvailableAdvicesRequest {
	s.SchemaTableName = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) Validate() error {
	return dara.Validate(s)
}
