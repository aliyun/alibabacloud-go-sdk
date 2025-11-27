// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateLocalExtendDiskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateLocalExtendDiskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateLocalExtendDiskResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateLocalExtendDiskResponseBody) *EvaluateLocalExtendDiskResponse
  GetBody() *EvaluateLocalExtendDiskResponseBody 
}

type EvaluateLocalExtendDiskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateLocalExtendDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateLocalExtendDiskResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateLocalExtendDiskResponse) GoString() string {
  return s.String()
}

func (s *EvaluateLocalExtendDiskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateLocalExtendDiskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateLocalExtendDiskResponse) GetBody() *EvaluateLocalExtendDiskResponseBody  {
  return s.Body
}

func (s *EvaluateLocalExtendDiskResponse) SetHeaders(v map[string]*string) *EvaluateLocalExtendDiskResponse {
  s.Headers = v
  return s
}

func (s *EvaluateLocalExtendDiskResponse) SetStatusCode(v int32) *EvaluateLocalExtendDiskResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponse) SetBody(v *EvaluateLocalExtendDiskResponseBody) *EvaluateLocalExtendDiskResponse {
  s.Body = v
  return s
}

func (s *EvaluateLocalExtendDiskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

