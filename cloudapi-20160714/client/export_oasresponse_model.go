// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOASResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportOASResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportOASResponse
  GetStatusCode() *int32 
  SetBody(v *ExportOASResponseBody) *ExportOASResponse
  GetBody() *ExportOASResponseBody 
}

type ExportOASResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportOASResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportOASResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportOASResponse) GoString() string {
  return s.String()
}

func (s *ExportOASResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportOASResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportOASResponse) GetBody() *ExportOASResponseBody  {
  return s.Body
}

func (s *ExportOASResponse) SetHeaders(v map[string]*string) *ExportOASResponse {
  s.Headers = v
  return s
}

func (s *ExportOASResponse) SetStatusCode(v int32) *ExportOASResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportOASResponse) SetBody(v *ExportOASResponseBody) *ExportOASResponse {
  s.Body = v
  return s
}

func (s *ExportOASResponse) Validate() error {
  return dara.Validate(s)
}

