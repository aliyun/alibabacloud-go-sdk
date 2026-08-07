// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVoiceAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *UpdateVoiceAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateVoiceAccessProfileRequest
	GetInstanceId() *string
	SetNlsEngine(v string) *UpdateVoiceAccessProfileRequest
	GetNlsEngine() *string
	SetProfile(v *UpdateVoiceAccessProfileRequestProfile) *UpdateVoiceAccessProfileRequest
	GetProfile() *UpdateVoiceAccessProfileRequestProfile
}

type UpdateVoiceAccessProfileRequest struct {
	// 接入配置ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 目前支持IFLYTEK、VOLC
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// 配置
	Profile *UpdateVoiceAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
}

func (s UpdateVoiceAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateVoiceAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateVoiceAccessProfileRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateVoiceAccessProfileRequest) GetProfile() *UpdateVoiceAccessProfileRequestProfile {
	return s.Profile
}

func (s *UpdateVoiceAccessProfileRequest) SetAccessProfileId(v string) *UpdateVoiceAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequest) SetInstanceId(v string) *UpdateVoiceAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequest) SetNlsEngine(v string) *UpdateVoiceAccessProfileRequest {
	s.NlsEngine = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequest) SetProfile(v *UpdateVoiceAccessProfileRequestProfile) *UpdateVoiceAccessProfileRequest {
	s.Profile = v
	return s
}

func (s *UpdateVoiceAccessProfileRequest) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVoiceAccessProfileRequestProfile struct {
	// example:
	//
	// ****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// a9872e2342952e248727798f642936c7
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// c0358c6e51c1013b446fdeb21a3a5d2e
	ApiSecret *string `json:"ApiSecret,omitempty" xml:"ApiSecret,omitempty"`
	// example:
	//
	// 9479688350
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// DW0yKRHQEe1nAd8c
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 暂无使用
	AsrAppKey *string `json:"AsrAppKey,omitempty" xml:"AsrAppKey,omitempty"`
	// example:
	//
	// sci_r3b3e62udqcujnkerrorqztnpu
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// example:
	//
	// y5MZfFdW6yBZgJdKonHZBA
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// example:
	//
	// 暂无使用
	TtsApiKey *string `json:"TtsApiKey,omitempty" xml:"TtsApiKey,omitempty"`
}

func (s UpdateVoiceAccessProfileRequestProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileRequestProfile) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetAccessKey() *string {
	return s.AccessKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetApiSecret() *string {
	return s.ApiSecret
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetAppId() *string {
	return s.AppId
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetAsrAppKey() *string {
	return s.AsrAppKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetSecretId() *string {
	return s.SecretId
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetSecretKey() *string {
	return s.SecretKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) GetTtsApiKey() *string {
	return s.TtsApiKey
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetAccessKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.AccessKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetApiKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.ApiKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetApiSecret(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.ApiSecret = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetAppId(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.AppId = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetAppKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.AppKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetAsrAppKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.AsrAppKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetSecretId(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.SecretId = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetSecretKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.SecretKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) SetTtsApiKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.TtsApiKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
