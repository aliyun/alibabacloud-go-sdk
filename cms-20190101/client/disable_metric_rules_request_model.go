// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DisableMetricRulesRequest
	GetDryRun() *bool
	SetRegionId(v string) *DisableMetricRulesRequest
	GetRegionId() *string
	SetRuleId(v []*string) *DisableMetricRulesRequest
	GetRuleId() []*string
}

type DisableMetricRulesRequest struct {
	// Specifies whether to perform a dry run of the request. Valid values:
	//
	// - **true**: sends a dry run request and does not bind a security group to the instance. Check items include whether required parameters are specified, the request format, and service limits. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// This parameter is required.
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DisableMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableMetricRulesRequest) GetRuleId() []*string {
	return s.RuleId
}

func (s *DisableMetricRulesRequest) SetDryRun(v bool) *DisableMetricRulesRequest {
	s.DryRun = &v
	return s
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
