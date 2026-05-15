// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCodeBreak(v bool) *AiccsSmartCallRequest
	GetActionCodeBreak() *bool
	SetActionCodeTimeBreak(v int32) *AiccsSmartCallRequest
	GetActionCodeTimeBreak() *int32
	SetAsrAlsAmId(v string) *AiccsSmartCallRequest
	GetAsrAlsAmId() *string
	SetAsrBaseId(v string) *AiccsSmartCallRequest
	GetAsrBaseId() *string
	SetAsrModelId(v string) *AiccsSmartCallRequest
	GetAsrModelId() *string
	SetAsrVocabularyId(v string) *AiccsSmartCallRequest
	GetAsrVocabularyId() *string
	SetBackgroundFileCode(v string) *AiccsSmartCallRequest
	GetBackgroundFileCode() *string
	SetBackgroundSpeed(v int32) *AiccsSmartCallRequest
	GetBackgroundSpeed() *int32
	SetBackgroundVolume(v int32) *AiccsSmartCallRequest
	GetBackgroundVolume() *int32
	SetCalledNumber(v string) *AiccsSmartCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *AiccsSmartCallRequest
	GetCalledShowNumber() *string
	SetDynamicId(v string) *AiccsSmartCallRequest
	GetDynamicId() *string
	SetEarlyMediaAsr(v bool) *AiccsSmartCallRequest
	GetEarlyMediaAsr() *bool
	SetEnableITN(v bool) *AiccsSmartCallRequest
	GetEnableITN() *bool
	SetMuteTime(v int32) *AiccsSmartCallRequest
	GetMuteTime() *int32
	SetOutId(v string) *AiccsSmartCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *AiccsSmartCallRequest
	GetOwnerId() *int64
	SetPauseTime(v int32) *AiccsSmartCallRequest
	GetPauseTime() *int32
	SetPlayTimes(v int32) *AiccsSmartCallRequest
	GetPlayTimes() *int32
	SetProdCode(v string) *AiccsSmartCallRequest
	GetProdCode() *string
	SetRecordFlag(v bool) *AiccsSmartCallRequest
	GetRecordFlag() *bool
	SetResourceOwnerAccount(v string) *AiccsSmartCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AiccsSmartCallRequest
	GetResourceOwnerId() *int64
	SetSessionTimeout(v int32) *AiccsSmartCallRequest
	GetSessionTimeout() *int32
	SetSpeed(v int32) *AiccsSmartCallRequest
	GetSpeed() *int32
	SetTtsConf(v bool) *AiccsSmartCallRequest
	GetTtsConf() *bool
	SetTtsSpeed(v int32) *AiccsSmartCallRequest
	GetTtsSpeed() *int32
	SetTtsStyle(v string) *AiccsSmartCallRequest
	GetTtsStyle() *string
	SetTtsVolume(v int32) *AiccsSmartCallRequest
	GetTtsVolume() *int32
	SetVoiceCode(v string) *AiccsSmartCallRequest
	GetVoiceCode() *string
	SetVoiceCodeParam(v string) *AiccsSmartCallRequest
	GetVoiceCodeParam() *string
	SetVolume(v int32) *AiccsSmartCallRequest
	GetVolume() *int32
}

