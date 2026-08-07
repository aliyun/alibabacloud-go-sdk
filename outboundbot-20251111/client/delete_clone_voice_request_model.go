// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloneVoiceId(v string) *DeleteCloneVoiceRequest
	GetCloneVoiceId() *string
	SetInstanceId(v string) *DeleteCloneVoiceRequest
	GetInstanceId() *string
}

type DeleteCloneVoiceRequest struct {
	// 克隆音色ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloneVoiceRequest) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *DeleteCloneVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloneVoiceRequest) SetCloneVoiceId(v string) *DeleteCloneVoiceRequest {
	s.CloneVoiceId = &v
	return s
}

func (s *DeleteCloneVoiceRequest) SetInstanceId(v string) *DeleteCloneVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
