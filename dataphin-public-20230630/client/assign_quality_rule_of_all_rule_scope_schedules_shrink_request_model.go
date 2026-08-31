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
	SetOpUserId(v string) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest
	GetOpUserId() *string
}

type AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest struct {
	// The assignment binding command.
	//
	// This parameter is required.
	AssignCommandShrink *string `json:"AssignCommand,omitempty" xml:"AssignCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) SetAssignCommandShrink(v string) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest {
	s.AssignCommandShrink = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) SetOpTenantId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) SetOpUserId(v string) *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
