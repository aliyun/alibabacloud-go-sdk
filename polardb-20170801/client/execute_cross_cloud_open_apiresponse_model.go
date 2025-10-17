// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCrossCloudOpenAPIResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteCrossCloudOpenAPIResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteCrossCloudOpenAPIResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteCrossCloudOpenAPIResponseBody) *ExecuteCrossCloudOpenAPIResponse
  GetBody() *ExecuteCrossCloudOpenAPIResponseBody 
}

type ExecuteCrossCloudOpenAPIResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteCrossCloudOpenAPIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteCrossCloudOpenAPIResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCrossCloudOpenAPIResponse) GoString() string {
  return s.String()
}

func (s *ExecuteCrossCloudOpenAPIResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteCrossCloudOpenAPIResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteCrossCloudOpenAPIResponse) GetBody() *ExecuteCrossCloudOpenAPIResponseBody  {
  return s.Body
}

func (s *ExecuteCrossCloudOpenAPIResponse) SetHeaders(v map[string]*string) *ExecuteCrossCloudOpenAPIResponse {
  s.Headers = v
  return s
}

func (s *ExecuteCrossCloudOpenAPIResponse) SetStatusCode(v int32) *ExecuteCrossCloudOpenAPIResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteCrossCloudOpenAPIResponse) SetBody(v *ExecuteCrossCloudOpenAPIResponseBody) *ExecuteCrossCloudOpenAPIResponse {
  s.Body = v
  return s
}

func (s *ExecuteCrossCloudOpenAPIResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

