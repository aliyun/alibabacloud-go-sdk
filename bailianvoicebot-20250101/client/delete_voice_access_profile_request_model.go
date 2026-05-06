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
	SetBusinessUnitId(v string) *DeleteVoiceAccessProfileRequest
	GetBusinessUnitId() *string
}

type DeleteVoiceAccessProfileRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
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

func (s *DeleteVoiceAccessProfileRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *DeleteVoiceAccessProfileRequest) SetAccessProfileId(v string) *DeleteVoiceAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *DeleteVoiceAccessProfileRequest) SetBusinessUnitId(v string) *DeleteVoiceAccessProfileRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *DeleteVoiceAccessProfileRequest) Validate() error {
	return dara.Validate(s)
}
