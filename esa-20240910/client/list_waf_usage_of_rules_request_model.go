// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafUsageOfRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhase(v string) *ListWafUsageOfRulesRequest
	GetPhase() *string
	SetSiteId(v int64) *ListWafUsageOfRulesRequest
	GetSiteId() *int64
}

type ListWafUsageOfRulesRequest struct {
	// Name of the WAF operation phase.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
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

func (s *ListWafUsageOfRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafUsageOfRulesRequest) GetSiteId() *int64 {
	return s.SiteId
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
