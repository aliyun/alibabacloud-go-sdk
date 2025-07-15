// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnalysisTagDetailByTaskIdResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetCode() *string 
  SetData(v string) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportAnalysisTagDetailByTaskIdResponseBody
  GetSuccess() *bool 
}

type ExportAnalysisTagDetailByTaskIdResponseBody struct {
  // example:
  // 
  // NoData
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // http://www.example.com/xxx.xlsx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 1813ceee-7fe5-41b4-87e5-982a4d18cca5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportAnalysisTagDetailByTaskIdResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportAnalysisTagDetailByTaskIdResponseBody) GoString() string {
  return s.String()
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetCode(v string) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.Code = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetData(v string) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.Data = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetHttpStatusCode(v int32) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetMessage(v string) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.Message = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetRequestId(v string) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) SetSuccess(v bool) *ExportAnalysisTagDetailByTaskIdResponseBody {
  s.Success = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) Validate() error {
  return dara.Validate(s)
}

