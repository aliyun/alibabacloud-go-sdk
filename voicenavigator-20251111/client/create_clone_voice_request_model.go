// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloneVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *CreateCloneVoiceRequest
	GetFileKey() *string
	SetInstanceId(v string) *CreateCloneVoiceRequest
	GetInstanceId() *string
	SetModel(v string) *CreateCloneVoiceRequest
	GetModel() *string
}

type CreateCloneVoiceRequest struct {
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *CreateCloneVoiceRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *CreateCloneVoiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloneVoiceRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCloneVoiceRequest) SetFileKey(v string) *CreateCloneVoiceRequest {
	s.FileKey = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetInstanceId(v string) *CreateCloneVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloneVoiceRequest) SetModel(v string) *CreateCloneVoiceRequest {
	s.Model = &v
	return s
}

func (s *CreateCloneVoiceRequest) Validate() error {
	return dara.Validate(s)
}
