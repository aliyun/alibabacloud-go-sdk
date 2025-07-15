// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomCallTaggingResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportCustomCallTaggingResponseBody
  GetCode() *string 
  SetData(v string) *ExportCustomCallTaggingResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportCustomCallTaggingResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportCustomCallTaggingResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportCustomCallTaggingResponseBody
  GetRequestId() *string 
}

type ExportCustomCallTaggingResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // http://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-test/tagging.xlsx ?Expires=1610910578&amp;OSSAccessKeyId=****&amp;Signature=****
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // BA03159C-E808-4FF1-B27E-A61B6E888D7F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportCustomCallTaggingResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomCallTaggingResponseBody) GoString() string {
  return s.String()
}

func (s *ExportCustomCallTaggingResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportCustomCallTaggingResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportCustomCallTaggingResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportCustomCallTaggingResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportCustomCallTaggingResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportCustomCallTaggingResponseBody) SetCode(v string) *ExportCustomCallTaggingResponseBody {
  s.Code = &v
  return s
}

func (s *ExportCustomCallTaggingResponseBody) SetData(v string) *ExportCustomCallTaggingResponseBody {
  s.Data = &v
  return s
}

func (s *ExportCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *ExportCustomCallTaggingResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportCustomCallTaggingResponseBody) SetMessage(v string) *ExportCustomCallTaggingResponseBody {
  s.Message = &v
  return s
}

func (s *ExportCustomCallTaggingResponseBody) SetRequestId(v string) *ExportCustomCallTaggingResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportCustomCallTaggingResponseBody) Validate() error {
  return dara.Validate(s)
}

