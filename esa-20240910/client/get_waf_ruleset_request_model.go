// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetWafRulesetRequest
	GetId() *int64
	SetPhase(v string) *GetWafRulesetRequest
	GetPhase() *string
	SetSiteId(v int64) *GetWafRulesetRequest
	GetSiteId() *int64
}

type GetWafRulesetRequest struct {
	// The ID of the WAF ruleset. You can call the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The WAF rule execution phase. Valid values:
	//
	// - http_whitelist: whitelist rules
	//
	// - http_custom: custom rules
	//
	// - http_managed: managed rules
	//
	// - http_anti_scan: scan protection rules
	//
	// - http_ratelimit: rate limiting rules
	//
	// - ip_access_rule: IP access rules
	//
	// - http_bot: advanced mode bots
	//
	// - http_security_level_rule: security rules
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *GetWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *GetWafRulesetRequest) GetPhase() *string {
	return s.Phase
}

func (s *GetWafRulesetRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetWafRulesetRequest) SetId(v int64) *GetWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *GetWafRulesetRequest) SetPhase(v string) *GetWafRulesetRequest {
	s.Phase = &v
	return s
}

func (s *GetWafRulesetRequest) SetSiteId(v int64) *GetWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *GetWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
