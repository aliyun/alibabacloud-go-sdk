// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEnterpriseAcceleratePolicyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableEnterpriseAcceleratePolicyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableEnterpriseAcceleratePolicyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableEnterpriseAcceleratePolicyResponseBody) *EnableEnterpriseAcceleratePolicyResponse
  GetBody() *EnableEnterpriseAcceleratePolicyResponseBody 
}

type EnableEnterpriseAcceleratePolicyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableEnterpriseAcceleratePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableEnterpriseAcceleratePolicyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableEnterpriseAcceleratePolicyResponse) GoString() string {
  return s.String()
}

func (s *EnableEnterpriseAcceleratePolicyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableEnterpriseAcceleratePolicyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableEnterpriseAcceleratePolicyResponse) GetBody() *EnableEnterpriseAcceleratePolicyResponseBody  {
  return s.Body
}

func (s *EnableEnterpriseAcceleratePolicyResponse) SetHeaders(v map[string]*string) *EnableEnterpriseAcceleratePolicyResponse {
  s.Headers = v
  return s
}

func (s *EnableEnterpriseAcceleratePolicyResponse) SetStatusCode(v int32) *EnableEnterpriseAcceleratePolicyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableEnterpriseAcceleratePolicyResponse) SetBody(v *EnableEnterpriseAcceleratePolicyResponseBody) *EnableEnterpriseAcceleratePolicyResponse {
  s.Body = v
  return s
}

func (s *EnableEnterpriseAcceleratePolicyResponse) Validate() error {
  return dara.Validate(s)
}

