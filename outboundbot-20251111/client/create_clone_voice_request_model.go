// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *CreateCloneVoiceRequest
	GetFileKey() *string
	SetInstanceId(v string) *CreateCloneVoiceRequest
	GetInstanceId() *string
	SetModel(v string) *CreateCloneVoiceRequest
	GetModel() *string
}

type CreateCloneVoiceRequest struct {
	// 文件Key
	//
	// example:
	//
	// voice_clone/upload/d25ace5f-c8c6-45af-a5b1-8fd6b8595747/019FDB17-4901-17A9-99D6-27B77BC047C0_record.wav
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 模型名称
	//
	// example:
	//
	// CosyVoice
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s CreateCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *CreateCloneVoiceRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *CreateCloneVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloneVoiceRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCloneVoiceRequest) SetFileKey(v string) *CreateCloneVoiceRequest {
	s.FileKey = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetInstanceId(v string) *CreateCloneVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetModel(v string) *CreateCloneVoiceRequest {
	s.Model = &v
	return s
}

func (s *CreateCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
