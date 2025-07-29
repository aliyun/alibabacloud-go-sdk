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
}

type ExportAnalysisTagDetailByTaskIdResponseBody struct {
  // example:
  // 
  // xx
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // http://www.example.com/xxxx.xlsx
  Data *string `json:"data,omitempty" xml:"data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // example:
  // 
  // ok
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *ExportAnalysisTagDetailByTaskIdResponseBody) Validate() error {
  return dara.Validate(s)
}

