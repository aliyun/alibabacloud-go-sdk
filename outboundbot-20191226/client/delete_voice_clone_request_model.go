// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteVoiceCloneRequest
	GetInstanceId() *string
	SetVoiceId(v string) *DeleteVoiceCloneRequest
	GetVoiceId() *string
}

type DeleteVoiceCloneRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the cloned voice to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// cosyvoice-v2-067dc4b2e9d34dbeb5de20b5b093d1e3
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s DeleteVoiceCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceCloneRequest) GoString() string {
	return s.String()
}

func (s *DeleteVoiceCloneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVoiceCloneRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *DeleteVoiceCloneRequest) SetInstanceId(v string) *DeleteVoiceCloneRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVoiceCloneRequest) SetVoiceId(v string) *DeleteVoiceCloneRequest {
	s.VoiceId = &v
	return s
}

func (s *DeleteVoiceCloneRequest) Validate() error {
	return dara.Validate(s)
}
