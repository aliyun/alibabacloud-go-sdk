// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomSourceAnalysisTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportCustomSourceAnalysisTaskResponseBody
  GetCode() *string 
  SetData(v string) *ExportCustomSourceAnalysisTaskResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportCustomSourceAnalysisTaskResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportCustomSourceAnalysisTaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportCustomSourceAnalysisTaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportCustomSourceAnalysisTaskResponseBody
  GetSuccess() *bool 
}

type ExportCustomSourceAnalysisTaskResponseBody struct {
  // example:
  // 
  // NoData
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // http://www.example.com/xxx.jsonLine
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

func (s ExportCustomSourceAnalysisTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomSourceAnalysisTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetCode(v string) *ExportCustomSourceAnalysisTaskResponseBody {
  s.Code = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetData(v string) *ExportCustomSourceAnalysisTaskResponseBody {
  s.Data = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *ExportCustomSourceAnalysisTaskResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetMessage(v string) *ExportCustomSourceAnalysisTaskResponseBody {
  s.Message = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetRequestId(v string) *ExportCustomSourceAnalysisTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) SetSuccess(v bool) *ExportCustomSourceAnalysisTaskResponseBody {
  s.Success = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

