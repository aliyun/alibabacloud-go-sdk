// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchRepeatJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *CreateBatchRepeatJobRequest
	GetCallingNumber() []*string
	SetDescription(v string) *CreateBatchRepeatJobRequest
	GetDescription() *string
	SetFilterStatus(v string) *CreateBatchRepeatJobRequest
	GetFilterStatus() *string
	SetFlashSmsExtras(v string) *CreateBatchRepeatJobRequest
	GetFlashSmsExtras() *string
	SetInstanceId(v string) *CreateBatchRepeatJobRequest
	GetInstanceId() *string
	SetMinConcurrency(v int64) *CreateBatchRepeatJobRequest
	GetMinConcurrency() *int64
	SetName(v string) *CreateBatchRepeatJobRequest
	GetName() *string
	SetPriority(v string) *CreateBatchRepeatJobRequest
	GetPriority() *string
	SetRecallCallingNumber(v []*string) *CreateBatchRepeatJobRequest
	GetRecallCallingNumber() []*string
	SetRecallStrategyJson(v string) *CreateBatchRepeatJobRequest
	GetRecallStrategyJson() *string
	SetRingingDuration(v int64) *CreateBatchRepeatJobRequest
	GetRingingDuration() *int64
	SetScriptId(v string) *CreateBatchRepeatJobRequest
	GetScriptId() *string
	SetSourceGroupId(v string) *CreateBatchRepeatJobRequest
	GetSourceGroupId() *string
	SetStrategyJson(v string) *CreateBatchRepeatJobRequest
	GetStrategyJson() *string
}

type CreateBatchRepeatJobRequest struct {
	// example:
	//
	// []
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Succeeded,NoInteraction,Failed,Cancelled
	FilterStatus   *string `json:"FilterStatus,omitempty" xml:"FilterStatus,omitempty"`
	FlashSmsExtras *string `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2bfa5ae4-7185-4227-a3b8-328f26f11be1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Priority            *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"emptyNumberIgnore\\":false,\\"inArrearsIgnore\\":false,\\"outOfServiceIgnore\\":false}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// example:
	//
	// 0fe7f71c-8771-42ef-9bb1-19aa16ae7120
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c4f8a3d3-2e94-4bd4-aef8-e35f663d4847
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// example:
	//
	// {\\"maxAttemptsPerDay\\":1,\\"minAttemptInterval\\":1,\\"routingStrategy\\":\\"LocalFirst\\"}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s CreateBatchRepeatJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRepeatJobRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchRepeatJobRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *CreateBatchRepeatJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchRepeatJobRequest) GetFilterStatus() *string {
	return s.FilterStatus
}

func (s *CreateBatchRepeatJobRequest) GetFlashSmsExtras() *string {
	return s.FlashSmsExtras
}

func (s *CreateBatchRepeatJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBatchRepeatJobRequest) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *CreateBatchRepeatJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateBatchRepeatJobRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateBatchRepeatJobRequest) GetRecallCallingNumber() []*string {
	return s.RecallCallingNumber
}

func (s *CreateBatchRepeatJobRequest) GetRecallStrategyJson() *string {
	return s.RecallStrategyJson
}

func (s *CreateBatchRepeatJobRequest) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *CreateBatchRepeatJobRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBatchRepeatJobRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *CreateBatchRepeatJobRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *CreateBatchRepeatJobRequest) SetCallingNumber(v []*string) *CreateBatchRepeatJobRequest {
	s.CallingNumber = v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetDescription(v string) *CreateBatchRepeatJobRequest {
	s.Description = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetFilterStatus(v string) *CreateBatchRepeatJobRequest {
	s.FilterStatus = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetFlashSmsExtras(v string) *CreateBatchRepeatJobRequest {
	s.FlashSmsExtras = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetInstanceId(v string) *CreateBatchRepeatJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetMinConcurrency(v int64) *CreateBatchRepeatJobRequest {
	s.MinConcurrency = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetName(v string) *CreateBatchRepeatJobRequest {
	s.Name = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetPriority(v string) *CreateBatchRepeatJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetRecallCallingNumber(v []*string) *CreateBatchRepeatJobRequest {
	s.RecallCallingNumber = v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetRecallStrategyJson(v string) *CreateBatchRepeatJobRequest {
	s.RecallStrategyJson = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetRingingDuration(v int64) *CreateBatchRepeatJobRequest {
	s.RingingDuration = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetScriptId(v string) *CreateBatchRepeatJobRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetSourceGroupId(v string) *CreateBatchRepeatJobRequest {
	s.SourceGroupId = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) SetStrategyJson(v string) *CreateBatchRepeatJobRequest {
	s.StrategyJson = &v
	return s
}

func (s *CreateBatchRepeatJobRequest) Validate() error {
	return dara.Validate(s)
}
