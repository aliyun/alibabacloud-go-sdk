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
  // The de-identified data, which is described in a JSON string. The JSON string contains the following parameters:
  // 
  // 	- **dataHeaderList**: the names of columns that contain the de-identified data.
  // 
  // 	- **dataList**: the de-identified data. The column order of the de-identified data is the same as that indicated by the dataHeaderList parameter.
  // 
  // 	- **ruleList**: the IDs of sensitive data detection rules.
  // 
  // example:
  // 
  // {"dataHeaderList":["name","age"],"dataList":[["l***",18],["l***",17]],"ruleList":[1002,null]}
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The ID of the request, which is used to locate and troubleshoot issues.
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

