// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportWorkflowsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportWorkflowsResponse
  GetStatusCode() *int32 
  SetBody(v []byte) *ExportWorkflowsResponse
  GetBody() []byte 
}

type ExportWorkflowsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body []byte `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportWorkflowsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowsResponse) GoString() string {
  return s.String()
}

func (s *ExportWorkflowsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportWorkflowsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportWorkflowsResponse) GetBody() []byte  {
  return s.Body
}

func (s *ExportWorkflowsResponse) SetHeaders(v map[string]*string) *ExportWorkflowsResponse {
  s.Headers = v
  return s
}

func (s *ExportWorkflowsResponse) SetStatusCode(v int32) *ExportWorkflowsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportWorkflowsResponse) SetBody(v []byte) *ExportWorkflowsResponse {
  s.Body = v
  return s
}

func (s *ExportWorkflowsResponse) Validate() error {
  return dara.Validate(s)
}

