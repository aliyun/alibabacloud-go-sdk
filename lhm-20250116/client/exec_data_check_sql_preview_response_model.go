// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckSqlPreviewResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckSqlPreviewResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckSqlPreviewResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckSqlPreviewResponseBody) *ExecDataCheckSqlPreviewResponse
  GetBody() *ExecDataCheckSqlPreviewResponseBody 
}

type ExecDataCheckSqlPreviewResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckSqlPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckSqlPreviewResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckSqlPreviewResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckSqlPreviewResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckSqlPreviewResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckSqlPreviewResponse) GetBody() *ExecDataCheckSqlPreviewResponseBody  {
  return s.Body
}

func (s *ExecDataCheckSqlPreviewResponse) SetHeaders(v map[string]*string) *ExecDataCheckSqlPreviewResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckSqlPreviewResponse) SetStatusCode(v int32) *ExecDataCheckSqlPreviewResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckSqlPreviewResponse) SetBody(v *ExecDataCheckSqlPreviewResponseBody) *ExecDataCheckSqlPreviewResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckSqlPreviewResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

