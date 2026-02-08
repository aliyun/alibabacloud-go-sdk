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
	// Recording file ID
	//
	// > Obtain this parameter from the Folder parameter in the GetJobDataUploadParams API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6f91885fa24b4c408d8f4eb392fd8ae6
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Recording file name
	//
	// This parameter is required.
	//
	// example:
	//
	// faaf8508-9542-4ac4-84a2-0ddcbb5f79a6 (2).json
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2bfa5ae4-7185-4227-a3b8-328f26f11be1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script text to be broadcast in the recording
	//
	// This parameter is required.
	//
	// example:
	//
	// 您好，您昨晚在李佳琦直播间下单的娜薇诗椰子水二代精华2支装还未付款，超时未支付会自动关闭订单，活动错过不再有奥。
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
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
