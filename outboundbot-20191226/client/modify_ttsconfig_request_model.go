// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTTSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *ModifyTTSConfigRequest
	GetAppKey() *string
	SetInstanceId(v string) *ModifyTTSConfigRequest
	GetInstanceId() *string
	SetNlsServiceType(v string) *ModifyTTSConfigRequest
	GetNlsServiceType() *string
	SetPitchRate(v string) *ModifyTTSConfigRequest
	GetPitchRate() *string
	SetScriptId(v string) *ModifyTTSConfigRequest
	GetScriptId() *string
	SetSpeechRate(v string) *ModifyTTSConfigRequest
	GetSpeechRate() *string
	SetVoice(v string) *ModifyTTSConfigRequest
	GetVoice() *string
	SetVolume(v string) *ModifyTTSConfigRequest
	GetVolume() *string
}

type ModifyTTSConfigRequest struct {
	// example:
	//
	// 99****Aw
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 27244bae-e446-4811-bb1d-f8a07b525af0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Managed
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// 语调 [-500,500]之间整数。默认值为0。
	//
	// 大于0表示升高音高。
	//
	// 小于0表示降低音高。
	PitchRate *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1f1a2ba0-b3e7-4ff9-baf1-6dc8aeac0791
	ScriptId   *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 100
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyTTSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ModifyTTSConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTTSConfigRequest) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *ModifyTTSConfigRequest) GetPitchRate() *string {
	return s.PitchRate
}

func (s *ModifyTTSConfigRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyTTSConfigRequest) GetSpeechRate() *string {
	return s.SpeechRate
}

func (s *ModifyTTSConfigRequest) GetVoice() *string {
	return s.Voice
}

func (s *ModifyTTSConfigRequest) GetVolume() *string {
	return s.Volume
}

func (s *ModifyTTSConfigRequest) SetAppKey(v string) *ModifyTTSConfigRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetInstanceId(v string) *ModifyTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetNlsServiceType(v string) *ModifyTTSConfigRequest {
	s.NlsServiceType = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetPitchRate(v string) *ModifyTTSConfigRequest {
	s.PitchRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetScriptId(v string) *ModifyTTSConfigRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetSpeechRate(v string) *ModifyTTSConfigRequest {
	s.SpeechRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVoice(v string) *ModifyTTSConfigRequest {
	s.Voice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVolume(v string) *ModifyTTSConfigRequest {
	s.Volume = &v
	return s
}

func (s *ModifyTTSConfigRequest) Validate() error {
	return dara.Validate(s)
}
