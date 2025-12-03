// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iYuqingFinanceEvent interface {
	dara.Model
	String() string
	GoString() string
	SetComprehensiveRisk(v float64) *YuqingFinanceEvent
	GetComprehensiveRisk() *float64
	SetEntityArea(v string) *YuqingFinanceEvent
	GetEntityArea() *string
	SetEntityCrn(v string) *YuqingFinanceEvent
	GetEntityCrn() *string
	SetEntityEmotionScore(v float64) *YuqingFinanceEvent
	GetEntityEmotionScore() *float64
	SetEntityId(v int64) *YuqingFinanceEvent
	GetEntityId() *int64
	SetEntityName(v string) *YuqingFinanceEvent
	GetEntityName() *string
	SetEntityRelevanceScore(v float64) *YuqingFinanceEvent
	GetEntityRelevanceScore() *float64
	SetEntityShowName(v string) *YuqingFinanceEvent
	GetEntityShowName() *string
	SetEntitySummary(v string) *YuqingFinanceEvent
	GetEntitySummary() *string
	SetEntityType(v string) *YuqingFinanceEvent
	GetEntityType() *string
	SetEventId(v string) *YuqingFinanceEvent
	GetEventId() *string
	SetEventLevel3Code(v int64) *YuqingFinanceEvent
	GetEventLevel3Code() *int64
	SetEventLevel3Name(v string) *YuqingFinanceEvent
	GetEventLevel3Name() *string
	SetEventTags(v string) *YuqingFinanceEvent
	GetEventTags() *string
	SetEventTime(v int64) *YuqingFinanceEvent
	GetEventTime() *int64
	SetSecurityAbbreviation(v string) *YuqingFinanceEvent
	GetSecurityAbbreviation() *string
	SetSecurityCategoryCodes(v []*string) *YuqingFinanceEvent
	GetSecurityCategoryCodes() []*string
	SetSecurityCodes(v []*string) *YuqingFinanceEvent
	GetSecurityCodes() []*string
	SetSecurityMarketsCodes(v []*string) *YuqingFinanceEvent
	GetSecurityMarketsCodes() []*string
	SetSpamScore(v float64) *YuqingFinanceEvent
	GetSpamScore() *float64
	SetUserSubscribeEntityTags(v []*string) *YuqingFinanceEvent
	GetUserSubscribeEntityTags() []*string
	SetUserSubscribeEventTags(v []*int64) *YuqingFinanceEvent
	GetUserSubscribeEventTags() []*int64
}

type YuqingFinanceEvent struct {
	ComprehensiveRisk       *float64  `json:"comprehensiveRisk,omitempty" xml:"comprehensiveRisk,omitempty"`
	EntityArea              *string   `json:"entityArea,omitempty" xml:"entityArea,omitempty"`
	EntityCrn               *string   `json:"entityCrn,omitempty" xml:"entityCrn,omitempty"`
	EntityEmotionScore      *float64  `json:"entityEmotionScore,omitempty" xml:"entityEmotionScore,omitempty"`
	EntityId                *int64    `json:"entityId,omitempty" xml:"entityId,omitempty"`
	EntityName              *string   `json:"entityName,omitempty" xml:"entityName,omitempty"`
	EntityRelevanceScore    *float64  `json:"entityRelevanceScore,omitempty" xml:"entityRelevanceScore,omitempty"`
	EntityShowName          *string   `json:"entityShowName,omitempty" xml:"entityShowName,omitempty"`
	EntitySummary           *string   `json:"entitySummary,omitempty" xml:"entitySummary,omitempty"`
	EntityType              *string   `json:"entityType,omitempty" xml:"entityType,omitempty"`
	EventId                 *string   `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventLevel3Code         *int64    `json:"eventLevel3Code,omitempty" xml:"eventLevel3Code,omitempty"`
	EventLevel3Name         *string   `json:"eventLevel3Name,omitempty" xml:"eventLevel3Name,omitempty"`
	EventTags               *string   `json:"eventTags,omitempty" xml:"eventTags,omitempty"`
	EventTime               *int64    `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	SecurityAbbreviation    *string   `json:"securityAbbreviation,omitempty" xml:"securityAbbreviation,omitempty"`
	SecurityCategoryCodes   []*string `json:"securityCategoryCodes,omitempty" xml:"securityCategoryCodes,omitempty" type:"Repeated"`
	SecurityCodes           []*string `json:"securityCodes,omitempty" xml:"securityCodes,omitempty" type:"Repeated"`
	SecurityMarketsCodes    []*string `json:"securityMarketsCodes,omitempty" xml:"securityMarketsCodes,omitempty" type:"Repeated"`
	SpamScore               *float64  `json:"spamScore,omitempty" xml:"spamScore,omitempty"`
	UserSubscribeEntityTags []*string `json:"userSubscribeEntityTags,omitempty" xml:"userSubscribeEntityTags,omitempty" type:"Repeated"`
	UserSubscribeEventTags  []*int64  `json:"userSubscribeEventTags,omitempty" xml:"userSubscribeEventTags,omitempty" type:"Repeated"`
}

