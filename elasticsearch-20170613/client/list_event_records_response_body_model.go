// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEventRecordsResponseBody
	GetRequestId() *string
	SetResult(v *ListEventRecordsResponseBodyResult) *ListEventRecordsResponseBody
	GetResult() *ListEventRecordsResponseBodyResult
}

type ListEventRecordsResponseBody struct {
	// example:
	//
	// 7F40EAA1-6F1D-4DD9-8DB8-C5F00C4E****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListEventRecordsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListEventRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventRecordsResponseBody) GetResult() *ListEventRecordsResponseBodyResult {
	return s.Result
}

func (s *ListEventRecordsResponseBody) SetRequestId(v string) *ListEventRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventRecordsResponseBody) SetResult(v *ListEventRecordsResponseBodyResult) *ListEventRecordsResponseBody {
	s.Result = v
	return s
}

func (s *ListEventRecordsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventRecordsResponseBodyResult struct {
	Result []*ListEventRecordsResponseBodyResultResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEventRecordsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponseBodyResult) GetResult() []*ListEventRecordsResponseBodyResultResult {
	return s.Result
}

func (s *ListEventRecordsResponseBodyResult) GetTotal() *string {
	return s.Total
}

func (s *ListEventRecordsResponseBodyResult) SetResult(v []*ListEventRecordsResponseBodyResultResult) *ListEventRecordsResponseBodyResult {
	s.Result = v
	return s
}

func (s *ListEventRecordsResponseBodyResult) SetTotal(v string) *ListEventRecordsResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListEventRecordsResponseBodyResult) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventRecordsResponseBodyResultResult struct {
	// example:
	//
	// false
	AutoAlarm *bool `json:"autoAlarm,omitempty" xml:"autoAlarm,omitempty"`
	// example:
	//
	// Instance.SpecModify
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// Info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// false
	MustOps *bool `json:"mustOps,omitempty" xml:"mustOps,omitempty"`
	// example:
	//
	// elasticsearch
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// example:
	//
	// 2025-05-08 18:41:01
	ScheduleExecuteTime *string `json:"scheduleExecuteTime,omitempty" xml:"scheduleExecuteTime,omitempty"`
	// example:
	//
	// 2025-05-08 18:41:01
	ScheduleFinishTime *string                                              `json:"scheduleFinishTime,omitempty" xml:"scheduleFinishTime,omitempty"`
	ShowContent        *ListEventRecordsResponseBodyResultResultShowContent `json:"showContent,omitempty" xml:"showContent,omitempty" type:"Struct"`
	// example:
	//
	// webConsole
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Executed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// UserOperator
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEventRecordsResponseBodyResultResult) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponseBodyResultResult) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponseBodyResultResult) GetAutoAlarm() *bool {
	return s.AutoAlarm
}

func (s *ListEventRecordsResponseBodyResultResult) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventRecordsResponseBodyResultResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListEventRecordsResponseBodyResultResult) GetLevel() *string {
	return s.Level
}

func (s *ListEventRecordsResponseBodyResultResult) GetMustOps() *bool {
	return s.MustOps
}

func (s *ListEventRecordsResponseBodyResultResult) GetProduct() *string {
	return s.Product
}

func (s *ListEventRecordsResponseBodyResultResult) GetScheduleExecuteTime() *string {
	return s.ScheduleExecuteTime
}

func (s *ListEventRecordsResponseBodyResultResult) GetScheduleFinishTime() *string {
	return s.ScheduleFinishTime
}

func (s *ListEventRecordsResponseBodyResultResult) GetShowContent() *ListEventRecordsResponseBodyResultResultShowContent {
	return s.ShowContent
}

func (s *ListEventRecordsResponseBodyResultResult) GetSource() *string {
	return s.Source
}

func (s *ListEventRecordsResponseBodyResultResult) GetStatus() *string {
	return s.Status
}

func (s *ListEventRecordsResponseBodyResultResult) GetType() *string {
	return s.Type
}

