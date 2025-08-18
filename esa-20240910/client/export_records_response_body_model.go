// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExportRecordsResponseBody
  GetContent() *string 
  SetRequestId(v string) *ExportRecordsResponseBody
  GetRequestId() *string 
}

type ExportRecordsResponseBody struct {
  // The exported DNS records.
  // 
  // example:
  // 
  // ;; site:example.com.\\n;; Exported:2024-01-24 15:54:35\\n\\n;; A Records\\na1.example.com. 30 IN A 1.1.1.1 direct\\na2.example.com. 30 IN A 1.1.1.1 direct\\na3.example.com. 30 IN A 1.1.1.1 direct\\n
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // C69B5894-D1BA-592C-95D0-DADBE7AEAC63
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportRecordsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportRecordsResponseBody) GetContent() *string  {
  return s.Content
}

func (s *ExportRecordsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportRecordsResponseBody) SetContent(v string) *ExportRecordsResponseBody {
  s.Content = &v
  return s
}

func (s *ExportRecordsResponseBody) SetRequestId(v string) *ExportRecordsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportRecordsResponseBody) Validate() error {
  return dara.Validate(s)
}

