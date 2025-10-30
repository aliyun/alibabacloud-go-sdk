// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpressionTestResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpressionTestResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpressionTestResponse
  GetStatusCode() *int32 
  SetBody(v *ExpressionTestResponseBody) *ExpressionTestResponse
  GetBody() *ExpressionTestResponseBody 
}

type ExpressionTestResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExpressionTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpressionTestResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpressionTestResponse) GoString() string {
  return s.String()
}

func (s *ExpressionTestResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpressionTestResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpressionTestResponse) GetBody() *ExpressionTestResponseBody  {
  return s.Body
}

func (s *ExpressionTestResponse) SetHeaders(v map[string]*string) *ExpressionTestResponse {
  s.Headers = v
  return s
}

func (s *ExpressionTestResponse) SetStatusCode(v int32) *ExpressionTestResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpressionTestResponse) SetBody(v *ExpressionTestResponseBody) *ExpressionTestResponse {
  s.Body = v
  return s
}

func (s *ExpressionTestResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

