// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedLogstashRestartTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EstimatedLogstashRestartTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EstimatedLogstashRestartTimeResponse
  GetStatusCode() *int32 
  SetBody(v *EstimatedLogstashRestartTimeResponseBody) *EstimatedLogstashRestartTimeResponse
  GetBody() *EstimatedLogstashRestartTimeResponseBody 
}

type EstimatedLogstashRestartTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EstimatedLogstashRestartTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EstimatedLogstashRestartTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponse) GoString() string {
  return s.String()
}

func (s *EstimatedLogstashRestartTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EstimatedLogstashRestartTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EstimatedLogstashRestartTimeResponse) GetBody() *EstimatedLogstashRestartTimeResponseBody  {
  return s.Body
}

func (s *EstimatedLogstashRestartTimeResponse) SetHeaders(v map[string]*string) *EstimatedLogstashRestartTimeResponse {
  s.Headers = v
  return s
}

func (s *EstimatedLogstashRestartTimeResponse) SetStatusCode(v int32) *EstimatedLogstashRestartTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *EstimatedLogstashRestartTimeResponse) SetBody(v *EstimatedLogstashRestartTimeResponseBody) *EstimatedLogstashRestartTimeResponse {
  s.Body = v
  return s
}

func (s *EstimatedLogstashRestartTimeResponse) Validate() error {
  return dara.Validate(s)
}

