// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserWafRulesetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserWafRulesetsShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListUserWafRulesetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserWafRulesetsShrinkRequest
	GetPageSize() *int32
	SetPhase(v string) *ListUserWafRulesetsShrinkRequest
	GetPhase() *string
	SetQueryArgsShrink(v string) *ListUserWafRulesetsShrinkRequest
	GetQueryArgsShrink() *string
}

type ListUserWafRulesetsShrinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies the execution phase of the WAF rule.
	//
	// - `http_whitelist`: whitelist rule
	//
	// - `http_custom`: custom rule
	//
	// - `http_managed`: managed rule
	//
	// - `http_anti_scan`: anti-scan rule
	//
	// - `http_ratelimit`: rate limit rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_bot`: advanced mode bot
	//
	// - `http_security_level_rule`: security rule
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Parameters for filtering and sorting the results.
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListUserWafRulesetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserWafRulesetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserWafRulesetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserWafRulesetsShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListUserWafRulesetsShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListUserWafRulesetsShrinkRequest) SetInstanceId(v string) *ListUserWafRulesetsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserWafRulesetsShrinkRequest) SetPageNumber(v int32) *ListUserWafRulesetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserWafRulesetsShrinkRequest) SetPageSize(v int32) *ListUserWafRulesetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserWafRulesetsShrinkRequest) SetPhase(v string) *ListUserWafRulesetsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListUserWafRulesetsShrinkRequest) SetQueryArgsShrink(v string) *ListUserWafRulesetsShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListUserWafRulesetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
