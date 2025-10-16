// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopListInfoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportDesktopListInfoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportDesktopListInfoResponse
  GetStatusCode() *int32 
  SetBody(v *ExportDesktopListInfoResponseBody) *ExportDesktopListInfoResponse
  GetBody() *ExportDesktopListInfoResponseBody 
}

type ExportDesktopListInfoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportDesktopListInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDesktopListInfoResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopListInfoResponse) GoString() string {
  return s.String()
}

func (s *ExportDesktopListInfoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportDesktopListInfoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportDesktopListInfoResponse) GetBody() *ExportDesktopListInfoResponseBody  {
  return s.Body
}

func (s *ExportDesktopListInfoResponse) SetHeaders(v map[string]*string) *ExportDesktopListInfoResponse {
  s.Headers = v
  return s
}

func (s *ExportDesktopListInfoResponse) SetStatusCode(v int32) *ExportDesktopListInfoResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportDesktopListInfoResponse) SetBody(v *ExportDesktopListInfoResponseBody) *ExportDesktopListInfoResponse {
  s.Body = v
  return s
}

func (s *ExportDesktopListInfoResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

