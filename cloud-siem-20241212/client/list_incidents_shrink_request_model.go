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
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// example:
	//
	// 1749090526055
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ECS unusual log in
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// example:
	//
	// [{\\"data_source\\":[\\"sas\\"]}]
	IncidentTags        *string `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	IncidentUuidsShrink *string `json:"IncidentUuids,omitempty" xml:"IncidentUuids,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// desc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// example:
	//
	// GmtModified
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
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
	// 6c740667-80b2-476d-8924-2e706feb****
	RelateAssetId  *string `json:"RelateAssetId,omitempty" xml:"RelateAssetId,omitempty"`
	RelateEntityId *string `json:"RelateEntityId,omitempty" xml:"RelateEntityId,omitempty"`
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
	// 1690102943000
	StartTime   *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
