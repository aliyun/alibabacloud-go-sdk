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
	// example:
	//
	// dataset-qt0n8246gs9nackg****
	DataSetId  *string   `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetIds []*string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty" type:"Repeated"`
	// example:
	//
	// lmftest
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 0
	DataSetStatus *int32 `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	// example:
	//
	// custom
	DataSetType *string `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// example:
	//
	// GmtCreate
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
