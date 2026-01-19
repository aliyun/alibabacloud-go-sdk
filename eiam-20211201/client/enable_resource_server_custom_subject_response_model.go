// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceServerCustomSubjectResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableResourceServerCustomSubjectResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableResourceServerCustomSubjectResponse
  GetStatusCode() *int32 
  SetBody(v *EnableResourceServerCustomSubjectResponseBody) *EnableResourceServerCustomSubjectResponse
  GetBody() *EnableResourceServerCustomSubjectResponseBody 
}

type EnableResourceServerCustomSubjectResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableResourceServerCustomSubjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceServerCustomSubjectResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceServerCustomSubjectResponse) GoString() string {
  return s.String()
}

func (s *EnableResourceServerCustomSubjectResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableResourceServerCustomSubjectResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableResourceServerCustomSubjectResponse) GetBody() *EnableResourceServerCustomSubjectResponseBody  {
  return s.Body
}

func (s *EnableResourceServerCustomSubjectResponse) SetHeaders(v map[string]*string) *EnableResourceServerCustomSubjectResponse {
  s.Headers = v
  return s
}

func (s *EnableResourceServerCustomSubjectResponse) SetStatusCode(v int32) *EnableResourceServerCustomSubjectResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableResourceServerCustomSubjectResponse) SetBody(v *EnableResourceServerCustomSubjectResponseBody) *EnableResourceServerCustomSubjectResponse {
  s.Body = v
  return s
}

func (s *EnableResourceServerCustomSubjectResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

