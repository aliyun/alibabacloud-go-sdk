// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloneVoiceId(v string) *UpdateCloneVoiceRequest
	GetCloneVoiceId() *string
	SetInstanceId(v string) *UpdateCloneVoiceRequest
	GetInstanceId() *string
	SetName(v string) *UpdateCloneVoiceRequest
	GetName() *string
}

type UpdateCloneVoiceRequest struct {
	// The UUID of the cloned voice.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cloned voice.
	//
	// example:
	//
	// TestClonedVoice
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloneVoiceRequest) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *UpdateCloneVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloneVoiceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCloneVoiceRequest) SetCloneVoiceId(v string) *UpdateCloneVoiceRequest {
	s.CloneVoiceId = &v
	return s
}

func (s *UpdateCloneVoiceRequest) SetInstanceId(v string) *UpdateCloneVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloneVoiceRequest) SetName(v string) *UpdateCloneVoiceRequest {
	s.Name = &v
	return s
}

func (s *UpdateCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
