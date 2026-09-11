// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMaintainWindowResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableMaintainWindowResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableMaintainWindowResponse
  GetStatusCode() *int32 
  SetBody(v *EnableMaintainWindowResponseBody) *EnableMaintainWindowResponse
  GetBody() *EnableMaintainWindowResponseBody 
}

type EnableMaintainWindowResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMaintainWindowResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableMaintainWindowResponse) GoString() string {
  return s.String()
}

func (s *EnableMaintainWindowResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableMaintainWindowResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableMaintainWindowResponse) GetBody() *EnableMaintainWindowResponseBody  {
  return s.Body
}

func (s *EnableMaintainWindowResponse) SetHeaders(v map[string]*string) *EnableMaintainWindowResponse {
  s.Headers = v
  return s
}

func (s *EnableMaintainWindowResponse) SetStatusCode(v int32) *EnableMaintainWindowResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableMaintainWindowResponse) SetBody(v *EnableMaintainWindowResponseBody) *EnableMaintainWindowResponse {
  s.Body = v
  return s
}

func (s *EnableMaintainWindowResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

