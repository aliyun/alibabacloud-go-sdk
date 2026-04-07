// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateCloneVoiceRequest
	GetBusinessUnitId() *string
	SetFileKey(v string) *CreateCloneVoiceRequest
	GetFileKey() *string
	SetModel(v string) *CreateCloneVoiceRequest
	GetModel() *string
}

type CreateCloneVoiceRequest struct {
	// example:
	//
	// llm-xdne77rxe14ziszr
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	FileKey        *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// CosyVoice
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s CreateCloneVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloneVoiceRequest) GoString() string {
	return s.String()
}

func (s *CreateCloneVoiceRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateCloneVoiceRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *CreateCloneVoiceRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCloneVoiceRequest) SetBusinessUnitId(v string) *CreateCloneVoiceRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetFileKey(v string) *CreateCloneVoiceRequest {
	s.FileKey = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetModel(v string) *CreateCloneVoiceRequest {
	s.Model = &v
	return s
}

func (s *CreateCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
