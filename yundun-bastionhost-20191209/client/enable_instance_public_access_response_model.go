// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstancePublicAccessResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInstancePublicAccessResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInstancePublicAccessResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInstancePublicAccessResponseBody) *EnableInstancePublicAccessResponse
  GetBody() *EnableInstancePublicAccessResponseBody 
}

type EnableInstancePublicAccessResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstancePublicAccessResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInstancePublicAccessResponse) GoString() string {
  return s.String()
}

func (s *EnableInstancePublicAccessResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInstancePublicAccessResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInstancePublicAccessResponse) GetBody() *EnableInstancePublicAccessResponseBody  {
  return s.Body
}

func (s *EnableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *EnableInstancePublicAccessResponse {
  s.Headers = v
  return s
}

func (s *EnableInstancePublicAccessResponse) SetStatusCode(v int32) *EnableInstancePublicAccessResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInstancePublicAccessResponse) SetBody(v *EnableInstancePublicAccessResponseBody) *EnableInstancePublicAccessResponse {
  s.Body = v
  return s
}

func (s *EnableInstancePublicAccessResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

