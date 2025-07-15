// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDoNotCallNumbersResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportDoNotCallNumbersResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportDoNotCallNumbersResponse
  GetStatusCode() *int32 
  SetBody(v *ExportDoNotCallNumbersResponseBody) *ExportDoNotCallNumbersResponse
  GetBody() *ExportDoNotCallNumbersResponseBody 
}

type ExportDoNotCallNumbersResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportDoNotCallNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDoNotCallNumbersResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportDoNotCallNumbersResponse) GoString() string {
  return s.String()
}

func (s *ExportDoNotCallNumbersResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportDoNotCallNumbersResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportDoNotCallNumbersResponse) GetBody() *ExportDoNotCallNumbersResponseBody  {
  return s.Body
}

func (s *ExportDoNotCallNumbersResponse) SetHeaders(v map[string]*string) *ExportDoNotCallNumbersResponse {
  s.Headers = v
  return s
}

func (s *ExportDoNotCallNumbersResponse) SetStatusCode(v int32) *ExportDoNotCallNumbersResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportDoNotCallNumbersResponse) SetBody(v *ExportDoNotCallNumbersResponseBody) *ExportDoNotCallNumbersResponse {
  s.Body = v
  return s
}

func (s *ExportDoNotCallNumbersResponse) Validate() error {
  return dara.Validate(s)
}

