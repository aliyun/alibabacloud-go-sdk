// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptVoiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeScriptVoiceConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeScriptVoiceConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeScriptVoiceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeScriptVoiceConfigResponseBody
	GetRequestId() *string
	SetScriptVoiceConfig(v *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) *DescribeScriptVoiceConfigResponseBody
	GetScriptVoiceConfig() *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig
	SetSuccess(v bool) *DescribeScriptVoiceConfigResponseBody
	GetSuccess() *bool
}

type DescribeScriptVoiceConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScriptVoiceConfig *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig `json:"ScriptVoiceConfig,omitempty" xml:"ScriptVoiceConfig,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeScriptVoiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptVoiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScriptVoiceConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeScriptVoiceConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeScriptVoiceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScriptVoiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScriptVoiceConfigResponseBody) GetScriptVoiceConfig() *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	return s.ScriptVoiceConfig
}

func (s *DescribeScriptVoiceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeScriptVoiceConfigResponseBody) SetCode(v string) *DescribeScriptVoiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) SetHttpStatusCode(v int32) *DescribeScriptVoiceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) SetMessage(v string) *DescribeScriptVoiceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) SetRequestId(v string) *DescribeScriptVoiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) SetScriptVoiceConfig(v *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) *DescribeScriptVoiceConfigResponseBody {
	s.ScriptVoiceConfig = v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) SetSuccess(v bool) *DescribeScriptVoiceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBody) Validate() error {
	if s.ScriptVoiceConfig != nil {
		if err := s.ScriptVoiceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig struct {
	// example:
	//
	// 291cfc6a-8703-4bdd-a99d-9cba32d5288a
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// example:
	//
	// 947e0875-b5d4-4b33-b18c-7b2cf85bcb4f
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 2a07b634-e15d-445f-bbcb-fc4ea2df7b87
	ScriptVoiceConfigId    *string `json:"ScriptVoiceConfigId,omitempty" xml:"ScriptVoiceConfigId,omitempty"`
	ScriptWaveformRelation *string `json:"ScriptWaveformRelation,omitempty" xml:"ScriptWaveformRelation,omitempty"`
	// example:
	//
	// DIALOGUE_FLOW
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// WAVEFORM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptVoiceConfigId() *string {
	return s.ScriptVoiceConfigId
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetScriptWaveformRelation() *string {
	return s.ScriptWaveformRelation
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetSource() *string {
	return s.Source
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) GetType() *string {
	return s.Type
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetInstanceId(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptContent(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptContent = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptId(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptId = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptVoiceConfigId(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptVoiceConfigId = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetScriptWaveformRelation(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.ScriptWaveformRelation = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetSource(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.Source = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) SetType(v string) *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig {
	s.Type = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponseBodyScriptVoiceConfig) Validate() error {
	return dara.Validate(s)
}
