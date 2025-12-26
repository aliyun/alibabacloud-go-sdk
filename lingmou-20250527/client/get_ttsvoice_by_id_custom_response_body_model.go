// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTTSVoiceByIdCustomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTTSVoiceByIdCustomResponseBody
	GetCode() *string
	SetData(v *GetTTSVoiceByIdCustomResponseBodyData) *GetTTSVoiceByIdCustomResponseBody
	GetData() *GetTTSVoiceByIdCustomResponseBodyData
	SetHttpStatusCode(v int32) *GetTTSVoiceByIdCustomResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTTSVoiceByIdCustomResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTTSVoiceByIdCustomResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTTSVoiceByIdCustomResponseBody
	GetSuccess() *bool
}

type GetTTSVoiceByIdCustomResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTTSVoiceByIdCustomResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 7239F9E5-A4DB-55BA-B701-0CE47DBD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTTSVoiceByIdCustomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTTSVoiceByIdCustomResponseBody) GoString() string {
	return s.String()
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetData() *GetTTSVoiceByIdCustomResponseBodyData {
	return s.Data
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTTSVoiceByIdCustomResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetCode(v string) *GetTTSVoiceByIdCustomResponseBody {
	s.Code = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetData(v *GetTTSVoiceByIdCustomResponseBodyData) *GetTTSVoiceByIdCustomResponseBody {
	s.Data = v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetHttpStatusCode(v int32) *GetTTSVoiceByIdCustomResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetMessage(v string) *GetTTSVoiceByIdCustomResponseBody {
	s.Message = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetRequestId(v string) *GetTTSVoiceByIdCustomResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) SetSuccess(v bool) *GetTTSVoiceByIdCustomResponseBody {
	s.Success = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTTSVoiceByIdCustomResponseBodyData struct {
	// example:
	//
	// https://xxx-aliyuncs.com/material/INPUT_TTS_VOICE/Mt.CQEG75L4FWIU2/TestTTSVoiceName.mp3?Expires=1764262805&OSSAccessKeyId=LTAI5tK3WcKwKtAyaTSe*****&Signature=D%2Fld6gp9Zh6TsGRU9cd6GD2pFY0%3D
	AudioURL *string `json:"audioURL,omitempty" xml:"audioURL,omitempty"`
	// example:
	//
	// CHECKING。
	CensorStatus *string `json:"censorStatus,omitempty" xml:"censorStatus,omitempty"`
	// example:
	//
	// false。
	Common *bool `json:"common,omitempty" xml:"common,omitempty"`
	// example:
	//
	// 2025-11-28T10:11:28
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// This is a testTTSVoice。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorDetail *string `json:"errorDetail,omitempty" xml:"errorDetail,omitempty"`
	// example:
	//
	// FEMALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// M1lhKArheOyYdeYyb*****
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2025-11-28T13:41:31
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// TestTTSVoiceName。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 100。
	RemainSeconds *int64 `json:"remainSeconds,omitempty" xml:"remainSeconds,omitempty"`
	// example:
	//
	// SUCCESS。
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Text   *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// avatar-2464b55a65794e75a20fe07dde2*****
	VoiceKey *string `json:"voiceKey,omitempty" xml:"voiceKey,omitempty"`
}

func (s GetTTSVoiceByIdCustomResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTTSVoiceByIdCustomResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetAudioURL() *string {
	return s.AudioURL
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetCensorStatus() *string {
	return s.CensorStatus
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetCommon() *bool {
	return s.Common
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetGender() *string {
	return s.Gender
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetRemainSeconds() *int64 {
	return s.RemainSeconds
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetText() *string {
	return s.Text
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) GetVoiceKey() *string {
	return s.VoiceKey
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetAudioURL(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetCensorStatus(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.CensorStatus = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetCommon(v bool) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Common = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetCreateTime(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetDescription(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetErrorDetail(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.ErrorDetail = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetGender(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetId(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetLanguage(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetModifiedTime(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetName(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetRemainSeconds(v int64) *GetTTSVoiceByIdCustomResponseBodyData {
	s.RemainSeconds = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetStatus(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetText(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.Text = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) SetVoiceKey(v string) *GetTTSVoiceByIdCustomResponseBodyData {
	s.VoiceKey = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponseBodyData) Validate() error {
	return dara.Validate(s)
}