func (s *ListEventRecordsResponseBodyResultResult) SetAutoAlarm(v bool) *ListEventRecordsResponseBodyResultResult {
	s.AutoAlarm = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetDisplayName(v string) *ListEventRecordsResponseBodyResultResult {
	s.DisplayName = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetDryRun(v bool) *ListEventRecordsResponseBodyResultResult {
	s.DryRun = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetLevel(v string) *ListEventRecordsResponseBodyResultResult {
	s.Level = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetMustOps(v bool) *ListEventRecordsResponseBodyResultResult {
	s.MustOps = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetProduct(v string) *ListEventRecordsResponseBodyResultResult {
	s.Product = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetScheduleExecuteTime(v string) *ListEventRecordsResponseBodyResultResult {
	s.ScheduleExecuteTime = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetScheduleFinishTime(v string) *ListEventRecordsResponseBodyResultResult {
	s.ScheduleFinishTime = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetShowContent(v *ListEventRecordsResponseBodyResultResultShowContent) *ListEventRecordsResponseBodyResultResult {
	s.ShowContent = v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetSource(v string) *ListEventRecordsResponseBodyResultResult {
	s.Source = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetStatus(v string) *ListEventRecordsResponseBodyResultResult {
	s.Status = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) SetType(v string) *ListEventRecordsResponseBodyResultResult {
	s.Type = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResult) Validate() error {
	if s.ShowContent != nil {
		if err := s.ShowContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventRecordsResponseBodyResultResultShowContent struct {
	ActionSuggest *ListEventRecordsResponseBodyResultResultShowContentActionSuggest `json:"actionSuggest,omitempty" xml:"actionSuggest,omitempty" type:"Struct"`
	// example:
	//
	// Instance.SpecModify
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// Executed
	EventStatus *string `json:"eventStatus,omitempty" xml:"eventStatus,omitempty"`
	// example:
	//
	// 2025-05-08 18:31:01
	EventTime *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	// example:
	//
	// 2025-05-08 18:41:01
	ExecuteFinishTime *string `json:"executeFinishTime,omitempty" xml:"executeFinishTime,omitempty"`
	// example:
	//
	// 2025-05-08 18:31:00
	ExecuteStartTime *string `json:"executeStartTime,omitempty" xml:"executeStartTime,omitempty"`
	// example:
	//
	// es-cn-a5cb2dece****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// f183728022a1c769e97b4*****
	OpsChangeId *string `json:"opsChangeId,omitempty" xml:"opsChangeId,omitempty"`
}

func (s ListEventRecordsResponseBodyResultResultShowContent) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponseBodyResultResultShowContent) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetActionSuggest() *ListEventRecordsResponseBodyResultResultShowContentActionSuggest {
	return s.ActionSuggest
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetDesc() *string {
	return s.Desc
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetEventStatus() *string {
	return s.EventStatus
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetEventTime() *string {
	return s.EventTime
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetExecuteFinishTime() *string {
	return s.ExecuteFinishTime
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetExecuteStartTime() *string {
	return s.ExecuteStartTime
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) GetOpsChangeId() *string {
	return s.OpsChangeId
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetActionSuggest(v *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) *ListEventRecordsResponseBodyResultResultShowContent {
	s.ActionSuggest = v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetDesc(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.Desc = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetEventStatus(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.EventStatus = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetEventTime(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.EventTime = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetExecuteFinishTime(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.ExecuteFinishTime = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetExecuteStartTime(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetInstanceId(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.InstanceId = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) SetOpsChangeId(v string) *ListEventRecordsResponseBodyResultResultShowContent {
	s.OpsChangeId = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContent) Validate() error {
	if s.ActionSuggest != nil {
		if err := s.ActionSuggest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventRecordsResponseBodyResultResultShowContentActionSuggest struct {
	SuggestActions []*string `json:"suggestActions,omitempty" xml:"suggestActions,omitempty" type:"Repeated"`
	// example:
	//
	// “”
	SuggestText *string `json:"suggestText,omitempty" xml:"suggestText,omitempty"`
	// example:
	//
	// promptText
	SuggestType *string `json:"suggestType,omitempty" xml:"suggestType,omitempty"`
}

func (s ListEventRecordsResponseBodyResultResultShowContentActionSuggest) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsResponseBodyResultResultShowContentActionSuggest) GoString() string {
	return s.String()
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) GetSuggestActions() []*string {
	return s.SuggestActions
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) GetSuggestText() *string {
	return s.SuggestText
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) GetSuggestType() *string {
	return s.SuggestType
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) SetSuggestActions(v []*string) *ListEventRecordsResponseBodyResultResultShowContentActionSuggest {
	s.SuggestActions = v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) SetSuggestText(v string) *ListEventRecordsResponseBodyResultResultShowContentActionSuggest {
	s.SuggestText = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) SetSuggestType(v string) *ListEventRecordsResponseBodyResultResultShowContentActionSuggest {
	s.SuggestType = &v
	return s
}

func (s *ListEventRecordsResponseBodyResultResultShowContentActionSuggest) Validate() error {
	return dara.Validate(s)
}
