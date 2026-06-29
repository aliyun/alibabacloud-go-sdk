// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityWatchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityWatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityWatchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityWatchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityWatchTaskResponseBody
	GetSuccess() *bool
	SetWatchTaskInfo(v *GetQualityWatchTaskResponseBodyWatchTaskInfo) *GetQualityWatchTaskResponseBody
	GetWatchTaskInfo() *GetQualityWatchTaskResponseBodyWatchTaskInfo
}

type GetQualityWatchTaskResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the monitoring node task object.
	WatchTaskInfo *GetQualityWatchTaskResponseBodyWatchTaskInfo `json:"WatchTaskInfo,omitempty" xml:"WatchTaskInfo,omitempty" type:"Struct"`
}

func (s GetQualityWatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityWatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityWatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityWatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityWatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityWatchTaskResponseBody) GetWatchTaskInfo() *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	return s.WatchTaskInfo
}

func (s *GetQualityWatchTaskResponseBody) SetCode(v string) *GetQualityWatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityWatchTaskResponseBody) SetHttpStatusCode(v int32) *GetQualityWatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityWatchTaskResponseBody) SetMessage(v string) *GetQualityWatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityWatchTaskResponseBody) SetRequestId(v string) *GetQualityWatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityWatchTaskResponseBody) SetSuccess(v bool) *GetQualityWatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityWatchTaskResponseBody) SetWatchTaskInfo(v *GetQualityWatchTaskResponseBodyWatchTaskInfo) *GetQualityWatchTaskResponseBody {
	s.WatchTaskInfo = v
	return s
}

