// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateSampleRequest
	GetLang() *string
	SetClientFileName(v string) *CreateSampleRequest
	GetClientFileName() *string
	SetClientPath(v string) *CreateSampleRequest
	GetClientPath() *string
	SetFileType(v string) *CreateSampleRequest
	GetFileType() *string
	SetRegId(v string) *CreateSampleRequest
	GetRegId() *string
	SetSampleTag(v string) *CreateSampleRequest
	GetSampleTag() *string
	SetSampleType(v string) *CreateSampleRequest
	GetSampleType() *string
	SetSampleValues(v string) *CreateSampleRequest
	GetSampleValues() *string
	SetUploadType(v string) *CreateSampleRequest
	GetUploadType() *string
}

type CreateSampleRequest struct {
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
	// OSS client uploaded file name
	//
	// example:
	//
	// 样本文件.csv
	ClientFileName *string `json:"clientFileName,omitempty" xml:"clientFileName,omitempty"`
	// OSS client address
	//
	// example:
	//
	// sample/path
	ClientPath *string `json:"clientPath,omitempty" xml:"clientPath,omitempty"`
	// File type
	//
	// example:
	//
	// CSV
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample tag
	//
	// example:
	//
	// rm0102
	SampleTag *string `json:"sampleTag,omitempty" xml:"sampleTag,omitempty"`
	// Sample type
	//
	// example:
	//
	// PHONE
	SampleType *string `json:"sampleType,omitempty" xml:"sampleType,omitempty"`
	// Sample values
	//
	// example:
	//
	// 1777000000,1777000001
	SampleValues *string `json:"sampleValues,omitempty" xml:"sampleValues,omitempty"`
	// Upload type
	//
	// example:
	//
	// ANNEX
	UploadType *string `json:"uploadType,omitempty" xml:"uploadType,omitempty"`
}

func (s CreateSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSampleRequest) GetClientFileName() *string {
	return s.ClientFileName
}

func (s *CreateSampleRequest) GetClientPath() *string {
	return s.ClientPath
}

func (s *CreateSampleRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateSampleRequest) GetSampleTag() *string {
	return s.SampleTag
}

func (s *CreateSampleRequest) GetSampleType() *string {
	return s.SampleType
}

func (s *CreateSampleRequest) GetSampleValues() *string {
	return s.SampleValues
}

func (s *CreateSampleRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *CreateSampleRequest) SetLang(v string) *CreateSampleRequest {
	s.Lang = &v
	return s
}

func (s *CreateSampleRequest) SetClientFileName(v string) *CreateSampleRequest {
	s.ClientFileName = &v
	return s
}

func (s *CreateSampleRequest) SetClientPath(v string) *CreateSampleRequest {
	s.ClientPath = &v
	return s
}

func (s *CreateSampleRequest) SetFileType(v string) *CreateSampleRequest {
	s.FileType = &v
	return s
}

func (s *CreateSampleRequest) SetRegId(v string) *CreateSampleRequest {
	s.RegId = &v
	return s
}

func (s *CreateSampleRequest) SetSampleTag(v string) *CreateSampleRequest {
	s.SampleTag = &v
	return s
}

func (s *CreateSampleRequest) SetSampleType(v string) *CreateSampleRequest {
	s.SampleType = &v
	return s
}

func (s *CreateSampleRequest) SetSampleValues(v string) *CreateSampleRequest {
	s.SampleValues = &v
	return s
}

func (s *CreateSampleRequest) SetUploadType(v string) *CreateSampleRequest {
	s.UploadType = &v
	return s
}

func (s *CreateSampleRequest) Validate() error {
	return dara.Validate(s)
}
