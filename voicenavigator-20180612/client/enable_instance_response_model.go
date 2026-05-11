// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInstanceResponseBody) *EnableInstanceResponse
  GetBody() *EnableInstanceResponseBody 
}

type EnableInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceResponse) GoString() string {
  return s.String()
}

func (s *EnableInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInstanceResponse) GetBody() *EnableInstanceResponseBody  {
  return s.Body
}

func (s *EnableInstanceResponse) SetHeaders(v map[string]*string) *EnableInstanceResponse {
  s.Headers = v
  return s
}

func (s *EnableInstanceResponse) SetStatusCode(v int32) *EnableInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInstanceResponse) SetBody(v *EnableInstanceResponseBody) *EnableInstanceResponse {
  s.Body = v
  return s
}

func (s *EnableInstanceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

