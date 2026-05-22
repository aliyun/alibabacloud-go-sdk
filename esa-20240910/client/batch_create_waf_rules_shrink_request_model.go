// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateWafRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *BatchCreateWafRulesShrinkRequest
	GetConfigsShrink() *string
	SetPhase(v string) *BatchCreateWafRulesShrinkRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchCreateWafRulesShrinkRequest
	GetRulesetId() *int64
	SetSharedShrink(v string) *BatchCreateWafRulesShrinkRequest
	GetSharedShrink() *string
	SetSiteId(v int64) *BatchCreateWafRulesShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchCreateWafRulesShrinkRequest
	GetSiteVersion() *int32
}

type BatchCreateWafRulesShrinkRequest struct {
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// example:
	//
	// http_custom
	Phase        *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	RulesetId    *int64  `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchCreateWafRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *BatchCreateWafRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchCreateWafRulesShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchCreateWafRulesShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *BatchCreateWafRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchCreateWafRulesShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchCreateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetPhase(v string) *BatchCreateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetRulesetId(v int64) *BatchCreateWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteId(v int64) *BatchCreateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchCreateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
