// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileUrl(v string) *ImportFieldRequest
	GetFileUrl() *string
	SetLang(v string) *ImportFieldRequest
	GetLang() *string
	SetRegId(v string) *ImportFieldRequest
	GetRegId() *string
}

type ImportFieldRequest struct {
	// Attachment download URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// filed/data/text.csv
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s ImportFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportFieldRequest) GoString() string {
	return s.String()
}

func (s *ImportFieldRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *ImportFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *ImportFieldRequest) SetFileUrl(v string) *ImportFieldRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportFieldRequest) SetLang(v string) *ImportFieldRequest {
	s.Lang = &v
	return s
}

func (s *ImportFieldRequest) SetRegId(v string) *ImportFieldRequest {
	s.RegId = &v
	return s
}

func (s *ImportFieldRequest) Validate() error {
	return dara.Validate(s)
}
