// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomPrivacyPolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCustomPrivacyPolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCustomPrivacyPolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCustomPrivacyPolicyResponseBody) *EnableCustomPrivacyPolicyResponse
  GetBody() *EnableCustomPrivacyPolicyResponseBody 
}

type EnableCustomPrivacyPolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomPrivacyPolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomPrivacyPolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableCustomPrivacyPolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCustomPrivacyPolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCustomPrivacyPolicyResponse) GetBody() *EnableCustomPrivacyPolicyResponseBody  {
  return s.Body
}

func (s *EnableCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *EnableCustomPrivacyPolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableCustomPrivacyPolicyResponse) SetStatusCode(v int32) *EnableCustomPrivacyPolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCustomPrivacyPolicyResponse) SetBody(v *EnableCustomPrivacyPolicyResponseBody) *EnableCustomPrivacyPolicyResponse {
  s.Body = v
  return s
}

func (s *EnableCustomPrivacyPolicyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

