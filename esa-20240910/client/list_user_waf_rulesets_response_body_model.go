// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserWafRulesetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceUsage(v int64) *ListUserWafRulesetsResponseBody
	GetInstanceUsage() *int64
	SetPageNumber(v int32) *ListUserWafRulesetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserWafRulesetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserWafRulesetsResponseBody
	GetRequestId() *string
	SetRulesets(v []*ListUserWafRulesetsResponseBodyRulesets) *ListUserWafRulesetsResponseBody
	GetRulesets() []*ListUserWafRulesetsResponseBodyRulesets
	SetTotalCount(v int64) *ListUserWafRulesetsResponseBody
	GetTotalCount() *int64
}

type ListUserWafRulesetsResponseBody struct {
	// The instance usage.
	//
	// example:
	//
	// 10
	InstanceUsage *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of WAF ruleset objects.
	Rulesets []*ListUserWafRulesetsResponseBodyRulesets `json:"Rulesets,omitempty" xml:"Rulesets,omitempty" type:"Repeated"`
	// The total number of records after filtering.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserWafRulesetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsResponseBody) GetInstanceUsage() *int64 {
	return s.InstanceUsage
}

func (s *ListUserWafRulesetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserWafRulesetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserWafRulesetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserWafRulesetsResponseBody) GetRulesets() []*ListUserWafRulesetsResponseBodyRulesets {
	return s.Rulesets
}

func (s *ListUserWafRulesetsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserWafRulesetsResponseBody) SetInstanceUsage(v int64) *ListUserWafRulesetsResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListUserWafRulesetsResponseBody) SetPageNumber(v int32) *ListUserWafRulesetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserWafRulesetsResponseBody) SetPageSize(v int32) *ListUserWafRulesetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserWafRulesetsResponseBody) SetRequestId(v string) *ListUserWafRulesetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserWafRulesetsResponseBody) SetRulesets(v []*ListUserWafRulesetsResponseBodyRulesets) *ListUserWafRulesetsResponseBody {
	s.Rulesets = v
	return s
}

func (s *ListUserWafRulesetsResponseBody) SetTotalCount(v int64) *ListUserWafRulesetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserWafRulesetsResponseBody) Validate() error {
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

type ListUserWafRulesetsResponseBodyRulesets struct {
	// The WAF ruleset description.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The WAF ruleset ID.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The WAF ruleset name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The WAF rule execution phase. Possible values:
	//
	// - `http_whitelist`: Whitelist rule
	//
	// - `http_custom`: Custom rule
	//
	// - `http_managed`: Managed rule
	//
	// - `http_anti_scan`: Scan protection rule
	//
	// - `http_ratelimit`: Rate limit rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_bot`: Bot rule
	//
	// - `http_security_level_rule`: Security rule
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The WAF ruleset position.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// The WAF ruleset status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUserWafRulesetsResponseBodyRulesets) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsResponseBodyRulesets) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetDescription() *string {
	return s.Description
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetId() *int64 {
	return s.Id
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetName() *string {
	return s.Name
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetPhase() *string {
	return s.Phase
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetPosition() *int64 {
	return s.Position
}

func (s *ListUserWafRulesetsResponseBodyRulesets) GetStatus() *string {
	return s.Status
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetDescription(v string) *ListUserWafRulesetsResponseBodyRulesets {
	s.Description = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetId(v int64) *ListUserWafRulesetsResponseBodyRulesets {
	s.Id = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetName(v string) *ListUserWafRulesetsResponseBodyRulesets {
	s.Name = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetPhase(v string) *ListUserWafRulesetsResponseBodyRulesets {
	s.Phase = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetPosition(v int64) *ListUserWafRulesetsResponseBodyRulesets {
	s.Position = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) SetStatus(v string) *ListUserWafRulesetsResponseBodyRulesets {
	s.Status = &v
	return s
}

func (s *ListUserWafRulesetsResponseBodyRulesets) Validate() error {
	return dara.Validate(s)
}
