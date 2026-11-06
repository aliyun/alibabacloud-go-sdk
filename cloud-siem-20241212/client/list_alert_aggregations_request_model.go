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
	// The aggregation method.
	//
	// example:
	//
	// avg
	AggregationType *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	// The time when the alert ended.
	//
	// example:
	//
	// 1773936020000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The maximum number of entries to return for this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. If this is the first query or no subsequent query is required, you do not need to specify this parameter. If a subsequent query is required, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query filter conditions in JSON format. Valid values:
	//
	// - `{"Type":"maxCost", "Value":"100"}`: The top 100 queries that take the longest time.
	//
	// - `{"Type":"status","Value":"finished"}`: Completed queries.
	//
	// - `{"Type":"status","Value":"running"}`: Running queries.
	//
	// - `{"Type":"cost","Min":"30","Max":"50"}`: Custom duration queries. You can specify the minimum and maximum query durations. **Min*	- specifies the minimum duration, and **Max*	- specifies the maximum duration. Unit: milliseconds.
	//
	//     - If only **Min*	- is specified, the query returns queries whose execution duration is greater than the specified value.
	//
	//     - If only **Max*	- is specified, the query returns queries whose execution duration is less than the specified value.
	//
	//     - If both **Min*	- and **Max*	- are specified, the query returns queries whose execution duration is greater than or equal to the value of **Min*	- and less than or equal to the value of **Max**.
	//
	// example:
	//
	// {\\"Type\\":\\"cost\\",\\"Max\\":\\"200\\"}
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The unique identifier of the query view.
	//
	// example:
	//
	// qv-a1b2c3d4e5f6g7****
	QueryViewId *string `json:"QueryViewId,omitempty" xml:"QueryViewId,omitempty"`
	// The region where the data management center of Threat Analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The time when the alert first occurred.
	//
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
