// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeDataSourceParametersResponseBodyData) *DescribeDataSourceParametersResponseBody
	GetData() []*DescribeDataSourceParametersResponseBodyData
	SetRequestId(v string) *DescribeDataSourceParametersResponseBody
	GetRequestId() *string
}

type DescribeDataSourceParametersResponseBody struct {
	// The data returned.
	Data []*DescribeDataSourceParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataSourceParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBody) GetData() []*DescribeDataSourceParametersResponseBodyData {
	return s.Data
}

func (s *DescribeDataSourceParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourceParametersResponseBody) SetData(v []*DescribeDataSourceParametersResponseBodyData) *DescribeDataSourceParametersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataSourceParametersResponseBody) SetRequestId(v string) *DescribeDataSourceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceParametersResponseBodyData struct {
	// Indicates whether the edit operation is supported. Valid values:
	//
	// 	- **0**
	//
	// 	- **1**
	//
	// example:
	//
	// wafApi
	CanEditted *int32 `json:"CanEditted,omitempty" xml:"CanEditted,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- **qcloud**: Tencent Cloud
	//
	// 	- **aliyun**: Alibaba Cloud
	//
	// 	- **hcloud**: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **obs**: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- **wafApi**: download API of Tencent Cloud Web Application Firewall (WAF)
	//
	// 	- **ckafka**: Tencent Cloud TDMQ for CKafka
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// wafApi
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the modification operation is forbidden. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// wafApi
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The method that is used to check the parameter format.
	//
	// example:
	//
	// email
	FormatCheck *string `json:"FormatCheck,omitempty" xml:"FormatCheck,omitempty"`
	// The additional information.
	//
	// example:
	//
	// obs docment
	Hit *string `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The code of the parameter.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The parameter level. Valid values:
	//
	// 	- **1**: the parameters of the data source
	//
	// 	- **2**: the parameters of the log
	//
	// example:
	//
	// 1
	ParaLevel *int32 `json:"ParaLevel,omitempty" xml:"ParaLevel,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// region local
	ParaName *string `json:"ParaName,omitempty" xml:"ParaName,omitempty"`
	// The data type of the parameter.
	//
	// example:
	//
	// string
	ParaType *string `json:"ParaType,omitempty" xml:"ParaType,omitempty"`
	// The value of the parameter.
	ParamValue []*DescribeDataSourceParametersResponseBodyDataParamValue `json:"ParamValue,omitempty" xml:"ParamValue,omitempty" type:"Repeated"`
	// Indicates whether the parameter is required. Valid values:
	//
	// 	- **1**: required
	//
	// 	- **0**: optional
	//
	// example:
	//
	// string
	Required *int32 `json:"Required,omitempty" xml:"Required,omitempty"`
	// The note for the parameter value.
	//
	// example:
	//
	// obs bucket name
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeDataSourceParametersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBodyData) GetCanEditted() *int32 {
	return s.CanEditted
}

func (s *DescribeDataSourceParametersResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeDataSourceParametersResponseBodyData) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeDataSourceParametersResponseBodyData) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeDataSourceParametersResponseBodyData) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeDataSourceParametersResponseBodyData) GetFormatCheck() *string {
	return s.FormatCheck
}

func (s *DescribeDataSourceParametersResponseBodyData) GetHit() *string {
	return s.Hit
}

func (s *DescribeDataSourceParametersResponseBodyData) GetParaCode() *string {
	return s.ParaCode
}

func (s *DescribeDataSourceParametersResponseBodyData) GetParaLevel() *int32 {
	return s.ParaLevel
}

func (s *DescribeDataSourceParametersResponseBodyData) GetParaName() *string {
	return s.ParaName
}

func (s *DescribeDataSourceParametersResponseBodyData) GetParaType() *string {
	return s.ParaType
}

func (s *DescribeDataSourceParametersResponseBodyData) GetParamValue() []*DescribeDataSourceParametersResponseBodyDataParamValue {
	return s.ParamValue
}

func (s *DescribeDataSourceParametersResponseBodyData) GetRequired() *int32 {
	return s.Required
}

func (s *DescribeDataSourceParametersResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DescribeDataSourceParametersResponseBodyData) SetCanEditted(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.CanEditted = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetCloudCode(v string) *DescribeDataSourceParametersResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDataSourceType(v string) *DescribeDataSourceParametersResponseBodyData {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDefaultValue(v string) *DescribeDataSourceParametersResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDisabled(v bool) *DescribeDataSourceParametersResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetFormatCheck(v string) *DescribeDataSourceParametersResponseBodyData {
	s.FormatCheck = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetHit(v string) *DescribeDataSourceParametersResponseBodyData {
	s.Hit = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaCode(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaLevel(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.ParaLevel = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaName(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaName = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaType(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaType = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParamValue(v []*DescribeDataSourceParametersResponseBodyDataParamValue) *DescribeDataSourceParametersResponseBodyData {
	s.ParamValue = v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetRequired(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.Required = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetTitle(v string) *DescribeDataSourceParametersResponseBodyData {
	s.Title = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) Validate() error {
	if s.ParamValue != nil {
		for _, item := range s.ParamValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceParametersResponseBodyDataParamValue struct {
	// The display value.
	//
	// example:
	//
	// guangzhou
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The actual value.
	//
	// example:
	//
	// ap-guangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataSourceParametersResponseBodyDataParamValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBodyDataParamValue) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) GetLabel() *string {
	return s.Label
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) SetLabel(v string) *DescribeDataSourceParametersResponseBodyDataParamValue {
	s.Label = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) SetValue(v string) *DescribeDataSourceParametersResponseBodyDataParamValue {
	s.Value = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) Validate() error {
	return dara.Validate(s)
}
