// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateWafRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *BatchUpdateWafRulesShrinkRequest
	GetConfigsShrink() *string
	SetPhase(v string) *BatchUpdateWafRulesShrinkRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchUpdateWafRulesShrinkRequest
	GetRulesetId() *int64
	SetSharedShrink(v string) *BatchUpdateWafRulesShrinkRequest
	GetSharedShrink() *string
	SetSiteId(v int64) *BatchUpdateWafRulesShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchUpdateWafRulesShrinkRequest
	GetSiteVersion() *int32
}

type BatchUpdateWafRulesShrinkRequest struct {
	// The configurations of rules.
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The WAF rule category.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF ruleset, which can be obtained by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The configurations shared by multiple rules.
	//
	// example:
	//
	// 10000001
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version of the website.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchUpdateWafRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *BatchUpdateWafRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchUpdateWafRulesShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchUpdateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetPhase(v string) *BatchUpdateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetRulesetId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchUpdateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
