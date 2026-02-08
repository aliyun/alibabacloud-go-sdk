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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// ID of the recording to delete
	//
	// This parameter is required.
	//
	// example:
	//
	// df8216aa-d8f6-4501-864f-f8334d946561
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
