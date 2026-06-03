// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptWaveformsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryScriptWaveformsRequest
	GetInstanceId() *string
	SetScriptContent(v string) *QueryScriptWaveformsRequest
	GetScriptContent() *string
	SetScriptId(v string) *QueryScriptWaveformsRequest
	GetScriptId() *string
}

type QueryScriptWaveformsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// This parameter is required.
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s QueryScriptWaveformsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptWaveformsRequest) GoString() string {
	return s.String()
}

func (s *QueryScriptWaveformsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryScriptWaveformsRequest) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *QueryScriptWaveformsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *QueryScriptWaveformsRequest) SetInstanceId(v string) *QueryScriptWaveformsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryScriptWaveformsRequest) SetScriptContent(v string) *QueryScriptWaveformsRequest {
	s.ScriptContent = &v
	return s
}

func (s *QueryScriptWaveformsRequest) SetScriptId(v string) *QueryScriptWaveformsRequest {
	s.ScriptId = &v
	return s
}

func (s *QueryScriptWaveformsRequest) Validate() error {
	return dara.Validate(s)
}
