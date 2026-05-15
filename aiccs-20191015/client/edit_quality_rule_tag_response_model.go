// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleTagResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditQualityRuleTagResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditQualityRuleTagResponse
  GetStatusCode() *int32 
  SetBody(v *EditQualityRuleTagResponseBody) *EditQualityRuleTagResponse
  GetBody() *EditQualityRuleTagResponseBody 
}

type EditQualityRuleTagResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditQualityRuleTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditQualityRuleTagResponse) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleTagResponse) GoString() string {
  return s.String()
}

func (s *EditQualityRuleTagResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditQualityRuleTagResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditQualityRuleTagResponse) GetBody() *EditQualityRuleTagResponseBody  {
  return s.Body
}

func (s *EditQualityRuleTagResponse) SetHeaders(v map[string]*string) *EditQualityRuleTagResponse {
  s.Headers = v
  return s
}

func (s *EditQualityRuleTagResponse) SetStatusCode(v int32) *EditQualityRuleTagResponse {
  s.StatusCode = &v
  return s
}

func (s *EditQualityRuleTagResponse) SetBody(v *EditQualityRuleTagResponseBody) *EditQualityRuleTagResponse {
  s.Body = v
  return s
}

func (s *EditQualityRuleTagResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

