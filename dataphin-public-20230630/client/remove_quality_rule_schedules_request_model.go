// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQualityRuleSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveQualityRuleSchedulesRequest
	GetOpTenantId() *int64
	SetRemoveCommand(v *RemoveQualityRuleSchedulesRequestRemoveCommand) *RemoveQualityRuleSchedulesRequest
	GetRemoveCommand() *RemoveQualityRuleSchedulesRequestRemoveCommand
}

type RemoveQualityRuleSchedulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveQualityRuleSchedulesRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveQualityRuleSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveQualityRuleSchedulesRequest) GoString() string {
	return s.String()
}

func (s *RemoveQualityRuleSchedulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveQualityRuleSchedulesRequest) GetRemoveCommand() *RemoveQualityRuleSchedulesRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveQualityRuleSchedulesRequest) SetOpTenantId(v int64) *RemoveQualityRuleSchedulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveQualityRuleSchedulesRequest) SetRemoveCommand(v *RemoveQualityRuleSchedulesRequestRemoveCommand) *RemoveQualityRuleSchedulesRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveQualityRuleSchedulesRequest) Validate() error {
	if s.RemoveCommand != nil {
		if err := s.RemoveCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveQualityRuleSchedulesRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// This parameter is required.
	ScheduleIdList []*int64 `json:"ScheduleIdList,omitempty" xml:"ScheduleIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s RemoveQualityRuleSchedulesRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveQualityRuleSchedulesRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) GetRuleId() *int64 {
	return s.RuleId
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) GetScheduleIdList() []*int64 {
	return s.ScheduleIdList
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) SetRuleId(v int64) *RemoveQualityRuleSchedulesRequestRemoveCommand {
	s.RuleId = &v
	return s
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) SetScheduleIdList(v []*int64) *RemoveQualityRuleSchedulesRequestRemoveCommand {
	s.ScheduleIdList = v
	return s
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) SetWatchId(v int64) *RemoveQualityRuleSchedulesRequestRemoveCommand {
	s.WatchId = &v
	return s
}

func (s *RemoveQualityRuleSchedulesRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}
