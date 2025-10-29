// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPlaylistResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditPlaylistResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditPlaylistResponse
  GetStatusCode() *int32 
  SetBody(v *EditPlaylistResponseBody) *EditPlaylistResponse
  GetBody() *EditPlaylistResponseBody 
}

type EditPlaylistResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditPlaylistResponse) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistResponse) GoString() string {
  return s.String()
}

func (s *EditPlaylistResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditPlaylistResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditPlaylistResponse) GetBody() *EditPlaylistResponseBody  {
  return s.Body
}

func (s *EditPlaylistResponse) SetHeaders(v map[string]*string) *EditPlaylistResponse {
  s.Headers = v
  return s
}

func (s *EditPlaylistResponse) SetStatusCode(v int32) *EditPlaylistResponse {
  s.StatusCode = &v
  return s
}

func (s *EditPlaylistResponse) SetBody(v *EditPlaylistResponseBody) *EditPlaylistResponse {
  s.Body = v
  return s
}

func (s *EditPlaylistResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

