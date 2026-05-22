// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafTemplateRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListWafTemplateRulesShrinkRequest
	GetInstanceId() *string
	SetPhase(v string) *ListWafTemplateRulesShrinkRequest
	GetPhase() *string
	SetQueryArgsShrink(v string) *ListWafTemplateRulesShrinkRequest
	GetQueryArgsShrink() *string
	SetSiteId(v int64) *ListWafTemplateRulesShrinkRequest
	GetSiteId() *int64
}

type ListWafTemplateRulesShrinkRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// WAF operation phase, used to filter template rules for a specific phase.
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Query parameters, used to filter template rules based on conditions such as rule type.
	//
	// example:
	//
	// http_anti_scan
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafTemplateRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWafTemplateRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafTemplateRulesShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListWafTemplateRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafTemplateRulesShrinkRequest) SetInstanceId(v string) *ListWafTemplateRulesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWafTemplateRulesShrinkRequest) SetPhase(v string) *ListWafTemplateRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesShrinkRequest) SetQueryArgsShrink(v string) *ListWafTemplateRulesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafTemplateRulesShrinkRequest) SetSiteId(v int64) *ListWafTemplateRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafTemplateRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
