// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportIntervenesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportIntervenesResponseBody
  GetCode() *string 
  SetData(v *ExportIntervenesResponseBodyData) *ExportIntervenesResponseBody
  GetData() *ExportIntervenesResponseBodyData 
  SetHttpStatusCode(v int32) *ExportIntervenesResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportIntervenesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportIntervenesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportIntervenesResponseBody
  GetSuccess() *bool 
}

type ExportIntervenesResponseBody struct {
  // example:
  // 
  // 0
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *ExportIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ExportIntervenesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportIntervenesResponseBody) GoString() string {
  return s.String()
}

func (s *ExportIntervenesResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportIntervenesResponseBody) GetData() *ExportIntervenesResponseBodyData  {
  return s.Data
}

func (s *ExportIntervenesResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportIntervenesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportIntervenesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportIntervenesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportIntervenesResponseBody) SetCode(v string) *ExportIntervenesResponseBody {
  s.Code = &v
  return s
}

func (s *ExportIntervenesResponseBody) SetData(v *ExportIntervenesResponseBodyData) *ExportIntervenesResponseBody {
  s.Data = v
  return s
}

func (s *ExportIntervenesResponseBody) SetHttpStatusCode(v int32) *ExportIntervenesResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportIntervenesResponseBody) SetMessage(v string) *ExportIntervenesResponseBody {
  s.Message = &v
  return s
}

func (s *ExportIntervenesResponseBody) SetRequestId(v string) *ExportIntervenesResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportIntervenesResponseBody) SetSuccess(v bool) *ExportIntervenesResponseBody {
  s.Success = &v
  return s
}

func (s *ExportIntervenesResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExportIntervenesResponseBodyData struct {
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // http://xxx/xxx.xls
  FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ExportIntervenesResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportIntervenesResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportIntervenesResponseBodyData) GetCode() *int32  {
  return s.Code
}

func (s *ExportIntervenesResponseBodyData) GetFileUrl() *string  {
  return s.FileUrl
}

func (s *ExportIntervenesResponseBodyData) SetCode(v int32) *ExportIntervenesResponseBodyData {
  s.Code = &v
  return s
}

func (s *ExportIntervenesResponseBodyData) SetFileUrl(v string) *ExportIntervenesResponseBodyData {
  s.FileUrl = &v
  return s
}

func (s *ExportIntervenesResponseBodyData) Validate() error {
  return dara.Validate(s)
}

