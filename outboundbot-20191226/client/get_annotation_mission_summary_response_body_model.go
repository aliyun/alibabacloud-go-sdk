// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAnnotationMissionSummaryResponseBody
	GetCode() *string
	SetData(v *GetAnnotationMissionSummaryResponseBodyData) *GetAnnotationMissionSummaryResponseBody
	GetData() *GetAnnotationMissionSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetAnnotationMissionSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAnnotationMissionSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAnnotationMissionSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAnnotationMissionSummaryResponseBody
	GetSuccess() *bool
}

type GetAnnotationMissionSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAnnotationMissionSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// You are not authorized to perform this action. CDR:View privileges are required.
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

func (s GetAnnotationMissionSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAnnotationMissionSummaryResponseBody) GetData() *GetAnnotationMissionSummaryResponseBodyData {
	return s.Data
}

func (s *GetAnnotationMissionSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAnnotationMissionSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAnnotationMissionSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAnnotationMissionSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAnnotationMissionSummaryResponseBody) SetCode(v string) *GetAnnotationMissionSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) SetData(v *GetAnnotationMissionSummaryResponseBodyData) *GetAnnotationMissionSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) SetHttpStatusCode(v int32) *GetAnnotationMissionSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) SetMessage(v string) *GetAnnotationMissionSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) SetRequestId(v string) *GetAnnotationMissionSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) SetSuccess(v bool) *GetAnnotationMissionSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAnnotationMissionSummaryResponseBodyData struct {
	// example:
	//
	// b3f2c931-5180-43ca-b4aa-2baee2d73c8b
	AnnotationMissionId *string                                                       `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AsrSummaryInfo      *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo    `json:"AsrSummaryInfo,omitempty" xml:"AsrSummaryInfo,omitempty" type:"Struct"`
	IntentSummaryInfo   *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo `json:"IntentSummaryInfo,omitempty" xml:"IntentSummaryInfo,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this action. CDR:View privileges are required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success        *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TagSummaryInfo *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo `json:"TagSummaryInfo,omitempty" xml:"TagSummaryInfo,omitempty" type:"Struct"`
}

