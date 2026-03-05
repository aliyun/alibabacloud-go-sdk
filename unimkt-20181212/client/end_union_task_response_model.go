// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndUnionTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndUnionTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndUnionTaskResponse
  GetStatusCode() *int32 
  SetBody(v *EndUnionTaskResponseBody) *EndUnionTaskResponse
  GetBody() *EndUnionTaskResponseBody 
}

type EndUnionTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndUnionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndUnionTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s EndUnionTaskResponse) GoString() string {
  return s.String()
}

func (s *EndUnionTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndUnionTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndUnionTaskResponse) GetBody() *EndUnionTaskResponseBody  {
  return s.Body
}

func (s *EndUnionTaskResponse) SetHeaders(v map[string]*string) *EndUnionTaskResponse {
  s.Headers = v
  return s
}

func (s *EndUnionTaskResponse) SetStatusCode(v int32) *EndUnionTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *EndUnionTaskResponse) SetBody(v *EndUnionTaskResponseBody) *EndUnionTaskResponse {
  s.Body = v
  return s
}

func (s *EndUnionTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

