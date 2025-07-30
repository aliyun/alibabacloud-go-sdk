// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConditionalAccessPolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableConditionalAccessPolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableConditionalAccessPolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableConditionalAccessPolicyResponseBody) *EnableConditionalAccessPolicyResponse
  GetBody() *EnableConditionalAccessPolicyResponseBody 
}

type EnableConditionalAccessPolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableConditionalAccessPolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableConditionalAccessPolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableConditionalAccessPolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableConditionalAccessPolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableConditionalAccessPolicyResponse) GetBody() *EnableConditionalAccessPolicyResponseBody  {
  return s.Body
}

func (s *EnableConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *EnableConditionalAccessPolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableConditionalAccessPolicyResponse) SetStatusCode(v int32) *EnableConditionalAccessPolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableConditionalAccessPolicyResponse) SetBody(v *EnableConditionalAccessPolicyResponseBody) *EnableConditionalAccessPolicyResponse {
  s.Body = v
  return s
}

func (s *EnableConditionalAccessPolicyResponse) Validate() error {
  return dara.Validate(s)
}

