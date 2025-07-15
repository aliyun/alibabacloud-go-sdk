// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOASResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExportOASResponseBody
  GetData() *string 
  SetRequestId(v string) *ExportOASResponseBody
  GetRequestId() *string 
}

type ExportOASResponseBody struct {
  // The Base64-encoded data of the exported Swagger file. You can obtain the file by using Base64 decoding.
  // 
  // example:
  // 
  // UEsDBBQACAAIAABc8FgAAAAAAAAAAAAAAAA...
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // CEF72CEB-54B6-4AE8-B225-F876xxxxxxxx
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportOASResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportOASResponseBody) GoString() string {
  return s.String()
}

func (s *ExportOASResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportOASResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportOASResponseBody) SetData(v string) *ExportOASResponseBody {
  s.Data = &v
  return s
}

func (s *ExportOASResponseBody) SetRequestId(v string) *ExportOASResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportOASResponseBody) Validate() error {
  return dara.Validate(s)
}

