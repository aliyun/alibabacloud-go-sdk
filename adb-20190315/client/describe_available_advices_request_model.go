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
	// The date when the suggestion is generated. Specify the date in the yyyyMMdd format. The date must be in UTC.
	//
	// >  Suggestions are generated after analysis after midnight every day. You must specify a date that is at least one day earlier than the current date. For example, if the current date is 20240627, you must specify 20240626 or an earlier date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20221124
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The type of the suggestion. Valid values:
	//
	// 	- **INDEX**: index optimization.
	//
	// 	- **TIERING**: hot and cold data optimization.
	//
	// example:
	//
	// Index
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp198m028ih55****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The keyword that is used to query information by table name.
	//
	// example:
	//
	// you_table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The display language of the suggestion. Default value: zh. Valid values:
	//
	// 	- **zh**: simplified Chinese
	//
	// 	- **en**: English
	//
	// 	- **ja**: Japanese
	//
	// 	- **zh-tw**: traditional Chinese
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON format. Example: `[{"Field":"SchemaName","Type":"Asc"}]`.
	//
	// 	- `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     	- `SchemaName`: the name of the database.
	//
	//     	- `TableName`: the name of the table.
	//
	//     	- `Benefit`: the expected benefits of the applied optimization suggestion.
	//
	// 	- `Type` specifies the sorting order. Valid values:
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// >  If you do not specify this parameter, the query results are sorted in descending order based on the Benefit field.
	//
	// example:
	//
	// [{"Field":"Benefit","Type":"Desc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table in the **DatabaseName.TableName*	- format.
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
