// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCrossAccountManagementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCrossAccountManagementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCrossAccountManagementResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCrossAccountManagementResponseBody) *EnableCrossAccountManagementResponse
  GetBody() *EnableCrossAccountManagementResponseBody 
}

type EnableCrossAccountManagementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCrossAccountManagementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCrossAccountManagementResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCrossAccountManagementResponse) GoString() string {
  return s.String()
}

func (s *EnableCrossAccountManagementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCrossAccountManagementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCrossAccountManagementResponse) GetBody() *EnableCrossAccountManagementResponseBody  {
  return s.Body
}

func (s *EnableCrossAccountManagementResponse) SetHeaders(v map[string]*string) *EnableCrossAccountManagementResponse {
  s.Headers = v
  return s
}

func (s *EnableCrossAccountManagementResponse) SetStatusCode(v int32) *EnableCrossAccountManagementResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCrossAccountManagementResponse) SetBody(v *EnableCrossAccountManagementResponseBody) *EnableCrossAccountManagementResponse {
  s.Body = v
  return s
}

func (s *EnableCrossAccountManagementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

