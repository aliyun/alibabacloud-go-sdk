// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOssCheckStatResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportOssCheckStatResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportOssCheckStatResponse
  GetStatusCode() *int32 
  SetBody(v *ExportOssCheckStatResponseBody) *ExportOssCheckStatResponse
  GetBody() *ExportOssCheckStatResponseBody 
}

type ExportOssCheckStatResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportOssCheckStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportOssCheckStatResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportOssCheckStatResponse) GoString() string {
  return s.String()
}

func (s *ExportOssCheckStatResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportOssCheckStatResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportOssCheckStatResponse) GetBody() *ExportOssCheckStatResponseBody  {
  return s.Body
}

func (s *ExportOssCheckStatResponse) SetHeaders(v map[string]*string) *ExportOssCheckStatResponse {
  s.Headers = v
  return s
}

func (s *ExportOssCheckStatResponse) SetStatusCode(v int32) *ExportOssCheckStatResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportOssCheckStatResponse) SetBody(v *ExportOssCheckStatResponseBody) *ExportOssCheckStatResponse {
  s.Body = v
  return s
}

func (s *ExportOssCheckStatResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

