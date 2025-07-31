// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKeywordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportKeywordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportKeywordResponse
  GetStatusCode() *int32 
  SetBody(v *ExportKeywordResponseBody) *ExportKeywordResponse
  GetBody() *ExportKeywordResponseBody 
}

type ExportKeywordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportKeywordResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportKeywordResponse) GoString() string {
  return s.String()
}

func (s *ExportKeywordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportKeywordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportKeywordResponse) GetBody() *ExportKeywordResponseBody  {
  return s.Body
}

func (s *ExportKeywordResponse) SetHeaders(v map[string]*string) *ExportKeywordResponse {
  s.Headers = v
  return s
}

func (s *ExportKeywordResponse) SetStatusCode(v int32) *ExportKeywordResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportKeywordResponse) SetBody(v *ExportKeywordResponseBody) *ExportKeywordResponse {
  s.Body = v
  return s
}

func (s *ExportKeywordResponse) Validate() error {
  return dara.Validate(s)
}

