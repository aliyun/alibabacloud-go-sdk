// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAuthorizationRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAuthorizationRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAuthorizationRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAuthorizationRuleResponseBody) *EnableAuthorizationRuleResponse
  GetBody() *EnableAuthorizationRuleResponseBody 
}

type EnableAuthorizationRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAuthorizationRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAuthorizationRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableAuthorizationRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAuthorizationRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAuthorizationRuleResponse) GetBody() *EnableAuthorizationRuleResponseBody  {
  return s.Body
}

func (s *EnableAuthorizationRuleResponse) SetHeaders(v map[string]*string) *EnableAuthorizationRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableAuthorizationRuleResponse) SetStatusCode(v int32) *EnableAuthorizationRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAuthorizationRuleResponse) SetBody(v *EnableAuthorizationRuleResponseBody) *EnableAuthorizationRuleResponse {
  s.Body = v
  return s
}

func (s *EnableAuthorizationRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

