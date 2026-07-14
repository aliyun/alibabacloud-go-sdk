// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNotifyPolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableNotifyPolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableNotifyPolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableNotifyPolicyResponseBody) *EnableNotifyPolicyResponse
  GetBody() *EnableNotifyPolicyResponseBody 
}

type EnableNotifyPolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableNotifyPolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableNotifyPolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableNotifyPolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableNotifyPolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableNotifyPolicyResponse) GetBody() *EnableNotifyPolicyResponseBody  {
  return s.Body
}

func (s *EnableNotifyPolicyResponse) SetHeaders(v map[string]*string) *EnableNotifyPolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableNotifyPolicyResponse) SetStatusCode(v int32) *EnableNotifyPolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableNotifyPolicyResponse) SetBody(v *EnableNotifyPolicyResponseBody) *EnableNotifyPolicyResponse {
  s.Body = v
  return s
}

func (s *EnableNotifyPolicyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

