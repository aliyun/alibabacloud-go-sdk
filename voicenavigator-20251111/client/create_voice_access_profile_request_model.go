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
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VOLC
	NlsEngine *string                                 `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	Profile   *CreateVoiceAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
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

func (s *CreateVoiceAccessProfileRequestProfile) SetTtsApiKey(v string) *CreateVoiceAccessProfileRequestProfile {
	s.TtsApiKey = &v
	return s
}

func (s *CreateVoiceAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
