// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkWarehouseBatchSQLResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSparkWarehouseBatchSQLResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSparkWarehouseBatchSQLResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSparkWarehouseBatchSQLResponseBody) *ExecuteSparkWarehouseBatchSQLResponse
  GetBody() *ExecuteSparkWarehouseBatchSQLResponseBody 
}

type ExecuteSparkWarehouseBatchSQLResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSparkWarehouseBatchSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSparkWarehouseBatchSQLResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkWarehouseBatchSQLResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) GetBody() *ExecuteSparkWarehouseBatchSQLResponseBody  {
  return s.Body
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) SetHeaders(v map[string]*string) *ExecuteSparkWarehouseBatchSQLResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) SetStatusCode(v int32) *ExecuteSparkWarehouseBatchSQLResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) SetBody(v *ExecuteSparkWarehouseBatchSQLResponseBody) *ExecuteSparkWarehouseBatchSQLResponse {
  s.Body = v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

