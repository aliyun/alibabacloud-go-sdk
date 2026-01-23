// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignCommandShrink(v string) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest
	GetAssignCommandShrink() *string
	SetOpTenantId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest
	GetOpTenantId() *int64
}

type AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest struct {
	// This parameter is required.
	AssignCommandShrink *string `json:"AssignCommand,omitempty" xml:"AssignCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) GetAssignCommandShrink() *string {
	return s.AssignCommandShrink
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) SetAssignCommandShrink(v string) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest {
	s.AssignCommandShrink = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) SetOpTenantId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
