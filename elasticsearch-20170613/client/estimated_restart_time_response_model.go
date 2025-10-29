// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedRestartTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EstimatedRestartTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EstimatedRestartTimeResponse
  GetStatusCode() *int32 
  SetBody(v *EstimatedRestartTimeResponseBody) *EstimatedRestartTimeResponse
  GetBody() *EstimatedRestartTimeResponseBody 
}

type EstimatedRestartTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EstimatedRestartTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EstimatedRestartTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s EstimatedRestartTimeResponse) GoString() string {
  return s.String()
}

func (s *EstimatedRestartTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EstimatedRestartTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EstimatedRestartTimeResponse) GetBody() *EstimatedRestartTimeResponseBody  {
  return s.Body
}

func (s *EstimatedRestartTimeResponse) SetHeaders(v map[string]*string) *EstimatedRestartTimeResponse {
  s.Headers = v
  return s
}

func (s *EstimatedRestartTimeResponse) SetStatusCode(v int32) *EstimatedRestartTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *EstimatedRestartTimeResponse) SetBody(v *EstimatedRestartTimeResponseBody) *EstimatedRestartTimeResponse {
  s.Body = v
  return s
}

func (s *EstimatedRestartTimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

