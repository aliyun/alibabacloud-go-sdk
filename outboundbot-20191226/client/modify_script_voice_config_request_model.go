// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptVoiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyScriptVoiceConfigRequest
	GetInstanceId() *string
	SetScriptId(v string) *ModifyScriptVoiceConfigRequest
	GetScriptId() *string
	SetScriptVoiceConfigId(v string) *ModifyScriptVoiceConfigRequest
	GetScriptVoiceConfigId() *string
	SetScriptWaveformRelation(v string) *ModifyScriptVoiceConfigRequest
	GetScriptWaveformRelation() *string
	SetType(v string) *ModifyScriptVoiceConfigRequest
	GetType() *string
}

type ModifyScriptVoiceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1d7a26e0-628b-4c3c-9918-7f2e23273f54
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e15cc646-50e5-4bc0-87ec-e4f2d1063b90
	ScriptVoiceConfigId    *string `json:"ScriptVoiceConfigId,omitempty" xml:"ScriptVoiceConfigId,omitempty"`
	ScriptWaveformRelation *string `json:"ScriptWaveformRelation,omitempty" xml:"ScriptWaveformRelation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WAVEFORM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyScriptVoiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptVoiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyScriptVoiceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScriptVoiceConfigRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyScriptVoiceConfigRequest) GetScriptVoiceConfigId() *string {
	return s.ScriptVoiceConfigId
}

func (s *ModifyScriptVoiceConfigRequest) GetScriptWaveformRelation() *string {
	return s.ScriptWaveformRelation
}

func (s *ModifyScriptVoiceConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyScriptVoiceConfigRequest) SetInstanceId(v string) *ModifyScriptVoiceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyScriptVoiceConfigRequest) SetScriptId(v string) *ModifyScriptVoiceConfigRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyScriptVoiceConfigRequest) SetScriptVoiceConfigId(v string) *ModifyScriptVoiceConfigRequest {
	s.ScriptVoiceConfigId = &v
	return s
}

func (s *ModifyScriptVoiceConfigRequest) SetScriptWaveformRelation(v string) *ModifyScriptVoiceConfigRequest {
	s.ScriptWaveformRelation = &v
	return s
}

func (s *ModifyScriptVoiceConfigRequest) SetType(v string) *ModifyScriptVoiceConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyScriptVoiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
