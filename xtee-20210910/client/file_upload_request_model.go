// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *FileUploadRequest
	GetFileName() *string
	SetFileUrl(v string) *FileUploadRequest
	GetFileUrl() *string
	SetLang(v string) *FileUploadRequest
	GetLang() *string
	SetTab(v string) *FileUploadRequest
	GetTab() *string
}

type FileUploadRequest struct {
	// File name.
	//
	// example:
	//
	// P4911_2707.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File URL
	//
	// example:
	//
	// https://res-v1.cupl-fdfs.com/direct/79886bdc-9855-4ff4-aa34-eb5b21cd43a7
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Scenario.
	//
	// example:
	//
	// FNNCIEA
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s FileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s FileUploadRequest) GoString() string {
	return s.String()
}

func (s *FileUploadRequest) GetFileName() *string {
	return s.FileName
}

func (s *FileUploadRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *FileUploadRequest) GetLang() *string {
	return s.Lang
}

func (s *FileUploadRequest) GetTab() *string {
	return s.Tab
}

func (s *FileUploadRequest) SetFileName(v string) *FileUploadRequest {
	s.FileName = &v
	return s
}

func (s *FileUploadRequest) SetFileUrl(v string) *FileUploadRequest {
	s.FileUrl = &v
	return s
}

func (s *FileUploadRequest) SetLang(v string) *FileUploadRequest {
	s.Lang = &v
	return s
}

func (s *FileUploadRequest) SetTab(v string) *FileUploadRequest {
	s.Tab = &v
	return s
}

func (s *FileUploadRequest) Validate() error {
	return dara.Validate(s)
}
