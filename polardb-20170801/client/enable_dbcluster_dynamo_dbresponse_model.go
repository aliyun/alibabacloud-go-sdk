// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDynamoDBResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBClusterDynamoDBResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBClusterDynamoDBResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBClusterDynamoDBResponseBody) *EnableDBClusterDynamoDBResponse
  GetBody() *EnableDBClusterDynamoDBResponseBody 
}

type EnableDBClusterDynamoDBResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBClusterDynamoDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBClusterDynamoDBResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDynamoDBResponse) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDynamoDBResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBClusterDynamoDBResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBClusterDynamoDBResponse) GetBody() *EnableDBClusterDynamoDBResponseBody  {
  return s.Body
}

func (s *EnableDBClusterDynamoDBResponse) SetHeaders(v map[string]*string) *EnableDBClusterDynamoDBResponse {
  s.Headers = v
  return s
}

func (s *EnableDBClusterDynamoDBResponse) SetStatusCode(v int32) *EnableDBClusterDynamoDBResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBClusterDynamoDBResponse) SetBody(v *EnableDBClusterDynamoDBResponseBody) *EnableDBClusterDynamoDBResponse {
  s.Body = v
  return s
}

func (s *EnableDBClusterDynamoDBResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

