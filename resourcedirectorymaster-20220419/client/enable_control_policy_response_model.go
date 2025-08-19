// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableControlPolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableControlPolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableControlPolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableControlPolicyResponseBody) *EnableControlPolicyResponse
  GetBody() *EnableControlPolicyResponseBody 
}

type EnableControlPolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableControlPolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableControlPolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableControlPolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableControlPolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableControlPolicyResponse) GetBody() *EnableControlPolicyResponseBody  {
  return s.Body
}

func (s *EnableControlPolicyResponse) SetHeaders(v map[string]*string) *EnableControlPolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableControlPolicyResponse) SetStatusCode(v int32) *EnableControlPolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableControlPolicyResponse) SetBody(v *EnableControlPolicyResponseBody) *EnableControlPolicyResponse {
  s.Body = v
  return s
}

func (s *EnableControlPolicyResponse) Validate() error {
  return dara.Validate(s)
}

