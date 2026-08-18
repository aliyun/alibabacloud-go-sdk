// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCasesResponseBody
	GetCode() *string
	SetData(v *ListCasesResponseBodyData) *ListCasesResponseBody
	GetData() *ListCasesResponseBodyData
	SetHttpStatusCode(v int32) *ListCasesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCasesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCasesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCasesResponseBody
	GetSuccess() *bool
}

type ListCasesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paged data.
	Data *ListCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pass-through parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C377C5FF-4F94-1B23-89D0-50C560623EE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCasesResponseBody) GetData() *ListCasesResponseBodyData {
	return s.Data
}

func (s *ListCasesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCasesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCasesResponseBody) SetCode(v string) *ListCasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCasesResponseBody) SetData(v *ListCasesResponseBodyData) *ListCasesResponseBody {
	s.Data = v
	return s
}

func (s *ListCasesResponseBody) SetHttpStatusCode(v int32) *ListCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCasesResponseBody) SetMessage(v string) *ListCasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCasesResponseBody) SetParams(v []*string) *ListCasesResponseBody {
	s.Params = v
	return s
}

func (s *ListCasesResponseBody) SetRequestId(v string) *ListCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCasesResponseBody) SetSuccess(v bool) *ListCasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCasesResponseBodyData struct {
	// The list of cases.
	List []*ListCasesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyData) GetList() []*ListCasesResponseBodyDataList {
	return s.List
}

func (s *ListCasesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCasesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCasesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCasesResponseBodyData) SetList(v []*ListCasesResponseBodyDataList) *ListCasesResponseBodyData {
	s.List = v
	return s
}

