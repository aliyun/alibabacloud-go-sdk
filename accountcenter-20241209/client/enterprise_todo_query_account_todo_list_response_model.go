// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseTodoQueryAccountTodoListResponseBody) *EnterpriseTodoQueryAccountTodoListResponse
  GetBody() *EnterpriseTodoQueryAccountTodoListResponseBody 
}

type EnterpriseTodoQueryAccountTodoListResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseTodoQueryAccountTodoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) GetBody() *EnterpriseTodoQueryAccountTodoListResponseBody  {
  return s.Body
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetBody(v *EnterpriseTodoQueryAccountTodoListResponseBody) *EnterpriseTodoQueryAccountTodoListResponse {
  s.Body = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

