// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v []*string) *ListAlertsRequest
	GetAlertLevel() []*string
	SetAlertUuid(v string) *ListAlertsRequest
	GetAlertUuid() *string
	SetEndTime(v int64) *ListAlertsRequest
	GetEndTime() *int64
	SetLang(v string) *ListAlertsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListAlertsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertsRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListAlertsRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListAlertsRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListAlertsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertsRequest
	GetPageSize() *int32
	SetQueryCondition(v string) *ListAlertsRequest
	GetQueryCondition() *string
	SetQueryViewId(v string) *ListAlertsRequest
	GetQueryViewId() *string
	SetRegionId(v string) *ListAlertsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListAlertsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAlertsRequest
	GetRoleType() *int32
	SetStartTime(v int64) *ListAlertsRequest
	GetStartTime() *int64
}

type ListAlertsRequest struct {
	// The threat level of the alert. Valid values:
	//
	// - 5: critical.
	//
	// - 4: high-risk.
	//
	// - 3: medium-risk.
	//
	// - 2: low-risk.
	//
	// - 1: informational.
	AlertLevel []*string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty" type:"Repeated"`
	// The alert ID associated with the event.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The end time of the alert.
	//
	// example:
	//
	// 1766801904000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request or if no more results exist. If more results exist, set this parameter to the NextToken value returned in the previous API call.
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
	// GmtModified
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// The page number.
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
	// The query filter condition in JSON format. Valid values:
	//
	// - `{"Type":"maxCost", "Value":"100"}`: the top 100 queries with the longest execution duration.
	//
	// - `{"Type":"status","Value":"finished"}`: completed queries.
	//
	// - `{"Type":"status","Value":"running"}`: running queries.
	//
	// - `{"Type":"cost","Min":"30","Max":"50"}`: queries with a custom execution duration range. You can specify the minimum and maximum execution duration. **Min*	- specifies the minimum execution duration. **Max*	- specifies the maximum execution duration. Unit: milliseconds (ms).
	//
	//     - If only **Min*	- is specified, queries with an execution duration greater than this value are returned.
	//
	//     - If only **Max*	- is specified, queries with an execution duration less than this value are returned.
	//
	//     - If both **Min*	- and **Max*	- are specified, queries with an execution duration greater than or equal to **Min*	- and less than or equal to **Max*	- are returned.
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
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are located in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets are located outside China.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member accounts in the resource folder.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The time when the alert first occurred.
	//
	// example:
	//
	// 2025-09-30T02:23:00Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) GetAlertLevel() []*string {
	return s.AlertLevel
}

func (s *ListAlertsRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListAlertsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAlertsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAlertsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListAlertsRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListAlertsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *ListAlertsRequest) GetQueryViewId() *string {
	return s.QueryViewId
}

func (s *ListAlertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAlertsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAlertsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAlertsRequest) SetAlertLevel(v []*string) *ListAlertsRequest {
	s.AlertLevel = v
	return s
}

func (s *ListAlertsRequest) SetAlertUuid(v string) *ListAlertsRequest {
	s.AlertUuid = &v
	return s
}

func (s *ListAlertsRequest) SetEndTime(v int64) *ListAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertsRequest) SetLang(v string) *ListAlertsRequest {
	s.Lang = &v
	return s
}

func (s *ListAlertsRequest) SetMaxResults(v int32) *ListAlertsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertsRequest) SetNextToken(v string) *ListAlertsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertsRequest) SetOrderDirection(v string) *ListAlertsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListAlertsRequest) SetOrderFieldName(v string) *ListAlertsRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListAlertsRequest) SetPageNumber(v int32) *ListAlertsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertsRequest) SetPageSize(v int32) *ListAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertsRequest) SetQueryCondition(v string) *ListAlertsRequest {
	s.QueryCondition = &v
	return s
}

func (s *ListAlertsRequest) SetQueryViewId(v string) *ListAlertsRequest {
	s.QueryViewId = &v
	return s
}

func (s *ListAlertsRequest) SetRegionId(v string) *ListAlertsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlertsRequest) SetRoleFor(v int64) *ListAlertsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAlertsRequest) SetRoleType(v int32) *ListAlertsRequest {
	s.RoleType = &v
	return s
}

func (s *ListAlertsRequest) SetStartTime(v int64) *ListAlertsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlertsRequest) Validate() error {
	return dara.Validate(s)
}
