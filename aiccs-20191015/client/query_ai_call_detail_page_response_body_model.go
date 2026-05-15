// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallDetailPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAiCallDetailPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAiCallDetailPageResponseBody
	GetCode() *string
	SetData(v *QueryAiCallDetailPageResponseBodyData) *QueryAiCallDetailPageResponseBody
	GetData() *QueryAiCallDetailPageResponseBodyData
	SetMessage(v string) *QueryAiCallDetailPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAiCallDetailPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAiCallDetailPageResponseBody
	GetSuccess() *bool
}

type QueryAiCallDetailPageResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAiCallDetailPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A57441B2-8EB6-5B93-9F37-0A51B8E2C9F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiCallDetailPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAiCallDetailPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAiCallDetailPageResponseBody) GetData() *QueryAiCallDetailPageResponseBodyData {
	return s.Data
}

func (s *QueryAiCallDetailPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAiCallDetailPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiCallDetailPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiCallDetailPageResponseBody) SetAccessDeniedDetail(v string) *QueryAiCallDetailPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) SetCode(v string) *QueryAiCallDetailPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) SetData(v *QueryAiCallDetailPageResponseBodyData) *QueryAiCallDetailPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) SetMessage(v string) *QueryAiCallDetailPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) SetRequestId(v string) *QueryAiCallDetailPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) SetSuccess(v bool) *QueryAiCallDetailPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiCallDetailPageResponseBodyData struct {
	List []*QueryAiCallDetailPageResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAiCallDetailPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageResponseBodyData) GetList() []*QueryAiCallDetailPageResponseBodyDataList {
	return s.List
}

func (s *QueryAiCallDetailPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryAiCallDetailPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryAiCallDetailPageResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryAiCallDetailPageResponseBodyData) SetList(v []*QueryAiCallDetailPageResponseBodyDataList) *QueryAiCallDetailPageResponseBodyData {
	s.List = v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyData) SetPageNo(v int64) *QueryAiCallDetailPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyData) SetPageSize(v int64) *QueryAiCallDetailPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyData) SetTotal(v int64) *QueryAiCallDetailPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyData) Validate() error {
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

type QueryAiCallDetailPageResponseBodyDataList struct {
	// example:
	//
	// 12121211111*****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 49
	BranchId *int64 `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// example:
	//
	// 示例值
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// example:
	//
	// 9
	BranchVersionId *int64 `json:"BranchVersionId,omitempty" xml:"BranchVersionId,omitempty"`
	// example:
	//
	// 示例值
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// example:
	//
	// 0537101****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 1748948749000
	CallingTime *int64 `json:"CallingTime,omitempty" xml:"CallingTime,omitempty"`
	// example:
	//
	// 100
	ConversationDuration *int64 `json:"ConversationDuration,omitempty" xml:"ConversationDuration,omitempty"`
	// example:
	//
	// 示例值示例值
	ConversationRecord *string `json:"ConversationRecord,omitempty" xml:"ConversationRecord,omitempty"`
	// example:
	//
	// 10
	ConversationTurnCount *int64 `json:"ConversationTurnCount,omitempty" xml:"ConversationTurnCount,omitempty"`
	// example:
	//
	// 12121211111*****
	DetailId *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	// example:
	//
	// 21
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// example:
	//
	// 1748948749000
	ImportedTime *int64 `json:"ImportedTime,omitempty" xml:"ImportedTime,omitempty"`
	// example:
	//
	// A
	MajorIntent *string `json:"MajorIntent,omitempty" xml:"MajorIntent,omitempty"`
	// example:
	//
	// 示例值
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// outId
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// https://*******
	RecordingFilePath *string `json:"RecordingFilePath,omitempty" xml:"RecordingFilePath,omitempty"`
	// example:
	//
	// 51
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 12121211111*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// example:
	//
	// 55
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s QueryAiCallDetailPageResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetBatchId() *string {
	return s.BatchId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetBranchId() *int64 {
	return s.BranchId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetBranchName() *string {
	return s.BranchName
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetBranchVersionId() *int64 {
	return s.BranchVersionId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetCallingTime() *int64 {
	return s.CallingTime
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetConversationDuration() *int64 {
	return s.ConversationDuration
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetConversationRecord() *string {
	return s.ConversationRecord
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetConversationTurnCount() *int64 {
	return s.ConversationTurnCount
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetDetailId() *string {
	return s.DetailId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetFailedReason() *string {
	return s.FailedReason
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetImportedTime() *int64 {
	return s.ImportedTime
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetMajorIntent() *string {
	return s.MajorIntent
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetOptions() *string {
	return s.Options
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetOutId() *string {
	return s.OutId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetRecordingFilePath() *string {
	return s.RecordingFilePath
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetStatus() *int64 {
	return s.Status
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetVersionName() *string {
	return s.VersionName
}

func (s *QueryAiCallDetailPageResponseBodyDataList) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetBatchId(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.BatchId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetBranchId(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.BranchId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetBranchName(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.BranchName = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetBranchVersionId(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.BranchVersionId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetCallResult(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.CallResult = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetCalledNumber(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetCallingTime(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.CallingTime = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetConversationDuration(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.ConversationDuration = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetConversationRecord(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.ConversationRecord = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetConversationTurnCount(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.ConversationTurnCount = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetDetailId(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.DetailId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetEncryptionType(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.EncryptionType = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetFailedReason(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.FailedReason = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetImportedTime(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.ImportedTime = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetMajorIntent(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.MajorIntent = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetOptions(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.Options = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetOutId(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.OutId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetRecordingFilePath(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.RecordingFilePath = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetStatus(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetTaskId(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetVersionName(v string) *QueryAiCallDetailPageResponseBodyDataList {
	s.VersionName = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) SetVersionNo(v int64) *QueryAiCallDetailPageResponseBodyDataList {
	s.VersionNo = &v
	return s
}

func (s *QueryAiCallDetailPageResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
