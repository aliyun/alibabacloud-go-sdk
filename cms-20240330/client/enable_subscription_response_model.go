// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSubscriptionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSubscriptionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSubscriptionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSubscriptionResponseBody) *EnableSubscriptionResponse
  GetBody() *EnableSubscriptionResponseBody 
}

type EnableSubscriptionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSubscriptionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSubscriptionResponse) GoString() string {
  return s.String()
}

func (s *EnableSubscriptionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSubscriptionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSubscriptionResponse) GetBody() *EnableSubscriptionResponseBody  {
  return s.Body
}

func (s *EnableSubscriptionResponse) SetHeaders(v map[string]*string) *EnableSubscriptionResponse {
  s.Headers = v
  return s
}

func (s *EnableSubscriptionResponse) SetStatusCode(v int32) *EnableSubscriptionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSubscriptionResponse) SetBody(v *EnableSubscriptionResponseBody) *EnableSubscriptionResponse {
  s.Body = v
  return s
}

func (s *EnableSubscriptionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

