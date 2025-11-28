// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTTSVoiceCustomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateTTSVoiceCustomRequest
	GetFileName() *string
	SetGender(v string) *CreateTTSVoiceCustomRequest
	GetGender() *string
	SetName(v string) *CreateTTSVoiceCustomRequest
	GetName() *string
	SetOssKey(v string) *CreateTTSVoiceCustomRequest
	GetOssKey() *string
}

type CreateTTSVoiceCustomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestTTSVoiceName.mp3
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// FEMALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestTTSVoiceName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// material/INPUT_TRAIN_AUDIO/Mt.CN2VNOPRC5QU2
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateTTSVoiceCustomRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTTSVoiceCustomRequest) GoString() string {
	return s.String()
}

func (s *CreateTTSVoiceCustomRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateTTSVoiceCustomRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateTTSVoiceCustomRequest) GetName() *string {
	return s.Name
}

func (s *CreateTTSVoiceCustomRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateTTSVoiceCustomRequest) SetFileName(v string) *CreateTTSVoiceCustomRequest {
	s.FileName = &v
	return s
}

func (s *CreateTTSVoiceCustomRequest) SetGender(v string) *CreateTTSVoiceCustomRequest {
	s.Gender = &v
	return s
}

func (s *CreateTTSVoiceCustomRequest) SetName(v string) *CreateTTSVoiceCustomRequest {
	s.Name = &v
	return s
}

func (s *CreateTTSVoiceCustomRequest) SetOssKey(v string) *CreateTTSVoiceCustomRequest {
	s.OssKey = &v
	return s
}

func (s *CreateTTSVoiceCustomRequest) Validate() error {
	return dara.Validate(s)
}
