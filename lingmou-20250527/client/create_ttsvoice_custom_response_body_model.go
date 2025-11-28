// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTTSVoiceCustomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTTSVoiceCustomResponseBody
	GetCode() *string
	SetData(v *CreateTTSVoiceCustomResponseBodyData) *CreateTTSVoiceCustomResponseBody
	GetData() *CreateTTSVoiceCustomResponseBodyData
	SetHttpStatusCode(v int32) *CreateTTSVoiceCustomResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTTSVoiceCustomResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTTSVoiceCustomResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTTSVoiceCustomResponseBody
	GetSuccess() *bool
}

type CreateTTSVoiceCustomResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateTTSVoiceCustomResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A7174E51-3523-5AEB-AE18-1D877FF22497
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTTSVoiceCustomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTTSVoiceCustomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTTSVoiceCustomResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTTSVoiceCustomResponseBody) GetData() *CreateTTSVoiceCustomResponseBodyData {
	return s.Data
}

func (s *CreateTTSVoiceCustomResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTTSVoiceCustomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTTSVoiceCustomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTTSVoiceCustomResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTTSVoiceCustomResponseBody) SetCode(v string) *CreateTTSVoiceCustomResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) SetData(v *CreateTTSVoiceCustomResponseBodyData) *CreateTTSVoiceCustomResponseBody {
	s.Data = v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) SetHttpStatusCode(v int32) *CreateTTSVoiceCustomResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) SetMessage(v string) *CreateTTSVoiceCustomResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) SetRequestId(v string) *CreateTTSVoiceCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) SetSuccess(v bool) *CreateTTSVoiceCustomResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTTSVoiceCustomResponseBodyData struct {
	// example:
	//
	// https://xxx-aliyuncs.com/material/INPUT_TRAIN_AUDIO/Mt.CQEG75L4FWIU2/TestTTSVoiceName.mp3?Expires=1764262805&OSSAccessKeyId=LTAI5tK3WcKwKtAyaTSeq7sx&Signature=D%2Fld6gp9Zh6TsGRU9cd6GD2pFY0%3D
	AudioURL *string `json:"audioURL,omitempty" xml:"audioURL,omitempty"`
	// example:
	//
	// CHECKING
	CensorStatus *string `json:"censorStatus,omitempty" xml:"censorStatus,omitempty"`
	// example:
	//
	// false
	Common *bool `json:"common,omitempty" xml:"common,omitempty"`
	// example:
	//
	// 2024-10-10T07:48:31Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// This is a testTTSVoice。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Error
	ErrorDetail *string `json:"errorDetail,omitempty" xml:"errorDetail,omitempty"`
	// example:
	//
	// FEMALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// M1lhKArheOyYdeYybDFqS1-Q
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2024-10-10T07:48:31Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// TestTTSVoiceName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 100
	PitchRate *float32 `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	// example:
	//
	// 100
	RemainSeconds *int64 `json:"remainSeconds,omitempty" xml:"remainSeconds,omitempty"`
	// example:
	//
	// 100
	SpeechRate *float32 `json:"speechRate,omitempty" xml:"speechRate,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Text   *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// testTTSVoice
	VoiceKey *string `json:"voiceKey,omitempty" xml:"voiceKey,omitempty"`
}

func (s CreateTTSVoiceCustomResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTTSVoiceCustomResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetAudioURL() *string {
	return s.AudioURL
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetCensorStatus() *string {
	return s.CensorStatus
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetCommon() *bool {
	return s.Common
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetGender() *string {
	return s.Gender
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetPitchRate() *float32 {
	return s.PitchRate
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetRemainSeconds() *int64 {
	return s.RemainSeconds
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetSpeechRate() *float32 {
	return s.SpeechRate
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetText() *string {
	return s.Text
}

func (s *CreateTTSVoiceCustomResponseBodyData) GetVoiceKey() *string {
	return s.VoiceKey
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetAudioURL(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetCensorStatus(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.CensorStatus = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetCommon(v bool) *CreateTTSVoiceCustomResponseBodyData {
	s.Common = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetCreateTime(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetDescription(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetErrorDetail(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.ErrorDetail = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetGender(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Gender = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetId(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetLanguage(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Language = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetModifiedTime(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetName(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetPitchRate(v float32) *CreateTTSVoiceCustomResponseBodyData {
	s.PitchRate = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetRemainSeconds(v int64) *CreateTTSVoiceCustomResponseBodyData {
	s.RemainSeconds = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetSpeechRate(v float32) *CreateTTSVoiceCustomResponseBodyData {
	s.SpeechRate = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetStatus(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetText(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.Text = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) SetVoiceKey(v string) *CreateTTSVoiceCustomResponseBodyData {
	s.VoiceKey = &v
	return s
}

func (s *CreateTTSVoiceCustomResponseBodyData) Validate() error {
	return dara.Validate(s)
}
