// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartVoiceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVoiceType(v string) *ListSmartVoiceGroupsRequest
	GetVoiceType() *string
}

type ListSmartVoiceGroupsRequest struct {
	// The voice type. If you do not specify this parameter, the operation returns voices from the small model by default. Set this parameter to `CosyVoice` to retrieve voices from the Cosyvoice large model.
	//
	// example:
	//
	// CosyVoice
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s ListSmartVoiceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmartVoiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSmartVoiceGroupsRequest) GetVoiceType() *string {
	return s.VoiceType
}

func (s *ListSmartVoiceGroupsRequest) SetVoiceType(v string) *ListSmartVoiceGroupsRequest {
	s.VoiceType = &v
	return s
}

func (s *ListSmartVoiceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
