// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScriptResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportScriptResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportScriptResponse
  GetStatusCode() *int32 
  SetBody(v *ExportScriptResponseBody) *ExportScriptResponse
  GetBody() *ExportScriptResponseBody 
}

type ExportScriptResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportScriptResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportScriptResponse) GoString() string {
  return s.String()
}

func (s *ExportScriptResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportScriptResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportScriptResponse) GetBody() *ExportScriptResponseBody  {
  return s.Body
}

func (s *ExportScriptResponse) SetHeaders(v map[string]*string) *ExportScriptResponse {
  s.Headers = v
  return s
}

func (s *ExportScriptResponse) SetStatusCode(v int32) *ExportScriptResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportScriptResponse) SetBody(v *ExportScriptResponseBody) *ExportScriptResponse {
  s.Body = v
  return s
}

func (s *ExportScriptResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

