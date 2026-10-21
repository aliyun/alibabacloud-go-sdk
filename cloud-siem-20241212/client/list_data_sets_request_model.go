// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetId(v string) *ListDataSetsRequest
	GetDataSetId() *string
	SetDataSetIds(v []*string) *ListDataSetsRequest
	GetDataSetIds() []*string
	SetDataSetName(v string) *ListDataSetsRequest
	GetDataSetName() *string
	SetDataSetStatus(v int32) *ListDataSetsRequest
	GetDataSetStatus() *int32
	SetDataSetType(v string) *ListDataSetsRequest
	GetDataSetType() *string
	SetLang(v string) *ListDataSetsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataSetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetsRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListDataSetsRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListDataSetsRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListDataSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDataSetsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSetsRequest
	GetRoleFor() *int64
}

type ListDataSetsRequest struct {
	// The dataset ID.
	//
	// example:
	//
	// dataset-qt0n8246gs9nackg****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// The list of dataset IDs.
	DataSetIds []*string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty" type:"Repeated"`
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

func (s ListDataSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetsRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetsRequest) GetDataSetIds() []*string {
	return s.DataSetIds
}

func (s *ListDataSetsRequest) GetDataSetName() *string {
	return s.DataSetName
}

func (s *ListDataSetsRequest) GetDataSetStatus() *int32 {
	return s.DataSetStatus
}

func (s *ListDataSetsRequest) GetDataSetType() *string {
	return s.DataSetType
}

func (s *ListDataSetsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSetsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDataSetsRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListDataSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSetsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSetsRequest) SetDataSetId(v string) *ListDataSetsRequest {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetsRequest) SetDataSetIds(v []*string) *ListDataSetsRequest {
	s.DataSetIds = v
	return s
}

func (s *ListDataSetsRequest) SetDataSetName(v string) *ListDataSetsRequest {
	s.DataSetName = &v
	return s
}

func (s *ListDataSetsRequest) SetDataSetStatus(v int32) *ListDataSetsRequest {
	s.DataSetStatus = &v
	return s
}

func (s *ListDataSetsRequest) SetDataSetType(v string) *ListDataSetsRequest {
	s.DataSetType = &v
	return s
}

func (s *ListDataSetsRequest) SetLang(v string) *ListDataSetsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSetsRequest) SetMaxResults(v int32) *ListDataSetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSetsRequest) SetNextToken(v string) *ListDataSetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSetsRequest) SetOrderDirection(v string) *ListDataSetsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDataSetsRequest) SetOrderFieldName(v string) *ListDataSetsRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListDataSetsRequest) SetPageNumber(v int32) *ListDataSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetsRequest) SetPageSize(v int32) *ListDataSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSetsRequest) SetRegionId(v string) *ListDataSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSetsRequest) SetRoleFor(v int64) *ListDataSetsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSetsRequest) Validate() error {
	return dara.Validate(s)
}