func (s GetAnnotationMissionSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetAsrSummaryInfo() *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	return s.AsrSummaryInfo
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetIntentSummaryInfo() *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	return s.IntentSummaryInfo
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetAnnotationMissionSummaryResponseBodyData) GetTagSummaryInfo() *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo {
	return s.TagSummaryInfo
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetAnnotationMissionId(v string) *GetAnnotationMissionSummaryResponseBodyData {
	s.AnnotationMissionId = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetAsrSummaryInfo(v *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) *GetAnnotationMissionSummaryResponseBodyData {
	s.AsrSummaryInfo = v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetIntentSummaryInfo(v *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) *GetAnnotationMissionSummaryResponseBodyData {
	s.IntentSummaryInfo = v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetMessage(v string) *GetAnnotationMissionSummaryResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetSuccess(v bool) *GetAnnotationMissionSummaryResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) SetTagSummaryInfo(v *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) *GetAnnotationMissionSummaryResponseBodyData {
	s.TagSummaryInfo = v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyData) Validate() error {
	if s.AsrSummaryInfo != nil {
		if err := s.AsrSummaryInfo.Validate(); err != nil {
			return err
		}
	}
	if s.IntentSummaryInfo != nil {
		if err := s.IntentSummaryInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TagSummaryInfo != nil {
		if err := s.TagSummaryInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo struct {
	// example:
	//
	// 1
	AddCustomizationDataCount *int32 `json:"AddCustomizationDataCount,omitempty" xml:"AddCustomizationDataCount,omitempty"`
	// example:
	//
	// 1
	AddVocabularyDataCount *int32 `json:"AddVocabularyDataCount,omitempty" xml:"AddVocabularyDataCount,omitempty"`
	// example:
	//
	// 1
	AnnotationInvalid *int32 `json:"AnnotationInvalid,omitempty" xml:"AnnotationInvalid,omitempty"`
	// example:
	//
	// 1
	CharacterCorrectRate *int32 `json:"CharacterCorrectRate,omitempty" xml:"CharacterCorrectRate,omitempty"`
	// example:
	//
	// 1
	CharacterErrorRate *int32 `json:"CharacterErrorRate,omitempty" xml:"CharacterErrorRate,omitempty"`
	// example:
	//
	// 1
	ChatTotalCount *int32 `json:"ChatTotalCount,omitempty" xml:"ChatTotalCount,omitempty"`
	// example:
	//
	// 1
	NoAnnotation *int32 `json:"NoAnnotation,omitempty" xml:"NoAnnotation,omitempty"`
	// example:
	//
	// 1
	SentenceErrorRate *int32 `json:"SentenceErrorRate,omitempty" xml:"SentenceErrorRate,omitempty"`
	// example:
	//
	// 1
	WordErrorRate *int32 `json:"WordErrorRate,omitempty" xml:"WordErrorRate,omitempty"`
}

func (s GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetAddCustomizationDataCount() *int32 {
	return s.AddCustomizationDataCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetAddVocabularyDataCount() *int32 {
	return s.AddVocabularyDataCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetAnnotationInvalid() *int32 {
	return s.AnnotationInvalid
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetCharacterCorrectRate() *int32 {
	return s.CharacterCorrectRate
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetCharacterErrorRate() *int32 {
	return s.CharacterErrorRate
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetChatTotalCount() *int32 {
	return s.ChatTotalCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetNoAnnotation() *int32 {
	return s.NoAnnotation
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetSentenceErrorRate() *int32 {
	return s.SentenceErrorRate
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) GetWordErrorRate() *int32 {
	return s.WordErrorRate
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetAddCustomizationDataCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.AddCustomizationDataCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetAddVocabularyDataCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.AddVocabularyDataCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetAnnotationInvalid(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.AnnotationInvalid = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetCharacterCorrectRate(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.CharacterCorrectRate = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetCharacterErrorRate(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.CharacterErrorRate = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetChatTotalCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.ChatTotalCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetNoAnnotation(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.NoAnnotation = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetSentenceErrorRate(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.SentenceErrorRate = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) SetWordErrorRate(v int32) *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo {
	s.WordErrorRate = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataAsrSummaryInfo) Validate() error {
	return dara.Validate(s)
}

type GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo struct {
	// example:
	//
	// 1
	AnnotationCorrectCount *int32 `json:"AnnotationCorrectCount,omitempty" xml:"AnnotationCorrectCount,omitempty"`
	// example:
	//
	// 1
	AnnotationInvalid *int32 `json:"AnnotationInvalid,omitempty" xml:"AnnotationInvalid,omitempty"`
	// example:
	//
	// 1
	ChatTotalCount *int32 `json:"ChatTotalCount,omitempty" xml:"ChatTotalCount,omitempty"`
	// example:
	//
	// 1
	IntentUserSayCount *int32 `json:"IntentUserSayCount,omitempty" xml:"IntentUserSayCount,omitempty"`
	// example:
	//
	// 1
	IntentionNotCoveredCount *int32 `json:"IntentionNotCoveredCount,omitempty" xml:"IntentionNotCoveredCount,omitempty"`
	// example:
	//
	// 1
	MatchErrorCount *int32 `json:"MatchErrorCount,omitempty" xml:"MatchErrorCount,omitempty"`
	// example:
	//
	// 1
	NoAnnotation *int32 `json:"NoAnnotation,omitempty" xml:"NoAnnotation,omitempty"`
	// example:
	//
	// 1
	TranslationUnrecognizedCount *int32 `json:"TranslationUnrecognizedCount,omitempty" xml:"TranslationUnrecognizedCount,omitempty"`
}

func (s GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetAnnotationCorrectCount() *int32 {
	return s.AnnotationCorrectCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetAnnotationInvalid() *int32 {
	return s.AnnotationInvalid
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetChatTotalCount() *int32 {
	return s.ChatTotalCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetIntentUserSayCount() *int32 {
	return s.IntentUserSayCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetIntentionNotCoveredCount() *int32 {
	return s.IntentionNotCoveredCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetMatchErrorCount() *int32 {
	return s.MatchErrorCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetNoAnnotation() *int32 {
	return s.NoAnnotation
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) GetTranslationUnrecognizedCount() *int32 {
	return s.TranslationUnrecognizedCount
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetAnnotationCorrectCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.AnnotationCorrectCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetAnnotationInvalid(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.AnnotationInvalid = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetChatTotalCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.ChatTotalCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetIntentUserSayCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.IntentUserSayCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetIntentionNotCoveredCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.IntentionNotCoveredCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetMatchErrorCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.MatchErrorCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetNoAnnotation(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.NoAnnotation = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) SetTranslationUnrecognizedCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo {
	s.TranslationUnrecognizedCount = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataIntentSummaryInfo) Validate() error {
	return dara.Validate(s)
}

type GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo struct {
	TagSummaryInfoDetailList []*GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList `json:"TagSummaryInfoDetailList,omitempty" xml:"TagSummaryInfoDetailList,omitempty" type:"Repeated"`
}

func (s GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) GetTagSummaryInfoDetailList() []*GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList {
	return s.TagSummaryInfoDetailList
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) SetTagSummaryInfoDetailList(v []*GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo {
	s.TagSummaryInfoDetailList = v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfo) Validate() error {
	if s.TagSummaryInfoDetailList != nil {
		for _, item := range s.TagSummaryInfoDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList struct {
	// example:
	//
	// 3
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) GetCount() *int32 {
	return s.Count
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) GetName() *string {
	return s.Name
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) SetCount(v int32) *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList {
	s.Count = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) SetName(v string) *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList {
	s.Name = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponseBodyDataTagSummaryInfoTagSummaryInfoDetailList) Validate() error {
	return dara.Validate(s)
}
