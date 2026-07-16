// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordCnamesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportRecordCnamesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportRecordCnamesResponse
  GetStatusCode() *int32 
  SetBody(v *ExportRecordCnamesResponseBody) *ExportRecordCnamesResponse
  GetBody() *ExportRecordCnamesResponseBody 
}

type ExportRecordCnamesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportRecordCnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRecordCnamesResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordCnamesResponse) GoString() string {
  return s.String()
}

func (s *ExportRecordCnamesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportRecordCnamesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportRecordCnamesResponse) GetBody() *ExportRecordCnamesResponseBody  {
  return s.Body
}

func (s *ExportRecordCnamesResponse) SetHeaders(v map[string]*string) *ExportRecordCnamesResponse {
  s.Headers = v
  return s
}

func (s *ExportRecordCnamesResponse) SetStatusCode(v int32) *ExportRecordCnamesResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportRecordCnamesResponse) SetBody(v *ExportRecordCnamesResponseBody) *ExportRecordCnamesResponse {
  s.Body = v
  return s
}

func (s *ExportRecordCnamesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

