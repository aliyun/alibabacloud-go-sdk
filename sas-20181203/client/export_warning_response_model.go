// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWarningResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportWarningResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportWarningResponse
  GetStatusCode() *int32 
  SetBody(v *ExportWarningResponseBody) *ExportWarningResponse
  GetBody() *ExportWarningResponseBody 
}

type ExportWarningResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportWarningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportWarningResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportWarningResponse) GoString() string {
  return s.String()
}

func (s *ExportWarningResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportWarningResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportWarningResponse) GetBody() *ExportWarningResponseBody  {
  return s.Body
}

func (s *ExportWarningResponse) SetHeaders(v map[string]*string) *ExportWarningResponse {
  s.Headers = v
  return s
}

func (s *ExportWarningResponse) SetStatusCode(v int32) *ExportWarningResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportWarningResponse) SetBody(v *ExportWarningResponseBody) *ExportWarningResponse {
  s.Body = v
  return s
}

func (s *ExportWarningResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

