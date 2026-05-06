// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *SearchSampleRequest
	GetKeyword() *string
	SetLang(v string) *SearchSampleRequest
	GetLang() *string
	SetRegId(v string) *SearchSampleRequest
	GetRegId() *string
	SetTab(v string) *SearchSampleRequest
	GetTab() *string
	SetType(v string) *SearchSampleRequest
	GetType() *string
	SetUploadTimeEnd(v string) *SearchSampleRequest
	GetUploadTimeEnd() *string
	SetUploadTimeStart(v string) *SearchSampleRequest
	GetUploadTimeStart() *string
}

type SearchSampleRequest struct {
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// example:
	//
	// 2023-03-09 14:45:26
	UploadTimeEnd *string `json:"UploadTimeEnd,omitempty" xml:"UploadTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-09 14:12:12
	UploadTimeStart *string `json:"UploadTimeStart,omitempty" xml:"UploadTimeStart,omitempty"`
}

func (s SearchSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchSampleRequest) GoString() string {
	return s.String()
}

func (s *SearchSampleRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *SearchSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *SearchSampleRequest) GetTab() *string {
	return s.Tab
}

func (s *SearchSampleRequest) GetType() *string {
	return s.Type
}

func (s *SearchSampleRequest) GetUploadTimeEnd() *string {
	return s.UploadTimeEnd
}

func (s *SearchSampleRequest) GetUploadTimeStart() *string {
	return s.UploadTimeStart
}

func (s *SearchSampleRequest) SetKeyword(v string) *SearchSampleRequest {
	s.Keyword = &v
	return s
}

func (s *SearchSampleRequest) SetLang(v string) *SearchSampleRequest {
	s.Lang = &v
	return s
}

func (s *SearchSampleRequest) SetRegId(v string) *SearchSampleRequest {
	s.RegId = &v
	return s
}

func (s *SearchSampleRequest) SetTab(v string) *SearchSampleRequest {
	s.Tab = &v
	return s
}

func (s *SearchSampleRequest) SetType(v string) *SearchSampleRequest {
	s.Type = &v
	return s
}

func (s *SearchSampleRequest) SetUploadTimeEnd(v string) *SearchSampleRequest {
	s.UploadTimeEnd = &v
	return s
}

func (s *SearchSampleRequest) SetUploadTimeStart(v string) *SearchSampleRequest {
	s.UploadTimeStart = &v
	return s
}

func (s *SearchSampleRequest) Validate() error {
	return dara.Validate(s)
}
