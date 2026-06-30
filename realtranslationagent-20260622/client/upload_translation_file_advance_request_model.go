// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadTranslationFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *UploadTranslationFileAdvanceRequest
	GetAPIKey() *string
	SetFileObject(v io.Reader) *UploadTranslationFileAdvanceRequest
	GetFileObject() io.Reader
	SetFileName(v string) *UploadTranslationFileAdvanceRequest
	GetFileName() *string
}

type UploadTranslationFileAdvanceRequest struct {
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// This parameter is required.
	FileObject io.Reader `json:"File,omitempty" xml:"File,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s UploadTranslationFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadTranslationFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadTranslationFileAdvanceRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *UploadTranslationFileAdvanceRequest) GetFileObject() io.Reader {
	return s.FileObject
}

func (s *UploadTranslationFileAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadTranslationFileAdvanceRequest) SetAPIKey(v string) *UploadTranslationFileAdvanceRequest {
	s.APIKey = &v
	return s
}

func (s *UploadTranslationFileAdvanceRequest) SetFileObject(v io.Reader) *UploadTranslationFileAdvanceRequest {
	s.FileObject = v
	return s
}

func (s *UploadTranslationFileAdvanceRequest) SetFileName(v string) *UploadTranslationFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadTranslationFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
