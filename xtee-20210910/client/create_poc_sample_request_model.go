// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreatePocSampleRequest
	GetFileName() *string
	SetFileUrl(v string) *CreatePocSampleRequest
	GetFileUrl() *string
	SetLang(v string) *CreatePocSampleRequest
	GetLang() *string
	SetRegId(v string) *CreatePocSampleRequest
	GetRegId() *string
	SetRemark(v string) *CreatePocSampleRequest
	GetRemark() *string
	SetSampleName(v string) *CreatePocSampleRequest
	GetSampleName() *string
	SetTab(v string) *CreatePocSampleRequest
	GetTab() *string
	SetType(v string) *CreatePocSampleRequest
	GetType() *string
}

type CreatePocSampleRequest struct {
	// example:
	//
	// P4911_2707.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	// cs-pub
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// SampleNameTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// INTERNET
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePocSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePocSampleRequest) GoString() string {
	return s.String()
}

func (s *CreatePocSampleRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreatePocSampleRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreatePocSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePocSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreatePocSampleRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreatePocSampleRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *CreatePocSampleRequest) GetTab() *string {
	return s.Tab
}

func (s *CreatePocSampleRequest) GetType() *string {
	return s.Type
}

func (s *CreatePocSampleRequest) SetFileName(v string) *CreatePocSampleRequest {
	s.FileName = &v
	return s
}

func (s *CreatePocSampleRequest) SetFileUrl(v string) *CreatePocSampleRequest {
	s.FileUrl = &v
	return s
}

func (s *CreatePocSampleRequest) SetLang(v string) *CreatePocSampleRequest {
	s.Lang = &v
	return s
}

func (s *CreatePocSampleRequest) SetRegId(v string) *CreatePocSampleRequest {
	s.RegId = &v
	return s
}

func (s *CreatePocSampleRequest) SetRemark(v string) *CreatePocSampleRequest {
	s.Remark = &v
	return s
}

func (s *CreatePocSampleRequest) SetSampleName(v string) *CreatePocSampleRequest {
	s.SampleName = &v
	return s
}

func (s *CreatePocSampleRequest) SetTab(v string) *CreatePocSampleRequest {
	s.Tab = &v
	return s
}

func (s *CreatePocSampleRequest) SetType(v string) *CreatePocSampleRequest {
	s.Type = &v
	return s
}

func (s *CreatePocSampleRequest) Validate() error {
	return dara.Validate(s)
}
