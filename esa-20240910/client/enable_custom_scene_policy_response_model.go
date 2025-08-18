// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomScenePolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCustomScenePolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCustomScenePolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCustomScenePolicyResponseBody) *EnableCustomScenePolicyResponse
  GetBody() *EnableCustomScenePolicyResponseBody 
}

type EnableCustomScenePolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomScenePolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomScenePolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableCustomScenePolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCustomScenePolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCustomScenePolicyResponse) GetBody() *EnableCustomScenePolicyResponseBody  {
  return s.Body
}

func (s *EnableCustomScenePolicyResponse) SetHeaders(v map[string]*string) *EnableCustomScenePolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableCustomScenePolicyResponse) SetStatusCode(v int32) *EnableCustomScenePolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCustomScenePolicyResponse) SetBody(v *EnableCustomScenePolicyResponseBody) *EnableCustomScenePolicyResponse {
  s.Body = v
  return s
}

func (s *EnableCustomScenePolicyResponse) Validate() error {
  return dara.Validate(s)
}

