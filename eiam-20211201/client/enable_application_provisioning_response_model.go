// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationProvisioningResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationProvisioningResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationProvisioningResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationProvisioningResponseBody) *EnableApplicationProvisioningResponse
  GetBody() *EnableApplicationProvisioningResponseBody 
}

type EnableApplicationProvisioningResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationProvisioningResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationProvisioningResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationProvisioningResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationProvisioningResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationProvisioningResponse) GetBody() *EnableApplicationProvisioningResponseBody  {
  return s.Body
}

func (s *EnableApplicationProvisioningResponse) SetHeaders(v map[string]*string) *EnableApplicationProvisioningResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationProvisioningResponse) SetStatusCode(v int32) *EnableApplicationProvisioningResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationProvisioningResponse) SetBody(v *EnableApplicationProvisioningResponseBody) *EnableApplicationProvisioningResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationProvisioningResponse) Validate() error {
  return dara.Validate(s)
}

