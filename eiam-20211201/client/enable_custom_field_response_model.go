// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomFieldResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCustomFieldResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCustomFieldResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCustomFieldResponseBody) *EnableCustomFieldResponse
  GetBody() *EnableCustomFieldResponseBody 
}

type EnableCustomFieldResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCustomFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomFieldResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomFieldResponse) GoString() string {
  return s.String()
}

func (s *EnableCustomFieldResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCustomFieldResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCustomFieldResponse) GetBody() *EnableCustomFieldResponseBody  {
  return s.Body
}

func (s *EnableCustomFieldResponse) SetHeaders(v map[string]*string) *EnableCustomFieldResponse {
  s.Headers = v
  return s
}

func (s *EnableCustomFieldResponse) SetStatusCode(v int32) *EnableCustomFieldResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCustomFieldResponse) SetBody(v *EnableCustomFieldResponseBody) *EnableCustomFieldResponse {
  s.Body = v
  return s
}

func (s *EnableCustomFieldResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

