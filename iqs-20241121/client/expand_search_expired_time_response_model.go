// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandSearchExpiredTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpandSearchExpiredTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpandSearchExpiredTimeResponse
  GetStatusCode() *int32 
  SetBody(v *OperationResult) *ExpandSearchExpiredTimeResponse
  GetBody() *OperationResult 
}

type ExpandSearchExpiredTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *OperationResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandSearchExpiredTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpandSearchExpiredTimeResponse) GoString() string {
  return s.String()
}

func (s *ExpandSearchExpiredTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpandSearchExpiredTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpandSearchExpiredTimeResponse) GetBody() *OperationResult  {
  return s.Body
}

func (s *ExpandSearchExpiredTimeResponse) SetHeaders(v map[string]*string) *ExpandSearchExpiredTimeResponse {
  s.Headers = v
  return s
}

func (s *ExpandSearchExpiredTimeResponse) SetStatusCode(v int32) *ExpandSearchExpiredTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpandSearchExpiredTimeResponse) SetBody(v *OperationResult) *ExpandSearchExpiredTimeResponse {
  s.Body = v
  return s
}

func (s *ExpandSearchExpiredTimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

