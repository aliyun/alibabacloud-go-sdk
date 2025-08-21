// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DeleteSchedulerRuleRequest
	GetResourceGroupId() *string
	SetRuleName(v string) *DeleteSchedulerRuleRequest
	GetRuleName() *string
}

type DeleteSchedulerRuleRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the scheduling rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteSchedulerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteSchedulerRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteSchedulerRuleRequest) SetResourceGroupId(v string) *DeleteSchedulerRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSchedulerRuleRequest) SetRuleName(v string) *DeleteSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteSchedulerRuleRequest) Validate() error {
	return dara.Validate(s)
}
