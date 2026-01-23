// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateQualityRuleSwitchRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateQualityRuleSwitchRequestUpdateCommand) *UpdateQualityRuleSwitchRequest
	GetUpdateCommand() *UpdateQualityRuleSwitchRequestUpdateCommand
}

type UpdateQualityRuleSwitchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateQualityRuleSwitchRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateQualityRuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleSwitchRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateQualityRuleSwitchRequest) GetUpdateCommand() *UpdateQualityRuleSwitchRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateQualityRuleSwitchRequest) SetOpTenantId(v int64) *UpdateQualityRuleSwitchRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateQualityRuleSwitchRequest) SetUpdateCommand(v *UpdateQualityRuleSwitchRequestUpdateCommand) *UpdateQualityRuleSwitchRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateQualityRuleSwitchRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateQualityRuleSwitchRequestUpdateCommand struct {
	// This parameter is required.
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// This parameter is required.
	RuleIdList []*int64 `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
}

func (s UpdateQualityRuleSwitchRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleSwitchRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleSwitchRequestUpdateCommand) GetOpen() *bool {
	return s.Open
}

func (s *UpdateQualityRuleSwitchRequestUpdateCommand) GetRuleIdList() []*int64 {
	return s.RuleIdList
}

func (s *UpdateQualityRuleSwitchRequestUpdateCommand) SetOpen(v bool) *UpdateQualityRuleSwitchRequestUpdateCommand {
	s.Open = &v
	return s
}

func (s *UpdateQualityRuleSwitchRequestUpdateCommand) SetRuleIdList(v []*int64) *UpdateQualityRuleSwitchRequestUpdateCommand {
	s.RuleIdList = v
	return s
}

func (s *UpdateQualityRuleSwitchRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
