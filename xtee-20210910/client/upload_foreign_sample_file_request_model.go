// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadForeignSampleFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFile(v string) *UploadForeignSampleFileRequest
	GetFile() *string
	SetLang(v string) *UploadForeignSampleFileRequest
	GetLang() *string
	SetRegId(v string) *UploadForeignSampleFileRequest
	GetRegId() *string
	SetTab(v string) *UploadForeignSampleFileRequest
	GetTab() *string
}

type UploadForeignSampleFileRequest struct {
	File *string `json:"File,omitempty" xml:"File,omitempty"`
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
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s UploadForeignSampleFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileRequest) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileRequest) GetFile() *string {
	return s.File
}

func (s *UploadForeignSampleFileRequest) GetLang() *string {
	return s.Lang
}

func (s *UploadForeignSampleFileRequest) GetRegId() *string {
	return s.RegId
}

func (s *UploadForeignSampleFileRequest) GetTab() *string {
	return s.Tab
}

func (s *UploadForeignSampleFileRequest) SetFile(v string) *UploadForeignSampleFileRequest {
	s.File = &v
	return s
}

func (s *UploadForeignSampleFileRequest) SetLang(v string) *UploadForeignSampleFileRequest {
	s.Lang = &v
	return s
}

func (s *UploadForeignSampleFileRequest) SetRegId(v string) *UploadForeignSampleFileRequest {
	s.RegId = &v
	return s
}

func (s *UploadForeignSampleFileRequest) SetTab(v string) *UploadForeignSampleFileRequest {
	s.Tab = &v
	return s
}

func (s *UploadForeignSampleFileRequest) Validate() error {
	return dara.Validate(s)
}
