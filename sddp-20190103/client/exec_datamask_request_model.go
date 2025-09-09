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
  // The sensitive data to be de-identified. The value is a JSON string that contains the following parameters:
  // 
  // 	- **dataHeaderList**: the names of the columns in which data needs to be de-identified. Specify the column names in accordance with the order of data that needs to be de-identified.
  // 
  // 	- **dataList**: the data that needs to be de-identified.
  // 
  // 	- **ruleList**: the IDs of sensitive data detection rules used to detect data that needs to be de-identified. Specify the rule IDs in accordance with the order of data that needs to be de-identified. Each ID identifies a sensitive data detection rule that is used to detect a type of sensitive data. You can call the [DescribeRules](~~DescribeRules~~) operation to query the IDs of sensitive data detection rules.
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
  // 1
  FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
  // The language of the content within the request and response. Default value: **zh_cn**. Valid values:
  // 
  // 	- **zh_cn**: Simplified Chinese
  // 
  // 	- **en_us**: English
  // 
  // example:
  // 
  // zh_cn
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The ID of the de-identification template. The ID is generated after you create the de-identification template in the [Data Security Center (DSC) console](https://yundun.console.aliyun.com/?\\&p=sddpnext#/sddp/dm/template). You can choose **Data desensitization*	- > **Desensitization Template*	- in the left-side navigation pane and obtain the ID of the de-identification template from the **Desensitization Template*	- page.
  // 
  // 	- If you select **Field name*	- as the matching mode of the template, DSC matches data based on the columns specified by the **dataHeaderList*	- parameter in the **Data*	- parameter.
  // 
  // 	- If you select **Sensitive type*	- as the matching mode of the template, DSC matches data based on the sensitive data detection rules specified by the **ruleList*	- parameter in the **Data*	- parameter.
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

