// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSdlProtectedAssetResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSdlProtectedAssetResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSdlProtectedAssetResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSdlProtectedAssetResponseBody) *EnableSdlProtectedAssetResponse
  GetBody() *EnableSdlProtectedAssetResponseBody 
}

type EnableSdlProtectedAssetResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSdlProtectedAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSdlProtectedAssetResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSdlProtectedAssetResponse) GoString() string {
  return s.String()
}

func (s *EnableSdlProtectedAssetResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSdlProtectedAssetResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSdlProtectedAssetResponse) GetBody() *EnableSdlProtectedAssetResponseBody  {
  return s.Body
}

func (s *EnableSdlProtectedAssetResponse) SetHeaders(v map[string]*string) *EnableSdlProtectedAssetResponse {
  s.Headers = v
  return s
}

func (s *EnableSdlProtectedAssetResponse) SetStatusCode(v int32) *EnableSdlProtectedAssetResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSdlProtectedAssetResponse) SetBody(v *EnableSdlProtectedAssetResponseBody) *EnableSdlProtectedAssetResponse {
  s.Body = v
  return s
}

func (s *EnableSdlProtectedAssetResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

