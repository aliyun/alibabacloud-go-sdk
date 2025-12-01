// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupLogResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBackupLogResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBackupLogResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBackupLogResponseBody) *EnableBackupLogResponse
  GetBody() *EnableBackupLogResponseBody 
}

type EnableBackupLogResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBackupLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBackupLogResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupLogResponse) GoString() string {
  return s.String()
}

func (s *EnableBackupLogResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBackupLogResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBackupLogResponse) GetBody() *EnableBackupLogResponseBody  {
  return s.Body
}

func (s *EnableBackupLogResponse) SetHeaders(v map[string]*string) *EnableBackupLogResponse {
  s.Headers = v
  return s
}

func (s *EnableBackupLogResponse) SetStatusCode(v int32) *EnableBackupLogResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBackupLogResponse) SetBody(v *EnableBackupLogResponseBody) *EnableBackupLogResponse {
  s.Body = v
  return s
}

func (s *EnableBackupLogResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

