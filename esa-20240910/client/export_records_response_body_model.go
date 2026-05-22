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
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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

