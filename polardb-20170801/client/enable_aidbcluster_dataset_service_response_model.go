// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAIDBClusterDatasetServiceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAIDBClusterDatasetServiceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAIDBClusterDatasetServiceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAIDBClusterDatasetServiceResponseBody) *EnableAIDBClusterDatasetServiceResponse
  GetBody() *EnableAIDBClusterDatasetServiceResponseBody 
}

type EnableAIDBClusterDatasetServiceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAIDBClusterDatasetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAIDBClusterDatasetServiceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAIDBClusterDatasetServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableAIDBClusterDatasetServiceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAIDBClusterDatasetServiceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAIDBClusterDatasetServiceResponse) GetBody() *EnableAIDBClusterDatasetServiceResponseBody  {
  return s.Body
}

func (s *EnableAIDBClusterDatasetServiceResponse) SetHeaders(v map[string]*string) *EnableAIDBClusterDatasetServiceResponse {
  s.Headers = v
  return s
}

func (s *EnableAIDBClusterDatasetServiceResponse) SetStatusCode(v int32) *EnableAIDBClusterDatasetServiceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceResponse) SetBody(v *EnableAIDBClusterDatasetServiceResponseBody) *EnableAIDBClusterDatasetServiceResponse {
  s.Body = v
  return s
}

func (s *EnableAIDBClusterDatasetServiceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

