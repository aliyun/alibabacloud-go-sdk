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
	// OSS path of the file.
	//
	// example:
	//
	// saf/cpoc/953c883cde33b2e21d722eb661d26375/1779172027996_自动回溯测试 2605191.csv
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// Set the language type for requests and received messages. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// area encoding.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// scenario.
	//
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
