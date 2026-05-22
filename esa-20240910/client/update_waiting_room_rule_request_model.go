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
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq "video.example.com")
	//
	// This parameter is required.
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding global configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room bypass rule to be updated, which can be obtained by calling the [ListWaitingRoomRules](https://help.aliyun.com/document_detail/2850279.html) interface.
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
