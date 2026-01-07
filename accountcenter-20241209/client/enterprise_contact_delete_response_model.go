// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactDeleteResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseContactDeleteResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseContactDeleteResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseContactDeleteResponseBody) *EnterpriseContactDeleteResponse
  GetBody() *EnterpriseContactDeleteResponseBody 
}

type EnterpriseContactDeleteResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseContactDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseContactDeleteResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactDeleteResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseContactDeleteResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseContactDeleteResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseContactDeleteResponse) GetBody() *EnterpriseContactDeleteResponseBody  {
  return s.Body
}

func (s *EnterpriseContactDeleteResponse) SetHeaders(v map[string]*string) *EnterpriseContactDeleteResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseContactDeleteResponse) SetStatusCode(v int32) *EnterpriseContactDeleteResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseContactDeleteResponse) SetBody(v *EnterpriseContactDeleteResponseBody) *EnterpriseContactDeleteResponse {
  s.Body = v
  return s
}

func (s *EnterpriseContactDeleteResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

