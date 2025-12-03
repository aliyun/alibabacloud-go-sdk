// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueBackupResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHBaseueBackupResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHBaseueBackupResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHBaseueBackupResponseBody) *EnableHBaseueBackupResponse
  GetBody() *EnableHBaseueBackupResponseBody 
}

type EnableHBaseueBackupResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHBaseueBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHBaseueBackupResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueBackupResponse) GoString() string {
  return s.String()
}

func (s *EnableHBaseueBackupResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHBaseueBackupResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHBaseueBackupResponse) GetBody() *EnableHBaseueBackupResponseBody  {
  return s.Body
}

func (s *EnableHBaseueBackupResponse) SetHeaders(v map[string]*string) *EnableHBaseueBackupResponse {
  s.Headers = v
  return s
}

func (s *EnableHBaseueBackupResponse) SetStatusCode(v int32) *EnableHBaseueBackupResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHBaseueBackupResponse) SetBody(v *EnableHBaseueBackupResponseBody) *EnableHBaseueBackupResponse {
  s.Body = v
  return s
}

func (s *EnableHBaseueBackupResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

