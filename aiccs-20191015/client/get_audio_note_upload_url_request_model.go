// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioNoteUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetAudioNoteUploadUrlRequest
	GetFileName() *string
	SetFileType(v string) *GetAudioNoteUploadUrlRequest
	GetFileType() *string
}

type GetAudioNoteUploadUrlRequest struct {
	// The name of the file to upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-audio
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// example:
	//
	// wav
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetAudioNoteUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAudioNoteUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAudioNoteUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetAudioNoteUploadUrlRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetAudioNoteUploadUrlRequest) SetFileName(v string) *GetAudioNoteUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetAudioNoteUploadUrlRequest) SetFileType(v string) *GetAudioNoteUploadUrlRequest {
	s.FileType = &v
	return s
}

func (s *GetAudioNoteUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
