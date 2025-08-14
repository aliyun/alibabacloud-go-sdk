// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocEvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateFormat(v string) *CreatePocEvRequest
	GetDateFormat() *string
	SetFileName(v string) *CreatePocEvRequest
	GetFileName() *string
	SetFileType(v string) *CreatePocEvRequest
	GetFileType() *string
	SetFileUrl(v string) *CreatePocEvRequest
	GetFileUrl() *string
	SetLang(v string) *CreatePocEvRequest
	GetLang() *string
	SetRegId(v string) *CreatePocEvRequest
	GetRegId() *string
	SetServiceCode(v string) *CreatePocEvRequest
	GetServiceCode() *string
	SetServiceName(v string) *CreatePocEvRequest
	GetServiceName() *string
	SetTab(v string) *CreatePocEvRequest
	GetTab() *string
	SetTaskName(v string) *CreatePocEvRequest
	GetTaskName() *string
	SetType(v string) *CreatePocEvRequest
	GetType() *string
}

type CreatePocEvRequest struct {
	// Date format type
	//
	// example:
	//
	// yyyyMMdd
	DateFormat *string `json:"DateFormat,omitempty" xml:"DateFormat,omitempty"`
	// File name.
	//
	// > The file name must end with txt or sql. For example, test.txt, test.sql.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File type
	//
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// File URL.
	//
	// example:
	//
	// saf/cpoc/953c883cde33b2e21d722eb661d26375/测试文件模板-通用.csv
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
	// Service code.
	//
	// example:
	//
	// anti_fraud_v2
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Service name.
	//
	// example:
	//
	// 注册风险
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Scenario.
	//
	// example:
	//
	// INTERNET
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// Task name.
	//
	// example:
	//
	// er-log-s3
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Access type.
	//
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePocEvRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePocEvRequest) GoString() string {
	return s.String()
}

func (s *CreatePocEvRequest) GetDateFormat() *string {
	return s.DateFormat
}

func (s *CreatePocEvRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreatePocEvRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreatePocEvRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreatePocEvRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePocEvRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreatePocEvRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreatePocEvRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreatePocEvRequest) GetTab() *string {
	return s.Tab
}

func (s *CreatePocEvRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreatePocEvRequest) GetType() *string {
	return s.Type
}

func (s *CreatePocEvRequest) SetDateFormat(v string) *CreatePocEvRequest {
	s.DateFormat = &v
	return s
}

func (s *CreatePocEvRequest) SetFileName(v string) *CreatePocEvRequest {
	s.FileName = &v
	return s
}

func (s *CreatePocEvRequest) SetFileType(v string) *CreatePocEvRequest {
	s.FileType = &v
	return s
}

func (s *CreatePocEvRequest) SetFileUrl(v string) *CreatePocEvRequest {
	s.FileUrl = &v
	return s
}

func (s *CreatePocEvRequest) SetLang(v string) *CreatePocEvRequest {
	s.Lang = &v
	return s
}

func (s *CreatePocEvRequest) SetRegId(v string) *CreatePocEvRequest {
	s.RegId = &v
	return s
}

func (s *CreatePocEvRequest) SetServiceCode(v string) *CreatePocEvRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreatePocEvRequest) SetServiceName(v string) *CreatePocEvRequest {
	s.ServiceName = &v
	return s
}

func (s *CreatePocEvRequest) SetTab(v string) *CreatePocEvRequest {
	s.Tab = &v
	return s
}

func (s *CreatePocEvRequest) SetTaskName(v string) *CreatePocEvRequest {
	s.TaskName = &v
	return s
}

func (s *CreatePocEvRequest) SetType(v string) *CreatePocEvRequest {
	s.Type = &v
	return s
}

func (s *CreatePocEvRequest) Validate() error {
	return dara.Validate(s)
}
