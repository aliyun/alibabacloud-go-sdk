// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *SearchRecursionZonesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionZonesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *SearchRecursionZonesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionZonesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchRecursionZonesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchRecursionZonesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchRecursionZonesResponseBody
	GetTotalPages() *int32
	SetZones(v *SearchRecursionZonesResponseBodyZones) *SearchRecursionZonesResponseBody
	GetZones() *SearchRecursionZonesResponseBodyZones
}

type SearchRecursionZonesResponseBody struct {
	// example:
	//
	// 20
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 123
	TotalPages *int32                                 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Zones      *SearchRecursionZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s SearchRecursionZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionZonesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionZonesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchRecursionZonesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchRecursionZonesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchRecursionZonesResponseBody) GetZones() *SearchRecursionZonesResponseBodyZones {
	return s.Zones
}

func (s *SearchRecursionZonesResponseBody) SetMaxResults(v int32) *SearchRecursionZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetNextToken(v string) *SearchRecursionZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetPageNumber(v int32) *SearchRecursionZonesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetPageSize(v int32) *SearchRecursionZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetRequestId(v string) *SearchRecursionZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetTotalItems(v int32) *SearchRecursionZonesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetTotalPages(v int32) *SearchRecursionZonesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchRecursionZonesResponseBody) SetZones(v *SearchRecursionZonesResponseBodyZones) *SearchRecursionZonesResponseBody {
	s.Zones = v
	return s
}

func (s *SearchRecursionZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchRecursionZonesResponseBodyZones struct {
	Zone []*SearchRecursionZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s SearchRecursionZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBodyZones) GetZone() []*SearchRecursionZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *SearchRecursionZonesResponseBodyZones) SetZone(v []*SearchRecursionZonesResponseBodyZonesZone) *SearchRecursionZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *SearchRecursionZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}

type SearchRecursionZonesResponseBodyZonesZone struct {
	// example:
	//
	// 2022-10-17T06:13Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1749694625000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// ***
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// example:
	//
	// USER
	CreatorType     *string                                                   `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	EffectiveScopes *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty" type:"Struct"`
	// example:
	//
	// record
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// example:
	//
	// 20
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// example:
	//
	// 107
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 2024-08-13T01:44Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1639621006000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 169439170000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// ixiqiu.cn
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s SearchRecursionZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetCreator() *string {
	return s.Creator
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetCreatorType() *string {
	return s.CreatorType
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetEffectiveScopes() *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes {
	return s.EffectiveScopes
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchRecursionZonesResponseBodyZonesZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetCreateTime(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetCreateTimestamp(v int64) *SearchRecursionZonesResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetCreator(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetCreatorSubType(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.CreatorSubType = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetCreatorType(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.CreatorType = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetEffectiveScopes(v *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) *SearchRecursionZonesResponseBodyZonesZone {
	s.EffectiveScopes = v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetProxyPattern(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.ProxyPattern = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetRecordCount(v int32) *SearchRecursionZonesResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetRemark(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetUpdateTime(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetUpdateTimestamp(v int64) *SearchRecursionZonesResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetZoneId(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) SetZoneName(v string) *SearchRecursionZonesResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZone) Validate() error {
	return dara.Validate(s)
}

type SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes struct {
	EffectiveScope []*SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope `json:"EffectiveScope,omitempty" xml:"EffectiveScope,omitempty" type:"Repeated"`
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) GetEffectiveScope() []*SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	return s.EffectiveScope
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) SetEffectiveScope(v []*SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes {
	s.EffectiveScope = v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopes) Validate() error {
	return dara.Validate(s)
}

type SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope struct {
	// example:
	//
	// account
	EffectiveType *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	// example:
	//
	// [20003]
	Scopes *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Struct"`
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) GetScopes() *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes {
	return s.Scopes
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) SetEffectiveType(v string) *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	s.EffectiveType = &v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) SetScopes(v *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope {
	s.Scopes = v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScope) Validate() error {
	return dara.Validate(s)
}

type SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes struct {
	Scope []*string `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) GetScope() []*string {
	return s.Scope
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) SetScope(v []*string) *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes {
	s.Scope = v
	return s
}

func (s *SearchRecursionZonesResponseBodyZonesZoneEffectiveScopesEffectiveScopeScopes) Validate() error {
	return dara.Validate(s)
}
