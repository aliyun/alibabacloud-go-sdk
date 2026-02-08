// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTTSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyTTSConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyTTSConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyTTSConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyTTSConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyTTSConfigResponseBody
	GetSuccess() *bool
	SetTTSConfig(v *ModifyTTSConfigResponseBodyTTSConfig) *ModifyTTSConfigResponseBody
	GetTTSConfig() *ModifyTTSConfigResponseBodyTTSConfig
}

type ModifyTTSConfigResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API prompt message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// f765d3ee-ec03-4765-b235-6877501d99d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TTS Configuration Information
	TTSConfig *ModifyTTSConfigResponseBodyTTSConfig `json:"TTSConfig,omitempty" xml:"TTSConfig,omitempty" type:"Struct"`
}

func (s ModifyTTSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyTTSConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyTTSConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyTTSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTTSConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyTTSConfigResponseBody) GetTTSConfig() *ModifyTTSConfigResponseBodyTTSConfig {
	return s.TTSConfig
}

func (s *ModifyTTSConfigResponseBody) SetCode(v string) *ModifyTTSConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) SetHttpStatusCode(v int32) *ModifyTTSConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) SetMessage(v string) *ModifyTTSConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) SetRequestId(v string) *ModifyTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) SetSuccess(v bool) *ModifyTTSConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyTTSConfigResponseBody) SetTTSConfig(v *ModifyTTSConfigResponseBodyTTSConfig) *ModifyTTSConfigResponseBody {
	s.TTSConfig = v
	return s
}

func (s *ModifyTTSConfigResponseBody) Validate() error {
	if s.TTSConfig != nil {
		if err := s.TTSConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyTTSConfigResponseBodyTTSConfig struct {
	// Instance ID
	//
	// example:
	//
	// 291cfc6a-8703-4bdd-a99d-9cba32d5288a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 语调
	//
	// [-500,500]之间整数。默认值为0。
	//
	// 大于0表示升高音高。
	//
	// 小于0表示降低音高。
	//
	// example:
	//
	// 0
	PitchRate *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 947e0875-b5d4-4b33-b18c-7b2cf85bcb4f
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Speech rate
	//
	// example:
	//
	// 50
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Configuration ID
	//
	// example:
	//
	// 2a07b634-e15d-445f-bbcb-fc4ea2df7b87
	TTSConfigId *string `json:"TTSConfigId,omitempty" xml:"TTSConfigId,omitempty"`
	// Speech synthesis model
	//
	// example:
	//
	// xiaoyun
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// Volume
	//
	// example:
	//
	// 100
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyTTSConfigResponseBodyTTSConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigResponseBodyTTSConfig) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetPitchRate() *string {
	return s.PitchRate
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetSpeechRate() *string {
	return s.SpeechRate
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetTTSConfigId() *string {
	return s.TTSConfigId
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetVoice() *string {
	return s.Voice
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) GetVolume() *string {
	return s.Volume
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetInstanceId(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.InstanceId = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetPitchRate(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.PitchRate = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetScriptId(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.ScriptId = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetSpeechRate(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.SpeechRate = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetTTSConfigId(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.TTSConfigId = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetVoice(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.Voice = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) SetVolume(v string) *ModifyTTSConfigResponseBodyTTSConfig {
	s.Volume = &v
	return s
}

func (s *ModifyTTSConfigResponseBodyTTSConfig) Validate() error {
	return dara.Validate(s)
}
