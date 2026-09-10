// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecWorkflowConnectivityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecWorkflowConnectivityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecWorkflowConnectivityResponse
  GetStatusCode() *int32 
  SetBody(v *ExecWorkflowConnectivityResponseBody) *ExecWorkflowConnectivityResponse
  GetBody() *ExecWorkflowConnectivityResponseBody 
}

type ExecWorkflowConnectivityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecWorkflowConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecWorkflowConnectivityResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecWorkflowConnectivityResponse) GoString() string {
  return s.String()
}

func (s *ExecWorkflowConnectivityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecWorkflowConnectivityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecWorkflowConnectivityResponse) GetBody() *ExecWorkflowConnectivityResponseBody  {
  return s.Body
}

func (s *ExecWorkflowConnectivityResponse) SetHeaders(v map[string]*string) *ExecWorkflowConnectivityResponse {
  s.Headers = v
  return s
}

func (s *ExecWorkflowConnectivityResponse) SetStatusCode(v int32) *ExecWorkflowConnectivityResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecWorkflowConnectivityResponse) SetBody(v *ExecWorkflowConnectivityResponseBody) *ExecWorkflowConnectivityResponse {
  s.Body = v
  return s
}

func (s *ExecWorkflowConnectivityResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

