// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFileName(v string) *CreateAudioFileRequest
	GetAudioFileName() *string
	SetInstanceId(v string) *CreateAudioFileRequest
	GetInstanceId() *string
	SetName(v string) *CreateAudioFileRequest
	GetName() *string
	SetOssFileKey(v string) *CreateAudioFileRequest
	GetOssFileKey() *string
	SetUsage(v string) *CreateAudioFileRequest
	GetUsage() *string
}

type CreateAudioFileRequest struct {
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
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test/test-file.wav
	OssFileKey *string `json:"OssFileKey,omitempty" xml:"OssFileKey,omitempty"`
	Usage      *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s CreateAudioFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioFileRequest) GoString() string {
	return s.String()
}

func (s *CreateAudioFileRequest) GetAudioFileName() *string {
	return s.AudioFileName
}

func (s *CreateAudioFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAudioFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateAudioFileRequest) GetOssFileKey() *string {
	return s.OssFileKey
}

func (s *CreateAudioFileRequest) GetUsage() *string {
	return s.Usage
}

func (s *CreateAudioFileRequest) SetAudioFileName(v string) *CreateAudioFileRequest {
	s.AudioFileName = &v
	return s
}

func (s *CreateAudioFileRequest) SetInstanceId(v string) *CreateAudioFileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAudioFileRequest) SetName(v string) *CreateAudioFileRequest {
	s.Name = &v
	return s
}

func (s *CreateAudioFileRequest) SetOssFileKey(v string) *CreateAudioFileRequest {
	s.OssFileKey = &v
	return s
}

func (s *CreateAudioFileRequest) SetUsage(v string) *CreateAudioFileRequest {
	s.Usage = &v
	return s
}

func (s *CreateAudioFileRequest) Validate() error {
	return dara.Validate(s)
}
