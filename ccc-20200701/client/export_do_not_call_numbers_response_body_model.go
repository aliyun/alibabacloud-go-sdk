// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDoNotCallNumbersResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportDoNotCallNumbersResponseBody
  GetCode() *string 
  SetData(v string) *ExportDoNotCallNumbersResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportDoNotCallNumbersResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportDoNotCallNumbersResponseBody
  GetMessage() *string 
  SetParams(v []*string) *ExportDoNotCallNumbersResponseBody
  GetParams() []*string 
  SetRequestId(v string) *ExportDoNotCallNumbersResponseBody
  GetRequestId() *string 
}

type ExportDoNotCallNumbersResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // https://****.oss-cn-shanghai.aliyuncs.com/ccc-test/blacklist.xlsx?Expires=3294624578&OSSAccessKeyId=****&Signature=****
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // BA03159C-E808-4FF1-B27E-A61B6E888D7F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportDoNotCallNumbersResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportDoNotCallNumbersResponseBody) GoString() string {
  return s.String()
}

func (s *ExportDoNotCallNumbersResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportDoNotCallNumbersResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportDoNotCallNumbersResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportDoNotCallNumbersResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportDoNotCallNumbersResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *ExportDoNotCallNumbersResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportDoNotCallNumbersResponseBody) SetCode(v string) *ExportDoNotCallNumbersResponseBody {
  s.Code = &v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) SetData(v string) *ExportDoNotCallNumbersResponseBody {
  s.Data = &v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) SetHttpStatusCode(v int32) *ExportDoNotCallNumbersResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) SetMessage(v string) *ExportDoNotCallNumbersResponseBody {
  s.Message = &v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) SetParams(v []*string) *ExportDoNotCallNumbersResponseBody {
  s.Params = v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) SetRequestId(v string) *ExportDoNotCallNumbersResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportDoNotCallNumbersResponseBody) Validate() error {
  return dara.Validate(s)
}

