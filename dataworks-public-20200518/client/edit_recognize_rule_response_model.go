// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRecognizeRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditRecognizeRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditRecognizeRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EditRecognizeRuleResponseBody) *EditRecognizeRuleResponse
  GetBody() *EditRecognizeRuleResponseBody 
}

type EditRecognizeRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditRecognizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditRecognizeRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EditRecognizeRuleResponse) GoString() string {
  return s.String()
}

func (s *EditRecognizeRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditRecognizeRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditRecognizeRuleResponse) GetBody() *EditRecognizeRuleResponseBody  {
  return s.Body
}

func (s *EditRecognizeRuleResponse) SetHeaders(v map[string]*string) *EditRecognizeRuleResponse {
  s.Headers = v
  return s
}

func (s *EditRecognizeRuleResponse) SetStatusCode(v int32) *EditRecognizeRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EditRecognizeRuleResponse) SetBody(v *EditRecognizeRuleResponseBody) *EditRecognizeRuleResponse {
  s.Body = v
  return s
}

func (s *EditRecognizeRuleResponse) Validate() error {
  return dara.Validate(s)
}

