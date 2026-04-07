// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *DeleteCloneVoiceRequest
	GetBusinessUnitId() *string
	SetCloneVoiceId(v string) *DeleteCloneVoiceRequest
	GetCloneVoiceId() *string
}

type DeleteCloneVoiceRequest struct {
	// example:
	//
	// llm-zzu528i29ecnprcl
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 8ee1160a-6999-478f-8df6-f33ef21f27d5
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
}

func (s DeleteCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloneVoiceRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *DeleteCloneVoiceRequest) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *DeleteCloneVoiceRequest) SetBusinessUnitId(v string) *DeleteCloneVoiceRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *DeleteCloneVoiceRequest) SetCloneVoiceId(v string) *DeleteCloneVoiceRequest {
	s.CloneVoiceId = &v
	return s
}

func (s *DeleteCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
