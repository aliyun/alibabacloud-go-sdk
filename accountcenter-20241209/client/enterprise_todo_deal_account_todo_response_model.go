// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoDealAccountTodoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseTodoDealAccountTodoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseTodoDealAccountTodoResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseTodoDealAccountTodoResponseBody) *EnterpriseTodoDealAccountTodoResponse
  GetBody() *EnterpriseTodoDealAccountTodoResponseBody 
}

type EnterpriseTodoDealAccountTodoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseTodoDealAccountTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoDealAccountTodoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseTodoDealAccountTodoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseTodoDealAccountTodoResponse) GetBody() *EnterpriseTodoDealAccountTodoResponseBody  {
  return s.Body
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetHeaders(v map[string]*string) *EnterpriseTodoDealAccountTodoResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetStatusCode(v int32) *EnterpriseTodoDealAccountTodoResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetBody(v *EnterpriseTodoDealAccountTodoResponseBody) *EnterpriseTodoDealAccountTodoResponse {
  s.Body = v
  return s
}

func (s *EnterpriseTodoDealAccountTodoResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

