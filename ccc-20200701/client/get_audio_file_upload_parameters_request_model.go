// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileUploadParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFileName(v string) *GetAudioFileUploadParametersRequest
	GetAudioFileName() *string
	SetInstanceId(v string) *GetAudioFileUploadParametersRequest
	GetInstanceId() *string
}

type GetAudioFileUploadParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-file.wav
	AudioFileName *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAudioFileUploadParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileUploadParametersRequest) GoString() string {
	return s.String()
}

func (s *GetAudioFileUploadParametersRequest) GetAudioFileName() *string {
	return s.AudioFileName
}

func (s *GetAudioFileUploadParametersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAudioFileUploadParametersRequest) SetAudioFileName(v string) *GetAudioFileUploadParametersRequest {
	s.AudioFileName = &v
	return s
}

func (s *GetAudioFileUploadParametersRequest) SetInstanceId(v string) *GetAudioFileUploadParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAudioFileUploadParametersRequest) Validate() error {
	return dara.Validate(s)
}
