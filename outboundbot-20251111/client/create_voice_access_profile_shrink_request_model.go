// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateVoiceAccessProfileShrinkRequest
	GetInstanceId() *string
	SetNlsEngine(v string) *CreateVoiceAccessProfileShrinkRequest
	GetNlsEngine() *string
	SetProfileShrink(v string) *CreateVoiceAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type CreateVoiceAccessProfileShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The voice service provider.
	//
	// BAILIAN: Bailian.
	//
	// VOLC: Doubao.
	//
	// IFLYTEK: iFLYTEK.
	//
	// TENCENT: Tencent.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The provider configuration information.
	ProfileShrink *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s CreateVoiceAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVoiceAccessProfileShrinkRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateVoiceAccessProfileShrinkRequest) GetProfileShrink() *string {
	return s.ProfileShrink
}

func (s *CreateVoiceAccessProfileShrinkRequest) SetInstanceId(v string) *CreateVoiceAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVoiceAccessProfileShrinkRequest) SetNlsEngine(v string) *CreateVoiceAccessProfileShrinkRequest {
	s.NlsEngine = &v
	return s
}

func (s *CreateVoiceAccessProfileShrinkRequest) SetProfileShrink(v string) *CreateVoiceAccessProfileShrinkRequest {
	s.ProfileShrink = &v
	return s
}

func (s *CreateVoiceAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
