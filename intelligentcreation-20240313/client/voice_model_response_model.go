// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetResourceTypeDesc(v string) *VoiceModelResponse
	GetResourceTypeDesc() *string
	SetTtsVersion(v int32) *VoiceModelResponse
	GetTtsVersion() *int32
	SetUseScene(v string) *VoiceModelResponse
	GetUseScene() *string
	SetVoiceDesc(v string) *VoiceModelResponse
	GetVoiceDesc() *string
	SetVoiceGender(v string) *VoiceModelResponse
	GetVoiceGender() *string
	SetVoiceId(v int64) *VoiceModelResponse
	GetVoiceId() *int64
	SetVoiceLanguage(v string) *VoiceModelResponse
	GetVoiceLanguage() *string
	SetVoiceModel(v string) *VoiceModelResponse
	GetVoiceModel() *string
	SetVoiceName(v string) *VoiceModelResponse
	GetVoiceName() *string
	SetVoiceType(v string) *VoiceModelResponse
	GetVoiceType() *string
	SetVoiceUrl(v string) *VoiceModelResponse
	GetVoiceUrl() *string
}

type VoiceModelResponse struct {
	ResourceTypeDesc *string `json:"resourceTypeDesc,omitempty" xml:"resourceTypeDesc,omitempty"`
	TtsVersion       *int32  `json:"ttsVersion,omitempty" xml:"ttsVersion,omitempty"`
	UseScene         *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceDesc        *string `json:"voiceDesc,omitempty" xml:"voiceDesc,omitempty"`
	VoiceGender      *string `json:"voiceGender,omitempty" xml:"voiceGender,omitempty"`
	VoiceId          *int64  `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
	VoiceLanguage    *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceModel       *string `json:"voiceModel,omitempty" xml:"voiceModel,omitempty"`
	VoiceName        *string `json:"voiceName,omitempty" xml:"voiceName,omitempty"`
	VoiceType        *string `json:"voiceType,omitempty" xml:"voiceType,omitempty"`
	VoiceUrl         *string `json:"voiceUrl,omitempty" xml:"voiceUrl,omitempty"`
}

func (s VoiceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s VoiceModelResponse) GoString() string {
	return s.String()
}

func (s *VoiceModelResponse) GetResourceTypeDesc() *string {
	return s.ResourceTypeDesc
}

func (s *VoiceModelResponse) GetTtsVersion() *int32 {
	return s.TtsVersion
}

func (s *VoiceModelResponse) GetUseScene() *string {
	return s.UseScene
}

func (s *VoiceModelResponse) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *VoiceModelResponse) GetVoiceGender() *string {
	return s.VoiceGender
}

func (s *VoiceModelResponse) GetVoiceId() *int64 {
	return s.VoiceId
}

func (s *VoiceModelResponse) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *VoiceModelResponse) GetVoiceModel() *string {
	return s.VoiceModel
}

func (s *VoiceModelResponse) GetVoiceName() *string {
	return s.VoiceName
}

func (s *VoiceModelResponse) GetVoiceType() *string {
	return s.VoiceType
}

func (s *VoiceModelResponse) GetVoiceUrl() *string {
	return s.VoiceUrl
}

func (s *VoiceModelResponse) SetResourceTypeDesc(v string) *VoiceModelResponse {
	s.ResourceTypeDesc = &v
	return s
}

func (s *VoiceModelResponse) SetTtsVersion(v int32) *VoiceModelResponse {
	s.TtsVersion = &v
	return s
}

func (s *VoiceModelResponse) SetUseScene(v string) *VoiceModelResponse {
	s.UseScene = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceDesc(v string) *VoiceModelResponse {
	s.VoiceDesc = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceGender(v string) *VoiceModelResponse {
	s.VoiceGender = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceId(v int64) *VoiceModelResponse {
	s.VoiceId = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceLanguage(v string) *VoiceModelResponse {
	s.VoiceLanguage = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceModel(v string) *VoiceModelResponse {
	s.VoiceModel = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceName(v string) *VoiceModelResponse {
	s.VoiceName = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceType(v string) *VoiceModelResponse {
	s.VoiceType = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceUrl(v string) *VoiceModelResponse {
	s.VoiceUrl = &v
	return s
}

func (s *VoiceModelResponse) Validate() error {
	return dara.Validate(s)
}
