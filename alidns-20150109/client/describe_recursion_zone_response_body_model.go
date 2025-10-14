// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeRecursionZoneResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeRecursionZoneResponseBody
	GetCreateTimestamp() *int64
	SetCreator(v string) *DescribeRecursionZoneResponseBody
	GetCreator() *string
	SetCreatorSubType(v string) *DescribeRecursionZoneResponseBody
	GetCreatorSubType() *string
	SetCreatorType(v string) *DescribeRecursionZoneResponseBody
	GetCreatorType() *string
	SetEffectiveScopes(v *DescribeRecursionZoneResponseBodyEffectiveScopes) *DescribeRecursionZoneResponseBody
	GetEffectiveScopes() *DescribeRecursionZoneResponseBodyEffectiveScopes
	SetProxyPattern(v string) *DescribeRecursionZoneResponseBody
	GetProxyPattern() *string
	SetRecordCount(v int32) *DescribeRecursionZoneResponseBody
	GetRecordCount() *int32
	SetRemark(v string) *DescribeRecursionZoneResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeRecursionZoneResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *DescribeRecursionZoneResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeRecursionZoneResponseBody
	GetUpdateTimestamp() *int64
	SetUserId(v string) *DescribeRecursionZoneResponseBody
	GetUserId() *string
	SetZoneId(v string) *DescribeRecursionZoneResponseBody
	GetZoneId() *string
	SetZoneName(v string) *DescribeRecursionZoneResponseBody
	GetZoneName() *string
}

type DescribeRecursionZoneResponseBody struct {
	// example:
	//
	// 2018-06-06T11:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533773400000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// SOAR
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// example:
	//
	// USER
	CreatorType     *string                                           `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	EffectiveScopes *DescribeRecursionZoneResponseBodyEffectiveScopes `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty" type:"Struct"`
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
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2018-01-03T08:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 1527690629357
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 169438909000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// cheng.suow.cc
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeRecursionZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRecursionZoneResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeRecursionZoneResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *DescribeRecursionZoneResponseBody) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *DescribeRecursionZoneResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeRecursionZoneResponseBody) GetEffectiveScopes() *DescribeRecursionZoneResponseBodyEffectiveScopes {
	return s.EffectiveScopes
}

func (s *DescribeRecursionZoneResponseBody) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *DescribeRecursionZoneResponseBody) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *DescribeRecursionZoneResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeRecursionZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecursionZoneResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeRecursionZoneResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeRecursionZoneResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeRecursionZoneResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRecursionZoneResponseBody) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeRecursionZoneResponseBody) SetCreateTime(v string) *DescribeRecursionZoneResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetCreateTimestamp(v int64) *DescribeRecursionZoneResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetCreator(v string) *DescribeRecursionZoneResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetCreatorSubType(v string) *DescribeRecursionZoneResponseBody {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetCreatorType(v string) *DescribeRecursionZoneResponseBody {
	s.CreatorType = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetEffectiveScopes(v *DescribeRecursionZoneResponseBodyEffectiveScopes) *DescribeRecursionZoneResponseBody {
	s.EffectiveScopes = v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetProxyPattern(v string) *DescribeRecursionZoneResponseBody {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetRecordCount(v int32) *DescribeRecursionZoneResponseBody {
	s.RecordCount = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetRemark(v string) *DescribeRecursionZoneResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetRequestId(v string) *DescribeRecursionZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetUpdateTime(v string) *DescribeRecursionZoneResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetUpdateTimestamp(v int64) *DescribeRecursionZoneResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetUserId(v string) *DescribeRecursionZoneResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetZoneId(v string) *DescribeRecursionZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) SetZoneName(v string) *DescribeRecursionZoneResponseBody {
	s.ZoneName = &v
	return s
}

func (s *DescribeRecursionZoneResponseBody) Validate() error {
	if s.EffectiveScopes != nil {
		if err := s.EffectiveScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecursionZoneResponseBodyEffectiveScopes struct {
	EffectiveScope []*DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope `json:"EffectiveScope,omitempty" xml:"EffectiveScope,omitempty" type:"Repeated"`
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopes) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopes) GetEffectiveScope() []*DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope {
	return s.EffectiveScope
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopes) SetEffectiveScope(v []*DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) *DescribeRecursionZoneResponseBodyEffectiveScopes {
	s.EffectiveScope = v
	return s
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopes) Validate() error {
	if s.EffectiveScope != nil {
		for _, item := range s.EffectiveScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope struct {
	// example:
	//
	// account
	EffectiveType *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	// example:
	//
	// [20003]
	Scopes *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Struct"`
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) GetScopes() *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes {
	return s.Scopes
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) SetEffectiveType(v string) *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope {
	s.EffectiveType = &v
	return s
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) SetScopes(v *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope {
	s.Scopes = v
	return s
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScope) Validate() error {
	if s.Scopes != nil {
		if err := s.Scopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes struct {
	Scope []*string `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) GetScope() []*string {
	return s.Scope
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) SetScope(v []*string) *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes {
	s.Scope = v
	return s
}

func (s *DescribeRecursionZoneResponseBodyEffectiveScopesEffectiveScopeScopes) Validate() error {
	return dara.Validate(s)
}
