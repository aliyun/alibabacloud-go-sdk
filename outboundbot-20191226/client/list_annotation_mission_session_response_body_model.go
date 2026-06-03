// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAnnotationMissionSessionResponseBody
	GetCode() *string
	SetData(v *ListAnnotationMissionSessionResponseBodyData) *ListAnnotationMissionSessionResponseBody
	GetData() *ListAnnotationMissionSessionResponseBodyData
	SetHttpStatusCode(v int32) *ListAnnotationMissionSessionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAnnotationMissionSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAnnotationMissionSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnnotationMissionSessionResponseBody
	GetSuccess() *bool
}

type ListAnnotationMissionSessionResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAnnotationMissionSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// bp.java.nopowerContact
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAnnotationMissionSessionResponseBody) GetData() *ListAnnotationMissionSessionResponseBodyData {
	return s.Data
}

func (s *ListAnnotationMissionSessionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAnnotationMissionSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnnotationMissionSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnnotationMissionSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnnotationMissionSessionResponseBody) SetCode(v string) *ListAnnotationMissionSessionResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) SetData(v *ListAnnotationMissionSessionResponseBodyData) *ListAnnotationMissionSessionResponseBody {
	s.Data = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) SetHttpStatusCode(v int32) *ListAnnotationMissionSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) SetMessage(v string) *ListAnnotationMissionSessionResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) SetRequestId(v string) *ListAnnotationMissionSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) SetSuccess(v bool) *ListAnnotationMissionSessionResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAnnotationMissionSessionResponseBodyData struct {
	// example:
	//
	// b3f2c931-5180-43ca-b4aa-2baee2d73c8b
	AnnotationMissionId          *string                                                                     `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionList []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList `json:"AnnotationMissionSessionList,omitempty" xml:"AnnotationMissionSessionList,omitempty" type:"Repeated"`
	// example:
	//
	// bp.java.nopowerContact
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyData) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyData) GetAnnotationMissionSessionList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	return s.AnnotationMissionSessionList
}

func (s *ListAnnotationMissionSessionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListAnnotationMissionSessionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnnotationMissionSessionResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAnnotationMissionSessionResponseBodyData) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyData {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyData) SetAnnotationMissionSessionList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) *ListAnnotationMissionSessionResponseBodyData {
	s.AnnotationMissionSessionList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyData) SetMessage(v string) *ListAnnotationMissionSessionResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyData) SetSuccess(v bool) *ListAnnotationMissionSessionResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyData) SetTotalCount(v int64) *ListAnnotationMissionSessionResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyData) Validate() error {
	if s.AnnotationMissionSessionList != nil {
		for _, item := range s.AnnotationMissionSessionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList struct {
	AnnotationMissionChatList []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList `json:"AnnotationMissionChatList,omitempty" xml:"AnnotationMissionChatList,omitempty" type:"Repeated"`
	// example:
	//
	// 64ba5ac9-a4e1-4333-b03c-9d4a588c9f6c
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// ID
	//
	// example:
	//
	// 64ba5ac9-a4e1-4333-b03c-9d4a588c9f6c
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// 1691117009000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	DebugConversation *bool `json:"DebugConversation,omitempty" xml:"DebugConversation,omitempty"`
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2f642da1-b00b-4dd6-ac7d-dcbeefd13ff3
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// 42da7cde-a5e8-45cc-b9d2-828711d95b30
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1677987898384
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 408d6c4d-23e2-41f6-bbdd-f919a8297aa4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 40669e52-c1c8-487f-9593-29749086bdc9
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetAnnotationMissionChatList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	return s.AnnotationMissionChatList
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetDebugConversation() *bool {
	return s.DebugConversation
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetJobId() *string {
	return s.JobId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) GetVersion() *int32 {
	return s.Version
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetAnnotationMissionChatList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.AnnotationMissionChatList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetAnnotationStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.AnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetDebugConversation(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.DebugConversation = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetJobGroupId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.JobGroupId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetJobId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.JobId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetScriptId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.ScriptId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.SessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) SetVersion(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList {
	s.Version = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionList) Validate() error {
	if s.AnnotationMissionChatList != nil {
		for _, item := range s.AnnotationMissionChatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList struct {
	AnnotationAsrResult                            *string                                                                                                                                            `json:"AnnotationAsrResult,omitempty" xml:"AnnotationAsrResult,omitempty"`
	AnnotationMissionChatCustomizationDataInfoList []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList `json:"AnnotationMissionChatCustomizationDataInfoList,omitempty" xml:"AnnotationMissionChatCustomizationDataInfoList,omitempty" type:"Repeated"`
	// chat id
	//
	// example:
	//
	// 40669e52-c1c8-487f-9593-29749086bdc9
	AnnotationMissionChatId                    *string                                                                                                                                        `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	AnnotationMissionChatIntentUserSayInfoList []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList `json:"AnnotationMissionChatIntentUserSayInfoList,omitempty" xml:"AnnotationMissionChatIntentUserSayInfoList,omitempty" type:"Repeated"`
	AnnotationMissionChatTagInfoList           []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList           `json:"AnnotationMissionChatTagInfoList,omitempty" xml:"AnnotationMissionChatTagInfoList,omitempty" type:"Repeated"`
	AnnotationMissionChatVocabularyInfoList    []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList    `json:"AnnotationMissionChatVocabularyInfoList,omitempty" xml:"AnnotationMissionChatVocabularyInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 03f56192-fa8a-40dc-9558-39b357f0618f
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// {\\"Answer\\": u\\"\\u53c2\\u8003\\u6587\\u6863\\uff1ahttps://help.aliyun.com/document_detail/215402.html\\", \\"QuestionId\\": 371159, \\"Uuid\\": \\"ac14000116781568127896224d0044\\"}
	Answer              *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	AsrAnnotationStatus *int32  `json:"AsrAnnotationStatus,omitempty" xml:"AsrAnnotationStatus,omitempty"`
	// example:
	//
	// 1682216045619
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentAnnotationStatus *int32  `json:"IntentAnnotationStatus,omitempty" xml:"IntentAnnotationStatus,omitempty"`
	// example:
	//
	// 1571649300000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 1682390676403
	OccurTime         *int64  `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	OriginalAsrResult *string `json:"OriginalAsrResult,omitempty" xml:"OriginalAsrResult,omitempty"`
	// example:
	//
	// 380578077
	SequenceId          *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	SubStatus           *int32  `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	TagAnnotationStatus *int32  `json:"TagAnnotationStatus,omitempty" xml:"TagAnnotationStatus,omitempty"`
	// example:
	//
	// 1
	TranslationError *int32 `json:"TranslationError,omitempty" xml:"TranslationError,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationAsrResult() *string {
	return s.AnnotationAsrResult
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatCustomizationDataInfoList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	return s.AnnotationMissionChatCustomizationDataInfoList
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatIntentUserSayInfoList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	return s.AnnotationMissionChatIntentUserSayInfoList
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatTagInfoList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	return s.AnnotationMissionChatTagInfoList
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatVocabularyInfoList() []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	return s.AnnotationMissionChatVocabularyInfoList
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAnswer() *string {
	return s.Answer
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetAsrAnnotationStatus() *int32 {
	return s.AsrAnnotationStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetIntentAnnotationStatus() *int32 {
	return s.IntentAnnotationStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetOccurTime() *int64 {
	return s.OccurTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetOriginalAsrResult() *string {
	return s.OriginalAsrResult
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetSequenceId() *string {
	return s.SequenceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetSubStatus() *int32 {
	return s.SubStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetTagAnnotationStatus() *int32 {
	return s.TagAnnotationStatus
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) GetTranslationError() *int32 {
	return s.TranslationError
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationAsrResult(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationAsrResult = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatCustomizationDataInfoList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatCustomizationDataInfoList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatIntentUserSayInfoList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatIntentUserSayInfoList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatTagInfoList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatTagInfoList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatVocabularyInfoList(v []*ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatVocabularyInfoList = v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAnswer(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.Answer = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetAsrAnnotationStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.AsrAnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetIntentAnnotationStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.IntentAnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetOccurTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.OccurTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetOriginalAsrResult(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.OriginalAsrResult = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetSequenceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.SequenceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetSubStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.SubStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetTagAnnotationStatus(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.TagAnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) SetTranslationError(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList {
	s.TranslationError = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatList) Validate() error {
	if s.AnnotationMissionChatCustomizationDataInfoList != nil {
		for _, item := range s.AnnotationMissionChatCustomizationDataInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AnnotationMissionChatIntentUserSayInfoList != nil {
		for _, item := range s.AnnotationMissionChatIntentUserSayInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AnnotationMissionChatTagInfoList != nil {
		for _, item := range s.AnnotationMissionChatTagInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AnnotationMissionChatVocabularyInfoList != nil {
		for _, item := range s.AnnotationMissionChatVocabularyInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList struct {
	// id
	AnnotationMissionChatCustomizationDataInfoId *string `json:"AnnotationMissionChatCustomizationDataInfoId,omitempty" xml:"AnnotationMissionChatCustomizationDataInfoId,omitempty"`
	// chat id
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// example:
	//
	// 2f03b24a-fda2-4501-90ba-0a9a59f8dd9d
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	Content                    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1682216045619
	CreateTime                   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomizationDataDescription *string `json:"CustomizationDataDescription,omitempty" xml:"CustomizationDataDescription,omitempty"`
	// id
	CustomizationDataId     *string `json:"CustomizationDataId,omitempty" xml:"CustomizationDataId,omitempty"`
	CustomizationDataName   *string `json:"CustomizationDataName,omitempty" xml:"CustomizationDataName,omitempty"`
	CustomizationDataWeight *int32  `json:"CustomizationDataWeight,omitempty" xml:"CustomizationDataWeight,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// d3fbfca8-4208-4d4b-a53a-c4dba5e43a66
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1679552585384
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionChatCustomizationDataInfoId() *string {
	return s.AnnotationMissionChatCustomizationDataInfoId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetContent() *string {
	return s.Content
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCreate() *bool {
	return s.Create
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataDescription() *string {
	return s.CustomizationDataDescription
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataId() *string {
	return s.CustomizationDataId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataName() *string {
	return s.CustomizationDataName
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataWeight() *int32 {
	return s.CustomizationDataWeight
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionChatCustomizationDataInfoId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionChatCustomizationDataInfoId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionChatId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetContent(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Content = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCreate(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Create = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataDescription(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataDescription = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataName(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataName = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataWeight(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataWeight = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetDelete(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Delete = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) Validate() error {
	return dara.Validate(s)
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList struct {
	// chat id
	//
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// id
	AnnotationMissionChatIntentUserSayInfoId *string `json:"AnnotationMissionChatIntentUserSayInfoId,omitempty" xml:"AnnotationMissionChatIntentUserSayInfoId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	BotId   *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// false
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1682216045619
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 1682216045619
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionChatIntentUserSayInfoId() *string {
	return s.AnnotationMissionChatIntentUserSayInfoId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetBotId() *string {
	return s.BotId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetContent() *string {
	return s.Content
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetCreate() *bool {
	return s.Create
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetDialogId() *int64 {
	return s.DialogId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionChatId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionChatIntentUserSayInfoId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionChatIntentUserSayInfoId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetBotId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.BotId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetContent(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Content = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetCreate(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Create = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetDelete(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Delete = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetDialogId(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.DialogId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetIntentId(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.IntentId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) Validate() error {
	return dara.Validate(s)
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList struct {
	// chat id
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// id
	AnnotationMissionChatTagInfoId *string `json:"AnnotationMissionChatTagInfoId,omitempty" xml:"AnnotationMissionChatTagInfoId,omitempty"`
	// example:
	//
	// e7272cbb-a60d-4b41-b5c4-8863edc0b8f7
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// tag id
	AnnotationMissionTagInfoId   *string `json:"AnnotationMissionTagInfoId,omitempty" xml:"AnnotationMissionTagInfoId,omitempty"`
	AnnotationMissionTagInfoName *string `json:"AnnotationMissionTagInfoName,omitempty" xml:"AnnotationMissionTagInfoName,omitempty"`
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1686797050000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 32be9d94-1346-4c4a-a4d0-ccd379f87013
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1679019660000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionChatTagInfoId() *string {
	return s.AnnotationMissionChatTagInfoId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionTagInfoId() *string {
	return s.AnnotationMissionTagInfoId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionTagInfoName() *string {
	return s.AnnotationMissionTagInfoName
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetCreate() *bool {
	return s.Create
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionChatId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionChatTagInfoId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionChatTagInfoId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionTagInfoId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionTagInfoId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionTagInfoName(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionTagInfoName = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetCreate(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.Create = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetDelete(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.Delete = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) Validate() error {
	return dara.Validate(s)
}

type ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList struct {
	// chat id
	//
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// id
	AnnotationMissionChatVocabularyInfoId *string `json:"AnnotationMissionChatVocabularyInfoId,omitempty" xml:"AnnotationMissionChatVocabularyInfoId,omitempty"`
	// example:
	//
	// b3f2c931-5180-43ca-b4aa-2baee2d73c8b
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1675218421941
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 818961ce-d9ba-4f88-89ca-2a04b24bdf01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1676272557653
	ModifiedTime          *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Vocabulary            *string `json:"Vocabulary,omitempty" xml:"Vocabulary,omitempty"`
	VocabularyDescription *string `json:"VocabularyDescription,omitempty" xml:"VocabularyDescription,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	VocabularyId     *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	VocabularyName   *string `json:"VocabularyName,omitempty" xml:"VocabularyName,omitempty"`
	VocabularyWeight *int32  `json:"VocabularyWeight,omitempty" xml:"VocabularyWeight,omitempty"`
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionChatVocabularyInfoId() *string {
	return s.AnnotationMissionChatVocabularyInfoId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetCreate() *bool {
	return s.Create
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabulary() *string {
	return s.Vocabulary
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyDescription() *string {
	return s.VocabularyDescription
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyName() *string {
	return s.VocabularyName
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyWeight() *int32 {
	return s.VocabularyWeight
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionChatId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionChatVocabularyInfoId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionChatVocabularyInfoId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetCreate(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Create = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetCreateTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetDelete(v bool) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Delete = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetInstanceId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetModifiedTime(v int64) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabulary(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Vocabulary = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyDescription(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyDescription = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyId(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyId = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyName(v string) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyName = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyWeight(v int32) *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyWeight = &v
	return s
}

func (s *ListAnnotationMissionSessionResponseBodyDataAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) Validate() error {
	return dara.Validate(s)
}
