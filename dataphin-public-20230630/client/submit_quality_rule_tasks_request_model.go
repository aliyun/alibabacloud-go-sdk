// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityRuleTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitQualityRuleTasksRequest
	GetOpTenantId() *int64
	SetSubmitCommand(v *SubmitQualityRuleTasksRequestSubmitCommand) *SubmitQualityRuleTasksRequest
	GetSubmitCommand() *SubmitQualityRuleTasksRequestSubmitCommand
}

type SubmitQualityRuleTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitCommand *SubmitQualityRuleTasksRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitQualityRuleTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitQualityRuleTasksRequest) GetSubmitCommand() *SubmitQualityRuleTasksRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitQualityRuleTasksRequest) SetOpTenantId(v int64) *SubmitQualityRuleTasksRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitQualityRuleTasksRequest) SetSubmitCommand(v *SubmitQualityRuleTasksRequestSubmitCommand) *SubmitQualityRuleTasksRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitQualityRuleTasksRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitQualityRuleTasksRequestSubmitCommand struct {
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// This parameter is required.
	IsTestRun *bool `json:"IsTestRun,omitempty" xml:"IsTestRun,omitempty"`
	// example:
	//
	// ds=${yyyyMMdd}
	PartitionExpression *string `json:"PartitionExpression,omitempty" xml:"PartitionExpression,omitempty"`
	// example:
	//
	// CUSTOM
	PartitionExpressionFrom *string `json:"PartitionExpressionFrom,omitempty" xml:"PartitionExpressionFrom,omitempty"`
	// example:
	//
	// 1
	ScheduleId *int64 `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// This parameter is required.
	WatchRuleIdList []*SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList `json:"WatchRuleIdList,omitempty" xml:"WatchRuleIdList,omitempty" type:"Repeated"`
}

func (s SubmitQualityRuleTasksRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetBizDate() *string {
	return s.BizDate
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetIsTestRun() *bool {
	return s.IsTestRun
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetPartitionExpression() *string {
	return s.PartitionExpression
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetPartitionExpressionFrom() *string {
	return s.PartitionExpressionFrom
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) GetWatchRuleIdList() []*SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList {
	return s.WatchRuleIdList
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetBizDate(v string) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.BizDate = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetIsTestRun(v bool) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.IsTestRun = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetPartitionExpression(v string) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.PartitionExpression = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetPartitionExpressionFrom(v string) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.PartitionExpressionFrom = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetScheduleId(v int64) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.ScheduleId = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) SetWatchRuleIdList(v []*SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) *SubmitQualityRuleTasksRequestSubmitCommand {
	s.WatchRuleIdList = v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommand) Validate() error {
	if s.WatchRuleIdList != nil {
		for _, item := range s.WatchRuleIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) GetWatchId() *int64 {
	return s.WatchId
}

func (s *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) SetRuleId(v int64) *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList {
	s.RuleId = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) SetWatchId(v int64) *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList {
	s.WatchId = &v
	return s
}

func (s *SubmitQualityRuleTasksRequestSubmitCommandWatchRuleIdList) Validate() error {
	return dara.Validate(s)
}