func (s YuqingFinanceEvent) String() string {
	return dara.Prettify(s)
}

func (s YuqingFinanceEvent) GoString() string {
	return s.String()
}

func (s *YuqingFinanceEvent) GetComprehensiveRisk() *float64 {
	return s.ComprehensiveRisk
}

func (s *YuqingFinanceEvent) GetEntityArea() *string {
	return s.EntityArea
}

func (s *YuqingFinanceEvent) GetEntityCrn() *string {
	return s.EntityCrn
}

func (s *YuqingFinanceEvent) GetEntityEmotionScore() *float64 {
	return s.EntityEmotionScore
}

func (s *YuqingFinanceEvent) GetEntityId() *int64 {
	return s.EntityId
}

func (s *YuqingFinanceEvent) GetEntityName() *string {
	return s.EntityName
}

func (s *YuqingFinanceEvent) GetEntityRelevanceScore() *float64 {
	return s.EntityRelevanceScore
}

func (s *YuqingFinanceEvent) GetEntityShowName() *string {
	return s.EntityShowName
}

func (s *YuqingFinanceEvent) GetEntitySummary() *string {
	return s.EntitySummary
}

func (s *YuqingFinanceEvent) GetEntityType() *string {
	return s.EntityType
}

func (s *YuqingFinanceEvent) GetEventId() *string {
	return s.EventId
}

func (s *YuqingFinanceEvent) GetEventLevel3Code() *int64 {
	return s.EventLevel3Code
}

func (s *YuqingFinanceEvent) GetEventLevel3Name() *string {
	return s.EventLevel3Name
}

func (s *YuqingFinanceEvent) GetEventTags() *string {
	return s.EventTags
}

func (s *YuqingFinanceEvent) GetEventTime() *int64 {
	return s.EventTime
}

func (s *YuqingFinanceEvent) GetSecurityAbbreviation() *string {
	return s.SecurityAbbreviation
}

func (s *YuqingFinanceEvent) GetSecurityCategoryCodes() []*string {
	return s.SecurityCategoryCodes
}

func (s *YuqingFinanceEvent) GetSecurityCodes() []*string {
	return s.SecurityCodes
}

func (s *YuqingFinanceEvent) GetSecurityMarketsCodes() []*string {
	return s.SecurityMarketsCodes
}

func (s *YuqingFinanceEvent) GetSpamScore() *float64 {
	return s.SpamScore
}

func (s *YuqingFinanceEvent) GetUserSubscribeEntityTags() []*string {
	return s.UserSubscribeEntityTags
}

func (s *YuqingFinanceEvent) GetUserSubscribeEventTags() []*int64 {
	return s.UserSubscribeEventTags
}

func (s *YuqingFinanceEvent) SetComprehensiveRisk(v float64) *YuqingFinanceEvent {
	s.ComprehensiveRisk = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityArea(v string) *YuqingFinanceEvent {
	s.EntityArea = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityCrn(v string) *YuqingFinanceEvent {
	s.EntityCrn = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityEmotionScore(v float64) *YuqingFinanceEvent {
	s.EntityEmotionScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityId(v int64) *YuqingFinanceEvent {
	s.EntityId = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityName(v string) *YuqingFinanceEvent {
	s.EntityName = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityRelevanceScore(v float64) *YuqingFinanceEvent {
	s.EntityRelevanceScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityShowName(v string) *YuqingFinanceEvent {
	s.EntityShowName = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntitySummary(v string) *YuqingFinanceEvent {
	s.EntitySummary = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityType(v string) *YuqingFinanceEvent {
	s.EntityType = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventId(v string) *YuqingFinanceEvent {
	s.EventId = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventLevel3Code(v int64) *YuqingFinanceEvent {
	s.EventLevel3Code = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventLevel3Name(v string) *YuqingFinanceEvent {
	s.EventLevel3Name = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventTags(v string) *YuqingFinanceEvent {
	s.EventTags = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventTime(v int64) *YuqingFinanceEvent {
	s.EventTime = &v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityAbbreviation(v string) *YuqingFinanceEvent {
	s.SecurityAbbreviation = &v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityCategoryCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityCategoryCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityMarketsCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityMarketsCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSpamScore(v float64) *YuqingFinanceEvent {
	s.SpamScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetUserSubscribeEntityTags(v []*string) *YuqingFinanceEvent {
	s.UserSubscribeEntityTags = v
	return s
}

func (s *YuqingFinanceEvent) SetUserSubscribeEventTags(v []*int64) *YuqingFinanceEvent {
	s.UserSubscribeEventTags = v
	return s
}

func (s *YuqingFinanceEvent) Validate() error {
	return dara.Validate(s)
}
