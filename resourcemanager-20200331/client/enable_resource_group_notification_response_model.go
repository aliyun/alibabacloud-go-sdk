// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceGroupNotificationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableResourceGroupNotificationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableResourceGroupNotificationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableResourceGroupNotificationResponseBody) *EnableResourceGroupNotificationResponse
  GetBody() *EnableResourceGroupNotificationResponseBody 
}

type EnableResourceGroupNotificationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableResourceGroupNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceGroupNotificationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceGroupNotificationResponse) GoString() string {
  return s.String()
}

func (s *EnableResourceGroupNotificationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableResourceGroupNotificationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableResourceGroupNotificationResponse) GetBody() *EnableResourceGroupNotificationResponseBody  {
  return s.Body
}

func (s *EnableResourceGroupNotificationResponse) SetHeaders(v map[string]*string) *EnableResourceGroupNotificationResponse {
  s.Headers = v
  return s
}

func (s *EnableResourceGroupNotificationResponse) SetStatusCode(v int32) *EnableResourceGroupNotificationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableResourceGroupNotificationResponse) SetBody(v *EnableResourceGroupNotificationResponseBody) *EnableResourceGroupNotificationResponse {
  s.Body = v
  return s
}

func (s *EnableResourceGroupNotificationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

