// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadTranslationFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *UploadTranslationFileRequest
	GetAPIKey() *string
	SetFile(v string) *UploadTranslationFileRequest
	GetFile() *string
	SetFileName(v string) *UploadTranslationFileRequest
	GetFileName() *string
}

type UploadTranslationFileRequest struct {
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// This parameter is required.
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s UploadTranslationFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadTranslationFileRequest) GoString() string {
	return s.String()
}

func (s *UploadTranslationFileRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *UploadTranslationFileRequest) GetFile() *string {
	return s.File
}

func (s *UploadTranslationFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadTranslationFileRequest) SetAPIKey(v string) *UploadTranslationFileRequest {
	s.APIKey = &v
	return s
}

func (s *UploadTranslationFileRequest) SetFile(v string) *UploadTranslationFileRequest {
	s.File = &v
	return s
}

func (s *UploadTranslationFileRequest) SetFileName(v string) *UploadTranslationFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadTranslationFileRequest) Validate() error {
	return dara.Validate(s)
}
