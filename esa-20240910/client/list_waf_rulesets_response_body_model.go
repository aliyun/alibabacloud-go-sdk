// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceUsage(v int64) *ListWafRulesetsResponseBody
	GetInstanceUsage() *int64
	SetPageNumber(v int32) *ListWafRulesetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListWafRulesetsResponseBody
	GetRequestId() *string
	SetRulesets(v []*ListWafRulesetsResponseBodyRulesets) *ListWafRulesetsResponseBody
	GetRulesets() []*ListWafRulesetsResponseBodyRulesets
	SetSiteUsage(v int64) *ListWafRulesetsResponseBody
	GetSiteUsage() *int64
	SetTotalCount(v int64) *ListWafRulesetsResponseBody
	GetTotalCount() *int64
}

type ListWafRulesetsResponseBody struct {
	// Number of WAF rule sets used by the instance in this WAF operation phase.
	//
	// example:
	//
	// 10
	InstanceUsage *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of rule set information, containing detailed information about the rule sets.
	Rulesets []*ListWafRulesetsResponseBodyRulesets `json:"Rulesets,omitempty" xml:"Rulesets,omitempty" type:"Repeated"`
	// Number of WAF rule sets used by the site in this WAF operation phase.
	//
	// example:
	//
	// 5
	SiteUsage *int64 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	// Total number of filtered records.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWafRulesetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponseBody) GetInstanceUsage() *int64 {
	return s.InstanceUsage
}

func (s *ListWafRulesetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafRulesetsResponseBody) GetRulesets() []*ListWafRulesetsResponseBodyRulesets {
	return s.Rulesets
}

func (s *ListWafRulesetsResponseBody) GetSiteUsage() *int64 {
	return s.SiteUsage
}

func (s *ListWafRulesetsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWafRulesetsResponseBody) SetInstanceUsage(v int64) *ListWafRulesetsResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetPageNumber(v int32) *ListWafRulesetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetPageSize(v int32) *ListWafRulesetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetRequestId(v string) *ListWafRulesetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetRulesets(v []*ListWafRulesetsResponseBodyRulesets) *ListWafRulesetsResponseBody {
	s.Rulesets = v
	return s
}

func (s *ListWafRulesetsResponseBody) SetSiteUsage(v int64) *ListWafRulesetsResponseBody {
	s.SiteUsage = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetTotalCount(v int64) *ListWafRulesetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWafRulesetsResponseBody) Validate() error {
	if s.Rulesets != nil {
		for _, item := range s.Rulesets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWafRulesetsResponseBodyRulesets struct {
	// List of match objects.
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// ID of the WAF rule set.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the rule set.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// WAF operation phase.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Status of the rule set.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Protection target type in http_bot.
	//
	// example:
	//
	// web
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// List of rule types.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// Last modification time of the rule set.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWafRulesetsResponseBodyRulesets) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsResponseBodyRulesets) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponseBodyRulesets) GetFields() []*string {
	return s.Fields
}

func (s *ListWafRulesetsResponseBodyRulesets) GetId() *int64 {
	return s.Id
}

func (s *ListWafRulesetsResponseBodyRulesets) GetName() *string {
	return s.Name
}

func (s *ListWafRulesetsResponseBodyRulesets) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesetsResponseBodyRulesets) GetStatus() *string {
	return s.Status
}

func (s *ListWafRulesetsResponseBodyRulesets) GetTarget() *string {
	return s.Target
}

func (s *ListWafRulesetsResponseBodyRulesets) GetTypes() []*string {
	return s.Types
}

func (s *ListWafRulesetsResponseBodyRulesets) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListWafRulesetsResponseBodyRulesets) SetFields(v []*string) *ListWafRulesetsResponseBodyRulesets {
	s.Fields = v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetId(v int64) *ListWafRulesetsResponseBodyRulesets {
	s.Id = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetName(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Name = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetPhase(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetStatus(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Status = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetTarget(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Target = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetTypes(v []*string) *ListWafRulesetsResponseBodyRulesets {
	s.Types = v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetUpdateTime(v string) *ListWafRulesetsResponseBodyRulesets {
	s.UpdateTime = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) Validate() error {
	return dara.Validate(s)
}
