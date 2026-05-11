// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndDialogueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndDialogueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndDialogueResponse
  GetStatusCode() *int32 
  SetBody(v *EndDialogueResponseBody) *EndDialogueResponse
  GetBody() *EndDialogueResponseBody 
}

type EndDialogueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndDialogueResponse) String() string {
  return dara.Prettify(s)
}

func (s EndDialogueResponse) GoString() string {
  return s.String()
}

func (s *EndDialogueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndDialogueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndDialogueResponse) GetBody() *EndDialogueResponseBody  {
  return s.Body
}

func (s *EndDialogueResponse) SetHeaders(v map[string]*string) *EndDialogueResponse {
  s.Headers = v
  return s
}

func (s *EndDialogueResponse) SetStatusCode(v int32) *EndDialogueResponse {
  s.StatusCode = &v
  return s
}

func (s *EndDialogueResponse) SetBody(v *EndDialogueResponseBody) *EndDialogueResponse {
  s.Body = v
  return s
}

func (s *EndDialogueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

