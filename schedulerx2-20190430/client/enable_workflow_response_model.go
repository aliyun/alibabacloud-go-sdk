// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWorkflowResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWorkflowResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWorkflowResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWorkflowResponseBody) *EnableWorkflowResponse
  GetBody() *EnableWorkflowResponseBody 
}

type EnableWorkflowResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWorkflowResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWorkflowResponse) GoString() string {
  return s.String()
}

func (s *EnableWorkflowResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWorkflowResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWorkflowResponse) GetBody() *EnableWorkflowResponseBody  {
  return s.Body
}

func (s *EnableWorkflowResponse) SetHeaders(v map[string]*string) *EnableWorkflowResponse {
  s.Headers = v
  return s
}

func (s *EnableWorkflowResponse) SetStatusCode(v int32) *EnableWorkflowResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWorkflowResponse) SetBody(v *EnableWorkflowResponseBody) *EnableWorkflowResponse {
  s.Body = v
  return s
}

func (s *EnableWorkflowResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

