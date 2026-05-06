// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *UploadSampleFileRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadSampleFileRequest
	GetFileUrl() *string
	SetLang(v string) *UploadSampleFileRequest
	GetLang() *string
	SetRegId(v string) *UploadSampleFileRequest
	GetRegId() *string
	SetTab(v string) *UploadSampleFileRequest
	GetTab() *string
	SetType(v string) *UploadSampleFileRequest
	GetType() *string
}

type UploadSampleFileRequest struct {
	// example:
	//
	// icekredit_202312_23a_1764640368_6908.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// saf/cpoc/34cd7959590ef568086035b956210495/1761186218068_XN_test JR142_1023_204794.csv
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// INTERNET
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UploadSampleFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileRequest) GoString() string {
	return s.String()
}

func (s *UploadSampleFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadSampleFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadSampleFileRequest) GetLang() *string {
	return s.Lang
}

func (s *UploadSampleFileRequest) GetRegId() *string {
	return s.RegId
}

func (s *UploadSampleFileRequest) GetTab() *string {
	return s.Tab
}

func (s *UploadSampleFileRequest) GetType() *string {
	return s.Type
}

func (s *UploadSampleFileRequest) SetFileName(v string) *UploadSampleFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadSampleFileRequest) SetFileUrl(v string) *UploadSampleFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadSampleFileRequest) SetLang(v string) *UploadSampleFileRequest {
	s.Lang = &v
	return s
}

func (s *UploadSampleFileRequest) SetRegId(v string) *UploadSampleFileRequest {
	s.RegId = &v
	return s
}

func (s *UploadSampleFileRequest) SetTab(v string) *UploadSampleFileRequest {
	s.Tab = &v
	return s
}

func (s *UploadSampleFileRequest) SetType(v string) *UploadSampleFileRequest {
	s.Type = &v
	return s
}

func (s *UploadSampleFileRequest) Validate() error {
	return dara.Validate(s)
}
