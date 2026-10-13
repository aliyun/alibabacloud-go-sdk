// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVoiceCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyVoiceCloneRequest
	GetInstanceId() *string
	SetStatus(v string) *ModifyVoiceCloneRequest
	GetStatus() *string
	SetVoiceId(v string) *ModifyVoiceCloneRequest
	GetVoiceId() *string
	SetVoiceName(v string) *ModifyVoiceCloneRequest
	GetVoiceName() *string
}

type ModifyVoiceCloneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cosyvoice-v2-067dc4b2e9d34dbeb5de20b5b093d1e3
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// example:
	//
	// 克隆音色1
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s ModifyVoiceCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVoiceCloneRequest) GoString() string {
	return s.String()
}

func (s *ModifyVoiceCloneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyVoiceCloneRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyVoiceCloneRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *ModifyVoiceCloneRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *ModifyVoiceCloneRequest) SetInstanceId(v string) *ModifyVoiceCloneRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyVoiceCloneRequest) SetStatus(v string) *ModifyVoiceCloneRequest {
	s.Status = &v
	return s
}

func (s *ModifyVoiceCloneRequest) SetVoiceId(v string) *ModifyVoiceCloneRequest {
	s.VoiceId = &v
	return s
}

func (s *ModifyVoiceCloneRequest) SetVoiceName(v string) *ModifyVoiceCloneRequest {
	s.VoiceName = &v
	return s
}

func (s *ModifyVoiceCloneRequest) Validate() error {
	return dara.Validate(s)
}
