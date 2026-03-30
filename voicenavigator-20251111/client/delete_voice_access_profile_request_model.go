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
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
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
