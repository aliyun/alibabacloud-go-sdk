// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *CreateSampleApiRequest
	GetDataType() *string
	SetDataValue(v string) *CreateSampleApiRequest
	GetDataValue() *string
	SetLang(v string) *CreateSampleApiRequest
	GetLang() *string
	SetRegId(v string) *CreateSampleApiRequest
	GetRegId() *string
	SetSampleBatchType(v string) *CreateSampleApiRequest
	GetSampleBatchType() *string
	SetServiceList(v string) *CreateSampleApiRequest
	GetServiceList() *string
}

type CreateSampleApiRequest struct {
	// Same as input parameter
	//
	// example:
	//
	// ip/accountID
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Specific data value
	//
	// example:
	//
	// 同参数
	DataValue *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
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
	// regionId
	//
	// example:
	//
	// cn-shanghai
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Sample batch type
	//
	// example:
	//
	// 白名单/黑名单/混合
	SampleBatchType *string `json:"SampleBatchType,omitempty" xml:"SampleBatchType,omitempty"`
	// Service list.
	//
	// example:
	//
	// 同参数
	ServiceList *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
}

func (s CreateSampleApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleApiRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleApiRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateSampleApiRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *CreateSampleApiRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSampleApiRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateSampleApiRequest) GetSampleBatchType() *string {
	return s.SampleBatchType
}

func (s *CreateSampleApiRequest) GetServiceList() *string {
	return s.ServiceList
}

func (s *CreateSampleApiRequest) SetDataType(v string) *CreateSampleApiRequest {
	s.DataType = &v
	return s
}

func (s *CreateSampleApiRequest) SetDataValue(v string) *CreateSampleApiRequest {
	s.DataValue = &v
	return s
}

func (s *CreateSampleApiRequest) SetLang(v string) *CreateSampleApiRequest {
	s.Lang = &v
	return s
}

func (s *CreateSampleApiRequest) SetRegId(v string) *CreateSampleApiRequest {
	s.RegId = &v
	return s
}

func (s *CreateSampleApiRequest) SetSampleBatchType(v string) *CreateSampleApiRequest {
	s.SampleBatchType = &v
	return s
}

func (s *CreateSampleApiRequest) SetServiceList(v string) *CreateSampleApiRequest {
	s.ServiceList = &v
	return s
}

func (s *CreateSampleApiRequest) Validate() error {
	return dara.Validate(s)
}
