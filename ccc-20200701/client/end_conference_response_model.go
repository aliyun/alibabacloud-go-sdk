// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndConferenceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndConferenceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndConferenceResponse
  GetStatusCode() *int32 
  SetBody(v *EndConferenceResponseBody) *EndConferenceResponse
  GetBody() *EndConferenceResponseBody 
}

type EndConferenceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndConferenceResponse) String() string {
  return dara.Prettify(s)
}

func (s EndConferenceResponse) GoString() string {
  return s.String()
}

func (s *EndConferenceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndConferenceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndConferenceResponse) GetBody() *EndConferenceResponseBody  {
  return s.Body
}

func (s *EndConferenceResponse) SetHeaders(v map[string]*string) *EndConferenceResponse {
  s.Headers = v
  return s
}

func (s *EndConferenceResponse) SetStatusCode(v int32) *EndConferenceResponse {
  s.StatusCode = &v
  return s
}

func (s *EndConferenceResponse) SetBody(v *EndConferenceResponseBody) *EndConferenceResponse {
  s.Body = v
  return s
}

func (s *EndConferenceResponse) Validate() error {
  return dara.Validate(s)
}

