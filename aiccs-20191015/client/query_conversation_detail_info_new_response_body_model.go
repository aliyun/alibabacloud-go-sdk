// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryConversationDetailInfoNewResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryConversationDetailInfoNewResponseBody
	GetCode() *string
	SetData(v *QueryConversationDetailInfoNewResponseBodyData) *QueryConversationDetailInfoNewResponseBody
	GetData() *QueryConversationDetailInfoNewResponseBodyData
	SetMessage(v string) *QueryConversationDetailInfoNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryConversationDetailInfoNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryConversationDetailInfoNewResponseBody
	GetSuccess() *bool
}

type QueryConversationDetailInfoNewResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryConversationDetailInfoNewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConversationDetailInfoNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryConversationDetailInfoNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryConversationDetailInfoNewResponseBody) GetData() *QueryConversationDetailInfoNewResponseBodyData {
	return s.Data
}

func (s *QueryConversationDetailInfoNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryConversationDetailInfoNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConversationDetailInfoNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryConversationDetailInfoNewResponseBody) SetAccessDeniedDetail(v string) *QueryConversationDetailInfoNewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) SetCode(v string) *QueryConversationDetailInfoNewResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) SetData(v *QueryConversationDetailInfoNewResponseBodyData) *QueryConversationDetailInfoNewResponseBody {
	s.Data = v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) SetMessage(v string) *QueryConversationDetailInfoNewResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) SetRequestId(v string) *QueryConversationDetailInfoNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) SetSuccess(v bool) *QueryConversationDetailInfoNewResponseBody {
	s.Success = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConversationDetailInfoNewResponseBodyData struct {
	// example:
	//
	// 1234******
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 123*******^213******
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// ANSWERED
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// example:
	//
	// 130********
	CalledPhone *string `json:"CalledPhone,omitempty" xml:"CalledPhone,omitempty"`
	// example:
	//
	// 0571*******
	CallerPhone *string `json:"CallerPhone,omitempty" xml:"CallerPhone,omitempty"`
	// example:
	//
	// 示例值
	ConversationRecord *string `json:"ConversationRecord,omitempty" xml:"ConversationRecord,omitempty"`
	// example:
	//
	// 23
	ConversationTurnCount *int64 `json:"ConversationTurnCount,omitempty" xml:"ConversationTurnCount,omitempty"`
	// example:
	//
	// 1234*******
	DetailId *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	// example:
	//
	// 30
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 28
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// 示例值示例值
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// example:
	//
	// 0
	HangupDirection *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	// example:
	//
	// 43
	ImportedTime *int64 `json:"ImportedTime,omitempty" xml:"ImportedTime,omitempty"`
	// example:
	//
	// A
	MajorIntent *string `json:"MajorIntent,omitempty" xml:"MajorIntent,omitempty"`
	// example:
	//
	// 示例值示例值
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 123***
	OutId      *string                                                     `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OutputTags []*QueryConversationDetailInfoNewResponseBodyDataOutputTags `json:"OutputTags,omitempty" xml:"OutputTags,omitempty" type:"Repeated"`
	// example:
	//
	// 12349732441
	PickUpTime *int64 `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	// example:
	//
	// recording.oss.file
	RecordingFileDownloadUrl *string `json:"RecordingFileDownloadUrl,omitempty" xml:"RecordingFileDownloadUrl,omitempty"`
	// example:
	//
	// 7
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// 1286987391
	StartCallTime *int64 `json:"StartCallTime,omitempty" xml:"StartCallTime,omitempty"`
	// example:
	//
	// 72
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 200005
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 示例值示例值
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
	// example:
	//
	// 138************
	TaskId    *string                                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Variables []*QueryConversationDetailInfoNewResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s QueryConversationDetailInfoNewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetBatchId() *string {
	return s.BatchId
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetCalledPhone() *string {
	return s.CalledPhone
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetCallerPhone() *string {
	return s.CallerPhone
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetConversationRecord() *string {
	return s.ConversationRecord
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetConversationTurnCount() *int64 {
	return s.ConversationTurnCount
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetDetailId() *string {
	return s.DetailId
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetFailedReason() *string {
	return s.FailedReason
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetHangupDirection() *string {
	return s.HangupDirection
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetImportedTime() *int64 {
	return s.ImportedTime
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetMajorIntent() *string {
	return s.MajorIntent
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetOptions() *string {
	return s.Options
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetOutId() *string {
	return s.OutId
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetOutputTags() []*QueryConversationDetailInfoNewResponseBodyDataOutputTags {
	return s.OutputTags
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetPickUpTime() *int64 {
	return s.PickUpTime
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetRecordingFileDownloadUrl() *string {
	return s.RecordingFileDownloadUrl
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetStartCallTime() *int64 {
	return s.StartCallTime
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetStatus() *int64 {
	return s.Status
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryConversationDetailInfoNewResponseBodyData) GetVariables() []*QueryConversationDetailInfoNewResponseBodyDataVariables {
	return s.Variables
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetBatchId(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetCallId(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.CallId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetCallResult(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.CallResult = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetCalledPhone(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.CalledPhone = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetCallerPhone(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.CallerPhone = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetConversationRecord(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.ConversationRecord = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetConversationTurnCount(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.ConversationTurnCount = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetDetailId(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.DetailId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetDuration(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.Duration = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetEncryptionType(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.EncryptionType = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetFailedReason(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.FailedReason = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetHangupDirection(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.HangupDirection = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetImportedTime(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.ImportedTime = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetMajorIntent(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.MajorIntent = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetOptions(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.Options = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetOutId(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.OutId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetOutputTags(v []*QueryConversationDetailInfoNewResponseBodyDataOutputTags) *QueryConversationDetailInfoNewResponseBodyData {
	s.OutputTags = v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetPickUpTime(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.PickUpTime = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetRecordingFileDownloadUrl(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.RecordingFileDownloadUrl = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetReleaseTime(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetStartCallTime(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.StartCallTime = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetStatus(v int64) *QueryConversationDetailInfoNewResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetStatusCode(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetStatusMsg(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.StatusMsg = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetTaskId(v string) *QueryConversationDetailInfoNewResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) SetVariables(v []*QueryConversationDetailInfoNewResponseBodyDataVariables) *QueryConversationDetailInfoNewResponseBodyData {
	s.Variables = v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyData) Validate() error {
	if s.OutputTags != nil {
		for _, item := range s.OutputTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryConversationDetailInfoNewResponseBodyDataOutputTags struct {
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutputTagDescription *string `json:"OutputTagDescription,omitempty" xml:"OutputTagDescription,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutputTagName *string `json:"OutputTagName,omitempty" xml:"OutputTagName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutputTagValue *string `json:"OutputTagValue,omitempty" xml:"OutputTagValue,omitempty"`
}

func (s QueryConversationDetailInfoNewResponseBodyDataOutputTags) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewResponseBodyDataOutputTags) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) GetId() *string {
	return s.Id
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) GetOutputTagDescription() *string {
	return s.OutputTagDescription
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) GetOutputTagName() *string {
	return s.OutputTagName
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) GetOutputTagValue() *string {
	return s.OutputTagValue
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) SetId(v string) *QueryConversationDetailInfoNewResponseBodyDataOutputTags {
	s.Id = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) SetOutputTagDescription(v string) *QueryConversationDetailInfoNewResponseBodyDataOutputTags {
	s.OutputTagDescription = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) SetOutputTagName(v string) *QueryConversationDetailInfoNewResponseBodyDataOutputTags {
	s.OutputTagName = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) SetOutputTagValue(v string) *QueryConversationDetailInfoNewResponseBodyDataOutputTags {
	s.OutputTagValue = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataOutputTags) Validate() error {
	return dara.Validate(s)
}

type QueryConversationDetailInfoNewResponseBodyDataVariables struct {
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// user name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// source
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// mike
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryConversationDetailInfoNewResponseBodyDataVariables) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetId() *string {
	return s.Id
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetKey() *string {
	return s.Key
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetName() *string {
	return s.Name
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetRequired() *bool {
	return s.Required
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetSource() *string {
	return s.Source
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) GetValue() *string {
	return s.Value
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetId(v string) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Id = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetKey(v string) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Key = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetName(v string) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Name = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetRequired(v bool) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Required = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetSource(v string) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Source = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) SetValue(v string) *QueryConversationDetailInfoNewResponseBodyDataVariables {
	s.Value = &v
	return s
}

func (s *QueryConversationDetailInfoNewResponseBodyDataVariables) Validate() error {
	return dara.Validate(s)
}
