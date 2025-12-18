// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluatePreConfigRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluatePreConfigRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluatePreConfigRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluatePreConfigRulesResponseBody) *EvaluatePreConfigRulesResponse
  GetBody() *EvaluatePreConfigRulesResponseBody 
}

type EvaluatePreConfigRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluatePreConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluatePreConfigRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesResponse) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluatePreConfigRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluatePreConfigRulesResponse) GetBody() *EvaluatePreConfigRulesResponseBody  {
  return s.Body
}

func (s *EvaluatePreConfigRulesResponse) SetHeaders(v map[string]*string) *EvaluatePreConfigRulesResponse {
  s.Headers = v
  return s
}

func (s *EvaluatePreConfigRulesResponse) SetStatusCode(v int32) *EvaluatePreConfigRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluatePreConfigRulesResponse) SetBody(v *EvaluatePreConfigRulesResponseBody) *EvaluatePreConfigRulesResponse {
  s.Body = v
  return s
}

func (s *EvaluatePreConfigRulesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

