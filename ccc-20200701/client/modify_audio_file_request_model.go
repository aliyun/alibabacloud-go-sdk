// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAudioFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFileName(v string) *ModifyAudioFileRequest
	GetAudioFileName() *string
	SetAudioResourceId(v string) *ModifyAudioFileRequest
	GetAudioResourceId() *string
	SetInstanceId(v string) *ModifyAudioFileRequest
	GetInstanceId() *string
	SetName(v string) *ModifyAudioFileRequest
	GetName() *string
	SetOssFileKey(v string) *ModifyAudioFileRequest
	GetOssFileKey() *string
	SetUsage(v string) *ModifyAudioFileRequest
	GetUsage() *string
}

type ModifyAudioFileRequest struct {
	// Name of the audio file to be modified. You can specify new content for the audio file here.
	//
	// This parameter is required.
	//
	// example:
	//
	// new-test-file.wav
	AudioFileName *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty"`
	// Audio resource ID, which uniquely identifies an audio file.
	//
	// This parameter is required.
	//
	// example:
	//
	// acc300c4-75c9-41ba-ba5e-2a365c96c248
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Display name of the audio file. It must be 1 to 32 characters in length. The display name cannot be changed when modifying the audio file, so you must provide the original display name here.
	//
	// This parameter is required.
	//
	// example:
	//
	// 欢迎语
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS key of the audio file to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test/new-test-file.wav
	OssFileKey *string `json:"OssFileKey,omitempty" xml:"OssFileKey,omitempty"`
	// Usage of the audio file. The default value is General (used in scenarios such as IVR). Other valid values include HoldMusic (hold music during call waiting).
	//
	// example:
	//
	// General
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ModifyAudioFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAudioFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyAudioFileRequest) GetAudioFileName() *string {
	return s.AudioFileName
}

func (s *ModifyAudioFileRequest) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *ModifyAudioFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAudioFileRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAudioFileRequest) GetOssFileKey() *string {
	return s.OssFileKey
}

func (s *ModifyAudioFileRequest) GetUsage() *string {
	return s.Usage
}

func (s *ModifyAudioFileRequest) SetAudioFileName(v string) *ModifyAudioFileRequest {
	s.AudioFileName = &v
	return s
}

func (s *ModifyAudioFileRequest) SetAudioResourceId(v string) *ModifyAudioFileRequest {
	s.AudioResourceId = &v
	return s
}

func (s *ModifyAudioFileRequest) SetInstanceId(v string) *ModifyAudioFileRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAudioFileRequest) SetName(v string) *ModifyAudioFileRequest {
	s.Name = &v
	return s
}

func (s *ModifyAudioFileRequest) SetOssFileKey(v string) *ModifyAudioFileRequest {
	s.OssFileKey = &v
	return s
}

func (s *ModifyAudioFileRequest) SetUsage(v string) *ModifyAudioFileRequest {
	s.Usage = &v
	return s
}

func (s *ModifyAudioFileRequest) Validate() error {
	return dara.Validate(s)
}
