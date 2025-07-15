// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGeneratedContentResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportGeneratedContentResponseBody
  GetCode() *string 
  SetData(v string) *ExportGeneratedContentResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportGeneratedContentResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportGeneratedContentResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportGeneratedContentResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportGeneratedContentResponseBody
  GetSuccess() *bool 
}

type ExportGeneratedContentResponseBody struct {
  // example:
  // 
  // NoData
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // xxx
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

func (s ExportGeneratedContentResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportGeneratedContentResponseBody) GoString() string {
  return s.String()
}

func (s *ExportGeneratedContentResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportGeneratedContentResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportGeneratedContentResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportGeneratedContentResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportGeneratedContentResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportGeneratedContentResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportGeneratedContentResponseBody) SetCode(v string) *ExportGeneratedContentResponseBody {
  s.Code = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) SetData(v string) *ExportGeneratedContentResponseBody {
  s.Data = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) SetHttpStatusCode(v int32) *ExportGeneratedContentResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) SetMessage(v string) *ExportGeneratedContentResponseBody {
  s.Message = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) SetRequestId(v string) *ExportGeneratedContentResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) SetSuccess(v bool) *ExportGeneratedContentResponseBody {
  s.Success = &v
  return s
}

func (s *ExportGeneratedContentResponseBody) Validate() error {
  return dara.Validate(s)
}