func (s *GetQualityWatchTaskResponseBody) Validate() error {
	if s.WatchTaskInfo != nil {
		if err := s.WatchTaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityWatchTaskResponseBodyWatchTaskInfo struct {
	// The business date.
	//
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The business date format.
	//
	// example:
	//
	// yyyy-MM-dd
	BizDateFormat *string `json:"BizDateFormat,omitempty" xml:"BizDateFormat,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The end time. Time format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The quality watchtask ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user ID of the last modifier.
	//
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user ID of the quality owner.
	//
	// example:
	//
	// 30012011
	QualityOwner *string `json:"QualityOwner,omitempty" xml:"QualityOwner,omitempty"`
	// The name of the quality owner.
	//
	// example:
	//
	// test
	QualityOwnerName *string `json:"QualityOwnerName,omitempty" xml:"QualityOwnerName,omitempty"`
	// The quality rule count information.
	RuleCountInfo *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo `json:"RuleCountInfo,omitempty" xml:"RuleCountInfo,omitempty" type:"Struct"`
	// The list of quality rule IDs.
	RuleIdList []*int64 `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
	// The start time. Time format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - NOT_RUN: not executed.
	//
	// - WAITING: waiting.
	//
	// - RUNNING: executing.
	//
	// - SUCCESS: executed successfully.
	//
	// - FAILED: execution failed.
	//
	// - CANCEL: canceled.
	//
	// - TIMEOUT: timed out.
	//
	// - OFFLINE: offline.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The monitored object ID.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetBizDate() *string {
	return s.BizDate
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetBizDateFormat() *string {
	return s.BizDateFormat
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetQualityOwnerName() *string {
	return s.QualityOwnerName
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetRuleCountInfo() *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo {
	return s.RuleCountInfo
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetRuleIdList() []*int64 {
	return s.RuleIdList
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetBizDate(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.BizDate = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetBizDateFormat(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.BizDateFormat = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetCreateTime(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetCreator(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetEndTime(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.EndTime = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetId(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.Id = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetModifier(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetModifyTime(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetQualityOwner(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.QualityOwner = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetQualityOwnerName(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.QualityOwnerName = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetRuleCountInfo(v *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.RuleCountInfo = v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetRuleIdList(v []*int64) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.RuleIdList = v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetStartTime(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.StartTime = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetStatus(v string) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.Status = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) SetWatchId(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfo {
	s.WatchId = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfo) Validate() error {
	if s.RuleCountInfo != nil {
		if err := s.RuleCountInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo struct {
	// The strong rule count.
	StrongRuleCount *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount `json:"StrongRuleCount,omitempty" xml:"StrongRuleCount,omitempty" type:"Struct"`
	// The validation rule count.
	ValidateRuleCount *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount `json:"ValidateRuleCount,omitempty" xml:"ValidateRuleCount,omitempty" type:"Struct"`
	// The weak rule count.
	WeakRuleCount *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount `json:"WeakRuleCount,omitempty" xml:"WeakRuleCount,omitempty" type:"Struct"`
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) GetStrongRuleCount() *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount {
	return s.StrongRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) GetValidateRuleCount() *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount {
	return s.ValidateRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) GetWeakRuleCount() *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount {
	return s.WeakRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) SetStrongRuleCount(v *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo {
	s.StrongRuleCount = v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) SetValidateRuleCount(v *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo {
	s.ValidateRuleCount = v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) SetWeakRuleCount(v *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo {
	s.WeakRuleCount = v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfo) Validate() error {
	if s.StrongRuleCount != nil {
		if err := s.StrongRuleCount.Validate(); err != nil {
			return err
		}
	}
	if s.ValidateRuleCount != nil {
		if err := s.ValidateRuleCount.Validate(); err != nil {
			return err
		}
	}
	if s.WeakRuleCount != nil {
		if err := s.WeakRuleCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount struct {
	// The number of rules that failed to execute.
	//
	// example:
	//
	// 0
	ErrorRuleCount *int64 `json:"ErrorRuleCount,omitempty" xml:"ErrorRuleCount,omitempty"`
	// The number of rules that have been executed.
	//
	// example:
	//
	// 1
	FinishedRuleCount *int64 `json:"FinishedRuleCount,omitempty" xml:"FinishedRuleCount,omitempty"`
	// The number of rules that were executed successfully.
	//
	// example:
	//
	// 1
	SuccessRuleCount *int64 `json:"SuccessRuleCount,omitempty" xml:"SuccessRuleCount,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 1
	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" xml:"TotalRuleCount,omitempty"`
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) SetErrorRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) SetFinishedRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) SetSuccessRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) SetTotalRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoStrongRuleCount) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount struct {
	// The number of rules that failed to execute.
	//
	// example:
	//
	// 0
	ErrorRuleCount *int64 `json:"ErrorRuleCount,omitempty" xml:"ErrorRuleCount,omitempty"`
	// The number of rules that have been executed.
	//
	// example:
	//
	// 1
	FinishedRuleCount *int64 `json:"FinishedRuleCount,omitempty" xml:"FinishedRuleCount,omitempty"`
	// The number of rules that were executed successfully.
	//
	// example:
	//
	// 1
	SuccessRuleCount *int64 `json:"SuccessRuleCount,omitempty" xml:"SuccessRuleCount,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 1
	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" xml:"TotalRuleCount,omitempty"`
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) SetErrorRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) SetFinishedRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) SetSuccessRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) SetTotalRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoValidateRuleCount) Validate() error {
	return dara.Validate(s)
}

type GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount struct {
	// The number of rules that failed to execute.
	//
	// example:
	//
	// 0
	ErrorRuleCount *int64 `json:"ErrorRuleCount,omitempty" xml:"ErrorRuleCount,omitempty"`
	// The number of rules that have been executed.
	//
	// example:
	//
	// 1
	FinishedRuleCount *int64 `json:"FinishedRuleCount,omitempty" xml:"FinishedRuleCount,omitempty"`
	// The number of rules that were executed successfully.
	//
	// example:
	//
	// 1
	SuccessRuleCount *int64 `json:"SuccessRuleCount,omitempty" xml:"SuccessRuleCount,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 1
	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" xml:"TotalRuleCount,omitempty"`
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) SetErrorRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) SetFinishedRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) SetSuccessRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) SetTotalRuleCount(v int64) *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *GetQualityWatchTaskResponseBodyWatchTaskInfoRuleCountInfoWeakRuleCount) Validate() error {
	return dara.Validate(s)
}
