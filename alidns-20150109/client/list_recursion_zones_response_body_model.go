// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecursionZonesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecursionZonesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListRecursionZonesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecursionZonesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRecursionZonesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListRecursionZonesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListRecursionZonesResponseBody
	GetTotalPages() *int32
	SetZones(v *ListRecursionZonesResponseBodyZones) *ListRecursionZonesResponseBody
	GetZones() *ListRecursionZonesResponseBodyZones
}

type ListRecursionZonesResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32                               `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Zones      *ListRecursionZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s ListRecursionZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecursionZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecursionZonesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecursionZonesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecursionZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecursionZonesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListRecursionZonesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListRecursionZonesResponseBody) GetZones() *ListRecursionZonesResponseBodyZones {
	return s.Zones
}

func (s *ListRecursionZonesResponseBody) SetMaxResults(v int32) *ListRecursionZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetNextToken(v string) *ListRecursionZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetPageNumber(v int32) *ListRecursionZonesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetPageSize(v int32) *ListRecursionZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetRequestId(v string) *ListRecursionZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetTotalItems(v int32) *ListRecursionZonesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetTotalPages(v int32) *ListRecursionZonesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListRecursionZonesResponseBody) SetZones(v *ListRecursionZonesResponseBodyZones) *ListRecursionZonesResponseBody {
	s.Zones = v
	return s
}

func (s *ListRecursionZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecursionZonesResponseBodyZones struct {
	Zone []*ListRecursionZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s ListRecursionZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBodyZones) GetZone() []*ListRecursionZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *ListRecursionZonesResponseBodyZones) SetZone(v []*ListRecursionZonesResponseBodyZonesZone) *ListRecursionZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *ListRecursionZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}

type ListRecursionZonesResponseBodyZonesZone struct {
	// example:
	//
	// 2021-03-08T05:45Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1729674680000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// 218497924149333932
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// example:
	//
	// USER
	CreatorType     *string                                                 `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	EffectiveScopes *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty" type:"Struct"`
	// example:
	//
	// record
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// example:
	//
	// 8
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// example:
	//
	// 107
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 2024-11-12T04:30Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1707189878000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// dfsdfsd
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s ListRecursionZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetCreator() *string {
	return s.Creator
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetEffectiveScopes() *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes {
	return s.EffectiveScopes
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetRemark() *string {
	return s.Remark
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListRecursionZonesResponseBodyZonesZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetCreateTime(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetCreateTimestamp(v int64) *ListRecursionZonesResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetCreator(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetCreatorSubType(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.CreatorSubType = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetCreatorType(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.CreatorType = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetEffectiveScopes(v *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) *ListRecursionZonesResponseBodyZonesZone {
	s.EffectiveScopes = v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetProxyPattern(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.ProxyPattern = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetRecordCount(v int32) *ListRecursionZonesResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetRemark(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetUpdateTime(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetUpdateTimestamp(v int64) *ListRecursionZonesResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetZoneId(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) SetZoneName(v string) *ListRecursionZonesResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZone) Validate() error {
	return dara.Validate(s)
}

type ListRecursionZonesResponseBodyZonesZoneEffectiveScopes struct {
	EffectiveScope []*ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope `json:"EffectiveScope,omitempty" xml:"EffectiveScope,omitempty" type:"Repeated"`
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) GetEffectiveScope() []*ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	return s.EffectiveScope
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) SetEffectiveScope(v []*ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes {
	s.EffectiveScope = v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopes) Validate() error {
	return dara.Validate(s)
}

type ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope struct {
	// example:
	//
	// account
	EffectiveType *string                                                                     `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	Scopes        *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Struct"`
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GetScopes() *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes {
	return s.Scopes
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) SetEffectiveType(v string) *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	s.EffectiveType = &v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) SetScopes(v *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	s.Scopes = v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) Validate() error {
	return dara.Validate(s)
}

type ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes struct {
	Scope []*string `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) GetScope() []*string {
	return s.Scope
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) SetScope(v []*string) *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes {
	s.Scope = v
	return s
}

func (s *ListRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) Validate() error {
	return dara.Validate(s)
}
