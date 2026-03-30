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
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VOLC
	NlsEngine     *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
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
