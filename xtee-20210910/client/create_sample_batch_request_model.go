// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateSampleBatchRequest
	GetLang() *string
	SetBatchName(v string) *CreateSampleBatchRequest
	GetBatchName() *string
	SetDataType(v string) *CreateSampleBatchRequest
	GetDataType() *string
	SetDataValue(v string) *CreateSampleBatchRequest
	GetDataValue() *string
	SetDescription(v string) *CreateSampleBatchRequest
	GetDescription() *string
	SetOssFileName(v string) *CreateSampleBatchRequest
	GetOssFileName() *string
	SetRegId(v string) *CreateSampleBatchRequest
	GetRegId() *string
	SetSampleBatchType(v string) *CreateSampleBatchRequest
	GetSampleBatchType() *string
	SetServiceList(v string) *CreateSampleBatchRequest
	GetServiceList() *string
}

type CreateSampleBatchRequest struct {
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
	// Content of the list entered in the text box
	//
	// example:
	//
	// 1770000000,1770000001
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
	// Specific type of sample list
	//
	// example:
	//
	// 白名单
	SampleBatchType *string `json:"sampleBatchType,omitempty" xml:"sampleBatchType,omitempty"`
	// Service list
	//
	// example:
	//
	// account_takeover
	ServiceList *string `json:"serviceList,omitempty" xml:"serviceList,omitempty"`
}

func (s CreateSampleBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleBatchRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSampleBatchRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *CreateSampleBatchRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateSampleBatchRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *CreateSampleBatchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSampleBatchRequest) GetOssFileName() *string {
	return s.OssFileName
}

func (s *CreateSampleBatchRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateSampleBatchRequest) GetSampleBatchType() *string {
	return s.SampleBatchType
}

func (s *CreateSampleBatchRequest) GetServiceList() *string {
	return s.ServiceList
}

func (s *CreateSampleBatchRequest) SetLang(v string) *CreateSampleBatchRequest {
	s.Lang = &v
	return s
}

func (s *CreateSampleBatchRequest) SetBatchName(v string) *CreateSampleBatchRequest {
	s.BatchName = &v
	return s
}

func (s *CreateSampleBatchRequest) SetDataType(v string) *CreateSampleBatchRequest {
	s.DataType = &v
	return s
}

func (s *CreateSampleBatchRequest) SetDataValue(v string) *CreateSampleBatchRequest {
	s.DataValue = &v
	return s
}

func (s *CreateSampleBatchRequest) SetDescription(v string) *CreateSampleBatchRequest {
	s.Description = &v
	return s
}

func (s *CreateSampleBatchRequest) SetOssFileName(v string) *CreateSampleBatchRequest {
	s.OssFileName = &v
	return s
}

func (s *CreateSampleBatchRequest) SetRegId(v string) *CreateSampleBatchRequest {
	s.RegId = &v
	return s
}

func (s *CreateSampleBatchRequest) SetSampleBatchType(v string) *CreateSampleBatchRequest {
	s.SampleBatchType = &v
	return s
}

func (s *CreateSampleBatchRequest) SetServiceList(v string) *CreateSampleBatchRequest {
	s.ServiceList = &v
	return s
}

func (s *CreateSampleBatchRequest) Validate() error {
	return dara.Validate(s)
}
