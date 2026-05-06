// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateVoiceAccessProfileShrinkRequest
	GetBusinessUnitId() *string
	SetNlsEngine(v string) *CreateVoiceAccessProfileShrinkRequest
	GetNlsEngine() *string
	SetProfileShrink(v string) *CreateVoiceAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type CreateVoiceAccessProfileShrinkRequest struct {
	// example:
	//
	// llm-xdne77rxe14ziszr
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
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

func (s *CreateVoiceAccessProfileShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateVoiceAccessProfileShrinkRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *CreateVoiceAccessProfileShrinkRequest) GetProfileShrink() *string {
	return s.ProfileShrink
}

func (s *CreateVoiceAccessProfileShrinkRequest) SetBusinessUnitId(v string) *CreateVoiceAccessProfileShrinkRequest {
	s.BusinessUnitId = &v
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
