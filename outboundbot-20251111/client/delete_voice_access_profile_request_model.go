// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *DeleteVoiceAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *DeleteVoiceAccessProfileRequest
	GetInstanceId() *string
}

type DeleteVoiceAccessProfileRequest struct {
	// 接入配置ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteVoiceAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteVoiceAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *DeleteVoiceAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVoiceAccessProfileRequest) SetAccessProfileId(v string) *DeleteVoiceAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *DeleteVoiceAccessProfileRequest) SetInstanceId(v string) *DeleteVoiceAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVoiceAccessProfileRequest) Validate() error {
	return dara.Validate(s)
}
