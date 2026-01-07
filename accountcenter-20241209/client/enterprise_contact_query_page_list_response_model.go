// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryPageListResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseContactQueryPageListResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseContactQueryPageListResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseContactQueryPageListResponseBody) *EnterpriseContactQueryPageListResponse
  GetBody() *EnterpriseContactQueryPageListResponseBody 
}

type EnterpriseContactQueryPageListResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseContactQueryPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseContactQueryPageListResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryPageListResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryPageListResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseContactQueryPageListResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseContactQueryPageListResponse) GetBody() *EnterpriseContactQueryPageListResponseBody  {
  return s.Body
}

func (s *EnterpriseContactQueryPageListResponse) SetHeaders(v map[string]*string) *EnterpriseContactQueryPageListResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseContactQueryPageListResponse) SetStatusCode(v int32) *EnterpriseContactQueryPageListResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponse) SetBody(v *EnterpriseContactQueryPageListResponseBody) *EnterpriseContactQueryPageListResponse {
  s.Body = v
  return s
}

func (s *EnterpriseContactQueryPageListResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

