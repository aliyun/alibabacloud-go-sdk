// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConversationShareResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableConversationShareResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableConversationShareResponse
  GetStatusCode() *int32 
  SetBody(v *EnableConversationShareResponseBody) *EnableConversationShareResponse
  GetBody() *EnableConversationShareResponseBody 
}

type EnableConversationShareResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableConversationShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableConversationShareResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableConversationShareResponse) GoString() string {
  return s.String()
}

func (s *EnableConversationShareResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableConversationShareResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableConversationShareResponse) GetBody() *EnableConversationShareResponseBody  {
  return s.Body
}

func (s *EnableConversationShareResponse) SetHeaders(v map[string]*string) *EnableConversationShareResponse {
  s.Headers = v
  return s
}

func (s *EnableConversationShareResponse) SetStatusCode(v int32) *EnableConversationShareResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableConversationShareResponse) SetBody(v *EnableConversationShareResponseBody) *EnableConversationShareResponse {
  s.Body = v
  return s
}

func (s *EnableConversationShareResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

