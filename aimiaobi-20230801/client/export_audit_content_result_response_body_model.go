// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAuditContentResultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportAuditContentResultResponseBody
  GetCode() *string 
  SetData(v string) *ExportAuditContentResultResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportAuditContentResultResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportAuditContentResultResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportAuditContentResultResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportAuditContentResultResponseBody
  GetSuccess() *bool 
}

type ExportAuditContentResultResponseBody struct {
  // example:
  // 
  // successful
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // http://www.example.com/xxx.docx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 94512A33-8EC1-5452-A793-5C91F18ED2F0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportAuditContentResultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportAuditContentResultResponseBody) GoString() string {
  return s.String()
}

func (s *ExportAuditContentResultResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportAuditContentResultResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportAuditContentResultResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportAuditContentResultResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportAuditContentResultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportAuditContentResultResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportAuditContentResultResponseBody) SetCode(v string) *ExportAuditContentResultResponseBody {
  s.Code = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) SetData(v string) *ExportAuditContentResultResponseBody {
  s.Data = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) SetHttpStatusCode(v int32) *ExportAuditContentResultResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) SetMessage(v string) *ExportAuditContentResultResponseBody {
  s.Message = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) SetRequestId(v string) *ExportAuditContentResultResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) SetSuccess(v bool) *ExportAuditContentResultResponseBody {
  s.Success = &v
  return s
}

func (s *ExportAuditContentResultResponseBody) Validate() error {
  return dara.Validate(s)
}

