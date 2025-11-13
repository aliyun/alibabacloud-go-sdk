// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnDisableScalingRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnDisableScalingRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnDisableScalingRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EnDisableScalingRulesResponseBody) *EnDisableScalingRulesResponse
  GetBody() *EnDisableScalingRulesResponseBody 
}

type EnDisableScalingRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnDisableScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnDisableScalingRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnDisableScalingRulesResponse) GoString() string {
  return s.String()
}

func (s *EnDisableScalingRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnDisableScalingRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnDisableScalingRulesResponse) GetBody() *EnDisableScalingRulesResponseBody  {
  return s.Body
}

func (s *EnDisableScalingRulesResponse) SetHeaders(v map[string]*string) *EnDisableScalingRulesResponse {
  s.Headers = v
  return s
}

func (s *EnDisableScalingRulesResponse) SetStatusCode(v int32) *EnDisableScalingRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnDisableScalingRulesResponse) SetBody(v *EnDisableScalingRulesResponseBody) *EnDisableScalingRulesResponse {
  s.Body = v
  return s
}

func (s *EnDisableScalingRulesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

