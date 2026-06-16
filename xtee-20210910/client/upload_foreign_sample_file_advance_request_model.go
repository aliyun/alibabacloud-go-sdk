// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadForeignSampleFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileObject(v io.Reader) *UploadForeignSampleFileAdvanceRequest
	GetFileObject() io.Reader
	SetLang(v string) *UploadForeignSampleFileAdvanceRequest
	GetLang() *string
	SetRegId(v string) *UploadForeignSampleFileAdvanceRequest
	GetRegId() *string
	SetTab(v string) *UploadForeignSampleFileAdvanceRequest
	GetTab() *string
}

type UploadForeignSampleFileAdvanceRequest struct {
	// OSS path of the file.
	//
	// example:
	//
	// saf/cpoc/953c883cde33b2e21d722eb661d26375/1779172027996_自动回溯测试 2605191.csv
	FileObject io.Reader `json:"File,omitempty" xml:"File,omitempty"`
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

func (s UploadForeignSampleFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileAdvanceRequest) GetFileObject() io.Reader {
	return s.FileObject
}

func (s *UploadForeignSampleFileAdvanceRequest) GetLang() *string {
	return s.Lang
}

func (s *UploadForeignSampleFileAdvanceRequest) GetRegId() *string {
	return s.RegId
}

func (s *UploadForeignSampleFileAdvanceRequest) GetTab() *string {
	return s.Tab
}

func (s *UploadForeignSampleFileAdvanceRequest) SetFileObject(v io.Reader) *UploadForeignSampleFileAdvanceRequest {
	s.FileObject = v
	return s
}

func (s *UploadForeignSampleFileAdvanceRequest) SetLang(v string) *UploadForeignSampleFileAdvanceRequest {
	s.Lang = &v
	return s
}

func (s *UploadForeignSampleFileAdvanceRequest) SetRegId(v string) *UploadForeignSampleFileAdvanceRequest {
	s.RegId = &v
	return s
}

func (s *UploadForeignSampleFileAdvanceRequest) SetTab(v string) *UploadForeignSampleFileAdvanceRequest {
	s.Tab = &v
	return s
}

func (s *UploadForeignSampleFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
