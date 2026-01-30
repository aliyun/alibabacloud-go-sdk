// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportRecordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportRecordResponse
  GetStatusCode() *int32 
  SetBody(v *ExportRecordResponseBody) *ExportRecordResponse
  GetBody() *ExportRecordResponseBody 
}

type ExportRecordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRecordResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordResponse) GoString() string {
  return s.String()
}

func (s *ExportRecordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportRecordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportRecordResponse) GetBody() *ExportRecordResponseBody  {
  return s.Body
}

func (s *ExportRecordResponse) SetHeaders(v map[string]*string) *ExportRecordResponse {
  s.Headers = v
  return s
}

func (s *ExportRecordResponse) SetStatusCode(v int32) *ExportRecordResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportRecordResponse) SetBody(v *ExportRecordResponseBody) *ExportRecordResponse {
  s.Body = v
  return s
}

func (s *ExportRecordResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

