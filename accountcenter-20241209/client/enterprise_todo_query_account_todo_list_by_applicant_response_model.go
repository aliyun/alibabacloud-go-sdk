// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListByApplicantResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListByApplicantResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) *EnterpriseTodoQueryAccountTodoListByApplicantResponse
  GetBody() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody 
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) GetBody() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody  {
  return s.Body
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetBody(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
  s.Body = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

