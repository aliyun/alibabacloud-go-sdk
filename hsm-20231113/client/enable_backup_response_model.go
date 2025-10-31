// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBackupResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBackupResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBackupResponseBody) *EnableBackupResponse
  GetBody() *EnableBackupResponseBody 
}

type EnableBackupResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBackupResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupResponse) GoString() string {
  return s.String()
}

func (s *EnableBackupResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBackupResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBackupResponse) GetBody() *EnableBackupResponseBody  {
  return s.Body
}

func (s *EnableBackupResponse) SetHeaders(v map[string]*string) *EnableBackupResponse {
  s.Headers = v
  return s
}

func (s *EnableBackupResponse) SetStatusCode(v int32) *EnableBackupResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBackupResponse) SetBody(v *EnableBackupResponseBody) *EnableBackupResponse {
  s.Body = v
  return s
}

func (s *EnableBackupResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

