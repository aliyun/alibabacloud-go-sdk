// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneDefensePolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSceneDefensePolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSceneDefensePolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSceneDefensePolicyResponseBody) *EnableSceneDefensePolicyResponse
  GetBody() *EnableSceneDefensePolicyResponseBody 
}

type EnableSceneDefensePolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSceneDefensePolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneDefensePolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableSceneDefensePolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSceneDefensePolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSceneDefensePolicyResponse) GetBody() *EnableSceneDefensePolicyResponseBody  {
  return s.Body
}

func (s *EnableSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *EnableSceneDefensePolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableSceneDefensePolicyResponse) SetStatusCode(v int32) *EnableSceneDefensePolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSceneDefensePolicyResponse) SetBody(v *EnableSceneDefensePolicyResponseBody) *EnableSceneDefensePolicyResponse {
  s.Body = v
  return s
}

func (s *EnableSceneDefensePolicyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

