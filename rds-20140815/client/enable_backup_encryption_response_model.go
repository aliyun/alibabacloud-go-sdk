// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupEncryptionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBackupEncryptionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBackupEncryptionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBackupEncryptionResponseBody) *EnableBackupEncryptionResponse
  GetBody() *EnableBackupEncryptionResponseBody 
}

type EnableBackupEncryptionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBackupEncryptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBackupEncryptionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupEncryptionResponse) GoString() string {
  return s.String()
}

func (s *EnableBackupEncryptionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBackupEncryptionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBackupEncryptionResponse) GetBody() *EnableBackupEncryptionResponseBody  {
  return s.Body
}

func (s *EnableBackupEncryptionResponse) SetHeaders(v map[string]*string) *EnableBackupEncryptionResponse {
  s.Headers = v
  return s
}

func (s *EnableBackupEncryptionResponse) SetStatusCode(v int32) *EnableBackupEncryptionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBackupEncryptionResponse) SetBody(v *EnableBackupEncryptionResponseBody) *EnableBackupEncryptionResponse {
  s.Body = v
  return s
}

func (s *EnableBackupEncryptionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

