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
	// The reason why the access request was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryAiCallDetailPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// A description of the status code.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A57441B2-8EB6-5B93-9F37-0A51B8E2C9F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values are:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
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
	// A list of task details.
	List []*QueryAiCallDetailPageResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 60
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
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
	// The batch ID.
	//
	// example:
	//
	// 1183**************
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 49
	BranchId *int64 `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// example:
	//
	// example
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// example:
	//
	// 9
	BranchVersionId *int64 `json:"BranchVersionId,omitempty" xml:"BranchVersionId,omitempty"`
	// The call result.
	//
	// example:
	//
	// 用户接通
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// The called number.
	//
	// example:
	//
	// 0537101****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The call time, formatted as a timestamp in milliseconds.
	//
	// example:
	//
	// 1748948749000
	CallingTime *int64 `json:"CallingTime,omitempty" xml:"CallingTime,omitempty"`
	// The conversation duration, in seconds.
	//
	// example:
	//
	// 100
	ConversationDuration *int64 `json:"ConversationDuration,omitempty" xml:"ConversationDuration,omitempty"`
	// The conversation record, formatted as a chronologically sorted JSON array. Each object has the following structure:
	//
	// ```json
	//
	// [
	//
	//     {
	//
	//         "content":"The content of the message.",
	//
	//         "role":"The role of the speaker.", // Valid values: user, assistant
	//
	//     }
	//
	// ]
	//
	// ```
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "content": "111您好，年龄222，性别男，我这边是**汽车的官方顾问，我们新出了一款车型为**；**已经上市了，售价**万元起，**分钟破*台，您看要不了解一下？",
	//
	//         "role": "assistant"
	//
	//     },
	//
	//     {
	//
	//         "content": "<客户打断>哎，你是谁？",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "<客户打断>你再说一遍。",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "哎，我没听清。",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "你在说什么？",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "您好，",
	//
	//         "role": "assistant",
	//
	//     },
	//
	//     {
	//
	//         "content": "我是**汽车总部销售服务顾问。",
	//
	//         "role": "assistant",
	//
	//     },
	//
	//     {
	//
	//         "content": "我们最近推出了一款新车**，想了解一下您是否对这款车型感兴趣？",
	//
	//         "role": "assistant",
	//
	//     },
	//
	//     {
	//
	//         "content": "<客户打断>哎，那我是谁？",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "你在说什么呢？",
	//
	//         "role": "user",
	//
	//     },
	//
	//     {
	//
	//         "content": "抱歉打扰了，111先生。",
	//
	//         "role": "assistant",
	//
	//     },
	//
	//     {
	//
	//         "content": "祝您生活愉快！再见！",
	//
	//         "role": "assistant",
	//
	//     }
	//
	// ]
	ConversationRecord *string `json:"ConversationRecord,omitempty" xml:"ConversationRecord,omitempty"`
	// The conversation turn count.
	//
	// example:
	//
	// 10
	ConversationTurnCount *int64 `json:"ConversationTurnCount,omitempty" xml:"ConversationTurnCount,omitempty"`
	// The task detail ID.
	//
	// example:
	//
	// 9662*************
	DetailId *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	// The encryption type. Valid values are: 0 (no encryption), 1 (MD5), 2 (SHA256), and 3 (SM3).
	//
	// example:
	//
	// 1
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The failure reason. Provided only if the call fails.
	//
	// example:
	//
	// 账户停机
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The import time, formatted as a timestamp in milliseconds.
	//
	// example:
	//
	// 1748948749000
	ImportedTime *int64 `json:"ImportedTime,omitempty" xml:"ImportedTime,omitempty"`
	// The major intent.
	//
	// example:
	//
	// A
	MajorIntent *string `json:"MajorIntent,omitempty" xml:"MajorIntent,omitempty"`
	// A JSON object of key-value pairs for runtime variables.
	//
	// example:
	//
	// {
	//
	//   "date": "666",
	//
	//   "phoneNumber": "777",
	//
	//   "distance": "555",
	//
	//   "mendian": "444",
	//
	//   "sex": "男",
	//
	//   "name": "111",
	//
	//   "age": "222"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// outId
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The recording file path. Provided only after the recording file is generated.
	//
	// example:
	//
	// https://*******
	RecordingFilePath *string `json:"RecordingFilePath,omitempty" xml:"RecordingFilePath,omitempty"`
	// The task detail status.
	//
	// - 0: Initializing
	//
	// - 1: Waiting to call
	//
	// - 2: Waiting to retry
	//
	// - 3: Calling
	//
	// - 4: Call ended
	//
	// - 5: Call failed
	//
	// Only statuses 4 and 5 are terminal states.
	//
	// example:
	//
	// 4
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1187**************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// example
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
