// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateVoiceAccessProfileRequest
	GetInstanceId() *string
	SetNlsEngine(v string) *CreateVoiceAccessProfileRequest
	GetNlsEngine() *string
	SetProfile(v *CreateVoiceAccessProfileRequestProfile) *CreateVoiceAccessProfileRequest
	GetProfile() *CreateVoiceAccessProfileRequestProfile
}

type CreateVoiceAccessProfileRequest struct {
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
	Profile *CreateVoiceAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
}

func (s CreateVoiceAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVoiceAccessProfileRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateVoiceAccessProfileRequest) GetProfile() *CreateVoiceAccessProfileRequestProfile {
	return s.Profile
}

func (s *CreateVoiceAccessProfileRequest) SetInstanceId(v string) *CreateVoiceAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVoiceAccessProfileRequest) SetNlsEngine(v string) *CreateVoiceAccessProfileRequest {
	s.NlsEngine = &v
	return s
}

func (s *CreateVoiceAccessProfileRequest) SetProfile(v *CreateVoiceAccessProfileRequestProfile) *CreateVoiceAccessProfileRequest {
	s.Profile = v
	return s
}

func (s *CreateVoiceAccessProfileRequest) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVoiceAccessProfileRequestProfile struct {
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

func (s CreateVoiceAccessProfileRequestProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileRequestProfile) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileRequestProfile) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateVoiceAccessProfileRequestProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateVoiceAccessProfileRequestProfile) GetApiSecret() *string {
	return s.ApiSecret
}

func (s *CreateVoiceAccessProfileRequestProfile) GetAppId() *string {
	return s.AppId
}

func (s *CreateVoiceAccessProfileRequestProfile) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateVoiceAccessProfileRequestProfile) GetAsrAppKey() *string {
	return s.AsrAppKey
}

func (s *CreateVoiceAccessProfileRequestProfile) GetSecretId() *string {
	return s.SecretId
}

func (s *CreateVoiceAccessProfileRequestProfile) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateVoiceAccessProfileRequestProfile) GetTtsApiKey() *string {
	return s.TtsApiKey
}

func (s *CreateVoiceAccessProfileRequestProfile) SetAccessKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.AccessKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetApiKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.ApiKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetApiSecret(v string) *CreateVoiceAccessProfileRequestProfile {
	s.ApiSecret = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetAppId(v string) *CreateVoiceAccessProfileRequestProfile {
	s.AppId = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetAppKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.AppKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetAsrAppKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.AsrAppKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetSecretId(v string) *CreateVoiceAccessProfileRequestProfile {
	s.SecretId = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetSecretKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.SecretKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) SetTtsApiKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.TtsApiKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
