// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptVoiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyScriptVoiceConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyScriptVoiceConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyScriptVoiceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScriptVoiceConfigResponseBody
	GetRequestId() *string
	SetScriptVoiceConfig(v *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) *ModifyScriptVoiceConfigResponseBody
	GetScriptVoiceConfig() *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig
	SetSuccess(v bool) *ModifyScriptVoiceConfigResponseBody
	GetSuccess() *bool
}

type ModifyScriptVoiceConfigResponseBody struct {
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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Script voice configuration information
	ScriptVoiceConfig *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig `json:"ScriptVoiceConfig,omitempty" xml:"ScriptVoiceConfig,omitempty" type:"Struct"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScriptVoiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptVoiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScriptVoiceConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyScriptVoiceConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyScriptVoiceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScriptVoiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScriptVoiceConfigResponseBody) GetScriptVoiceConfig() *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	return s.ScriptVoiceConfig
}

func (s *ModifyScriptVoiceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScriptVoiceConfigResponseBody) SetCode(v string) *ModifyScriptVoiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) SetHttpStatusCode(v int32) *ModifyScriptVoiceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) SetMessage(v string) *ModifyScriptVoiceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) SetRequestId(v string) *ModifyScriptVoiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) SetScriptVoiceConfig(v *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) *ModifyScriptVoiceConfigResponseBody {
	s.ScriptVoiceConfig = v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) SetSuccess(v bool) *ModifyScriptVoiceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBody) Validate() error {
	if s.ScriptVoiceConfig != nil {
		if err := s.ScriptVoiceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig struct {
	// Instance ID
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Text corresponding to the audio
	//
	// example:
	//
	// 你好
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// Scene ID
	//
	// example:
	//
	// 1d7a26e0-628b-4c3c-9918-7f2e23273f54
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Scene audio configuration ID
	//
	// example:
	//
	// 4ddea690-6534-4c88-9cbd-0b5882ec64c0
	ScriptVoiceConfigId *string `json:"ScriptVoiceConfigId,omitempty" xml:"ScriptVoiceConfigId,omitempty"`
	// Relational data of the recording correspondence
	//
	// example:
	//
	// [{"ScriptContent":"你好","ScriptWaveformId":"6facc560-9f25-420f-bb0b-99f4299ad0d5"},{"ScriptContent":"吗","ScriptWaveformId":"76f48266-e253-4f44-ab4f-f64f7d38f1b4"}]
	ScriptWaveformRelation *string `json:"ScriptWaveformRelation,omitempty" xml:"ScriptWaveformRelation,omitempty"`
	// Script source
	//
	// example:
	//
	// DIALOGUE_FLOW
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Broadcast type
	//
	// example:
	//
	// WAVEFORM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GoString() string {
	return s.String()
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptVoiceConfigId() *string {
	return s.ScriptVoiceConfigId
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptWaveformRelation() *string {
	return s.ScriptWaveformRelation
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetSource() *string {
	return s.Source
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) GetType() *string {
	return s.Type
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetInstanceId(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.InstanceId = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptContent(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptContent = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptId(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptId = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptVoiceConfigId(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptVoiceConfigId = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptWaveformRelation(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptWaveformRelation = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetSource(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.Source = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) SetType(v string) *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.Type = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponseBodyScriptVoiceConfig) Validate() error {
	return dara.Validate(s)
}
