// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityWatchTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListQualityWatchTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityWatchTasksResponseBody
	GetMessage() *string
	SetPageResult(v *ListQualityWatchTasksResponseBodyPageResult) *ListQualityWatchTasksResponseBody
	GetPageResult() *ListQualityWatchTasksResponseBodyPageResult
	SetRequestId(v string) *ListQualityWatchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityWatchTasksResponseBody
	GetSuccess() *bool
}

type ListQualityWatchTasksResponseBody struct {
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
	// The error details from the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListQualityWatchTasksResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityWatchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityWatchTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityWatchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityWatchTasksResponseBody) GetPageResult() *ListQualityWatchTasksResponseBodyPageResult {
	return s.PageResult
}

func (s *ListQualityWatchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityWatchTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityWatchTasksResponseBody) SetCode(v string) *ListQualityWatchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityWatchTasksResponseBody) SetHttpStatusCode(v int32) *ListQualityWatchTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityWatchTasksResponseBody) SetMessage(v string) *ListQualityWatchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityWatchTasksResponseBody) SetPageResult(v *ListQualityWatchTasksResponseBodyPageResult) *ListQualityWatchTasksResponseBody {
	s.PageResult = v
	return s
}

func (s *ListQualityWatchTasksResponseBody) SetRequestId(v string) *ListQualityWatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityWatchTasksResponseBody) SetSuccess(v bool) *ListQualityWatchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityWatchTasksResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchTasksResponseBodyPageResult struct {
	// The list of quality watchtasks.
	QualityWatchTaskList []*ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList `json:"QualityWatchTaskList,omitempty" xml:"QualityWatchTaskList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityWatchTasksResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResult) GetQualityWatchTaskList() []*ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	return s.QualityWatchTaskList
}

func (s *ListQualityWatchTasksResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityWatchTasksResponseBodyPageResult) SetQualityWatchTaskList(v []*ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) *ListQualityWatchTasksResponseBodyPageResult {
	s.QualityWatchTaskList = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResult) SetTotalCount(v int64) *ListQualityWatchTasksResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResult) Validate() error {
	if s.QualityWatchTaskList != nil {
		for _, item := range s.QualityWatchTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList struct {
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
	// The end time, in the yyyy-MM-dd HH:mm:ss format.
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
	RuleCountInfo *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo `json:"RuleCountInfo,omitempty" xml:"RuleCountInfo,omitempty" type:"Struct"`
	// The quality rule IDs.
	RuleIdList []*int64 `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
	// The start time, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task statuses. Valid values:
	//
	// - NOT_RUN: not executed.
	//
	// - WAITING: waiting.
	//
	// - RUNNING: running.
	//
	// - SUCCESS: succeeded.
	//
	// - FAILED: failed.
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

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetBizDateFormat() *string {
	return s.BizDateFormat
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetId() *int64 {
	return s.Id
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetModifier() *string {
	return s.Modifier
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetQualityOwnerName() *string {
	return s.QualityOwnerName
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetRuleCountInfo() *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo {
	return s.RuleCountInfo
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetRuleIdList() []*int64 {
	return s.RuleIdList
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) GetWatchId() *int64 {
	return s.WatchId
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetBizDate(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.BizDate = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetBizDateFormat(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.BizDateFormat = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetCreateTime(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetCreator(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.Creator = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetEndTime(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.EndTime = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetId(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.Id = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetModifier(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.Modifier = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetModifyTime(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetQualityOwner(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.QualityOwner = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetQualityOwnerName(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.QualityOwnerName = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetRuleCountInfo(v *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.RuleCountInfo = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetRuleIdList(v []*int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.RuleIdList = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetStartTime(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.StartTime = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetStatus(v string) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.Status = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) SetWatchId(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList {
	s.WatchId = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskList) Validate() error {
	if s.RuleCountInfo != nil {
		if err := s.RuleCountInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo struct {
	// The strong rule count.
	StrongRuleCount *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount `json:"StrongRuleCount,omitempty" xml:"StrongRuleCount,omitempty" type:"Struct"`
	// The validation rule count.
	ValidateRuleCount *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount `json:"ValidateRuleCount,omitempty" xml:"ValidateRuleCount,omitempty" type:"Struct"`
	// The weak rule count.
	WeakRuleCount *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount `json:"WeakRuleCount,omitempty" xml:"WeakRuleCount,omitempty" type:"Struct"`
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) GetStrongRuleCount() *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount {
	return s.StrongRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) GetValidateRuleCount() *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount {
	return s.ValidateRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) GetWeakRuleCount() *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount {
	return s.WeakRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) SetStrongRuleCount(v *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo {
	s.StrongRuleCount = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) SetValidateRuleCount(v *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo {
	s.ValidateRuleCount = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) SetWeakRuleCount(v *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo {
	s.WeakRuleCount = v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfo) Validate() error {
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

type ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount struct {
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

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) SetErrorRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) SetFinishedRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) SetSuccessRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) SetTotalRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoStrongRuleCount) Validate() error {
	return dara.Validate(s)
}

type ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount struct {
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

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) SetErrorRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) SetFinishedRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) SetSuccessRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) SetTotalRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoValidateRuleCount) Validate() error {
	return dara.Validate(s)
}

type ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount struct {
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

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) GetErrorRuleCount() *int64 {
	return s.ErrorRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) GetFinishedRuleCount() *int64 {
	return s.FinishedRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) GetSuccessRuleCount() *int64 {
	return s.SuccessRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) SetErrorRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount {
	s.ErrorRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) SetFinishedRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount {
	s.FinishedRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) SetSuccessRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount {
	s.SuccessRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) SetTotalRuleCount(v int64) *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount {
	s.TotalRuleCount = &v
	return s
}

func (s *ListQualityWatchTasksResponseBodyPageResultQualityWatchTaskListRuleCountInfoWeakRuleCount) Validate() error {
	return dara.Validate(s)
}
