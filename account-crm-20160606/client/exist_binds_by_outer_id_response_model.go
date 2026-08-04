// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistBindsByOuterIdResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExistBindsByOuterIdResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExistBindsByOuterIdResponse
  GetStatusCode() *int32 
  SetBody(v *ExistBindsByOuterIdResponseBody) *ExistBindsByOuterIdResponse
  GetBody() *ExistBindsByOuterIdResponseBody 
}

type ExistBindsByOuterIdResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExistBindsByOuterIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExistBindsByOuterIdResponse) String() string {
  return dara.Prettify(s)
}

func (s ExistBindsByOuterIdResponse) GoString() string {
  return s.String()
}

func (s *ExistBindsByOuterIdResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExistBindsByOuterIdResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExistBindsByOuterIdResponse) GetBody() *ExistBindsByOuterIdResponseBody  {
  return s.Body
}

func (s *ExistBindsByOuterIdResponse) SetHeaders(v map[string]*string) *ExistBindsByOuterIdResponse {
  s.Headers = v
  return s
}

func (s *ExistBindsByOuterIdResponse) SetStatusCode(v int32) *ExistBindsByOuterIdResponse {
  s.StatusCode = &v
  return s
}

func (s *ExistBindsByOuterIdResponse) SetBody(v *ExistBindsByOuterIdResponseBody) *ExistBindsByOuterIdResponse {
  s.Body = v
  return s
}

func (s *ExistBindsByOuterIdResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

