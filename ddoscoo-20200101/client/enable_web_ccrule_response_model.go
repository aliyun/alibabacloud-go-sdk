// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWebCCRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWebCCRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWebCCRuleResponseBody) *EnableWebCCRuleResponse
  GetBody() *EnableWebCCRuleResponseBody 
}

type EnableWebCCRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWebCCRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableWebCCRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWebCCRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWebCCRuleResponse) GetBody() *EnableWebCCRuleResponseBody  {
  return s.Body
}

func (s *EnableWebCCRuleResponse) SetHeaders(v map[string]*string) *EnableWebCCRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableWebCCRuleResponse) SetStatusCode(v int32) *EnableWebCCRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWebCCRuleResponse) SetBody(v *EnableWebCCRuleResponseBody) *EnableWebCCRuleResponse {
  s.Body = v
  return s
}

func (s *EnableWebCCRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

