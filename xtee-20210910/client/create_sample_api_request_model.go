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
	// Same as the request parameter.
	//
	// example:
	//
	// ip/accountID
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The specific data value.
	//
	// example:
	//
	// 同参数
	DataValue *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// The sample batch type.
	//
	// example:
	//
	// 白名单/黑名单/混合
	SampleBatchType *string `json:"SampleBatchType,omitempty" xml:"SampleBatchType,omitempty"`
	// The list of services.
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
