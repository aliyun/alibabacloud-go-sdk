// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVoiceAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetInstanceId() *string
	SetNlsEngine(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetNlsEngine() *string
	SetProfileShrink(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type UpdateVoiceAccessProfileShrinkRequest struct {
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
	// 目前支持IFLYTEK、VOLC
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// 配置
	ProfileShrink *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdateVoiceAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileShrinkRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateVoiceAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateVoiceAccessProfileShrinkRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *UpdateVoiceAccessProfileShrinkRequest) GetProfileShrink() *string {
	return s.ProfileShrink
}

func (s *UpdateVoiceAccessProfileShrinkRequest) SetAccessProfileId(v string) *UpdateVoiceAccessProfileShrinkRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateVoiceAccessProfileShrinkRequest) SetInstanceId(v string) *UpdateVoiceAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateVoiceAccessProfileShrinkRequest) SetNlsEngine(v string) *UpdateVoiceAccessProfileShrinkRequest {
	s.NlsEngine = &v
	return s
}

func (s *UpdateVoiceAccessProfileShrinkRequest) SetProfileShrink(v string) *UpdateVoiceAccessProfileShrinkRequest {
	s.ProfileShrink = &v
	return s
}

func (s *UpdateVoiceAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
