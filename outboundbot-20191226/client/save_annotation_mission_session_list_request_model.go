// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionSessionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SaveAnnotationMissionSessionListRequest
	GetAgentId() *string
	SetAgentKey(v string) *SaveAnnotationMissionSessionListRequest
	GetAgentKey() *string
	SetAnnotationMissionDataSourceType(v int64) *SaveAnnotationMissionSessionListRequest
	GetAnnotationMissionDataSourceType() *int64
	SetAnnotationMissionSessionList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) *SaveAnnotationMissionSessionListRequest
	GetAnnotationMissionSessionList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList
	SetAnnotationMissionSessionListJsonString(v string) *SaveAnnotationMissionSessionListRequest
	GetAnnotationMissionSessionListJsonString() *string
	SetEnvironment(v int64) *SaveAnnotationMissionSessionListRequest
	GetEnvironment() *int64
	SetUserNick(v string) *SaveAnnotationMissionSessionListRequest
	GetUserNick() *string
}

type SaveAnnotationMissionSessionListRequest struct {
	// example:
	//
	// 1198938
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	AnnotationMissionDataSourceType *int64                                                                 `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	AnnotationMissionSessionList    []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList `json:"AnnotationMissionSessionList,omitempty" xml:"AnnotationMissionSessionList,omitempty" type:"Repeated"`
	// example:
	//
	// []
	AnnotationMissionSessionListJsonString *string `json:"AnnotationMissionSessionListJsonString,omitempty" xml:"AnnotationMissionSessionListJsonString,omitempty"`
	// example:
	//
	// 0
	Environment *int64  `json:"Environment,omitempty" xml:"Environment,omitempty"`
	UserNick    *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequest) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SaveAnnotationMissionSessionListRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveAnnotationMissionSessionListRequest) GetAnnotationMissionDataSourceType() *int64 {
	return s.AnnotationMissionDataSourceType
}

func (s *SaveAnnotationMissionSessionListRequest) GetAnnotationMissionSessionList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	return s.AnnotationMissionSessionList
}

func (s *SaveAnnotationMissionSessionListRequest) GetAnnotationMissionSessionListJsonString() *string {
	return s.AnnotationMissionSessionListJsonString
}

func (s *SaveAnnotationMissionSessionListRequest) GetEnvironment() *int64 {
	return s.Environment
}

func (s *SaveAnnotationMissionSessionListRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *SaveAnnotationMissionSessionListRequest) SetAgentId(v string) *SaveAnnotationMissionSessionListRequest {
	s.AgentId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetAgentKey(v string) *SaveAnnotationMissionSessionListRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetAnnotationMissionDataSourceType(v int64) *SaveAnnotationMissionSessionListRequest {
	s.AnnotationMissionDataSourceType = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetAnnotationMissionSessionList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) *SaveAnnotationMissionSessionListRequest {
	s.AnnotationMissionSessionList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetAnnotationMissionSessionListJsonString(v string) *SaveAnnotationMissionSessionListRequest {
	s.AnnotationMissionSessionListJsonString = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetEnvironment(v int64) *SaveAnnotationMissionSessionListRequest {
	s.Environment = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) SetUserNick(v string) *SaveAnnotationMissionSessionListRequest {
	s.UserNick = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequest) Validate() error {
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

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList struct {
	AnnotationMissionChatList []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList `json:"AnnotationMissionChatList,omitempty" xml:"AnnotationMissionChatList,omitempty" type:"Repeated"`
	// example:
	//
	// c88cc004-de69-4eee-aa5f-2efed533a54e
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// 2023-04-14T02:01:23Z
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 29e669bd-a9d1-4529-98cd-c2b0549bcf53
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// 593aaf5e-1275-4add-9990-22696385dc6e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1683858248778
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// e4f32632-2e56-4399-9fec-47bdbaeefdf6
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// e6271044-b4b2-4ad8-ade4-c720be023538
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetAnnotationMissionChatList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	return s.AnnotationMissionChatList
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetJobId() *string {
	return s.JobId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetScriptId() *string {
	return s.ScriptId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) GetSessionId() *string {
	return s.SessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetAnnotationMissionChatList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.AnnotationMissionChatList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetAnnotationStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.AnnotationStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetJobGroupId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.JobGroupId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetJobId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.JobId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetScriptId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.ScriptId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) SetSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList {
	s.SessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionList) Validate() error {
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

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList struct {
	AnnotationAsrResult                            *string                                                                                                                                       `json:"AnnotationAsrResult,omitempty" xml:"AnnotationAsrResult,omitempty"`
	AnnotationMissionChatCustomizationDataInfoList []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList `json:"AnnotationMissionChatCustomizationDataInfoList,omitempty" xml:"AnnotationMissionChatCustomizationDataInfoList,omitempty" type:"Repeated"`
	// chat id
	AnnotationMissionChatId                    *string                                                                                                                                   `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	AnnotationMissionChatIntentUserSayInfoList []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList `json:"AnnotationMissionChatIntentUserSayInfoList,omitempty" xml:"AnnotationMissionChatIntentUserSayInfoList,omitempty" type:"Repeated"`
	AnnotationMissionChatTagInfoList           []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList           `json:"AnnotationMissionChatTagInfoList,omitempty" xml:"AnnotationMissionChatTagInfoList,omitempty" type:"Repeated"`
	AnnotationMissionChatVocabularyInfoList    []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList    `json:"AnnotationMissionChatVocabularyInfoList,omitempty" xml:"AnnotationMissionChatVocabularyInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// ddce607f-f537-4ebd-9914-cf45671defb9
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// {\\"Answer\\": u\\"\\u53c2\\u8003\\u6587\\u6863\\uff1ahttps://help.aliyun.com/document_detail/181325.html\\", \\"QuestionId\\": 372858, \\"Uuid\\": \\"ac14000516762684770197536d0044\\"}
	Answer              *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	AsrAnnotationStatus *int32  `json:"AsrAnnotationStatus,omitempty" xml:"AsrAnnotationStatus,omitempty"`
	// example:
	//
	// 1679629770336
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 5ec263fa-c8de-46f4-b844-6fb8275bb645
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentAnnotationStatus *int32  `json:"IntentAnnotationStatus,omitempty" xml:"IntentAnnotationStatus,omitempty"`
	// example:
	//
	// 1629360780000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 1682493047418
	OccurTime         *int64  `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	OriginalAsrResult *string `json:"OriginalAsrResult,omitempty" xml:"OriginalAsrResult,omitempty"`
	// example:
	//
	// 1475140522
	SequenceId *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	// example:
	//
	// 1
	SubStatus *int32 `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// example:
	//
	// 1
	TagAnnotationStatus *int32 `json:"TagAnnotationStatus,omitempty" xml:"TagAnnotationStatus,omitempty"`
	// example:
	//
	// 1
	TranslationError *int32 `json:"TranslationError,omitempty" xml:"TranslationError,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationAsrResult() *string {
	return s.AnnotationAsrResult
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatCustomizationDataInfoList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	return s.AnnotationMissionChatCustomizationDataInfoList
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatIntentUserSayInfoList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	return s.AnnotationMissionChatIntentUserSayInfoList
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatTagInfoList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	return s.AnnotationMissionChatTagInfoList
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionChatVocabularyInfoList() []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	return s.AnnotationMissionChatVocabularyInfoList
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAnswer() *string {
	return s.Answer
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetAsrAnnotationStatus() *int32 {
	return s.AsrAnnotationStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetIntentAnnotationStatus() *int32 {
	return s.IntentAnnotationStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetOccurTime() *int64 {
	return s.OccurTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetOriginalAsrResult() *string {
	return s.OriginalAsrResult
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetSequenceId() *string {
	return s.SequenceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetSubStatus() *int32 {
	return s.SubStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetTagAnnotationStatus() *int32 {
	return s.TagAnnotationStatus
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) GetTranslationError() *int32 {
	return s.TranslationError
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationAsrResult(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationAsrResult = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatCustomizationDataInfoList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatCustomizationDataInfoList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatIntentUserSayInfoList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatIntentUserSayInfoList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatTagInfoList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatTagInfoList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionChatVocabularyInfoList(v []*SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionChatVocabularyInfoList = v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnnotationStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AnnotationStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAnswer(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.Answer = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetAsrAnnotationStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.AsrAnnotationStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetIntentAnnotationStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.IntentAnnotationStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetOccurTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.OccurTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetOriginalAsrResult(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.OriginalAsrResult = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetSequenceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.SequenceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetSubStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.SubStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetTagAnnotationStatus(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.TagAnnotationStatus = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) SetTranslationError(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList {
	s.TranslationError = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatList) Validate() error {
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

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList struct {
	// id
	AnnotationMissionChatCustomizationDataInfoId *string `json:"AnnotationMissionChatCustomizationDataInfoId,omitempty" xml:"AnnotationMissionChatCustomizationDataInfoId,omitempty"`
	// chat id
	//
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// example:
	//
	// 03f56192-fa8a-40dc-9558-39b357f0618f
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	Content                    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1682316909210
	CreateTime                   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomizationDataDescription *string `json:"CustomizationDataDescription,omitempty" xml:"CustomizationDataDescription,omitempty"`
	// id
	//
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	CustomizationDataId     *string `json:"CustomizationDataId,omitempty" xml:"CustomizationDataId,omitempty"`
	CustomizationDataName   *string `json:"CustomizationDataName,omitempty" xml:"CustomizationDataName,omitempty"`
	CustomizationDataWeight *int32  `json:"CustomizationDataWeight,omitempty" xml:"CustomizationDataWeight,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// b598a47a-7c03-45fb-af6b-343950dd9512
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1673438100000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionChatCustomizationDataInfoId() *string {
	return s.AnnotationMissionChatCustomizationDataInfoId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetContent() *string {
	return s.Content
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCreate() *bool {
	return s.Create
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataDescription() *string {
	return s.CustomizationDataDescription
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataId() *string {
	return s.CustomizationDataId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataName() *string {
	return s.CustomizationDataName
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetCustomizationDataWeight() *int32 {
	return s.CustomizationDataWeight
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionChatCustomizationDataInfoId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionChatCustomizationDataInfoId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionChatId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetContent(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Content = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCreate(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Create = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataDescription(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataDescription = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataName(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataName = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetCustomizationDataWeight(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.CustomizationDataWeight = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetDelete(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.Delete = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatCustomizationDataInfoList) Validate() error {
	return dara.Validate(s)
}

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList struct {
	AnnotationMissionChatId                  *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	AnnotationMissionChatIntentUserSayInfoId *string `json:"AnnotationMissionChatIntentUserSayInfoId,omitempty" xml:"AnnotationMissionChatIntentUserSayInfoId,omitempty"`
	AnnotationMissionId                      *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId               *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	BotId                                    *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	Content                                  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Create                                   *bool   `json:"Create,omitempty" xml:"Create,omitempty"`
	CreateTime                               *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Delete                                   *bool   `json:"Delete,omitempty" xml:"Delete,omitempty"`
	DialogId                                 *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	InstanceId                               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentId                                 *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	ModifiedTime                             *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionChatIntentUserSayInfoId() *string {
	return s.AnnotationMissionChatIntentUserSayInfoId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetBotId() *string {
	return s.BotId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetContent() *string {
	return s.Content
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetCreate() *bool {
	return s.Create
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetDialogId() *int64 {
	return s.DialogId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetIntentId() *int64 {
	return s.IntentId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionChatId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionChatIntentUserSayInfoId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionChatIntentUserSayInfoId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetBotId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.BotId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetContent(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Content = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetCreate(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Create = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetDelete(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.Delete = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetDialogId(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.DialogId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetIntentId(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.IntentId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatIntentUserSayInfoList) Validate() error {
	return dara.Validate(s)
}

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList struct {
	// chat id
	//
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// id
	AnnotationMissionChatTagInfoId *string `json:"AnnotationMissionChatTagInfoId,omitempty" xml:"AnnotationMissionChatTagInfoId,omitempty"`
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// tag id
	//
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionTagInfoId *string `json:"AnnotationMissionTagInfoId,omitempty" xml:"AnnotationMissionTagInfoId,omitempty"`
	// example:
	//
	// -
	AnnotationMissionTagInfoName *string `json:"AnnotationMissionTagInfoName,omitempty" xml:"AnnotationMissionTagInfoName,omitempty"`
	// example:
	//
	// true
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1679710866060
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 32be9d94-1346-4c4a-a4d0-ccd379f87013
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1679710866060
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionChatTagInfoId() *string {
	return s.AnnotationMissionChatTagInfoId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionTagInfoId() *string {
	return s.AnnotationMissionTagInfoId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetAnnotationMissionTagInfoName() *string {
	return s.AnnotationMissionTagInfoName
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetCreate() *bool {
	return s.Create
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionChatId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionChatTagInfoId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionChatTagInfoId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionTagInfoId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionTagInfoId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetAnnotationMissionTagInfoName(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.AnnotationMissionTagInfoName = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetCreate(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.Create = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetDelete(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.Delete = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatTagInfoList) Validate() error {
	return dara.Validate(s)
}

type SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList struct {
	// chat id
	AnnotationMissionChatId *string `json:"AnnotationMissionChatId,omitempty" xml:"AnnotationMissionChatId,omitempty"`
	// id
	AnnotationMissionChatVocabularyInfoId *string `json:"AnnotationMissionChatVocabularyInfoId,omitempty" xml:"AnnotationMissionChatVocabularyInfoId,omitempty"`
	// example:
	//
	// 977a45dc-b636-4407-9e98-9f572c709ada
	AnnotationMissionId        *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// example:
	//
	// false
	Create *bool `json:"Create,omitempty" xml:"Create,omitempty"`
	// example:
	//
	// 1677552860720
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 77343553-cbc2-4487-a35c-869f1e86c573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1679283408230
	ModifiedTime          *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Vocabulary            *string `json:"Vocabulary,omitempty" xml:"Vocabulary,omitempty"`
	VocabularyDescription *string `json:"VocabularyDescription,omitempty" xml:"VocabularyDescription,omitempty"`
	VocabularyId          *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	VocabularyName        *string `json:"VocabularyName,omitempty" xml:"VocabularyName,omitempty"`
	VocabularyWeight      *int32  `json:"VocabularyWeight,omitempty" xml:"VocabularyWeight,omitempty"`
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionChatId() *string {
	return s.AnnotationMissionChatId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionChatVocabularyInfoId() *string {
	return s.AnnotationMissionChatVocabularyInfoId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetCreate() *bool {
	return s.Create
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabulary() *string {
	return s.Vocabulary
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyDescription() *string {
	return s.VocabularyDescription
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyName() *string {
	return s.VocabularyName
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) GetVocabularyWeight() *int32 {
	return s.VocabularyWeight
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionChatId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionChatId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionChatVocabularyInfoId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionChatVocabularyInfoId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetAnnotationMissionSessionId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetCreate(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Create = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetCreateTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.CreateTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetDelete(v bool) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Delete = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetInstanceId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetModifiedTime(v int64) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabulary(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.Vocabulary = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyDescription(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyDescription = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyId(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyName(v string) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyName = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) SetVocabularyWeight(v int32) *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList {
	s.VocabularyWeight = &v
	return s
}

func (s *SaveAnnotationMissionSessionListRequestAnnotationMissionSessionListAnnotationMissionChatListAnnotationMissionChatVocabularyInfoList) Validate() error {
	return dara.Validate(s)
}
