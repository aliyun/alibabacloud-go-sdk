// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignQualityRuleOfAllRuleScopeSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignCommand(v *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) *AssignQualityRuleOfAllRuleScopeSchedulesRequest
	GetAssignCommand() *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand
	SetOpTenantId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesRequest
	GetOpTenantId() *int64
}

type AssignQualityRuleOfAllRuleScopeSchedulesRequest struct {
	// This parameter is required.
	AssignCommand *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand `json:"AssignCommand,omitempty" xml:"AssignCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesRequest) GoString() string {
	return s.String()
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequest) GetAssignCommand() *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand {
	return s.AssignCommand
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequest) SetAssignCommand(v *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) *AssignQualityRuleOfAllRuleScopeSchedulesRequest {
	s.AssignCommand = v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequest) SetOpTenantId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequest) Validate() error {
	if s.AssignCommand != nil {
		if err := s.AssignCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand struct {
	// This parameter is required.
	RuleIdList []*int64 `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	ScheduleIdList []*int64 `json:"ScheduleIdList,omitempty" xml:"ScheduleIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) String() string {
	return dara.Prettify(s)
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) GoString() string {
	return s.String()
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) GetRuleIdList() []*int64 {
	return s.RuleIdList
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) GetScheduleIdList() []*int64 {
	return s.ScheduleIdList
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) SetRuleIdList(v []*int64) *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand {
	s.RuleIdList = v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) SetScheduleIdList(v []*int64) *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand {
	s.ScheduleIdList = v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) SetWatchId(v int64) *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand {
	s.WatchId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesRequestAssignCommand) Validate() error {
	return dara.Validate(s)
}
