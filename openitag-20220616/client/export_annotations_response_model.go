// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnnotationsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportAnnotationsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportAnnotationsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportAnnotationsResponseBody) *ExportAnnotationsResponse
  GetBody() *ExportAnnotationsResponseBody 
}

type ExportAnnotationsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportAnnotationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAnnotationsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportAnnotationsResponse) GoString() string {
  return s.String()
}

func (s *ExportAnnotationsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportAnnotationsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportAnnotationsResponse) GetBody() *ExportAnnotationsResponseBody  {
  return s.Body
}

func (s *ExportAnnotationsResponse) SetHeaders(v map[string]*string) *ExportAnnotationsResponse {
  s.Headers = v
  return s
}

func (s *ExportAnnotationsResponse) SetStatusCode(v int32) *ExportAnnotationsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportAnnotationsResponse) SetBody(v *ExportAnnotationsResponseBody) *ExportAnnotationsResponse {
  s.Body = v
  return s
}

func (s *ExportAnnotationsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

