// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHanaBackupPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHanaBackupPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHanaBackupPlanResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHanaBackupPlanResponseBody) *EnableHanaBackupPlanResponse
  GetBody() *EnableHanaBackupPlanResponseBody 
}

type EnableHanaBackupPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHanaBackupPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHanaBackupPlanResponse) GoString() string {
  return s.String()
}

func (s *EnableHanaBackupPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHanaBackupPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHanaBackupPlanResponse) GetBody() *EnableHanaBackupPlanResponseBody  {
  return s.Body
}

func (s *EnableHanaBackupPlanResponse) SetHeaders(v map[string]*string) *EnableHanaBackupPlanResponse {
  s.Headers = v
  return s
}

func (s *EnableHanaBackupPlanResponse) SetStatusCode(v int32) *EnableHanaBackupPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHanaBackupPlanResponse) SetBody(v *EnableHanaBackupPlanResponseBody) *EnableHanaBackupPlanResponse {
  s.Body = v
  return s
}

func (s *EnableHanaBackupPlanResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

