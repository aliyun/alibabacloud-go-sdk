// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DisableMetricRulesRequest
	GetRegionId() *string
	SetRuleId(v []*string) *DisableMetricRulesRequest
	GetRuleId() []*string
}

type DisableMetricRulesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule. Valid values of N: 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// detect_87****_HTTP_HttpLatency
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DisableMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableMetricRulesRequest) GetRuleId() []*string {
	return s.RuleId
}

func (s *DisableMetricRulesRequest) SetRegionId(v string) *DisableMetricRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DisableMetricRulesRequest) SetRuleId(v []*string) *DisableMetricRulesRequest {
	s.RuleId = v
	return s
}

func (s *DisableMetricRulesRequest) Validate() error {
	return dara.Validate(s)
}
