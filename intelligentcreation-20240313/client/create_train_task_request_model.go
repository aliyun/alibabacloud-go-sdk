// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunMainId(v string) *CreateTrainTaskRequest
	GetAliyunMainId() *string
	SetResSpecType(v string) *CreateTrainTaskRequest
	GetResSpecType() *string
	SetTaskType(v string) *CreateTrainTaskRequest
	GetTaskType() *string
	SetUseScene(v string) *CreateTrainTaskRequest
	GetUseScene() *string
	SetVoiceGender(v string) *CreateTrainTaskRequest
	GetVoiceGender() *string
	SetVoiceLanguage(v string) *CreateTrainTaskRequest
	GetVoiceLanguage() *string
	SetVoiceName(v string) *CreateTrainTaskRequest
	GetVoiceName() *string
	SetVoicePath(v string) *CreateTrainTaskRequest
	GetVoicePath() *string
}

type CreateTrainTaskRequest struct {
	// example:
	//
	// 13168123111
	AliyunMainId *string `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	// example:
	//
	// BASIC_MODEL
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// CopyAnchorAndVoice
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// realTimeInteractivity
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	// example:
	//
	// M
	VoiceGender *string `json:"voiceGender,omitempty" xml:"voiceGender,omitempty"`
	// example:
	//
	// zh
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceName     *string `json:"voiceName,omitempty" xml:"voiceName,omitempty"`
	// example:
	//
	// https://yic-pre/video/test-0513.mp3
	VoicePath *string `json:"voicePath,omitempty" xml:"voicePath,omitempty"`
}

func (s CreateTrainTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskRequest) GetAliyunMainId() *string {
	return s.AliyunMainId
}

func (s *CreateTrainTaskRequest) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *CreateTrainTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateTrainTaskRequest) GetUseScene() *string {
	return s.UseScene
}

func (s *CreateTrainTaskRequest) GetVoiceGender() *string {
	return s.VoiceGender
}

func (s *CreateTrainTaskRequest) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *CreateTrainTaskRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *CreateTrainTaskRequest) GetVoicePath() *string {
	return s.VoicePath
}

func (s *CreateTrainTaskRequest) SetAliyunMainId(v string) *CreateTrainTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *CreateTrainTaskRequest) SetResSpecType(v string) *CreateTrainTaskRequest {
	s.ResSpecType = &v
	return s
}

func (s *CreateTrainTaskRequest) SetTaskType(v string) *CreateTrainTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTrainTaskRequest) SetUseScene(v string) *CreateTrainTaskRequest {
	s.UseScene = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceGender(v string) *CreateTrainTaskRequest {
	s.VoiceGender = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceLanguage(v string) *CreateTrainTaskRequest {
	s.VoiceLanguage = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceName(v string) *CreateTrainTaskRequest {
	s.VoiceName = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoicePath(v string) *CreateTrainTaskRequest {
	s.VoicePath = &v
	return s
}

func (s *CreateTrainTaskRequest) Validate() error {
	return dara.Validate(s)
}
