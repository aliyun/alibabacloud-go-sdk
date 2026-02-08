// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadScriptRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UploadScriptRecordingRequest
	GetContent() *string
	SetFileId(v string) *UploadScriptRecordingRequest
	GetFileId() *string
	SetFileName(v string) *UploadScriptRecordingRequest
	GetFileName() *string
	SetInstanceId(v string) *UploadScriptRecordingRequest
	GetInstanceId() *string
	SetScriptId(v string) *UploadScriptRecordingRequest
	GetScriptId() *string
}

type UploadScriptRecordingRequest struct {
	// Text content of the recording file
	//
	// This parameter is required.
	//
	// example:
	//
	// 您好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Key of the file in OSS
	//
	// This parameter is required.
	//
	// example:
	//
	// cab_script_recording/upload/88a56c18-3dc8-4338-9116-911deb169780/hello.wav
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// File name
	//
	// This parameter is required.
	//
	// example:
	//
	// hello.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c209abb3-6804-4a75-b2c7-dd55c8c61b6a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s UploadScriptRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadScriptRecordingRequest) GoString() string {
	return s.String()
}

func (s *UploadScriptRecordingRequest) GetContent() *string {
	return s.Content
}

func (s *UploadScriptRecordingRequest) GetFileId() *string {
	return s.FileId
}

func (s *UploadScriptRecordingRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadScriptRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UploadScriptRecordingRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *UploadScriptRecordingRequest) SetContent(v string) *UploadScriptRecordingRequest {
	s.Content = &v
	return s
}

func (s *UploadScriptRecordingRequest) SetFileId(v string) *UploadScriptRecordingRequest {
	s.FileId = &v
	return s
}

func (s *UploadScriptRecordingRequest) SetFileName(v string) *UploadScriptRecordingRequest {
	s.FileName = &v
	return s
}

func (s *UploadScriptRecordingRequest) SetInstanceId(v string) *UploadScriptRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *UploadScriptRecordingRequest) SetScriptId(v string) *UploadScriptRecordingRequest {
	s.ScriptId = &v
	return s
}

func (s *UploadScriptRecordingRequest) Validate() error {
	return dara.Validate(s)
}