func (s *ListCasesResponseBodyData) SetPageNumber(v int32) *ListCasesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCasesResponseBodyData) SetPageSize(v int32) *ListCasesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCasesResponseBodyData) SetTotalCount(v int32) *ListCasesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCasesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCasesResponseBodyDataList struct {
	// The number of dial attempts.
	//
	// example:
	//
	// 1
	AttemptedCount *int32 `json:"AttemptedCount,omitempty" xml:"AttemptedCount,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 05923395478
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// The outbound campaign ID.
	//
	// example:
	//
	// 7607dae1-91ad-47ea-ad76-3d81ac34f729
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The name of the outbound campaign.
	//
	// example:
	//
	// e2d7a184-7d6c-45d4-ac24-34ab48f54600
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// The case ID.
	//
	// example:
	//
	// 00ed0dd9-c5a4-40e4-a8cd-822f0af859b9
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// The time when the case was created.
	//
	// example:
	//
	// 2025-07-27T11:25:15+08:00
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The custom variables in JSON string format.
	//
	// example:
	//
	// {}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The dialing time.
	//
	// example:
	//
	// 1786960840667
	DialingTime *int64 `json:"DialingTime,omitempty" xml:"DialingTime,omitempty"`
	// The disposition code.
	//
	// example:
	//
	// Answered
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// The disposition reason.
	//
	// example:
	//
	// Normal
	DispositionReason *string `json:"DispositionReason,omitempty" xml:"DispositionReason,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// e2d7a184-7d6c-45d4-ac24-34ab48f54669
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of labels.
	Labels []*ListCasesResponseBodyDataListLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The time of the last dial attempt.
	//
	// example:
	//
	// 1786960840667
	LastAttemptedTime *int64 `json:"LastAttemptedTime,omitempty" xml:"LastAttemptedTime,omitempty"`
	// The called number.
	//
	// example:
	//
	// 18512345678
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The priority of the case.
	//
	// example:
	//
	// Daily
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The reference ID.
	//
	// example:
	//
	// 5055-16-199313
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The ringing duration.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// The ringing time.
	//
	// example:
	//
	// 1786960840667
	RingingTime *int64 `json:"RingingTime,omitempty" xml:"RingingTime,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 8a988bd4-6c6e-45c6-b3a5-3def5ca3bc6f
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script.
	//
	// example:
	//
	// Scenario Name
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The session ID of the last call.
	//
	// example:
	//
	// SESSION_ID_312986372_7295954260941888
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The case state.
	//
	// example:
	//
	// Executing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The talk time.
	//
	// example:
	//
	// 10
	TalkTime *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
}

func (s ListCasesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyDataList) GetAttemptedCount() *int32 {
	return s.AttemptedCount
}

func (s *ListCasesResponseBodyDataList) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *ListCasesResponseBodyDataList) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCasesResponseBodyDataList) GetCampaignName() *string {
	return s.CampaignName
}

func (s *ListCasesResponseBodyDataList) GetCaseId() *string {
	return s.CaseId
}

func (s *ListCasesResponseBodyDataList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCasesResponseBodyDataList) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *ListCasesResponseBodyDataList) GetDialingTime() *int64 {
	return s.DialingTime
}

func (s *ListCasesResponseBodyDataList) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *ListCasesResponseBodyDataList) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *ListCasesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCasesResponseBodyDataList) GetLabels() []*ListCasesResponseBodyDataListLabels {
	return s.Labels
}

func (s *ListCasesResponseBodyDataList) GetLastAttemptedTime() *int64 {
	return s.LastAttemptedTime
}

func (s *ListCasesResponseBodyDataList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListCasesResponseBodyDataList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListCasesResponseBodyDataList) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListCasesResponseBodyDataList) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *ListCasesResponseBodyDataList) GetRingingTime() *int64 {
	return s.RingingTime
}

func (s *ListCasesResponseBodyDataList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCasesResponseBodyDataList) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListCasesResponseBodyDataList) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCasesResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListCasesResponseBodyDataList) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *ListCasesResponseBodyDataList) SetAttemptedCount(v int32) *ListCasesResponseBodyDataList {
	s.AttemptedCount = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCallerNumber(v string) *ListCasesResponseBodyDataList {
	s.CallerNumber = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCampaignId(v string) *ListCasesResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCampaignName(v string) *ListCasesResponseBodyDataList {
	s.CampaignName = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCaseId(v string) *ListCasesResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCreatedTime(v int64) *ListCasesResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCustomVariables(v string) *ListCasesResponseBodyDataList {
	s.CustomVariables = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetDialingTime(v int64) *ListCasesResponseBodyDataList {
	s.DialingTime = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetDispositionCode(v string) *ListCasesResponseBodyDataList {
	s.DispositionCode = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetDispositionReason(v string) *ListCasesResponseBodyDataList {
	s.DispositionReason = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetInstanceId(v string) *ListCasesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetLabels(v []*ListCasesResponseBodyDataListLabels) *ListCasesResponseBodyDataList {
	s.Labels = v
	return s
}

func (s *ListCasesResponseBodyDataList) SetLastAttemptedTime(v int64) *ListCasesResponseBodyDataList {
	s.LastAttemptedTime = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetPhoneNumber(v string) *ListCasesResponseBodyDataList {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetPriority(v int32) *ListCasesResponseBodyDataList {
	s.Priority = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetReferenceId(v string) *ListCasesResponseBodyDataList {
	s.ReferenceId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetRingingDuration(v int64) *ListCasesResponseBodyDataList {
	s.RingingDuration = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetRingingTime(v int64) *ListCasesResponseBodyDataList {
	s.RingingTime = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetScriptId(v string) *ListCasesResponseBodyDataList {
	s.ScriptId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetScriptName(v string) *ListCasesResponseBodyDataList {
	s.ScriptName = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetSessionId(v string) *ListCasesResponseBodyDataList {
	s.SessionId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetState(v string) *ListCasesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetTalkTime(v int64) *ListCasesResponseBodyDataList {
	s.TalkTime = &v
	return s
}

func (s *ListCasesResponseBodyDataList) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCasesResponseBodyDataListLabels struct {
	// The candidate values of the label.
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	// Indicates whether the item is collected.
	//
	// example:
	//
	// true
	Collected *bool `json:"Collected,omitempty" xml:"Collected,omitempty"`
	// The description of the label.
	//
	// example:
	//
	// Description content
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The matched value of the label.
	//
	// example:
	//
	// 1
	MatchedValue *string `json:"MatchedValue,omitempty" xml:"MatchedValue,omitempty"`
	// The name of the label.
	//
	// example:
	//
	// 软包装企业打电话_20251022_165548
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The system label.
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
}

func (s ListCasesResponseBodyDataListLabels) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBodyDataListLabels) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyDataListLabels) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *ListCasesResponseBodyDataListLabels) GetCollected() *bool {
	return s.Collected
}

func (s *ListCasesResponseBodyDataListLabels) GetDescription() *string {
	return s.Description
}

func (s *ListCasesResponseBodyDataListLabels) GetMatchedValue() *string {
	return s.MatchedValue
}

func (s *ListCasesResponseBodyDataListLabels) GetName() *string {
	return s.Name
}

func (s *ListCasesResponseBodyDataListLabels) GetSystem() *bool {
	return s.System
}

func (s *ListCasesResponseBodyDataListLabels) SetCandidateValues(v []*string) *ListCasesResponseBodyDataListLabels {
	s.CandidateValues = v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) SetCollected(v bool) *ListCasesResponseBodyDataListLabels {
	s.Collected = &v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) SetDescription(v string) *ListCasesResponseBodyDataListLabels {
	s.Description = &v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) SetMatchedValue(v string) *ListCasesResponseBodyDataListLabels {
	s.MatchedValue = &v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) SetName(v string) *ListCasesResponseBodyDataListLabels {
	s.Name = &v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) SetSystem(v bool) *ListCasesResponseBodyDataListLabels {
	s.System = &v
	return s
}

func (s *ListCasesResponseBodyDataListLabels) Validate() error {
	return dara.Validate(s)
}
