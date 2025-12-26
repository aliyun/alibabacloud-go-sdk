// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTTSVoiceByIdCustomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVoiceId(v string) *GetTTSVoiceByIdCustomRequest
	GetVoiceId() *string
}

type GetTTSVoiceByIdCustomRequest struct {
	// example:
	//
	// M1ScGtY****PBFEJHdUV1thQ
	VoiceId *string `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
}

func (s GetTTSVoiceByIdCustomRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTTSVoiceByIdCustomRequest) GoString() string {
	return s.String()
}

func (s *GetTTSVoiceByIdCustomRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *GetTTSVoiceByIdCustomRequest) SetVoiceId(v string) *GetTTSVoiceByIdCustomRequest {
	s.VoiceId = &v
	return s
}

func (s *GetTTSVoiceByIdCustomRequest) Validate() error {
	return dara.Validate(s)
}
