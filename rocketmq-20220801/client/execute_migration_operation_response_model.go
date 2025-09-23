// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMigrationOperationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteMigrationOperationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteMigrationOperationResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteMigrationOperationResponseBody) *ExecuteMigrationOperationResponse
  GetBody() *ExecuteMigrationOperationResponseBody 
}

type ExecuteMigrationOperationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteMigrationOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteMigrationOperationResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationResponse) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteMigrationOperationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteMigrationOperationResponse) GetBody() *ExecuteMigrationOperationResponseBody  {
  return s.Body
}

func (s *ExecuteMigrationOperationResponse) SetHeaders(v map[string]*string) *ExecuteMigrationOperationResponse {
  s.Headers = v
  return s
}

func (s *ExecuteMigrationOperationResponse) SetStatusCode(v int32) *ExecuteMigrationOperationResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteMigrationOperationResponse) SetBody(v *ExecuteMigrationOperationResponseBody) *ExecuteMigrationOperationResponse {
  s.Body = v
  return s
}

func (s *ExecuteMigrationOperationResponse) Validate() error {
  return dara.Validate(s)
}

