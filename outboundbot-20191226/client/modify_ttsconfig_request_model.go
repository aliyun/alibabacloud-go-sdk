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
	// AppKey for your Intelligent Speech Interaction project. Required only when NlsServiceType is Authorized.
	//
	// example:
	//
	// 99****Aw
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 27244bae-e446-4811-bb1d-f8a07b525af0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service type.
	//
	// Managed: The default public Intelligent Speech Interaction service for Outbound Bot.
	//
	// Authorized: A private Intelligent Speech Interaction service that you have purchased. To grant authorization, navigate to Scenario Management > Edit > Call Service > Custom Service.
	//
	// example:
	//
	// Managed
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// Pitch. An integer between -500 and 500. Default is 0.
	//
	// A value greater than 0 raises pitch.
	//
	// A value less than 0 lowers pitch.
	//
	// example:
	//
	// 0
	PitchRate *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1f1a2ba0-b3e7-4ff9-baf1-6dc8aeac0791
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Speech rate. An integer between -500 and 500. Default is 0.
	//
	// A value greater than 0 increases speech speed.
	//
	// A value less than 0 decreases speech speed.
	//
	// example:
	//
	// 0
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Voice model, such as aixia, siyue, or xiaoyun
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// Volume. An integer between 0 and 100. Default is 50.
	//
	// A value greater than 50 increases volume.
	//
	// A value less than 50 decreases volume.
	//
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
