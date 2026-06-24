// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDatamaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExecDatamaskResponseBody
  GetData() *string 
  SetRequestId(v string) *ExecDatamaskResponseBody
  GetRequestId() *string 
}

type ExecDatamaskResponseBody struct {
  // The data after it is masked. The data is a string in JSON format and includes the following fields:
  // 
  // - **dataHeaderList**: The column names of the masked data.
  // 
  // - **dataList**: The masked data. The order of the fields corresponds to the order of the column names.
  // 
  // - **ruleList**: The sensitive data type IDs. Each ID corresponds to the unique ID of a sensitive data type that is returned by the [DescribeRules](https://help.aliyun.com/document_detail/410141.html) operation.
  // 
  // example:
  // 
  // {"dataHeaderList":["name","age"],"dataList":[["l***",18],["l***",17]],"ruleList":[1002,null]}
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use this ID to troubleshoot issues.
  // 
  // example:
  // 
  // 813BA9FA-D062-42C4-8CD5-11A7640B96E6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecDatamaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecDatamaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecDatamaskResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExecDatamaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecDatamaskResponseBody) SetData(v string) *ExecDatamaskResponseBody {
  s.Data = &v
  return s
}

func (s *ExecDatamaskResponseBody) SetRequestId(v string) *ExecDatamaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecDatamaskResponseBody) Validate() error {
  return dara.Validate(s)
}

