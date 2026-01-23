// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityScheduleRequest
	GetOpTenantId() *int64
	SetUpsertCommand(v *UpsertQualityScheduleRequestUpsertCommand) *UpsertQualityScheduleRequest
	GetUpsertCommand() *UpsertQualityScheduleRequestUpsertCommand
}

type UpsertQualityScheduleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommand *UpsertQualityScheduleRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityScheduleRequest) GetUpsertCommand() *UpsertQualityScheduleRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityScheduleRequest) SetOpTenantId(v int64) *UpsertQualityScheduleRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityScheduleRequest) SetUpsertCommand(v *UpsertQualityScheduleRequestUpsertCommand) *UpsertQualityScheduleRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityScheduleRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityScheduleRequestUpsertCommand struct {
	// example:
	//
	// 	- 	- 1/	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ds=${yyyyMMdd}
	PartitionExpression *string `json:"PartitionExpression,omitempty" xml:"PartitionExpression,omitempty"`
	// example:
	//
	// CUSTOM
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// example:
	//
	// DAILY
	PeriodScheduleIntervalType *string   `json:"PeriodScheduleIntervalType,omitempty" xml:"PeriodScheduleIntervalType,omitempty"`
	PeriodScheduleParamList    []*string `json:"PeriodScheduleParamList,omitempty" xml:"PeriodScheduleParamList,omitempty" type:"Repeated"`
	// example:
	//
	// ONE_TASKS_FINISHED
	StaticTaskTriggerType *string   `json:"StaticTaskTriggerType,omitempty" xml:"StaticTaskTriggerType,omitempty"`
	TriggerNodeList       []*string `json:"TriggerNodeList,omitempty" xml:"TriggerNodeList,omitempty" type:"Repeated"`
	// example:
	//
	// STATIC_TASK_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERIOD_SCHEDULE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// TASK_REFERRED_PARTITION
	ValidatePartitionType *string `json:"ValidatePartitionType,omitempty" xml:"ValidatePartitionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s UpsertQualityScheduleRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetName() *string {
	return s.Name
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetPartitionExpression() *string {
	return s.PartitionExpression
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetPartitionType() *string {
	return s.PartitionType
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetPeriodScheduleIntervalType() *string {
	return s.PeriodScheduleIntervalType
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetPeriodScheduleParamList() []*string {
	return s.PeriodScheduleParamList
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetStaticTaskTriggerType() *string {
	return s.StaticTaskTriggerType
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetTriggerNodeList() []*string {
	return s.TriggerNodeList
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetType() *string {
	return s.Type
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetValidatePartitionType() *string {
	return s.ValidatePartitionType
}

func (s *UpsertQualityScheduleRequestUpsertCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetCronExpression(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.CronExpression = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetId(v int64) *UpsertQualityScheduleRequestUpsertCommand {
	s.Id = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetName(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.Name = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetPartitionExpression(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.PartitionExpression = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetPartitionType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.PartitionType = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetPeriodScheduleIntervalType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.PeriodScheduleIntervalType = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetPeriodScheduleParamList(v []*string) *UpsertQualityScheduleRequestUpsertCommand {
	s.PeriodScheduleParamList = v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetStaticTaskTriggerType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.StaticTaskTriggerType = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetTriggerNodeList(v []*string) *UpsertQualityScheduleRequestUpsertCommand {
	s.TriggerNodeList = v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetTriggerType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.TriggerType = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.Type = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetValidatePartitionType(v string) *UpsertQualityScheduleRequestUpsertCommand {
	s.ValidatePartitionType = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) SetWatchId(v int64) *UpsertQualityScheduleRequestUpsertCommand {
	s.WatchId = &v
	return s
}

func (s *UpsertQualityScheduleRequestUpsertCommand) Validate() error {
	return dara.Validate(s)
}
