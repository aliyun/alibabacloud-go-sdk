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
	SetBusinessUnitId(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetBusinessUnitId() *string
	SetNlsEngine(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetNlsEngine() *string
	SetProfileShrink(v string) *UpdateVoiceAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type UpdateVoiceAccessProfileShrinkRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// llm-xdne77rxe14ziszr
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine     *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
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

func (s *UpdateVoiceAccessProfileShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
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

func (s *UpdateVoiceAccessProfileShrinkRequest) SetBusinessUnitId(v string) *UpdateVoiceAccessProfileShrinkRequest {
	s.BusinessUnitId = &v
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
