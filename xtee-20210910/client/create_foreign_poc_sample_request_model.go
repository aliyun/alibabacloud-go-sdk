// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForeignPocSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFile(v string) *CreateForeignPocSampleRequest
	GetFile() *string
	SetLang(v string) *CreateForeignPocSampleRequest
	GetLang() *string
	SetRegId(v string) *CreateForeignPocSampleRequest
	GetRegId() *string
	SetRemark(v string) *CreateForeignPocSampleRequest
	GetRemark() *string
	SetSampleName(v string) *CreateForeignPocSampleRequest
	GetSampleName() *string
	SetTab(v string) *CreateForeignPocSampleRequest
	GetTab() *string
}

type CreateForeignPocSampleRequest struct {
	// OSS path of the file.
	//
	// example:
	//
	// saf/cpoc/953c883cde33b2e21d722eb661d26375/1779172027996_自动回溯测试 2605191.csv
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// Set the language type for requests and received messages. Default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Area encoding.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Remarks.
	//
	// example:
	//
	// nemo-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Sample Name.
	//
	// example:
	//
	// SampleNameTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// Scenario.
	//
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s CreateForeignPocSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForeignPocSampleRequest) GoString() string {
	return s.String()
}

func (s *CreateForeignPocSampleRequest) GetFile() *string {
	return s.File
}

func (s *CreateForeignPocSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateForeignPocSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateForeignPocSampleRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateForeignPocSampleRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *CreateForeignPocSampleRequest) GetTab() *string {
	return s.Tab
}

func (s *CreateForeignPocSampleRequest) SetFile(v string) *CreateForeignPocSampleRequest {
	s.File = &v
	return s
}

func (s *CreateForeignPocSampleRequest) SetLang(v string) *CreateForeignPocSampleRequest {
	s.Lang = &v
	return s
}

func (s *CreateForeignPocSampleRequest) SetRegId(v string) *CreateForeignPocSampleRequest {
	s.RegId = &v
	return s
}

func (s *CreateForeignPocSampleRequest) SetRemark(v string) *CreateForeignPocSampleRequest {
	s.Remark = &v
	return s
}

func (s *CreateForeignPocSampleRequest) SetSampleName(v string) *CreateForeignPocSampleRequest {
	s.SampleName = &v
	return s
}

func (s *CreateForeignPocSampleRequest) SetTab(v string) *CreateForeignPocSampleRequest {
	s.Tab = &v
	return s
}

func (s *CreateForeignPocSampleRequest) Validate() error {
	return dara.Validate(s)
}
