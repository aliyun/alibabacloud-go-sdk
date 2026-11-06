// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDocumentDBResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBClusterDocumentDBResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBClusterDocumentDBResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBClusterDocumentDBResponseBody) *EnableDBClusterDocumentDBResponse
  GetBody() *EnableDBClusterDocumentDBResponseBody 
}

type EnableDBClusterDocumentDBResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBClusterDocumentDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBClusterDocumentDBResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDocumentDBResponse) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDocumentDBResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBClusterDocumentDBResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBClusterDocumentDBResponse) GetBody() *EnableDBClusterDocumentDBResponseBody  {
  return s.Body
}

func (s *EnableDBClusterDocumentDBResponse) SetHeaders(v map[string]*string) *EnableDBClusterDocumentDBResponse {
  s.Headers = v
  return s
}

func (s *EnableDBClusterDocumentDBResponse) SetStatusCode(v int32) *EnableDBClusterDocumentDBResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBClusterDocumentDBResponse) SetBody(v *EnableDBClusterDocumentDBResponseBody) *EnableDBClusterDocumentDBResponse {
  s.Body = v
  return s
}

func (s *EnableDBClusterDocumentDBResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

