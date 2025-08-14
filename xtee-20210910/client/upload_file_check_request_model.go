// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UploadFileCheckRequest
	GetLang() *string
	SetBatchName(v string) *UploadFileCheckRequest
	GetBatchName() *string
	SetDataType(v string) *UploadFileCheckRequest
	GetDataType() *string
	SetOssFileName(v string) *UploadFileCheckRequest
	GetOssFileName() *string
	SetRegId(v string) *UploadFileCheckRequest
	GetRegId() *string
	SetSampleTagType(v string) *UploadFileCheckRequest
	GetSampleTagType() *string
	SetServiceList(v string) *UploadFileCheckRequest
	GetServiceList() *string
}

type UploadFileCheckRequest struct {
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
	// Sample batch name
	//
	// example:
	//
	// 白样本
	BatchName *string `json:"batchName,omitempty" xml:"batchName,omitempty"`
	// Data type
	//
	// example:
	//
	// mobile
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// Uploaded OSS address.
	//
	// example:
	//
	// sample/94b4193d321c44b44ee92b19984000a0/样本测试01/0d8dbc6809834d51b1d88a42295c803e/1753865835166.csv
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample type
	//
	// example:
	//
	// 白名单
	SampleTagType *string `json:"sampleTagType,omitempty" xml:"sampleTagType,omitempty"`
	// Service list
	//
	// example:
	//
	// account_abuse
	ServiceList *string `json:"serviceList,omitempty" xml:"serviceList,omitempty"`
}

func (s UploadFileCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadFileCheckRequest) GoString() string {
	return s.String()
}

func (s *UploadFileCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *UploadFileCheckRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *UploadFileCheckRequest) GetDataType() *string {
	return s.DataType
}

func (s *UploadFileCheckRequest) GetOssFileName() *string {
	return s.OssFileName
}

func (s *UploadFileCheckRequest) GetRegId() *string {
	return s.RegId
}

func (s *UploadFileCheckRequest) GetSampleTagType() *string {
	return s.SampleTagType
}

func (s *UploadFileCheckRequest) GetServiceList() *string {
	return s.ServiceList
}

func (s *UploadFileCheckRequest) SetLang(v string) *UploadFileCheckRequest {
	s.Lang = &v
	return s
}

func (s *UploadFileCheckRequest) SetBatchName(v string) *UploadFileCheckRequest {
	s.BatchName = &v
	return s
}

func (s *UploadFileCheckRequest) SetDataType(v string) *UploadFileCheckRequest {
	s.DataType = &v
	return s
}

func (s *UploadFileCheckRequest) SetOssFileName(v string) *UploadFileCheckRequest {
	s.OssFileName = &v
	return s
}

func (s *UploadFileCheckRequest) SetRegId(v string) *UploadFileCheckRequest {
	s.RegId = &v
	return s
}

func (s *UploadFileCheckRequest) SetSampleTagType(v string) *UploadFileCheckRequest {
	s.SampleTagType = &v
	return s
}

func (s *UploadFileCheckRequest) SetServiceList(v string) *UploadFileCheckRequest {
	s.ServiceList = &v
	return s
}

func (s *UploadFileCheckRequest) Validate() error {
	return dara.Validate(s)
}
