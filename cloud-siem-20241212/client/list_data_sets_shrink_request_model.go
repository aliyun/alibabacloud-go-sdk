// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetId(v string) *ListDataSetsShrinkRequest
	GetDataSetId() *string
	SetDataSetIdsShrink(v string) *ListDataSetsShrinkRequest
	GetDataSetIdsShrink() *string
	SetDataSetName(v string) *ListDataSetsShrinkRequest
	GetDataSetName() *string
	SetDataSetStatus(v int32) *ListDataSetsShrinkRequest
	GetDataSetStatus() *int32
	SetDataSetType(v string) *ListDataSetsShrinkRequest
	GetDataSetType() *string
	SetLang(v string) *ListDataSetsShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataSetsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetsShrinkRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListDataSetsShrinkRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListDataSetsShrinkRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListDataSetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetsShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDataSetsShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSetsShrinkRequest
	GetRoleFor() *int64
}

type ListDataSetsShrinkRequest struct {
	// The dataset ID.
	//
	// example:
	//
	// dataset-qt0n8246gs9nackg****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// The list of dataset IDs.
	DataSetIdsShrink *string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// lmftest
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// The dataset status. Valid values:
	//
	// - 0: deleted.
	//
	// - 1: enabled.
	//
	// example:
	//
	// 0
	DataSetStatus *int32 `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	// The dataset type. Valid values:
	//
	// - custom: custom dataset.
	//
	// - preset: predefined dataset.
	//
	// example:
	//
	// custom
	DataSetType *string `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// This parameter is required in practice. If you do not specify this parameter, the API returns a 400 error (The specified Lang can not be empty.). Valid values: zh and en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return when you use the NextToken-based pagination method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc*	- (default): ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// The field used for sorting. Valid values:
	//
	// - GmtCreate: creation time.
	//
	// - GmtModified: update time.
	//
	// example:
	//
	// GmtCreate
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: The assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataSetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetsShrinkRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetsShrinkRequest) GetDataSetIdsShrink() *string {
	return s.DataSetIdsShrink
}

func (s *ListDataSetsShrinkRequest) GetDataSetName() *string {
	return s.DataSetName
}

func (s *ListDataSetsShrinkRequest) GetDataSetStatus() *int32 {
	return s.DataSetStatus
}

func (s *ListDataSetsShrinkRequest) GetDataSetType() *string {
	return s.DataSetType
}

func (s *ListDataSetsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSetsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSetsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSetsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDataSetsShrinkRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListDataSetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSetsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSetsShrinkRequest) SetDataSetId(v string) *ListDataSetsShrinkRequest {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetDataSetIdsShrink(v string) *ListDataSetsShrinkRequest {
	s.DataSetIdsShrink = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetDataSetName(v string) *ListDataSetsShrinkRequest {
	s.DataSetName = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetDataSetStatus(v int32) *ListDataSetsShrinkRequest {
	s.DataSetStatus = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetDataSetType(v string) *ListDataSetsShrinkRequest {
	s.DataSetType = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetLang(v string) *ListDataSetsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetMaxResults(v int32) *ListDataSetsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetNextToken(v string) *ListDataSetsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetOrderDirection(v string) *ListDataSetsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetOrderFieldName(v string) *ListDataSetsShrinkRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetPageNumber(v int32) *ListDataSetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetPageSize(v int32) *ListDataSetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetRegionId(v string) *ListDataSetsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSetsShrinkRequest) SetRoleFor(v int64) *ListDataSetsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
