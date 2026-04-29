// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecallManagementTableResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportRecallManagementTableResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportRecallManagementTableResponse
  GetStatusCode() *int32 
  SetBody(v *ExportRecallManagementTableResponseBody) *ExportRecallManagementTableResponse
  GetBody() *ExportRecallManagementTableResponseBody 
}

type ExportRecallManagementTableResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRecallManagementTableResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportRecallManagementTableResponse) GoString() string {
  return s.String()
}

func (s *ExportRecallManagementTableResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportRecallManagementTableResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportRecallManagementTableResponse) GetBody() *ExportRecallManagementTableResponseBody  {
  return s.Body
}

func (s *ExportRecallManagementTableResponse) SetHeaders(v map[string]*string) *ExportRecallManagementTableResponse {
  s.Headers = v
  return s
}

func (s *ExportRecallManagementTableResponse) SetStatusCode(v int32) *ExportRecallManagementTableResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportRecallManagementTableResponse) SetBody(v *ExportRecallManagementTableResponseBody) *ExportRecallManagementTableResponse {
  s.Body = v
  return s
}

func (s *ExportRecallManagementTableResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

