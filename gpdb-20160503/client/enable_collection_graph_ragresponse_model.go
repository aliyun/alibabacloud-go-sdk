// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCollectionGraphRAGResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCollectionGraphRAGResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCollectionGraphRAGResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCollectionGraphRAGResponseBody) *EnableCollectionGraphRAGResponse
  GetBody() *EnableCollectionGraphRAGResponseBody 
}

type EnableCollectionGraphRAGResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCollectionGraphRAGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCollectionGraphRAGResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCollectionGraphRAGResponse) GoString() string {
  return s.String()
}

func (s *EnableCollectionGraphRAGResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCollectionGraphRAGResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCollectionGraphRAGResponse) GetBody() *EnableCollectionGraphRAGResponseBody  {
  return s.Body
}

func (s *EnableCollectionGraphRAGResponse) SetHeaders(v map[string]*string) *EnableCollectionGraphRAGResponse {
  s.Headers = v
  return s
}

func (s *EnableCollectionGraphRAGResponse) SetStatusCode(v int32) *EnableCollectionGraphRAGResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCollectionGraphRAGResponse) SetBody(v *EnableCollectionGraphRAGResponseBody) *EnableCollectionGraphRAGResponse {
  s.Body = v
  return s
}

func (s *EnableCollectionGraphRAGResponse) Validate() error {
  return dara.Validate(s)
}

