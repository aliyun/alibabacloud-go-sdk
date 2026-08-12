// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertAggregationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationType(v string) *ListAlertAggregationsRequest
	GetAggregationType() *string
	SetEndTime(v int64) *ListAlertAggregationsRequest
	GetEndTime() *int64
	SetLang(v string) *ListAlertAggregationsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListAlertAggregationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertAggregationsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListAlertAggregationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertAggregationsRequest
	GetPageSize() *int32
	SetQueryCondition(v string) *ListAlertAggregationsRequest
	GetQueryCondition() *string
	SetQueryViewId(v string) *ListAlertAggregationsRequest
	GetQueryViewId() *string
	SetRegionId(v string) *ListAlertAggregationsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListAlertAggregationsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAlertAggregationsRequest
	GetRoleType() *int32
	SetStartTime(v int64) *ListAlertAggregationsRequest
	GetStartTime() *int64
}

type ListAlertAggregationsRequest struct {
	// example:
	//
	// avg
	AggregationType *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	// example:
	//
	// 1773936020000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// {\\"Type\\":\\"cost\\",\\"Max\\":\\"200\\"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// example:
	//
	// qv-a1b2c3d4e5f6g7****
	QueryViewId *string `json:"QueryViewId,omitempty" xml:"QueryViewId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// example:
	//
	// 2026-02-04T08:36:26Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlertAggregationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsRequest) GetAggregationType() *string {
	return s.AggregationType
}

func (s *ListAlertAggregationsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAlertAggregationsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAlertAggregationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertAggregationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertAggregationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertAggregationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertAggregationsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *ListAlertAggregationsRequest) GetQueryViewId() *string {
	return s.QueryViewId
}

func (s *ListAlertAggregationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertAggregationsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAlertAggregationsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAlertAggregationsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAlertAggregationsRequest) SetAggregationType(v string) *ListAlertAggregationsRequest {
	s.AggregationType = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetEndTime(v int64) *ListAlertAggregationsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetLang(v string) *ListAlertAggregationsRequest {
	s.Lang = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetMaxResults(v int32) *ListAlertAggregationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetNextToken(v string) *ListAlertAggregationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetPageNumber(v int32) *ListAlertAggregationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetPageSize(v int32) *ListAlertAggregationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetQueryCondition(v string) *ListAlertAggregationsRequest {
	s.QueryCondition = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetQueryViewId(v string) *ListAlertAggregationsRequest {
	s.QueryViewId = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetRegionId(v string) *ListAlertAggregationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetRoleFor(v int64) *ListAlertAggregationsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetRoleType(v int32) *ListAlertAggregationsRequest {
	s.RoleType = &v
	return s
}

func (s *ListAlertAggregationsRequest) SetStartTime(v int64) *ListAlertAggregationsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlertAggregationsRequest) Validate() error {
	return dara.Validate(s)
}
