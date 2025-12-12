// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConditionBasicInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRange(v *ConditionBasicInfoCheckRange) *ConditionBasicInfo
	GetCheckRange() *ConditionBasicInfoCheckRange
	SetCid(v string) *ConditionBasicInfo
	GetCid() *string
	SetExclusion(v int32) *ConditionBasicInfo
	GetExclusion() *int32
	SetId(v int64) *ConditionBasicInfo
	GetId() *int64
	SetLambda(v string) *ConditionBasicInfo
	GetLambda() *string
	SetName(v string) *ConditionBasicInfo
	GetName() *string
	SetOperators(v []*OperatorBasicInfo) *ConditionBasicInfo
	GetOperators() []*OperatorBasicInfo
	SetRid(v string) *ConditionBasicInfo
	GetRid() *string
	SetUserGroup(v string) *ConditionBasicInfo
	GetUserGroup() *string
}

type ConditionBasicInfo struct {
	CheckRange *ConditionBasicInfoCheckRange `json:"Check_range,omitempty" xml:"Check_range,omitempty" type:"Struct"`
	Cid        *string                       `json:"Cid,omitempty" xml:"Cid,omitempty"`
	Exclusion  *int32                        `json:"Exclusion,omitempty" xml:"Exclusion,omitempty"`
	Id         *int64                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Lambda     *string                       `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Name       *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators  []*OperatorBasicInfo          `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	Rid        *string                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	UserGroup  *string                       `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s ConditionBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfo) GetCheckRange() *ConditionBasicInfoCheckRange {
	return s.CheckRange
}

func (s *ConditionBasicInfo) GetCid() *string {
	return s.Cid
}

func (s *ConditionBasicInfo) GetExclusion() *int32 {
	return s.Exclusion
}

func (s *ConditionBasicInfo) GetId() *int64 {
	return s.Id
}

func (s *ConditionBasicInfo) GetLambda() *string {
	return s.Lambda
}

func (s *ConditionBasicInfo) GetName() *string {
	return s.Name
}

func (s *ConditionBasicInfo) GetOperators() []*OperatorBasicInfo {
	return s.Operators
}

func (s *ConditionBasicInfo) GetRid() *string {
	return s.Rid
}

func (s *ConditionBasicInfo) GetUserGroup() *string {
	return s.UserGroup
}

func (s *ConditionBasicInfo) SetCheckRange(v *ConditionBasicInfoCheckRange) *ConditionBasicInfo {
	s.CheckRange = v
	return s
}

func (s *ConditionBasicInfo) SetCid(v string) *ConditionBasicInfo {
	s.Cid = &v
	return s
}

func (s *ConditionBasicInfo) SetExclusion(v int32) *ConditionBasicInfo {
	s.Exclusion = &v
	return s
}

func (s *ConditionBasicInfo) SetId(v int64) *ConditionBasicInfo {
	s.Id = &v
	return s
}

func (s *ConditionBasicInfo) SetLambda(v string) *ConditionBasicInfo {
	s.Lambda = &v
	return s
}

func (s *ConditionBasicInfo) SetName(v string) *ConditionBasicInfo {
	s.Name = &v
	return s
}

func (s *ConditionBasicInfo) SetOperators(v []*OperatorBasicInfo) *ConditionBasicInfo {
	s.Operators = v
	return s
}

func (s *ConditionBasicInfo) SetRid(v string) *ConditionBasicInfo {
	s.Rid = &v
	return s
}

func (s *ConditionBasicInfo) SetUserGroup(v string) *ConditionBasicInfo {
	s.UserGroup = &v
	return s
}

func (s *ConditionBasicInfo) Validate() error {
	if s.CheckRange != nil {
		if err := s.CheckRange.Validate(); err != nil {
			return err
		}
	}
	if s.Operators != nil {
		for _, item := range s.Operators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConditionBasicInfoCheckRange struct {
	Absolute            *bool                               `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	AllSentencesSatisfy *bool                               `json:"AllSentencesSatisfy,omitempty" xml:"AllSentencesSatisfy,omitempty"`
	Anchor              *ConditionBasicInfoCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	Range               *ConditionBasicInfoCheckRangeRange  `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Role                *string                             `json:"Role,omitempty" xml:"Role,omitempty"`
	RoleId              *int32                              `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ConditionBasicInfoCheckRange) String() string {
	return dara.Prettify(s)
}

