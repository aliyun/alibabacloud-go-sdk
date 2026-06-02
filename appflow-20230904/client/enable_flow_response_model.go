// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFlowResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableFlowResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableFlowResponse
  GetStatusCode() *int32 
  SetBody(v *EnableFlowResponseBody) *EnableFlowResponse
  GetBody() *EnableFlowResponseBody 
}

type EnableFlowResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFlowResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableFlowResponse) GoString() string {
  return s.String()
}

func (s *EnableFlowResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableFlowResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableFlowResponse) GetBody() *EnableFlowResponseBody  {
  return s.Body
}

func (s *EnableFlowResponse) SetHeaders(v map[string]*string) *EnableFlowResponse {
  s.Headers = v
  return s
}

func (s *EnableFlowResponse) SetStatusCode(v int32) *EnableFlowResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableFlowResponse) SetBody(v *EnableFlowResponseBody) *EnableFlowResponse {
  s.Body = v
  return s
}

func (s *EnableFlowResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

