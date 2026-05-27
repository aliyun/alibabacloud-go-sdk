// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateForeignPocSampleAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileObject(v io.Reader) *CreateForeignPocSampleAdvanceRequest
	GetFileObject() io.Reader
	SetLang(v string) *CreateForeignPocSampleAdvanceRequest
	GetLang() *string
	SetRegId(v string) *CreateForeignPocSampleAdvanceRequest
	GetRegId() *string
	SetRemark(v string) *CreateForeignPocSampleAdvanceRequest
	GetRemark() *string
	SetSampleName(v string) *CreateForeignPocSampleAdvanceRequest
	GetSampleName() *string
	SetTab(v string) *CreateForeignPocSampleAdvanceRequest
	GetTab() *string
}

type CreateForeignPocSampleAdvanceRequest struct {
	FileObject io.Reader `json:"File,omitempty" xml:"File,omitempty"`
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
	// nemo-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// SampleNameTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s CreateForeignPocSampleAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForeignPocSampleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateForeignPocSampleAdvanceRequest) GetFileObject() io.Reader {
	return s.FileObject
}

func (s *CreateForeignPocSampleAdvanceRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateForeignPocSampleAdvanceRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateForeignPocSampleAdvanceRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateForeignPocSampleAdvanceRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *CreateForeignPocSampleAdvanceRequest) GetTab() *string {
	return s.Tab
}

func (s *CreateForeignPocSampleAdvanceRequest) SetFileObject(v io.Reader) *CreateForeignPocSampleAdvanceRequest {
	s.FileObject = v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) SetLang(v string) *CreateForeignPocSampleAdvanceRequest {
	s.Lang = &v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) SetRegId(v string) *CreateForeignPocSampleAdvanceRequest {
	s.RegId = &v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) SetRemark(v string) *CreateForeignPocSampleAdvanceRequest {
	s.Remark = &v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) SetSampleName(v string) *CreateForeignPocSampleAdvanceRequest {
	s.SampleName = &v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) SetTab(v string) *CreateForeignPocSampleAdvanceRequest {
	s.Tab = &v
	return s
}

func (s *CreateForeignPocSampleAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
