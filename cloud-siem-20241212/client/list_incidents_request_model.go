// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIncidentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuid(v string) *ListIncidentsRequest
	GetAlertUuid() *string
	SetEndTime(v int64) *ListIncidentsRequest
	GetEndTime() *int64
	SetIncidentName(v string) *ListIncidentsRequest
	GetIncidentName() *string
	SetIncidentStatus(v int32) *ListIncidentsRequest
	GetIncidentStatus() *int32
	SetIncidentTags(v string) *ListIncidentsRequest
	GetIncidentTags() *string
	SetIncidentUuids(v []*string) *ListIncidentsRequest
	GetIncidentUuids() []*string
	SetLang(v string) *ListIncidentsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListIncidentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIncidentsRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListIncidentsRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListIncidentsRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListIncidentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIncidentsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListIncidentsRequest
	GetRegionId() *string
	SetRelateAssetId(v string) *ListIncidentsRequest
	GetRelateAssetId() *string
	SetRelateEntityId(v string) *ListIncidentsRequest
	GetRelateEntityId() *string
	SetRoleFor(v int64) *ListIncidentsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListIncidentsRequest
	GetRoleType() *int32
	SetStartTime(v int64) *ListIncidentsRequest
	GetStartTime() *int64
	SetThreatLevel(v []*string) *ListIncidentsRequest
	GetThreatLevel() []*string
}

type ListIncidentsRequest struct {
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
	IncidentTags  *string   `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	IncidentUuids []*string `json:"IncidentUuids,omitempty" xml:"IncidentUuids,omitempty" type:"Repeated"`
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

func (s ListIncidentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIncidentsRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentsRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListIncidentsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIncidentsRequest) GetIncidentName() *string {
	return s.IncidentName
}

func (s *ListIncidentsRequest) GetIncidentStatus() *int32 {
	return s.IncidentStatus
}

func (s *ListIncidentsRequest) GetIncidentTags() *string {
	return s.IncidentTags
}

func (s *ListIncidentsRequest) GetIncidentUuids() []*string {
	return s.IncidentUuids
}

func (s *ListIncidentsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListIncidentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIncidentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIncidentsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListIncidentsRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListIncidentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIncidentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIncidentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIncidentsRequest) GetRelateAssetId() *string {
	return s.RelateAssetId
}

func (s *ListIncidentsRequest) GetRelateEntityId() *string {
	return s.RelateEntityId
}

func (s *ListIncidentsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListIncidentsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListIncidentsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIncidentsRequest) GetThreatLevel() []*string {
	return s.ThreatLevel
}

func (s *ListIncidentsRequest) SetAlertUuid(v string) *ListIncidentsRequest {
	s.AlertUuid = &v
	return s
}

func (s *ListIncidentsRequest) SetEndTime(v int64) *ListIncidentsRequest {
	s.EndTime = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentName(v string) *ListIncidentsRequest {
	s.IncidentName = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentStatus(v int32) *ListIncidentsRequest {
	s.IncidentStatus = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentTags(v string) *ListIncidentsRequest {
	s.IncidentTags = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentUuids(v []*string) *ListIncidentsRequest {
	s.IncidentUuids = v
	return s
}

func (s *ListIncidentsRequest) SetLang(v string) *ListIncidentsRequest {
	s.Lang = &v
	return s
}

func (s *ListIncidentsRequest) SetMaxResults(v int32) *ListIncidentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIncidentsRequest) SetNextToken(v string) *ListIncidentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIncidentsRequest) SetOrderDirection(v string) *ListIncidentsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListIncidentsRequest) SetOrderFieldName(v string) *ListIncidentsRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListIncidentsRequest) SetPageNumber(v int32) *ListIncidentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentsRequest) SetPageSize(v int32) *ListIncidentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIncidentsRequest) SetRegionId(v string) *ListIncidentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIncidentsRequest) SetRelateAssetId(v string) *ListIncidentsRequest {
	s.RelateAssetId = &v
	return s
}

func (s *ListIncidentsRequest) SetRelateEntityId(v string) *ListIncidentsRequest {
	s.RelateEntityId = &v
	return s
}

func (s *ListIncidentsRequest) SetRoleFor(v int64) *ListIncidentsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListIncidentsRequest) SetRoleType(v int32) *ListIncidentsRequest {
	s.RoleType = &v
	return s
}

func (s *ListIncidentsRequest) SetStartTime(v int64) *ListIncidentsRequest {
	s.StartTime = &v
	return s
}

func (s *ListIncidentsRequest) SetThreatLevel(v []*string) *ListIncidentsRequest {
	s.ThreatLevel = v
	return s
}

func (s *ListIncidentsRequest) Validate() error {
	return dara.Validate(s)
}
