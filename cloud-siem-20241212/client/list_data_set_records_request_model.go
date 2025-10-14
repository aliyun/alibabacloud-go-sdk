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
	SetLang(v string) *ListDataSetRecordsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataSetRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetRecordsRequest
	GetNextToken() *string
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
	// This parameter is required.
	//
	// example:
	//
	// dataset-nhcrmjpx1zsorlaq****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
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

func (s ListDataSetRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSetRecordsRequest) GetDataSetId() *string {
	return s.DataSetId
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