type AiccsSmartCallRequest struct {
	// example:
	//
	// true
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// example:
	//
	// 120
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// example:
	//
	// 23387****
	AsrAlsAmId *string `json:"AsrAlsAmId,omitempty" xml:"AsrAlsAmId,omitempty"`
	// example:
	//
	// customer_service_8k
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// example:
	//
	// bf71664d30d2478fb8cb8c39c6b6****
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// example:
	//
	// 6689****
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// example:
	//
	// 1
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// example:
	//
	// 1
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1862222****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0571000****
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// example:
	//
	// 2234****
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// example:
	//
	// fasle
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// example:
	//
	// true
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// example:
	//
	// 10000
	MuteTime *int32 `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	// example:
	//
	// 222356****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 800
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// example:
	//
	// 1
	PlayTimes *int32 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// example:
	//
	// aiccs
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// example:
	//
	// 1
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// true
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// example:
	//
	// 100
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// example:
	//
	// xiaoyun
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2d4c-4e78-8d2a-afbb06cf****.wav,$name$
	VoiceCode      *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// example:
	//
	// 1
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AiccsSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallRequest) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallRequest) GetActionCodeBreak() *bool {
	return s.ActionCodeBreak
}

func (s *AiccsSmartCallRequest) GetActionCodeTimeBreak() *int32 {
	return s.ActionCodeTimeBreak
}

func (s *AiccsSmartCallRequest) GetAsrAlsAmId() *string {
	return s.AsrAlsAmId
}

func (s *AiccsSmartCallRequest) GetAsrBaseId() *string {
	return s.AsrBaseId
}

func (s *AiccsSmartCallRequest) GetAsrModelId() *string {
	return s.AsrModelId
}

func (s *AiccsSmartCallRequest) GetAsrVocabularyId() *string {
	return s.AsrVocabularyId
}

func (s *AiccsSmartCallRequest) GetBackgroundFileCode() *string {
	return s.BackgroundFileCode
}

func (s *AiccsSmartCallRequest) GetBackgroundSpeed() *int32 {
	return s.BackgroundSpeed
}

func (s *AiccsSmartCallRequest) GetBackgroundVolume() *int32 {
	return s.BackgroundVolume
}

func (s *AiccsSmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *AiccsSmartCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *AiccsSmartCallRequest) GetDynamicId() *string {
	return s.DynamicId
}

func (s *AiccsSmartCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *AiccsSmartCallRequest) GetEnableITN() *bool {
	return s.EnableITN
}

func (s *AiccsSmartCallRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *AiccsSmartCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *AiccsSmartCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AiccsSmartCallRequest) GetPauseTime() *int32 {
	return s.PauseTime
}

func (s *AiccsSmartCallRequest) GetPlayTimes() *int32 {
	return s.PlayTimes
}

func (s *AiccsSmartCallRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *AiccsSmartCallRequest) GetRecordFlag() *bool {
	return s.RecordFlag
}

func (s *AiccsSmartCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AiccsSmartCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AiccsSmartCallRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *AiccsSmartCallRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *AiccsSmartCallRequest) GetTtsConf() *bool {
	return s.TtsConf
}

func (s *AiccsSmartCallRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *AiccsSmartCallRequest) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *AiccsSmartCallRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *AiccsSmartCallRequest) GetVoiceCode() *string {
	return s.VoiceCode
}

func (s *AiccsSmartCallRequest) GetVoiceCodeParam() *string {
	return s.VoiceCodeParam
}

func (s *AiccsSmartCallRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *AiccsSmartCallRequest) SetActionCodeBreak(v bool) *AiccsSmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetActionCodeTimeBreak(v int32) *AiccsSmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrAlsAmId(v string) *AiccsSmartCallRequest {
	s.AsrAlsAmId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrBaseId(v string) *AiccsSmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrModelId(v string) *AiccsSmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetAsrVocabularyId(v string) *AiccsSmartCallRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundFileCode(v string) *AiccsSmartCallRequest {
	s.BackgroundFileCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundSpeed(v int32) *AiccsSmartCallRequest {
	s.BackgroundSpeed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetBackgroundVolume(v int32) *AiccsSmartCallRequest {
	s.BackgroundVolume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledNumber(v string) *AiccsSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetCalledShowNumber(v string) *AiccsSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *AiccsSmartCallRequest) SetDynamicId(v string) *AiccsSmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetEarlyMediaAsr(v bool) *AiccsSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *AiccsSmartCallRequest) SetEnableITN(v bool) *AiccsSmartCallRequest {
	s.EnableITN = &v
	return s
}

func (s *AiccsSmartCallRequest) SetMuteTime(v int32) *AiccsSmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetOutId(v string) *AiccsSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetOwnerId(v int64) *AiccsSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPauseTime(v int32) *AiccsSmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *AiccsSmartCallRequest) SetPlayTimes(v int32) *AiccsSmartCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *AiccsSmartCallRequest) SetProdCode(v string) *AiccsSmartCallRequest {
	s.ProdCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetRecordFlag(v bool) *AiccsSmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerAccount(v string) *AiccsSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AiccsSmartCallRequest) SetResourceOwnerId(v int64) *AiccsSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSessionTimeout(v int32) *AiccsSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *AiccsSmartCallRequest) SetSpeed(v int32) *AiccsSmartCallRequest {
	s.Speed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsConf(v bool) *AiccsSmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsSpeed(v int32) *AiccsSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsStyle(v string) *AiccsSmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *AiccsSmartCallRequest) SetTtsVolume(v int32) *AiccsSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCode(v string) *AiccsSmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVoiceCodeParam(v string) *AiccsSmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *AiccsSmartCallRequest) SetVolume(v int32) *AiccsSmartCallRequest {
	s.Volume = &v
	return s
}

func (s *AiccsSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
