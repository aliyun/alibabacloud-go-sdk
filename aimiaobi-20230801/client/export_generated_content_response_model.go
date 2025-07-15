// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGeneratedContentResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportGeneratedContentResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportGeneratedContentResponse
  GetStatusCode() *int32 
  SetBody(v *ExportGeneratedContentResponseBody) *ExportGeneratedContentResponse
  GetBody() *ExportGeneratedContentResponseBody 
}

type ExportGeneratedContentResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportGeneratedContentResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportGeneratedContentResponse) GoString() string {
  return s.String()
}

func (s *ExportGeneratedContentResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportGeneratedContentResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportGeneratedContentResponse) GetBody() *ExportGeneratedContentResponseBody  {
  return s.Body
}

func (s *ExportGeneratedContentResponse) SetHeaders(v map[string]*string) *ExportGeneratedContentResponse {
  s.Headers = v
  return s
}

func (s *ExportGeneratedContentResponse) SetStatusCode(v int32) *ExportGeneratedContentResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportGeneratedContentResponse) SetBody(v *ExportGeneratedContentResponseBody) *ExportGeneratedContentResponse {
  s.Body = v
  return s
}

func (s *ExportGeneratedContentResponse) Validate() error {
  return dara.Validate(s)
}

