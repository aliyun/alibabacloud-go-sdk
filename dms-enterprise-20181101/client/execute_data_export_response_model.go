// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataExportResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteDataExportResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteDataExportResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteDataExportResponseBody) *ExecuteDataExportResponse
  GetBody() *ExecuteDataExportResponseBody 
}

type ExecuteDataExportResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteDataExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteDataExportResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataExportResponse) GoString() string {
  return s.String()
}

func (s *ExecuteDataExportResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteDataExportResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteDataExportResponse) GetBody() *ExecuteDataExportResponseBody  {
  return s.Body
}

func (s *ExecuteDataExportResponse) SetHeaders(v map[string]*string) *ExecuteDataExportResponse {
  s.Headers = v
  return s
}

func (s *ExecuteDataExportResponse) SetStatusCode(v int32) *ExecuteDataExportResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteDataExportResponse) SetBody(v *ExecuteDataExportResponseBody) *ExecuteDataExportResponse {
  s.Body = v
  return s
}

func (s *ExecuteDataExportResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