func (s ConditionBasicInfoCheckRange) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRange) GetAbsolute() *bool {
	return s.Absolute
}

func (s *ConditionBasicInfoCheckRange) GetAllSentencesSatisfy() *bool {
	return s.AllSentencesSatisfy
}

func (s *ConditionBasicInfoCheckRange) GetAnchor() *ConditionBasicInfoCheckRangeAnchor {
	return s.Anchor
}

func (s *ConditionBasicInfoCheckRange) GetRange() *ConditionBasicInfoCheckRangeRange {
	return s.Range
}

func (s *ConditionBasicInfoCheckRange) GetRole() *string {
	return s.Role
}

func (s *ConditionBasicInfoCheckRange) GetRoleId() *int32 {
	return s.RoleId
}

func (s *ConditionBasicInfoCheckRange) SetAbsolute(v bool) *ConditionBasicInfoCheckRange {
	s.Absolute = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetAllSentencesSatisfy(v bool) *ConditionBasicInfoCheckRange {
	s.AllSentencesSatisfy = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetAnchor(v *ConditionBasicInfoCheckRangeAnchor) *ConditionBasicInfoCheckRange {
	s.Anchor = v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRange(v *ConditionBasicInfoCheckRangeRange) *ConditionBasicInfoCheckRange {
	s.Range = v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRole(v string) *ConditionBasicInfoCheckRange {
	s.Role = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) SetRoleId(v int32) *ConditionBasicInfoCheckRange {
	s.RoleId = &v
	return s
}

func (s *ConditionBasicInfoCheckRange) Validate() error {
	if s.Anchor != nil {
		if err := s.Anchor.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConditionBasicInfoCheckRangeAnchor struct {
	Cid      *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	HitTime  *int32  `json:"Hit_time,omitempty" xml:"Hit_time,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ConditionBasicInfoCheckRangeAnchor) String() string {
	return dara.Prettify(s)
}

func (s ConditionBasicInfoCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRangeAnchor) GetCid() *string {
	return s.Cid
}

func (s *ConditionBasicInfoCheckRangeAnchor) GetHitTime() *int32 {
	return s.HitTime
}

func (s *ConditionBasicInfoCheckRangeAnchor) GetLocation() *string {
	return s.Location
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetCid(v string) *ConditionBasicInfoCheckRangeAnchor {
	s.Cid = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetHitTime(v int32) *ConditionBasicInfoCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeAnchor) SetLocation(v string) *ConditionBasicInfoCheckRangeAnchor {
	s.Location = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeAnchor) Validate() error {
	return dara.Validate(s)
}

type ConditionBasicInfoCheckRangeRange struct {
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	To   *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ConditionBasicInfoCheckRangeRange) String() string {
	return dara.Prettify(s)
}

func (s ConditionBasicInfoCheckRangeRange) GoString() string {
	return s.String()
}

func (s *ConditionBasicInfoCheckRangeRange) GetFrom() *int32 {
	return s.From
}

func (s *ConditionBasicInfoCheckRangeRange) GetTo() *int32 {
	return s.To
}

func (s *ConditionBasicInfoCheckRangeRange) SetFrom(v int32) *ConditionBasicInfoCheckRangeRange {
	s.From = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeRange) SetTo(v int32) *ConditionBasicInfoCheckRangeRange {
	s.To = &v
	return s
}

func (s *ConditionBasicInfoCheckRangeRange) Validate() error {
	return dara.Validate(s)
}
