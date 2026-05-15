// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditQualityRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditQualityRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EditQualityRuleResponseBody) *EditQualityRuleResponse
  GetBody() *EditQualityRuleResponseBody 
}

type EditQualityRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditQualityRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleResponse) GoString() string {
  return s.String()
}

func (s *EditQualityRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditQualityRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditQualityRuleResponse) GetBody() *EditQualityRuleResponseBody  {
  return s.Body
}

func (s *EditQualityRuleResponse) SetHeaders(v map[string]*string) *EditQualityRuleResponse {
  s.Headers = v
  return s
}

func (s *EditQualityRuleResponse) SetStatusCode(v int32) *EditQualityRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EditQualityRuleResponse) SetBody(v *EditQualityRuleResponseBody) *EditQualityRuleResponse {
  s.Body = v
  return s
}

func (s *EditQualityRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

