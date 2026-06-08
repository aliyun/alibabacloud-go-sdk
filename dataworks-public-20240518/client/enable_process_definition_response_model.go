// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProcessDefinitionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableProcessDefinitionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableProcessDefinitionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableProcessDefinitionResponseBody) *EnableProcessDefinitionResponse
  GetBody() *EnableProcessDefinitionResponseBody 
}

type EnableProcessDefinitionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableProcessDefinitionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableProcessDefinitionResponse) GoString() string {
  return s.String()
}

func (s *EnableProcessDefinitionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableProcessDefinitionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableProcessDefinitionResponse) GetBody() *EnableProcessDefinitionResponseBody  {
  return s.Body
}

func (s *EnableProcessDefinitionResponse) SetHeaders(v map[string]*string) *EnableProcessDefinitionResponse {
  s.Headers = v
  return s
}

func (s *EnableProcessDefinitionResponse) SetStatusCode(v int32) *EnableProcessDefinitionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableProcessDefinitionResponse) SetBody(v *EnableProcessDefinitionResponseBody) *EnableProcessDefinitionResponse {
  s.Body = v
  return s
}

func (s *EnableProcessDefinitionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

