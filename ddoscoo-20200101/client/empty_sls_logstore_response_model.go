// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptySlsLogstoreResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EmptySlsLogstoreResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EmptySlsLogstoreResponse
  GetStatusCode() *int32 
  SetBody(v *EmptySlsLogstoreResponseBody) *EmptySlsLogstoreResponse
  GetBody() *EmptySlsLogstoreResponseBody 
}

type EmptySlsLogstoreResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EmptySlsLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmptySlsLogstoreResponse) String() string {
  return dara.Prettify(s)
}

func (s EmptySlsLogstoreResponse) GoString() string {
  return s.String()
}

func (s *EmptySlsLogstoreResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EmptySlsLogstoreResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EmptySlsLogstoreResponse) GetBody() *EmptySlsLogstoreResponseBody  {
  return s.Body
}

func (s *EmptySlsLogstoreResponse) SetHeaders(v map[string]*string) *EmptySlsLogstoreResponse {
  s.Headers = v
  return s
}

func (s *EmptySlsLogstoreResponse) SetStatusCode(v int32) *EmptySlsLogstoreResponse {
  s.StatusCode = &v
  return s
}

func (s *EmptySlsLogstoreResponse) SetBody(v *EmptySlsLogstoreResponseBody) *EmptySlsLogstoreResponse {
  s.Body = v
  return s
}

func (s *EmptySlsLogstoreResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

