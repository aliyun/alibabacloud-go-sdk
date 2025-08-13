// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolicyTypeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnablePolicyTypeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnablePolicyTypeResponse
  GetStatusCode() *int32 
  SetBody(v *EnablePolicyTypeResponseBody) *EnablePolicyTypeResponse
  GetBody() *EnablePolicyTypeResponseBody 
}

type EnablePolicyTypeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnablePolicyTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePolicyTypeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnablePolicyTypeResponse) GoString() string {
  return s.String()
}

func (s *EnablePolicyTypeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnablePolicyTypeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnablePolicyTypeResponse) GetBody() *EnablePolicyTypeResponseBody  {
  return s.Body
}

func (s *EnablePolicyTypeResponse) SetHeaders(v map[string]*string) *EnablePolicyTypeResponse {
  s.Headers = v
  return s
}

func (s *EnablePolicyTypeResponse) SetStatusCode(v int32) *EnablePolicyTypeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnablePolicyTypeResponse) SetBody(v *EnablePolicyTypeResponseBody) *EnablePolicyTypeResponse {
  s.Body = v
  return s
}

func (s *EnablePolicyTypeResponse) Validate() error {
  return dara.Validate(s)
}

