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
	// example:
	//
	// 8ee1160a-6999-478f-8df6-f33ef21f27d5
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
