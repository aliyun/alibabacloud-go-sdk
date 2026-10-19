// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowTemplateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportWorkflowTemplateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportWorkflowTemplateResponse
  GetStatusCode() *int32 
  SetBody(v *ExportWorkflowTemplateResponseBody) *ExportWorkflowTemplateResponse
  GetBody() *ExportWorkflowTemplateResponseBody 
}

type ExportWorkflowTemplateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportWorkflowTemplateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowTemplateResponse) GoString() string {
  return s.String()
}

func (s *ExportWorkflowTemplateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportWorkflowTemplateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportWorkflowTemplateResponse) GetBody() *ExportWorkflowTemplateResponseBody  {
  return s.Body
}

func (s *ExportWorkflowTemplateResponse) SetHeaders(v map[string]*string) *ExportWorkflowTemplateResponse {
  s.Headers = v
  return s
}

func (s *ExportWorkflowTemplateResponse) SetStatusCode(v int32) *ExportWorkflowTemplateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportWorkflowTemplateResponse) SetBody(v *ExportWorkflowTemplateResponseBody) *ExportWorkflowTemplateResponse {
  s.Body = v
  return s
}

func (s *ExportWorkflowTemplateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

