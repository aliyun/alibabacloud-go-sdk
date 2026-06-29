// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v string) *UpdateWaitingRoomRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateWaitingRoomRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateWaitingRoomRuleRequest
	GetRuleName() *string
	SetSiteId(v int64) *UpdateWaitingRoomRuleRequest
	GetSiteId() *int64
	SetWaitingRoomRuleId(v int64) *UpdateWaitingRoomRuleRequest
	GetWaitingRoomRuleId() *int64
}

type UpdateWaitingRoomRuleRequest struct {
	// The rule content. A conditional expression is used to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specified requests: set the value to a custom expression, for example, (http.host eq \\"video.example.com\\").
	//
	// For the complete syntax of rule expressions, refer to
	//
	// <props="china">https://www.alibabacloud.com/help/en/edge-security-acceleration/esa/user-guide/work-with-rules-engine/
	//
	// <props="intl">https://www.alibabacloud.com/help/edge-security-acceleration/esa/user-guide/work-with-rules-engine/
	//
	// This parameter is required.
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is not required when you add a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room bypass rule to update. You can obtain this ID after creating a rule by calling CreateWaitingRoomRule, or by calling the [ListWaitingRoomRules](https://help.aliyun.com/document_detail/2850279.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8987739839****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s UpdateWaitingRoomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateWaitingRoomRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateWaitingRoomRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateWaitingRoomRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWaitingRoomRuleRequest) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *UpdateWaitingRoomRuleRequest) SetRule(v string) *UpdateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleEnable(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleName(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetSiteId(v int64) *UpdateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetWaitingRoomRuleId(v int64) *UpdateWaitingRoomRuleRequest {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) Validate() error {
	return dara.Validate(s)
}
