// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPptArtifactResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportPptArtifactResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportPptArtifactResponse
  GetStatusCode() *int32 
  SetBody(v *ExportPptArtifactResponseBody) *ExportPptArtifactResponse
  GetBody() *ExportPptArtifactResponseBody 
}

type ExportPptArtifactResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportPptArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportPptArtifactResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportPptArtifactResponse) GoString() string {
  return s.String()
}

func (s *ExportPptArtifactResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportPptArtifactResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportPptArtifactResponse) GetBody() *ExportPptArtifactResponseBody  {
  return s.Body
}

func (s *ExportPptArtifactResponse) SetHeaders(v map[string]*string) *ExportPptArtifactResponse {
  s.Headers = v
  return s
}

func (s *ExportPptArtifactResponse) SetStatusCode(v int32) *ExportPptArtifactResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportPptArtifactResponse) SetBody(v *ExportPptArtifactResponseBody) *ExportPptArtifactResponse {
  s.Body = v
  return s
}

func (s *ExportPptArtifactResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

