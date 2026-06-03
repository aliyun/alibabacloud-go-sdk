// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptWaveformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *CreateScriptWaveformRequest
	GetFileId() *string
	SetFileName(v string) *CreateScriptWaveformRequest
	GetFileName() *string
	SetInstanceId(v string) *CreateScriptWaveformRequest
	GetInstanceId() *string
	SetScriptContent(v string) *CreateScriptWaveformRequest
	GetScriptContent() *string
	SetScriptId(v string) *CreateScriptWaveformRequest
	GetScriptId() *string
}

type CreateScriptWaveformRequest struct {
	// This parameter is required.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// This parameter is required.
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateScriptWaveformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptWaveformRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptWaveformRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateScriptWaveformRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateScriptWaveformRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptWaveformRequest) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *CreateScriptWaveformRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateScriptWaveformRequest) SetFileId(v string) *CreateScriptWaveformRequest {
	s.FileId = &v
	return s
}

func (s *CreateScriptWaveformRequest) SetFileName(v string) *CreateScriptWaveformRequest {
	s.FileName = &v
	return s
}

func (s *CreateScriptWaveformRequest) SetInstanceId(v string) *CreateScriptWaveformRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptWaveformRequest) SetScriptContent(v string) *CreateScriptWaveformRequest {
	s.ScriptContent = &v
	return s
}

func (s *CreateScriptWaveformRequest) SetScriptId(v string) *CreateScriptWaveformRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateScriptWaveformRequest) Validate() error {
	return dara.Validate(s)
}
