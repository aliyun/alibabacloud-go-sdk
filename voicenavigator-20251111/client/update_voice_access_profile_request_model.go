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
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VOLC
	NlsEngine *string                                 `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	Profile   *UpdateVoiceAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
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
	// HwRnTXgwnQOlsj68URDS5_VMm4Wtapq9
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// sk-12341e259b1049e8872b47981e545f78
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// c0358c6e51c1013b446fdeb21a3a1234
	ApiSecret *string `json:"ApiSecret,omitempty" xml:"ApiSecret,omitempty"`
	// example:
	//
	// 5b123bfb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2541370123
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AsrAppKey *string `json:"AsrAppKey,omitempty" xml:"AsrAppKey,omitempty"`
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

func (s *UpdateVoiceAccessProfileRequestProfile) SetTtsApiKey(v string) *UpdateVoiceAccessProfileRequestProfile {
	s.TtsApiKey = &v
	return s
}

func (s *UpdateVoiceAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
