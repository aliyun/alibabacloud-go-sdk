// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEventRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableEventRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableEventRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableEventRulesResponseBody) *EnableEventRulesResponse
  GetBody() *EnableEventRulesResponseBody 
}

type EnableEventRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableEventRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableEventRulesResponse) GoString() string {
  return s.String()
}

func (s *EnableEventRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableEventRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableEventRulesResponse) GetBody() *EnableEventRulesResponseBody  {
  return s.Body
}

func (s *EnableEventRulesResponse) SetHeaders(v map[string]*string) *EnableEventRulesResponse {
  s.Headers = v
  return s
}

func (s *EnableEventRulesResponse) SetStatusCode(v int32) *EnableEventRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableEventRulesResponse) SetBody(v *EnableEventRulesResponseBody) *EnableEventRulesResponse {
  s.Body = v
  return s
}

func (s *EnableEventRulesResponse) Validate() error {
  return dara.Validate(s)
}

