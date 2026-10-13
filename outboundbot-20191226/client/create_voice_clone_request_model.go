// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *CreateVoiceCloneRequest
	GetFileId() *string
	SetInstanceId(v string) *CreateVoiceCloneRequest
	GetInstanceId() *string
	SetModel(v string) *CreateVoiceCloneRequest
	GetModel() *string
	SetVoiceName(v string) *CreateVoiceCloneRequest
	GetVoiceName() *string
}

type CreateVoiceCloneRequest struct {
	// The Folder field returned after a file is uploaded by calling the GetJobDataUploadParams operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// voiceclone/recordings/demo.mp3
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model used for voice cloning.
	//
	// example:
	//
	// cosyvoice-v3.5-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the cloned voice. The names of enabled voices within the same instance must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cloned voice 1
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s CreateVoiceCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceCloneRequest) GoString() string {
	return s.String()
}

func (s *CreateVoiceCloneRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateVoiceCloneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVoiceCloneRequest) GetModel() *string {
	return s.Model
}

func (s *CreateVoiceCloneRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *CreateVoiceCloneRequest) SetFileId(v string) *CreateVoiceCloneRequest {
	s.FileId = &v
	return s
}

func (s *CreateVoiceCloneRequest) SetInstanceId(v string) *CreateVoiceCloneRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVoiceCloneRequest) SetModel(v string) *CreateVoiceCloneRequest {
	s.Model = &v
	return s
}

func (s *CreateVoiceCloneRequest) SetVoiceName(v string) *CreateVoiceCloneRequest {
	s.VoiceName = &v
	return s
}

func (s *CreateVoiceCloneRequest) Validate() error {
	return dara.Validate(s)
}
