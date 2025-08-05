// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBackupPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBackupPlanResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBackupPlanResponseBody) *EnableBackupPlanResponse
  GetBody() *EnableBackupPlanResponseBody 
}

type EnableBackupPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBackupPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupPlanResponse) GoString() string {
  return s.String()
}

func (s *EnableBackupPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBackupPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBackupPlanResponse) GetBody() *EnableBackupPlanResponseBody  {
  return s.Body
}

func (s *EnableBackupPlanResponse) SetHeaders(v map[string]*string) *EnableBackupPlanResponse {
  s.Headers = v
  return s
}

func (s *EnableBackupPlanResponse) SetStatusCode(v int32) *EnableBackupPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBackupPlanResponse) SetBody(v *EnableBackupPlanResponseBody) *EnableBackupPlanResponse {
  s.Body = v
  return s
}

func (s *EnableBackupPlanResponse) Validate() error {
  return dara.Validate(s)
}

