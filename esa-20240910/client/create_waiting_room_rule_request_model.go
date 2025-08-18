// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v string) *CreateWaitingRoomRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateWaitingRoomRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateWaitingRoomRuleRequest
	GetRuleName() *string
	SetSiteId(v int64) *CreateWaitingRoomRuleRequest
	GetSiteId() *int64
	SetWaitingRoomId(v string) *CreateWaitingRoomRuleRequest
	GetWaitingRoomId() *string
}

type CreateWaitingRoomRuleRequest struct {
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
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
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
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
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room to bypass.
	//
	// This parameter is required.
	//
	// example:
	//
	// 25133f536f1b1f6b6091f6a92c614dd4
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s CreateWaitingRoomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateWaitingRoomRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateWaitingRoomRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateWaitingRoomRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWaitingRoomRuleRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *CreateWaitingRoomRuleRequest) SetRule(v string) *CreateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleEnable(v string) *CreateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleName(v string) *CreateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetSiteId(v int64) *CreateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetWaitingRoomId(v string) *CreateWaitingRoomRuleRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) Validate() error {
	return dara.Validate(s)
}
