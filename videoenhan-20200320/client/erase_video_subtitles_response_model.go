// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoSubtitlesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EraseVideoSubtitlesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EraseVideoSubtitlesResponse
  GetStatusCode() *int32 
  SetBody(v *EraseVideoSubtitlesResponseBody) *EraseVideoSubtitlesResponse
  GetBody() *EraseVideoSubtitlesResponseBody 
}

type EraseVideoSubtitlesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EraseVideoSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EraseVideoSubtitlesResponse) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoSubtitlesResponse) GoString() string {
  return s.String()
}

func (s *EraseVideoSubtitlesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EraseVideoSubtitlesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EraseVideoSubtitlesResponse) GetBody() *EraseVideoSubtitlesResponseBody  {
  return s.Body
}

func (s *EraseVideoSubtitlesResponse) SetHeaders(v map[string]*string) *EraseVideoSubtitlesResponse {
  s.Headers = v
  return s
}

func (s *EraseVideoSubtitlesResponse) SetStatusCode(v int32) *EraseVideoSubtitlesResponse {
  s.StatusCode = &v
  return s
}

func (s *EraseVideoSubtitlesResponse) SetBody(v *EraseVideoSubtitlesResponseBody) *EraseVideoSubtitlesResponse {
  s.Body = v
  return s
}

func (s *EraseVideoSubtitlesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

