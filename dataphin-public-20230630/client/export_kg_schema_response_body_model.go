// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKgSchemaResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportKgSchemaResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *ExportKgSchemaResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportKgSchemaResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportKgSchemaResponseBody
  GetRequestId() *string 
  SetSchemaInfo(v *ExportKgSchemaResponseBodySchemaInfo) *ExportKgSchemaResponseBody
  GetSchemaInfo() *ExportKgSchemaResponseBodySchemaInfo 
  SetSuccess(v bool) *ExportKgSchemaResponseBody
  GetSuccess() *bool 
}

type ExportKgSchemaResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // internal error
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  SchemaInfo *ExportKgSchemaResponseBodySchemaInfo `json:"SchemaInfo,omitempty" xml:"SchemaInfo,omitempty" type:"Struct"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportKgSchemaResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportKgSchemaResponseBody) GoString() string {
  return s.String()
}

func (s *ExportKgSchemaResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportKgSchemaResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportKgSchemaResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportKgSchemaResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportKgSchemaResponseBody) GetSchemaInfo() *ExportKgSchemaResponseBodySchemaInfo  {
  return s.SchemaInfo
}

func (s *ExportKgSchemaResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportKgSchemaResponseBody) SetCode(v string) *ExportKgSchemaResponseBody {
  s.Code = &v
  return s
}

func (s *ExportKgSchemaResponseBody) SetHttpStatusCode(v int32) *ExportKgSchemaResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportKgSchemaResponseBody) SetMessage(v string) *ExportKgSchemaResponseBody {
  s.Message = &v
  return s
}

func (s *ExportKgSchemaResponseBody) SetRequestId(v string) *ExportKgSchemaResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportKgSchemaResponseBody) SetSchemaInfo(v *ExportKgSchemaResponseBodySchemaInfo) *ExportKgSchemaResponseBody {
  s.SchemaInfo = v
  return s
}

func (s *ExportKgSchemaResponseBody) SetSuccess(v bool) *ExportKgSchemaResponseBody {
  s.Success = &v
  return s
}

func (s *ExportKgSchemaResponseBody) Validate() error {
  if s.SchemaInfo != nil {
    if err := s.SchemaInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportKgSchemaResponseBodySchemaInfo struct {
  // example:
  // 
  // workspaceId: f1d4559a4db044158305e2d89bccf81f
  // 
  // name: jytest
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // example:
  // 
  // yaml
  OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
}

func (s ExportKgSchemaResponseBodySchemaInfo) String() string {
  return dara.Prettify(s)
}

func (s ExportKgSchemaResponseBodySchemaInfo) GoString() string {
  return s.String()
}

func (s *ExportKgSchemaResponseBodySchemaInfo) GetContent() *string  {
  return s.Content
}

func (s *ExportKgSchemaResponseBodySchemaInfo) GetOutputFormat() *string  {
  return s.OutputFormat
}

func (s *ExportKgSchemaResponseBodySchemaInfo) SetContent(v string) *ExportKgSchemaResponseBodySchemaInfo {
  s.Content = &v
  return s
}

func (s *ExportKgSchemaResponseBodySchemaInfo) SetOutputFormat(v string) *ExportKgSchemaResponseBodySchemaInfo {
  s.OutputFormat = &v
  return s
}

func (s *ExportKgSchemaResponseBodySchemaInfo) Validate() error {
  return dara.Validate(s)
}

