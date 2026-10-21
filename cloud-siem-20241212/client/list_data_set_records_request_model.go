// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetId(v string) *ListDataSetRecordsRequest
	GetDataSetId() *string
	SetFilter(v string) *ListDataSetRecordsRequest
	GetFilter() *string
	SetLang(v string) *ListDataSetRecordsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataSetRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetRecordsRequest
	GetNextToken() *string
	SetOrder(v string) *ListDataSetRecordsRequest
	GetOrder() *string
	SetOrderField(v string) *ListDataSetRecordsRequest
	GetOrderField() *string
	SetPageNumber(v int32) *ListDataSetRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDataSetRecordsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSetRecordsRequest
	GetRoleFor() *int64
}

type ListDataSetRecordsRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset-nhcrmjpx1zsorlaq****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// The JSON string for field fuzzy match. Example:
	//
	// {"field1":"value1","field2":"value2"}
	//
	// example:
	//
	// {"field1":"value1","field2":"value2"}
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return per query when you use the NextToken parameter. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort direction. Valid values:
	//
	// - "desc" (default): Descending order.
	//
	// - "asc": Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field by which to sort. Valid values:
	//
	// - "updatetime" (default): Update time.
	//
	// - "createtime": Creation time.
	//
	// example:
	//
	// updatetime
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center for threat analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataSetRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetRecordsRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetRecordsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDataSetRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSetRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSetRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSetRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataSetRecordsRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDataSetRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSetRecordsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSetRecordsRequest) SetDataSetId(v string) *ListDataSetRecordsRequest {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetFilter(v string) *ListDataSetRecordsRequest {
	s.Filter = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetLang(v string) *ListDataSetRecordsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetMaxResults(v int32) *ListDataSetRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetNextToken(v string) *ListDataSetRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetOrder(v string) *ListDataSetRecordsRequest {
	s.Order = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetOrderField(v string) *ListDataSetRecordsRequest {
	s.OrderField = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetPageNumber(v int32) *ListDataSetRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetPageSize(v int32) *ListDataSetRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetRegionId(v string) *ListDataSetRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSetRecordsRequest) SetRoleFor(v int64) *ListDataSetRecordsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSetRecordsRequest) Validate() error {
	return dara.Validate(s)
}
