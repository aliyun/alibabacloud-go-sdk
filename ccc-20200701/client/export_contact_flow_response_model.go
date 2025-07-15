// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportContactFlowResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportContactFlowResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportContactFlowResponse
  GetStatusCode() *int32 
  SetBody(v *ExportContactFlowResponseBody) *ExportContactFlowResponse
  GetBody() *ExportContactFlowResponseBody 
}

type ExportContactFlowResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportContactFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportContactFlowResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportContactFlowResponse) GoString() string {
  return s.String()
}

func (s *ExportContactFlowResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportContactFlowResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportContactFlowResponse) GetBody() *ExportContactFlowResponseBody  {
  return s.Body
}

func (s *ExportContactFlowResponse) SetHeaders(v map[string]*string) *ExportContactFlowResponse {
  s.Headers = v
  return s
}

func (s *ExportContactFlowResponse) SetStatusCode(v int32) *ExportContactFlowResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportContactFlowResponse) SetBody(v *ExportContactFlowResponseBody) *ExportContactFlowResponse {
  s.Body = v
  return s
}

func (s *ExportContactFlowResponse) Validate() error {
  return dara.Validate(s)
}

