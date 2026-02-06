// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptWaveformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteScriptWaveformRequest
	GetInstanceId() *string
	SetScriptId(v string) *DeleteScriptWaveformRequest
	GetScriptId() *string
	SetScriptWaveformId(v string) *DeleteScriptWaveformRequest
	GetScriptWaveformId() *string
}

type DeleteScriptWaveformRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	ScriptWaveformId *string `json:"ScriptWaveformId,omitempty" xml:"ScriptWaveformId,omitempty"`
}

func (s DeleteScriptWaveformRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptWaveformRequest) GoString() string {
	return s.String()
}

func (s *DeleteScriptWaveformRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteScriptWaveformRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteScriptWaveformRequest) GetScriptWaveformId() *string {
	return s.ScriptWaveformId
}

func (s *DeleteScriptWaveformRequest) SetInstanceId(v string) *DeleteScriptWaveformRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScriptWaveformRequest) SetScriptId(v string) *DeleteScriptWaveformRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteScriptWaveformRequest) SetScriptWaveformId(v string) *DeleteScriptWaveformRequest {
	s.ScriptWaveformId = &v
	return s
}

func (s *DeleteScriptWaveformRequest) Validate() error {
	return dara.Validate(s)
}
