// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBackupPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteBackupPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteBackupPlanResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteBackupPlanResponseBody) *ExecuteBackupPlanResponse
  GetBody() *ExecuteBackupPlanResponseBody 
}

type ExecuteBackupPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteBackupPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBackupPlanResponse) GoString() string {
  return s.String()
}

func (s *ExecuteBackupPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteBackupPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteBackupPlanResponse) GetBody() *ExecuteBackupPlanResponseBody  {
  return s.Body
}

func (s *ExecuteBackupPlanResponse) SetHeaders(v map[string]*string) *ExecuteBackupPlanResponse {
  s.Headers = v
  return s
}

func (s *ExecuteBackupPlanResponse) SetStatusCode(v int32) *ExecuteBackupPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteBackupPlanResponse) SetBody(v *ExecuteBackupPlanResponseBody) *ExecuteBackupPlanResponse {
  s.Body = v
  return s
}

func (s *ExecuteBackupPlanResponse) Validate() error {
  return dara.Validate(s)
}

