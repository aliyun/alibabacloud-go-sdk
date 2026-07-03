// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIncidentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuid(v string) *ListIncidentsShrinkRequest
	GetAlertUuid() *string
	SetEndTime(v int64) *ListIncidentsShrinkRequest
	GetEndTime() *int64
	SetIncidentName(v string) *ListIncidentsShrinkRequest
	GetIncidentName() *string
	SetIncidentStatus(v int32) *ListIncidentsShrinkRequest
	GetIncidentStatus() *int32
	SetIncidentTags(v string) *ListIncidentsShrinkRequest
	GetIncidentTags() *string
	SetIncidentUuidsShrink(v string) *ListIncidentsShrinkRequest
	GetIncidentUuidsShrink() *string
	SetLang(v string) *ListIncidentsShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListIncidentsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIncidentsShrinkRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListIncidentsShrinkRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListIncidentsShrinkRequest
	GetOrderFieldName() *string
	SetOwners(v []*string) *ListIncidentsShrinkRequest
	GetOwners() []*string
	SetPageNumber(v int32) *ListIncidentsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIncidentsShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListIncidentsShrinkRequest
	GetRegionId() *string
	SetRelateAssetId(v string) *ListIncidentsShrinkRequest
	GetRelateAssetId() *string
	SetRelateEntityId(v string) *ListIncidentsShrinkRequest
	GetRelateEntityId() *string
	SetRoleFor(v int64) *ListIncidentsShrinkRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListIncidentsShrinkRequest
	GetRoleType() *int32
	SetStartTime(v int64) *ListIncidentsShrinkRequest
	GetStartTime() *int64
	SetThreatLevel(v []*string) *ListIncidentsShrinkRequest
	GetThreatLevel() []*string
}

type ListIncidentsShrinkRequest struct {
	// The alert ID.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The end time as a timestamp in milliseconds (ms).
	//
	// example:
	//
	// 1749090526055
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the incident.
	//
	// example:
	//
	// ECS unusual log in
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The status of the incident. Valid values:
	//
	// - 0: unhandled.
	//
	// - 1: handling.
	//
	// - 5: handling failed.
	//
	// - 10: handled.
	//
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// The tags of the incident.
	//
	// example:
	//
	// [{\\"data_source\\":[\\"sas\\"]}]
	IncidentTags *string `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	// The list of incident UUIDs, separated by commas (,).
	IncidentUuidsShrink *string `json:"IncidentUuids,omitempty" xml:"IncidentUuids,omitempty"`
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
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort direction. Valid values:
	//
	// - **desc*	- (default): descending order.
	//
	// - **asc**: ascending order.
	//
	// example:
	//
	// desc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// The field name used to sort the list. Valid values:
	//
	// - GmtModified: sorts by incident update time (default).
	//
	// - ThreatScore: sorts by threat score.
	//
	// example:
	//
	// GmtModified
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// The UID of the account that owns the incident.
	Owners []*string `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
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
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the asset associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	RelateAssetId *string `json:"RelateAssetId,omitempty" xml:"RelateAssetId,omitempty"`
	// The ID of the entity associated with the incident.
	//
	// example:
	//
	// b920ed22259f5412099e97dfda96****
	RelateEntityId *string `json:"RelateEntityId,omitempty" xml:"RelateEntityId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
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
	// The start time as a timestamp in milliseconds (ms).
	//
	// example:
	//
	// 1690102943000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The threat level. Valid values:
	//
	// - 5: critical.
	//
	// - 4: high.
	//
	// - 3: medium.
	//
	// - 2: low.
	//
	// - 1: informational.
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListIncidentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIncidentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentsShrinkRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListIncidentsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIncidentsShrinkRequest) GetIncidentName() *string {
	return s.IncidentName
}

func (s *ListIncidentsShrinkRequest) GetIncidentStatus() *int32 {
	return s.IncidentStatus
}

func (s *ListIncidentsShrinkRequest) GetIncidentTags() *string {
	return s.IncidentTags
}

func (s *ListIncidentsShrinkRequest) GetIncidentUuidsShrink() *string {
	return s.IncidentUuidsShrink
}

func (s *ListIncidentsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListIncidentsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIncidentsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIncidentsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListIncidentsShrinkRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListIncidentsShrinkRequest) GetOwners() []*string {
	return s.Owners
}

func (s *ListIncidentsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIncidentsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIncidentsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIncidentsShrinkRequest) GetRelateAssetId() *string {
	return s.RelateAssetId
}

func (s *ListIncidentsShrinkRequest) GetRelateEntityId() *string {
	return s.RelateEntityId
}

func (s *ListIncidentsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListIncidentsShrinkRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListIncidentsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIncidentsShrinkRequest) GetThreatLevel() []*string {
	return s.ThreatLevel
}

func (s *ListIncidentsShrinkRequest) SetAlertUuid(v string) *ListIncidentsShrinkRequest {
	s.AlertUuid = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetEndTime(v int64) *ListIncidentsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetIncidentName(v string) *ListIncidentsShrinkRequest {
	s.IncidentName = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetIncidentStatus(v int32) *ListIncidentsShrinkRequest {
	s.IncidentStatus = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetIncidentTags(v string) *ListIncidentsShrinkRequest {
	s.IncidentTags = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetIncidentUuidsShrink(v string) *ListIncidentsShrinkRequest {
	s.IncidentUuidsShrink = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetLang(v string) *ListIncidentsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetMaxResults(v int32) *ListIncidentsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetNextToken(v string) *ListIncidentsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetOrderDirection(v string) *ListIncidentsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetOrderFieldName(v string) *ListIncidentsShrinkRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetOwners(v []*string) *ListIncidentsShrinkRequest {
	s.Owners = v
	return s
}

func (s *ListIncidentsShrinkRequest) SetPageNumber(v int32) *ListIncidentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetPageSize(v int32) *ListIncidentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetRegionId(v string) *ListIncidentsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetRelateAssetId(v string) *ListIncidentsShrinkRequest {
	s.RelateAssetId = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetRelateEntityId(v string) *ListIncidentsShrinkRequest {
	s.RelateEntityId = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetRoleFor(v int64) *ListIncidentsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetRoleType(v int32) *ListIncidentsShrinkRequest {
	s.RoleType = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetStartTime(v int64) *ListIncidentsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListIncidentsShrinkRequest) SetThreatLevel(v []*string) *ListIncidentsShrinkRequest {
	s.ThreatLevel = v
	return s
}

func (s *ListIncidentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
