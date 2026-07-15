// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafUsageOfRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListWafUsageOfRulesRequest
	GetInstanceId() *string
	SetPhase(v string) *ListWafUsageOfRulesRequest
	GetPhase() *string
	SetSiteId(v int64) *ListWafUsageOfRulesRequest
	GetSiteId() *int64
}

type ListWafUsageOfRulesRequest struct {
	// The WAF instance ID.
	//
	// If this parameter is left empty, the API returns an empty result. We recommend that you always specify this parameter.
	//
	// example:
	//
	// esa-site-b0wdsrlba3nk
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The phase in which the WAF rule runs. This parameter is required.
	//
	// Common values: http_custom, http_ratelimit, http_anti_scan, http_bot, http_managed, http_whitelist, and http_threat_intelligence.
	//
	// > Note: This parameter is required on the server side. If this parameter is not specified, the API returns InvalidParameter (400).
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafUsageOfRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafUsageOfRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWafUsageOfRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafUsageOfRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafUsageOfRulesRequest) SetInstanceId(v string) *ListWafUsageOfRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWafUsageOfRulesRequest) SetPhase(v string) *ListWafUsageOfRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafUsageOfRulesRequest) SetSiteId(v int64) *ListWafUsageOfRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafUsageOfRulesRequest) Validate() error {
	return dara.Validate(s)
}
