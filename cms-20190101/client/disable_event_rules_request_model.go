// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEventRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DisableEventRulesRequest
	GetRegionId() *string
	SetRuleNames(v []*string) *DisableEventRulesRequest
	GetRuleNames() []*string
}

type DisableEventRulesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ruleName1
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DisableEventRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableEventRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableEventRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableEventRulesRequest) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *DisableEventRulesRequest) SetRegionId(v string) *DisableEventRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DisableEventRulesRequest) SetRuleNames(v []*string) *DisableEventRulesRequest {
	s.RuleNames = v
	return s
}

func (s *DisableEventRulesRequest) Validate() error {
	return dara.Validate(s)
}
