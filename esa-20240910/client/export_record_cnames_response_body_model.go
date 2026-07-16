// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordCnamesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExportRecordCnamesResponseBody
  GetContent() *string 
  SetRequestId(v string) *ExportRecordCnamesResponseBody
  GetRequestId() *string 
}

type ExportRecordCnamesResponseBody struct {
  // The exported CNAME content.
  // 
  // example:
  // 
  // a.com a.com.cname.zone
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 2430E05E-1340-5773-B5E1-B743929F46F2
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportRecordCnamesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordCnamesResponseBody) GoString() string {
  return s.String()
}

func (s *ExportRecordCnamesResponseBody) GetContent() *string  {
  return s.Content
}

func (s *ExportRecordCnamesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportRecordCnamesResponseBody) SetContent(v string) *ExportRecordCnamesResponseBody {
  s.Content = &v
  return s
}

func (s *ExportRecordCnamesResponseBody) SetRequestId(v string) *ExportRecordCnamesResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportRecordCnamesResponseBody) Validate() error {
  return dara.Validate(s)
}

