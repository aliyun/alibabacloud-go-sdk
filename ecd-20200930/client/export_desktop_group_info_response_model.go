// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDesktopGroupInfoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportDesktopGroupInfoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportDesktopGroupInfoResponse
  GetStatusCode() *int32 
  SetBody(v *ExportDesktopGroupInfoResponseBody) *ExportDesktopGroupInfoResponse
  GetBody() *ExportDesktopGroupInfoResponseBody 
}

type ExportDesktopGroupInfoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportDesktopGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDesktopGroupInfoResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportDesktopGroupInfoResponse) GoString() string {
  return s.String()
}

func (s *ExportDesktopGroupInfoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportDesktopGroupInfoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportDesktopGroupInfoResponse) GetBody() *ExportDesktopGroupInfoResponseBody  {
  return s.Body
}

func (s *ExportDesktopGroupInfoResponse) SetHeaders(v map[string]*string) *ExportDesktopGroupInfoResponse {
  s.Headers = v
  return s
}

func (s *ExportDesktopGroupInfoResponse) SetStatusCode(v int32) *ExportDesktopGroupInfoResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportDesktopGroupInfoResponse) SetBody(v *ExportDesktopGroupInfoResponseBody) *ExportDesktopGroupInfoResponse {
  s.Body = v
  return s
}

func (s *ExportDesktopGroupInfoResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

