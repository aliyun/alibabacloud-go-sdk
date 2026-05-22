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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// http_anti_scan
	Phase  *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	SiteId *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
