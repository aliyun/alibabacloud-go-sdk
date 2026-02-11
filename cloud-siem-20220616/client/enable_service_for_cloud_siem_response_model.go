// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceForCloudSiemResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServiceForCloudSiemResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServiceForCloudSiemResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServiceForCloudSiemResponseBody) *EnableServiceForCloudSiemResponse
  GetBody() *EnableServiceForCloudSiemResponseBody 
}

type EnableServiceForCloudSiemResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServiceForCloudSiemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceForCloudSiemResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceForCloudSiemResponse) GoString() string {
  return s.String()
}

func (s *EnableServiceForCloudSiemResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServiceForCloudSiemResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServiceForCloudSiemResponse) GetBody() *EnableServiceForCloudSiemResponseBody  {
  return s.Body
}

func (s *EnableServiceForCloudSiemResponse) SetHeaders(v map[string]*string) *EnableServiceForCloudSiemResponse {
  s.Headers = v
  return s
}

func (s *EnableServiceForCloudSiemResponse) SetStatusCode(v int32) *EnableServiceForCloudSiemResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServiceForCloudSiemResponse) SetBody(v *EnableServiceForCloudSiemResponseBody) *EnableServiceForCloudSiemResponse {
  s.Body = v
  return s
}

func (s *EnableServiceForCloudSiemResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

