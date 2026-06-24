// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDatamaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExecDatamaskRequest
  GetData() *string 
  SetFeatureType(v int32) *ExecDatamaskRequest
  GetFeatureType() *int32 
  SetLang(v string) *ExecDatamaskRequest
  GetLang() *string 
  SetTemplateId(v int64) *ExecDatamaskRequest
  GetTemplateId() *int64 
}

type ExecDatamaskRequest struct {
  // The data that you want to mask. The data must be a string in JSON format and include the following fields:
  // 
  // - **dataHeaderList**: The column names of the data. The order of the column names must correspond to the order of the data that you want to mask.
  // 
  // - **dataList**: The data that you want to mask.
  // 
  // - **ruleList**: A list of sensitive data type IDs. The order of the IDs must correspond to the order of the data that you want to mask. Each ID is a number that represents a sensitive data type. You can call the [DescribeRules](https://help.aliyun.com/document_detail/410141.html) operation to obtain the IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"dataHeaderList":["name","age"],"dataList":[["lily",18],["lucy",17]],"ruleList":[1002,null]}
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // This parameter is deprecated.
  // 
  // example:
  // 
  // 2
  FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
  // The language of the request and response. Default value: **zh_cn**. Valid values:
  // 
  // - **zh_cn**: Simplified Chinese
  // 
  // - **en_us**: English (US)
  // 
  // example:
  // 
  // zh_cn
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The ID of the data masking template. A template ID is generated after you create a template in the [Data Security Center console](https://yundun.console.aliyun.com/?p=sddp#/dm/dmConfig/cn-zhangjiakou). You can find the **Template ID*	- on the **Data Masking*	- > **Masking Configuration*	- > **Masking Template*	- page.
  // 
  // - If the matching type of the data masking template is **Field Name**, the system matches the data against **dataHeaderList*	- in **Data**.
  // 
  // - If the matching type of the data masking template is **Sensitive Data Type**, the system matches the data against **ruleList*	- in **Data**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ExecDatamaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDatamaskRequest) GoString() string {
  return s.String()
}

func (s *ExecDatamaskRequest) GetData() *string  {
  return s.Data
}

func (s *ExecDatamaskRequest) GetFeatureType() *int32  {
  return s.FeatureType
}

func (s *ExecDatamaskRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecDatamaskRequest) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *ExecDatamaskRequest) SetData(v string) *ExecDatamaskRequest {
  s.Data = &v
  return s
}

func (s *ExecDatamaskRequest) SetFeatureType(v int32) *ExecDatamaskRequest {
  s.FeatureType = &v
  return s
}

func (s *ExecDatamaskRequest) SetLang(v string) *ExecDatamaskRequest {
  s.Lang = &v
  return s
}

func (s *ExecDatamaskRequest) SetTemplateId(v int64) *ExecDatamaskRequest {
  s.TemplateId = &v
  return s
}

func (s *ExecDatamaskRequest) Validate() error {
  return dara.Validate(s)
}

